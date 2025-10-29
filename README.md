# Noodle: Find information about catfishing domains

Noodle takes a URL and finds information about the page, site, and domain
so you can find associated domains.

## Why was it made?

I made noodle to investigate a business I thought was shady.
Using it, I was able to piece together a network of sites run by the same
organization funneling (desperate) people into a core business.

## Installing & Usage

To install, you must have a recent version of go:

```sh
go install -v github.com/josephlewis42/noodle@latest
```

To use, run `noodle` with a URL:

```sh
noodle https://example.com > report.md
```

## What does Noodle find?

* Comments in the HTML page.
* Analytics tracker IDs that can be correlated with other domains.
* Certificate information that can lead to additional domains.
* WHOIS information.
* `robots.txt` which can point to interesting pages.

**Future ideas**

* `.well-known` services https://en.wikipedia.org/wiki/List_of_/.well-known/_services_offered_by_webservers
* RSS
* URLs of sites linked to
* Privacy/security policy text that might be shared
* DNS records
* Fingerprints left by the tools used to build the site(s)
* Affiliate codes on links
* An interactive mode to explore sites

## LICENSE

Licensed under the MIT license.