package scraper

import (
	"context"
	"fmt"
	"io"
	"regexp"
)

var (
	googleAnalyticsOld = regexp.MustCompile(`UA-\d+(-\d+)?`)
	googleAnalyticsNew = regexp.MustCompile(`G-\w{8,10}`)
)

// Analytics looks for analytics tracker IDs on the page.
func Analytics(ctx context.Context, s Settings) error {
	s.Logger.Println("Report Analytics Trackers")
	defer s.Logger.Println("End Report Analytics Trackers")

	page, err := s.Client.Get(s.URL.String())
	if err != nil {
		return err
	}

	pageContents, err := io.ReadAll(page.Body)
	page.Body.Close()
	if err != nil {
		return err
	}

	trackers := []struct {
		name    string
		matcher *regexp.Regexp
	}{
		{"Google Analytics Old", googleAnalyticsOld},
		{"Google Analytics New", googleAnalyticsNew},
	}

	s.Reporter.Heading("Analytics Trackers")

	s.Reporter.Paragraph("These values are used by site owners to associate their website with trackers.")
	s.Reporter.Paragraph("Sites that share these values are often owned or managed by the same entity.")
	s.Reporter.Paragraph("Occasionally, misconfigured sites may make these values available to search engines.")

	for _, tracker := range trackers {
		results := tracker.matcher.FindAllString(string(pageContents), -1)

		s.Reporter.Table(func(w io.Writer) {
			fmt.Fprintf(w, "%s:\t%q\n", tracker.name, results)
		})
	}

	return nil
}
