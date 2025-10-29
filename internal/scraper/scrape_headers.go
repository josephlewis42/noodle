package scraper

import (
	"context"
	"fmt"
	"io"
)

// Headers looks for HTTP headers
func Headers(ctx context.Context, s Settings) error {
	s.Logger.Println("Report HTTP Headers")
	defer s.Logger.Println("End Report HTTP Headers")

	resp, err := s.Client.Get(s.URL.String())
	if err != nil {
		return err
	}

	s.Reporter.Heading("HTTP Response Headers")
	s.Reporter.Paragraph("This listing includes information sent from the site to you in the HTTP request.")
	s.Reporter.Paragraph("These values can tell you more about where and how the site is hosted.")

	// TODO make deterministic
	s.Reporter.Table(func(w io.Writer) {
		for k, v := range resp.Header {
			fmt.Fprintf(w, "%s:\t%q\n", k, v)
		}
	})
	return nil
}
