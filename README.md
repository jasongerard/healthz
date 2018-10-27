# healthz

healthz is a simple package to add liveness and readiness probes to your application, primarily for kubernetes support

Probes are defined as `type CheckFunc func() bool`. If this function returns true, the probe passes.

healthz will expose two HTTP endpoints:
- `/healthz/alive`
- `/healthz/ready`

By default, both use a simple check that always returns true. For passsing probes, the endpoints return 200 OK. For failing probes the 
endpoints return 503 Service Unavailable. 

## Example usage

The following code is from example/example.go


```
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jasongerard/healthz"
)

var ready = false

func main() {

	errs := make(chan error, 2)

	// setup signal handler (capture CTRL+C)
	go func() {
		signalChan := make(chan os.Signal)
		signal.Notify(signalChan, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-signalChan)
	}()

	// set up our healthcheck server
	go func() {
		healthz.ReadinessCheck(func() bool {
			return ready
		})

		mux := healthz.CreateMux()
		srv := &http.Server{
			Handler: mux,
			Addr:    ":8123",
		}

		errs <- srv.ListenAndServe()
	}()

	// set up useful stuff for your app to do here
	// simulate work
	time.Sleep(5 * time.Second)

	// set up is done, mark app as ready
	ready = true

	// block on errs channel until we get a SIGINT or the HTTP server exits
	fmt.Printf("program terminating: %v\n", <-errs)

	// any requests to http://localhost:8123/healthz/ready will return 503 for the first 5 seconds
	// and then begin returning 200 OK
	// run and test with `while true; do curl -I http://localhost:8123/healthz/ready; sleep 1; done`
}
```
