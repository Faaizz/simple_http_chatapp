package misc

import (
	"net/http"

	"github.com/aws/aws-xray-sdk-go/xray"
)

func ConfigXRay(xda string) error {
	return xray.Configure(xray.Config{
		DaemonAddr: xda,
	})
}

func XRayMw(next http.Handler) http.Handler {
	return xray.Handler(
		xray.NewFixedSegmentNamer("MyApp"),
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			}))
}
