package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"metrics/sdk/prometheus"
	"net/http"
	"strconv"
	"sync/atomic"
)

var gRequestCount atomic.Int64

func main() {
	gRequestCount.Store(0)
	metric1 := prometheus.NewGaugeVecMetric(
		"req_count",
		"the number of requests",
		[]string{"count"})
	prometheus.RegisterMetric(metric1)

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		gRequestCount.Add(1)
		count := gRequestCount.Load()
		strCount := strconv.FormatInt(count, 10)
		metric1.Reset()
		metric1.WithLabelValues(strCount)

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, I am service1!",
		})
	})
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(prometheus.RegistryInstance(), promhttp.HandlerOpts{})))

	r.Run(":8001")
}
