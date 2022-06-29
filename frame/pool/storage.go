package pool

type DataStorage interface {
	Set(key string, v *Value) error
	Get(key string) *Value
}

var _ DataStorage = &dataTrace{}

func NewDataStorage() DataStorage {
	return &dataTrace{pool: make(map[string]*Value, 0)}
}

type dataTrace struct {
	pool map[string]*Value
}

type Value struct {
	V    interface{}
	Step int
	URL  string
}

func (data *dataTrace) Set(key string, v *Value) error {
	data.pool[key] = v
	return nil
}

func (data dataTrace) Get(key string) *Value {
	v, ok := data.pool[key]
	if ok {
		return v
	}
	return nil
}
