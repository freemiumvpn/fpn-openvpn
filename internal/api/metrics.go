package api

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/freemiumvpn/fpn-openvpn-server/internal/openvpn/logs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricCreateSession = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fpn_openvpn_create_session",
		Help: "Sessions Created by Users",
	})

	metricDeleteSession = promauto.NewCounter(prometheus.CounterOpts{
		Name: "fpn_openvpn_delete_session",
		Help: "Sessions Deleted by Users",
	})

	metricActiveSessions = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "fpn_openvpn_active_sessions",
		Help: "Current Active VPN sessions",
	})
)

const OPENVPN_LOGS = "/tmp/openvpn-status.log"

func createActviceSessionsCounter() {
	for _ = range time.Tick(time.Minute) {
		l, err := logs.ParseLogs(OPENVPN_LOGS)
		if err != nil {
			spew.Dump("error", err)
			return
		}
		// metricActiveSessions.WithLabelValues()
		totalClients := float64(len(l.Clients))
		metricActiveSessions.Set(totalClients)
	}
}
