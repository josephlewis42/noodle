package scraper

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"time"
)

// Certificate adds information about the HTTPS certificate to the output.
func Certificate(ctx context.Context, s Settings) error {
	s.Logger.Println("Report Certificate Information")
	defer s.Logger.Println("End Report Certificate Information")

	resp, err := s.Client.Head(s.URL.String())
	if err != nil {
		return err
	}

	if resp.TLS == nil {
		return nil
	}

	s.Reporter.Heading("Certificate Information")

	s.Reporter.Paragraph("Information here is taken from the site's security certificate.")

	s.Reporter.List([]string{
		"Older or bad security may indicate sites without professional security people.",
		"The Issuer field indicates which organization attests to the certificate's validity and is often a parent company or hosting provider.",
		"The Subject field gives information about who the certificate was granted to.",
		"The Subject Alternative Names field lists additional sites the certificate is valid for; often other sites hosted by the same provider or owned by the same organization.",
	})

	s.Reporter.Table(func(w io.Writer) {

		fmt.Fprintf(w, "TLS Version:\t%q\n", tlsVersionToName(resp.TLS.Version))
		fmt.Fprintf(w, "Handshake complete?\t%v\n", resp.TLS.HandshakeComplete)
		fmt.Fprintf(w, "Cipher Suite:\t%v\n", tls.CipherSuiteName(resp.TLS.CipherSuite))

		knownInsecure := false
		for _, suite := range tls.InsecureCipherSuites() {
			if suite.ID == resp.TLS.CipherSuite {
				knownInsecure = true
				break
			}
		}
		fmt.Fprintf(w, "Known Insecure?\t%v\n", knownInsecure)
		fmt.Fprintf(w, "Negotiated Protocol:\t%v\n", resp.TLS.NegotiatedProtocol)
		fmt.Fprintf(w, "Server Name:\t%v\n", resp.TLS.ServerName)

		// Certificate, on the client side this is guaranteed to be populated.
		cert := resp.TLS.PeerCertificates[0]
		fmt.Fprintf(w, "Version:\t%v\n", cert.Version)
		fmt.Fprintf(w, "Serial Number:\t%v\n", cert.SerialNumber)
		fmt.Fprintf(w, "Certificate Signature Algorithm:\t%v\n", cert.SignatureAlgorithm)
		fmt.Fprintf(w, "Issuer:\t%v\n", cert.Issuer.String())
		fmt.Fprintf(w, "Validity:\t\n")
		fmt.Fprintf(w, "Validity Not Before:\t%v\n", cert.NotBefore.Format(time.RFC1123))
		fmt.Fprintf(w, "Validity Not After:\t%v\n", cert.NotAfter.Format(time.RFC1123))
		fmt.Fprintf(w, "Validity Duration:\t%v\n", cert.NotAfter.Sub(cert.NotBefore))
		fmt.Fprintf(w, "Subject:\t%v\n", cert.Subject)
		fmt.Fprintf(w, "Subject Public Key Algorithm:\t%v\n", cert.PublicKeyAlgorithm)

		// Look for Subject Alternative Names which may indicate other hosts owned
		// by the same entity.
		// https://tools.ietf.org/html/rfc5280#section-4.2.1.6
		// ASN1 ID: 2.5.29.17
		fmt.Fprintf(w, "Subject Alternative Names:\t%q\n", cert.DNSNames)
	})

	return nil
}

func tlsVersionToName(version uint16) string {
	switch version {
	case tls.VersionSSL30:
		return "SSLv3 (insecure)"
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("search '0x%04x' on https://golang.org/pkg/crypto/tls/#ConnectionState", version)
	}
}
