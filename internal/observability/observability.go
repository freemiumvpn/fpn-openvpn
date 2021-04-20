package observability

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

const METRICS_URL = "/__/metrics"

func New(ctx context.Context, address string) error {
	handler := http.NewServeMux()
	handler.Handle(METRICS_URL, promhttp.Handler())

	server := http.Server{
		Handler: handler,
		Addr:    address,
	}

	logrus.
		WithFields(logrus.Fields{
			"address": address,
			"metrics": METRICS_URL,
		}).
		Info("Starting Observability")

	return server.ListenAndServe()
}
