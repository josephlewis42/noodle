package scraper

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
)

// Robots looks for a robots.txt file with potentially private information.
func Robots(ctx context.Context, s Settings) error {
	s.Logger.Println("Robots Information")
	defer s.Logger.Println("End Robots Information")

	robots := &url.URL{}
	robots.Path = "/robots.txt"
	robots.Host = s.URL.Host
	robots.Scheme = "https"
	if s.URL.Scheme != "" {
		robots.Scheme = s.URL.Scheme
	}

	page, err := s.Client.Get(robots.String())
	if err != nil {
		s.Logger.Println(err.Error())
		return nil
	}

	pageContents, err := ioutil.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		return err
	}

	s.Reporter.Heading("Robots.txt")

	s.Reporter.Paragraph("The contents of the site's robots.txt file.")

	s.Reporter.List([]string{
		"Paths denied to all ('*') robots are used to hide things from search engines and automated tools.",
		fmt.Sprintf("Paths listed here are relative to 'https://%s'.", s.URL.Host),
		"Sometimes a link to a sitemap is included. Sitemaps are documents listing pages on a site for search engines.",
	})

	s.Reporter.Pre(pageContents)

	return nil
}
