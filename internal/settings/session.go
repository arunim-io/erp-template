package settings

import "time"

func DefaultSessionConfig() SessionConfig {
	return SessionConfig{Cookie: SessionCookieConfig{
		Age:                  time.Duration((2 * 7 * 24 * time.Hour).Seconds()),
		Name:                 "sessionHash",
		Path:                 "/",
		SameSite:             CookieSameSiteLax,
		Secure:               false,
		ExpireAtBrowserClose: false,
		SaveEveryRequest:     false,
	}}
}

// The settings for sessions.
type SessionConfig struct {
	Cookie SessionCookieConfig `koanf:"cookie"`
}

// The settings for session cookies.
type SessionCookieConfig struct {
	// the age of session cookies, in seconds.
	Age time.Duration `koanf:"age"`
	// The name of the cookie to use for sessions.
	Name string `koanf:"name"`
	// The path set on the session cookie.
	Path string `koanf:"path"`
	// The value of the SameSite flag on the session cookie. This flag prevents the cookie from being sent in cross-site requests thus preventing CSRF attacks and making some methods of stealing session cookie impossible.
	SameSite CookieSameSite `koanf:"same_site"`
	// Whether to use a secure cookie for the session cookie. If this is set to `true`, the cookie will be marked as "secure", which means browsers may ensure that the cookie is only sent under an HTTPS connection.
	Secure bool `koanf:"secure"`
	// Whether to expire the session when the user closes their browser.
	ExpireAtBrowserClose bool `koanf:"expire_at_browser_close"`
	// Whether to save the session data on every request. If this is `false` (default), then the session data will only be saved if it has been modified – that is, if any of its dictionary values have been assigned or deleted. Empty sessions won’t be created, even if this setting is active.
	SaveEveryRequest bool `koanf:"save_every_request"`
}

type CookieSameSite string

const (
	// prevents the cookie from being sent by the browser to the target site in all cross-site browsing context, even when following a regular link.
	CookieSameSiteStrict CookieSameSite = "Strict"
	// provides a balance between security and usability for websites that want to maintain user’s logged-in session after the user arrives from an external link.
	CookieSameSiteLax CookieSameSite = "Lax"
	// the session cookie will be sent with all same-site and cross-site requests.
	CookieSameSiteNone CookieSameSite = "None"
)
