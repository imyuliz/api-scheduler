package trace

type Trace interface {
	// context.Context
	WithTraceValue(key, value string)
	Key() string
	Value() string
}

func NewTrace(key, value string) Trace {
	return &trace{
		k: key,
		v: value,
	}
}

type trace struct {
	k string
	v string
}

func (t *trace) WithTraceValue(key, value string) {
	t.k = key
	t.v = value
}
func (t *trace) Key() string {
	return t.k
}

func (t *trace) Value() string {
	return t.v
}
