package middleware

import (
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
)

// Stats : stats middleware object
type Stats struct {
	Client influxdb2.Client
	Influx api.WriteAPIBlocking
}
