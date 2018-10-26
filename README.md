# healthz

healthz is a simple package to add liveness and readiness probes to your application, primarily for kubernetes support

Probes are defined as `type CheckFunc func() bool`. If this function returns true, the probe passes.

healthz will expose two HTTP endpoints:
- `/healthz/alive`
- `/healthz/ready`

By default, both use a simple check that always returns true. For passsing probes, the endpoints return 200 OK. For failing probes the 
endpoints return 503 Service Unavailable. 

## Eample usage

See example/example.go for example usage