package scraper

import (
	"context"
	"log"
	"net/http"
	"net/url"

	"github.com/josephlewis42/noodle/internal/reporter"
)

// Settings has the scrape settings.
type Settings struct {
	// URL contains the site to scrape.
	URL *url.URL

	// Client has the HTTP client used to scrape.
	Client *http.Client

	// Logger contains a logger that a scraper can use to report
	// progress information.
	Logger *log.Logger

	// Reporter is used by a scraper to report information.
	Reporter reporter.Reporter
}

// Scraper performs a single fetch/decode/report operation.
type Scraper func(ctx context.Context, s Settings) error
