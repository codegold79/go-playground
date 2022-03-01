# Collect and Check Errors

Since Go 1.13, we can use errors.Is() to check if an error has wrapped within it a certain error object.

I show that when errors are aggregated using the Kubernetes `k8s.io/apimachinery/pkg/util/errors` package, you can use errors.Is() with it.
