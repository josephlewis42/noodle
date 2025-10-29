package scraper

import (
	"context"
	"net"
	"sort"
)

// IPs looks up the site's IP addresses.
func IPs(ctx context.Context, s Settings) error {
	s.Logger.Println("Report IP Address")
	defer s.Logger.Println("End Report IP Address")

	ips, err := net.LookupIP(s.URL.Host)
	if err != nil {
		return err
	}

	sort.Slice(ips, func(i, j int) bool {
		return ips[i].String() < ips[j].String()
	})

	s.Reporter.Heading("IP Addresses")
	s.Reporter.Paragraph("IP addresses that serve traffic for the domain name.")
	s.Reporter.Paragraph("Visiting these IPs directly may give information about the hosting provider.")

	s.Reporter.List(ips)

	return nil
}
