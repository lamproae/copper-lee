package trace

//Tracer is the interface that decribes an object capable of tracing events throughout code.
import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...interface{})
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func Off() Tracer {
	return &nilTracer{}
}
