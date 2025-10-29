package scraper

import (
	"context"
	"fmt"
	"io"
	"time"
)

// Frontmatter adds report metadata to the output.
func Frontmatter(ctx context.Context, s Settings) error {
	s.Logger.Println("Report Metadata")
	defer s.Logger.Println("End Report Metadata")

	s.Reporter.Heading("Report Metadata")

	s.Reporter.Table(func(w io.Writer) {
		fmt.Fprintf(w, "URL:\t%s\n", s.URL.String())
		fmt.Fprintf(w, "Hostname:\t%s\n", s.URL.Hostname())
		fmt.Fprintf(w, "Report time:\t%s\n", time.Now().Format(time.RFC1123))
		fmt.Fprintf(w, "Archive:\thttps://web.archive.org/web/%s\n", s.URL)
	})

	return nil
}
