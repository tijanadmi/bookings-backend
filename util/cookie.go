package util

import (
	"net/http"
	"time"
)

func GetRefreshCookie(refreshToken string, cookieName string, cookiePath string, cookieDomain string, refreshExpiry time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Path:     cookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(refreshExpiry),
		MaxAge:   int(refreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   cookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func GetExpiredRefreshCookie(cookieName string, cookiePath string, cookieDomain string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Path:     cookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   cookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}
