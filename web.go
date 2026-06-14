package gophers

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const ContentTypeHeader = "Content-Type"
const CacheControlHeader = "Cache-Control"
const ContentTypeJson = "application/json"

type WebFunction func(response http.ResponseWriter, request *http.Request)

func SetCacheAge(response http.ResponseWriter, duration time.Duration) {
	response.Header().Set(CacheControlHeader, "max-age="+strconv.Itoa(int(duration.Seconds())))
}

func BuildUrlQueryParams(parameters map[string]string) string {
	var parts []string
	var first = true
	for key, value := range parameters {
		if first {
			parts = append(parts, "?")
			first = false
		} else {
			parts = append(parts, "&")
		}
		parts = append(parts, url.QueryEscape(key), "=", url.QueryEscape(value))
	}
	return strings.Join(parts, "")
}
