package scraper

import (
	"context"

	"golang.org/x/net/html"
)

// Comments adds information about the parsed web page
func Comments(ctx context.Context, s Settings) error {
	s.Logger.Println("Report Page Comments")
	defer s.Logger.Println("End Report Page Comments")

	page, err := s.Client.Get(s.URL.String())
	if err != nil {
		return err
	}

	tokenizer := html.NewTokenizer(page.Body)
	defer page.Body.Close()

	s.Reporter.Heading("HTML Comments")

	s.Reporter.Paragraph("This listing includes code or text left behind by developers or tools.")

	for {
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			goto end
		case tt == html.CommentToken:
			s.Reporter.Pre(tokenizer.Text())
		}
	}
end:
	return nil
}
