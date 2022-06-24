package handlers

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
)

var rateLimit = rate.NewLimiter(rate.Every(10*time.Second), 3) // 10 request every 10 seconds

func HandleAdapter(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		applog.Log.Infof("Requesting - Method: %s, URL %s", r.Method, r.URL)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		if rateLimitAllow() == false {
			applog.Log.Warnf("Rate limit to requests was exceed")
			errs.HttpToManyRequests(w)
			return
		}

		hf(w, r)
	}
}
