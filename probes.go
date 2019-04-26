package healthz

import (
	"net/http"
)

// CheckFunc is a function that returns true if the probe is passing
type CheckFunc func() bool

func defaultCheckFunc() bool {
	return true
}

var (
	liveness  CheckFunc = defaultCheckFunc
	readiness CheckFunc = defaultCheckFunc
)

// LivenessCheck sets the probe for the /healthz/alive endpoint
func LivenessCheck(fn CheckFunc) {
	if fn == nil {
		return
	}

	liveness = fn
}

// ReadinessCheck sets the probe for the /healthz/ready endpoint
func ReadinessCheck(fn CheckFunc) {
	if fn == nil {
		return
	}

	readiness = fn
}

func checkHandler(fn CheckFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if fn() {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
}

// CreateMux creates an *http.ServeMux with handlers for /healthz/alive and /healthz/ready.
// The endpoints will return 200 OK if their respective CheckFunc returns true and 503 Service Unavailable if it returns false.
// If no CheckFunc was specified for the corresponding endpoint, the default probe is used and always returns true
func CreateMux() *http.ServeMux {

	router := http.NewServeMux()

	router.Handle("/healthz/alive", checkHandler(liveness))
	router.Handle("/healthz/ready", checkHandler(readiness))

	return router

}
