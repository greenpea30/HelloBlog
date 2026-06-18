package zjulogin

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var executionRegexp = regexp.MustCompile(`name="execution" value="([^"]+)"`)

func findSubmatch(re *regexp.Regexp, s string) string {
	m := re.FindStringSubmatch(s)
	if len(m) > 1 {
		return m[1]
	}
	return ""
}

func closeResponseBody(res *http.Response) {
	if res != nil && res.Body != nil {
		_, _ = io.Copy(io.Discard, res.Body)
		res.Body.Close()
	}
}

func resolveLocation(baseURL string, location string) string {
	if strings.HasPrefix(location, "http://") || strings.HasPrefix(location, "https://") {
		return location
	}
	base, err := url.Parse(baseURL)
	if err != nil {
		return location
	}
	resolved, err := base.Parse(location)
	if err != nil {
		return location
	}
	return resolved.String()
}
