package gophers

import (
	"errors"
	"net"
	"net/http"
	"time"
)

type WebRetry struct {
	AttemptLimit int
	Delay        time.Duration
}

const webRetryDefaultAttemptLimit = 4

func (me WebRetry) isNetworkError(err error) bool {
	if err == nil {
		return false
	}
	var netErr net.Error
	return errors.As(err, &netErr)
}

func (me WebRetry) Run(client *http.Client, requestFactory RequestFactory) (*http.Response, error) {
	var latestError error
	for attempt := 0; attempt < me.GetAttemptLimit(); attempt++ {
		var request = requestFactory()
		var response, currentError = client.Do(request)
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
	return webRetryDefaultAttemptLimit
}
