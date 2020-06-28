package nr_yml

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHostAndQueueSize(t *testing.T) {
	withContents(`
production:
  infinite_tracing:
    trace_observer:
      host: my-trace-observer
    span_events:
      queue_size: 340382
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.Equal(t, "my-trace-observer", cfg.InfiniteTracing.TraceObserver.Host)
		assert.Equal(t, 340382, cfg.InfiniteTracing.SpanEvents.QueueSize)
	})
}

func TestPort(t *testing.T) {
	withContents(`
production:
  infinite_tracing:
    trace_observer:
      port: 3402
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.Equal(t, 3402, cfg.InfiniteTracing.TraceObserver.Port)
	})
}

func TestTraceObserverHostNoOverwrite(t *testing.T) {
	withContents(`
production:
  infinite_tracing:
    trace_observer:
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		cfg.InfiniteTracing.TraceObserver.Host = "previous-host"
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.Equal(t, "previous-host", cfg.InfiniteTracing.TraceObserver.Host)
	})
}

func TestSpanEventsNoOverwrite(t *testing.T) {
	withContents(`
production:
  infinite_tracing:
    span_events:
`, t, func(filename string, t *testing.T) {
		cfg := new(newrelic.Config)
		cfg.InfiniteTracing.SpanEvents.QueueSize = 340382
		ConfigFromYamlFile(filename)(cfg)
		assert.NoError(t, cfg.Error)
		assert.Equal(t, 340382, cfg.InfiniteTracing.SpanEvents.QueueSize)
	})
}
