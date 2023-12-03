package metrics

const (
	httpHandlerLabel = "handler"
	httpCodeLabel    = "code"
	httpMethodLabel  = "method"
)

var (
	HttpResponseTime = newHistogramVec(
		"response_time_seconds",
		"Histogram of application RT for any kind of requests seconds",
		TimeBucketsMedium,
		httpHandlerLabel, httpCodeLabel, httpMethodLabel,
	)

	HttpRequestTotalCollector = newCounterVec(
		"request_total",
		"Counter of http requests for any HTTP based requests",
		httpHandlerLabel, httpCodeLabel, httpMethodLabel,
	)
)

// nolint: gochecknoinits
func init() {
	mustRegister(HttpResponseTime, HttpRequestTotalCollector)
}
