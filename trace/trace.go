package trace

import (
	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/stats/view"
	"google.golang.org/api/option"
)

type Trace struct {
	exp *stackdriver.Exporter
}

func New(opts ...option.ClientOption) (*Trace, error) {
	exp, err := stackdriver.NewExporter(stackdriver.Options{
		MonitoringClientOptions: opts,
		TraceClientOptions:      opts,
	})
	if err != nil {
		return nil, err
	}
	return &Trace{
		exp: exp,
	}, nil
}

func (t *Trace) Register(views ...*view.View) error {
	if err := view.Register(views...); err != nil {
		return err
	}
	view.RegisterExporter(t.exp)
	return nil
}

func (t *Trace) Flush() {
	t.exp.Flush()
}

func (t *Trace) Exporter() *stackdriver.Exporter {
	return t.exp
}
