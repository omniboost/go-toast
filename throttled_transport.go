package toast

import (
	"net/http"

	"golang.org/x/time/rate"
)

// ThrottledTransport Rate Limited HTTP Client
type ThrottledTransport struct {
	roundTripperWrap http.RoundTripper
	limiters         []*rate.Limiter
}

func (c *ThrottledTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for _, l := range c.limiters {
		err := l.Wait(r.Context()) // This is a blocking call. Honors the rate limit
		if err != nil {
			return nil, err
		}
	}

	return c.roundTripperWrap.RoundTrip(r)
}

// NewThrottledTransport wraps transportWrap with a rate limitter
// examle usage:
// client := http.DefaultClient
// client.Transport = NewThrottledTransport(10*time.Seconds, 60, http.DefaultTransport) allows 60 requests every 10 seconds
func NewThrottledTransport(limiters []*rate.Limiter, transportWrap http.RoundTripper) http.RoundTripper {
	return &ThrottledTransport{
		roundTripperWrap: transportWrap,
		limiters:         limiters,
	}
}
