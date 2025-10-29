package scraper

import "context"

// WellKnown reports information about Well-Known Uniform Resource Identifiers
// as defined in RFC 8615: https://tools.ietf.org/html/rfc8615.
func WellKnown(ctx context.Context, s Settings) error {

	// https://en.wikipedia.org/wiki/List_of_/.well-known/_services_offered_by_webservers

	// acme-challenge	Automated Certificate Management Environment (ACME)	RFC8555	2019-03-01
	// apple-app-site-association	An Apple service that enables secure data exchange between iOS and a website.	web
	// apple-developer-merchantid-domain-association	Apple Pay	web
	// assetlinks.json	AssetLinks protocol used to identify one or more digital assets (such as web sites or mobile apps) that are related to the hosting web site in some fashion.	Google	2015-09-28
	// caldav	Locating Services for Calendaring Extensions to WebDAV (CalDAV) and vCard Extensions to WebDAV (CardDAV)	RFC6764
	// carddav	Locating Services for Calendaring Extensions to WebDAV (CalDAV) and vCard Extensions to WebDAV (CardDAV)	RFC6764
	// change-password	Helps password managers find the URL for the change password section.	web
	// openid-configuration	OpenID Connect	OpenID	2013-08-27
	// security.txt	Standard to help organizations define the process for security researchers to disclose security vulnerabilities	web	2018-08-20
	// webfinger	WebFinger	RFC7033	2013-03-15, 2013-09-06

	return nil
}
