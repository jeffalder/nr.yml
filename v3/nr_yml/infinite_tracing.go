package nr_yml

import "github.com/newrelic/go-agent/v3/newrelic"

// Don't use this; it's only exported for the yaml parser
type TraceObserverYaml struct {
	Host *string `yaml:"host"`
	// why is this signed and 32-bit? TCP ports are limited to uint16.
	Port *uint16 `yaml:"port"`
}

func (traceObserverYaml TraceObserverYaml) update(cfg *newrelic.Config) {
	if traceObserverYaml.Host != nil {
		cfg.InfiniteTracing.TraceObserver.Host = *traceObserverYaml.Host
	}
	if traceObserverYaml.Port != nil {
		cfg.InfiniteTracing.TraceObserver.Port = int(*traceObserverYaml.Port)
	}
}

// Don't use this; it's only exported for the yaml parser
type EventsYaml struct {
	QueueSize *uint32 `yaml:"queue_size"`
}

func (spanEventsYaml EventsYaml) updateSpanEvents(cfg *newrelic.Config) {
	if spanEventsYaml.QueueSize != nil {
		cfg.InfiniteTracing.SpanEvents.QueueSize = int(*spanEventsYaml.QueueSize)
	}
}

// Don't use this; it's only exported for the yaml parser
type InfiniteTracingYaml struct {
	TraceObserver *TraceObserverYaml `yaml:"trace_observer"`
	SpanEvents    *EventsYaml        `yaml:"span_events"`
}

func (tracingYaml InfiniteTracingYaml) update(cfg *newrelic.Config) {
	if tracingYaml.TraceObserver != nil {
		tracingYaml.TraceObserver.update(cfg)
	}
	if tracingYaml.SpanEvents != nil {
		tracingYaml.SpanEvents.updateSpanEvents(cfg)
	}
}
