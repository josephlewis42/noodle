package scraper

import (
	"context"
	"fmt"

	"github.com/domainr/whois"
)

// Whois performs a whois lookup on a URL.
func Whois(ctx context.Context, s Settings) error {
	s.Logger.Println("Running whois lookup")
	defer s.Logger.Println("Finshed whois lookup")

	request, err := whois.NewRequest(s.URL.Host)
	if err != nil {
		return err
	}

	// Can happen if URL is malformed.
	if request.Query == "" {
		return fmt.Errorf("hostname was blank, is the URL malformed?")
	}

	response, err := whois.DefaultClient.Fetch(request)
	if err != nil {
		return err
	}

	s.Reporter.Heading("Whois")
	s.Reporter.Paragraph("Whois returns information about the owner of the domain name.")
	s.Reporter.List([]string{
		"Domain owners are required to provide contact information to their domain registry.",
		"Often, this information is replaced in whois lookups by the name of the registry to protect privacy.",
	})
	s.Reporter.Pre(response.Body)

	return nil
}
