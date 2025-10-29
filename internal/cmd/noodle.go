package cmd

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gregjones/httpcache"
	"github.com/josephlewis42/noodle/internal/reporter"
	"github.com/josephlewis42/noodle/internal/scraper"
	"github.com/spf13/cobra"
)

func Noodle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "noodle URL",
		Short: "Noodle looks for hidden information on websites.",
		Long:  `Noodle discovers information about websites.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			parsedURL, err := url.Parse(args[0])
			if err != nil {
				return err
			}

			// Set up transport.
			client := &http.Client{}
			client.Transport = httpcache.NewMemoryCacheTransport()

			settings := scraper.Settings{
				URL:    parsedURL,
				Client: client,
				Logger: log.New(cmd.ErrOrStderr(), "[noodle] ", log.Lmsgprefix),
				Reporter: &reporter.Markdown{
					Writer: cmd.OutOrStdout(),
				},
			}

			ctx, cancel := context.WithTimeout(cmd.Context(), 1*time.Minute)
			defer cancel()

			scrapers := []scraper.Scraper{
				scraper.Frontmatter,
				scraper.Whois,
				scraper.Analytics,
				scraper.Headers,
				scraper.IPs,
				scraper.Certificate,
				scraper.Robots,
				scraper.Comments,
			}

			for _, s := range scrapers {
				if err := s(ctx, settings); err != nil {
					return err
				}
			}

			return nil
		},
	}

	return cmd
}
