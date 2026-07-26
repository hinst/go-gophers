package gophers

import (
	"errors"
	"net"
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

type WebRetry struct {
	AttemptLimit int
	Delay        time.Duration
}

func (me WebRetry) isNetworkError(err error) bool {
	if err == nil {
		return false
	}
	var netErr net.Error
	return errors.As(err, &netErr)
}

func (me WebRetry) Run(client *http.Client, request *http.Request) (*http.Response, error) {
	var latestError error
	for attempt := 0; attempt < me.GetAttemptLimit(); attempt++ {
		response, currentError := client.Do(request)
		if currentError == nil {
			return response, nil
		}
		latestError = currentError
		if !me.isNetworkError(currentError) {
			break
		}
		var isLastAttempt = attempt == me.GetAttemptLimit()-1
		if !isLastAttempt {
			time.Sleep(me.GetCurrentDelay(attempt))
		}
	}
	return nil, latestError
}

func (me WebRetry) GetCurrentDelay(attempt int) time.Duration {
	var delay = me.GetDelay()
	for range attempt {
		delay *= 2
	}
	return delay
}

func (me WebRetry) GetDelay() time.Duration {
	if me.Delay > 0 {
		return me.Delay
	}
	return 2 * time.Second
}

func (me WebRetry) GetAttemptLimit() int {
	if me.AttemptLimit > 0 {
		return me.AttemptLimit
	}
	return 4
}
