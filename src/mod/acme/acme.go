package acme

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-acme/lego/v4/certcrypto"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/go-acme/lego/v4/challenge/http01"
	"github.com/go-acme/lego/v4/lego"
	"github.com/go-acme/lego/v4/registration"
	"imuslab.com/zoraxy/mod/utils"
)

// ACMEUser represents a user in the ACME system.
type ACMEUser struct {
	Email        string
	Registration *registration.Resource
	key          crypto.PrivateKey
}

// GetEmail returns the email of the ACMEUser.
func (u *ACMEUser) GetEmail() string {
	return u.Email
}

// GetRegistration returns the registration resource of the ACMEUser.
func (u ACMEUser) GetRegistration() *registration.Resource {
	return u.Registration
}

// GetPrivateKey returns the private key of the ACMEUser.
func (u *ACMEUser) GetPrivateKey() crypto.PrivateKey {
	return u.key
}

// ACMEHandler handles ACME-related operations.
type ACMEHandler struct {
	DefaultAcmeServer string
	Port              string
}

// NewACME creates a new ACMEHandler instance.
func NewACME(acmeServer string, port string) *ACMEHandler {
	return &ACMEHandler{
		DefaultAcmeServer: acmeServer,
		Port:              port,
	}
}

// ObtainCert obtains a certificate for the specified domains.
func (a *ACMEHandler) ObtainCert(domains []string, certificateName string, email string, ca string) (bool, error) {
	log.Println("[ACME] Obtaining certificate...")

	// generate private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Println(err)
		return false, err
	}

	// create a admin user for our new generation
	adminUser := ACMEUser{
		Email: email,
		key:   privateKey,
	}

	// create config
	config := lego.NewConfig(&adminUser)

	// setup who is the issuer and the key type
	config.CADirURL = a.DefaultAcmeServer

	//Overwrite the CADir URL if set
	if ca != "" {
		caLinkOverwrite, err := loadCAApiServerFromName(ca)
		if err == nil {
			config.CADirURL = caLinkOverwrite
			log.Println("[INFO] Using " + caLinkOverwrite + " for CA Directory URL")
		} else {
			return false, errors.New("CA " + ca + " is not supported. Please contribute to the source code and add this CA's directory link.")
		}
	}

	config.Certificate.KeyType = certcrypto.RSA2048

	client, err := lego.NewClient(config)
	if err != nil {
		log.Println(err)
		return false, err
	}

	// setup how to receive challenge
	err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", a.Port))
	if err != nil {
		log.Println(err)
		return false, err
	}

	// New users will need to register
	reg, err := client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	if err != nil {
		log.Println(err)
		return false, err
	}
	adminUser.Registration = reg

	// obtain the certificate
	request := certificate.ObtainRequest{
		Domains: domains,
		Bundle:  true,
	}
	certificates, err := client.Certificate.Obtain(request)
	if err != nil {
		log.Println(err)
		return false, err
	}

	// Each certificate comes back with the cert bytes, the bytes of the client's
	// private key, and a certificate URL.
	err = ioutil.WriteFile("./conf/certs/"+certificateName+".crt", certificates.Certificate, 0777)
	if err != nil {
		log.Println(err)
		return false, err
	}
	err = ioutil.WriteFile("./conf/certs/"+certificateName+".key", certificates.PrivateKey, 0777)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, nil
}

// CheckCertificate returns a list of domains that are in expired certificates.
// It will return all domains that is in expired certificates
// *** if there is a vaild certificate contains the domain and there is a expired certificate contains the same domain
// it will said expired as well!
func (a *ACMEHandler) CheckCertificate() []string {
	// read from dir
	filenames, err := os.ReadDir("./conf/certs/")

	expiredCerts := []string{}

	if err != nil {
		log.Println(err)
		return []string{}
	}

	for _, filename := range filenames {
		certFilepath := filepath.Join("./conf/certs/", filename.Name())

		certBytes, err := os.ReadFile(certFilepath)
		if err != nil {
			// Unable to load this file
			continue
		} else {
			// Cert loaded. Check its expiry time
			block, _ := pem.Decode(certBytes)
			if block != nil {
				cert, err := x509.ParseCertificate(block.Bytes)
				if err == nil {
					elapsed := time.Since(cert.NotAfter)
					if elapsed > 0 {
						// if it is expired then add it in
						// make sure it's uniqueless
						for _, dnsName := range cert.DNSNames {
							if !contains(expiredCerts, dnsName) {
								expiredCerts = append(expiredCerts, dnsName)
							}
						}
						if !contains(expiredCerts, cert.Subject.CommonName) {
							expiredCerts = append(expiredCerts, cert.Subject.CommonName)
						}
					}
				}
			}
		}
	}

	return expiredCerts
}

// return the current port number
func (a *ACMEHandler) Getport() string {
	return a.Port
}

// contains checks if a string is present in a slice.
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// HandleGetExpiredDomains handles the HTTP GET request to retrieve the list of expired domains.
// It calls the CheckCertificate method to obtain the expired domains and sends a JSON response
// containing the list of expired domains.
func (a *ACMEHandler) HandleGetExpiredDomains(w http.ResponseWriter, r *http.Request) {
	type ExpiredDomains struct {
		Domain []string `json:"domain"`
	}

	info := ExpiredDomains{
		Domain: a.CheckCertificate(),
	}

	js, _ := json.MarshalIndent(info, "", " ")
	utils.SendJSONResponse(w, string(js))
}

// HandleRenewCertificate handles the HTTP GET request to renew a certificate for the provided domains.
// It retrieves the domains and filename parameters from the request, calls the ObtainCert method
// to renew the certificate, and sends a JSON response indicating the result of the renewal process.
func (a *ACMEHandler) HandleRenewCertificate(w http.ResponseWriter, r *http.Request) {
	domainPara, err := utils.PostPara(r, "domains")
	if err != nil {
		utils.SendErrorResponse(w, jsonEscape(err.Error()))
		return
	}

	filename, err := utils.PostPara(r, "filename")
	if err != nil {
		utils.SendErrorResponse(w, jsonEscape(err.Error()))
		return
	}

	email, err := utils.PostPara(r, "email")
	if err != nil {
		utils.SendErrorResponse(w, jsonEscape(err.Error()))
		return
	}

	ca, err := utils.PostPara(r, "ca")
	if err != nil {
		log.Println("CA not set. Using default (Let's Encrypt)")
		ca = "Let's Encrypt"
	}

	domains := strings.Split(domainPara, ",")
	result, err := a.ObtainCert(domains, filename, email, ca)
	if err != nil {
		utils.SendErrorResponse(w, jsonEscape(err.Error()))
		return
	}
	utils.SendJSONResponse(w, strconv.FormatBool(result))
}

// Escape JSON string
func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		log.Println("Unable to escape json data: " + err.Error())
		return i
	}
	s := string(b)
	return s[1 : len(s)-1]
}

// Helper function to check if a port is in use
func IsPortInUse(port int) bool {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return true // Port is in use
	}
	defer listener.Close()
	return false // Port is not in use
}