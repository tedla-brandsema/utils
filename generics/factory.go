package generics 

type Factory struct {
	registry map[string]func() any 
}

func NewFactory() *Factory {
	return &Factory{
        registry: map[string]func() any{},
    }
}

func (f *Factory) Register(name string, ctor func() any) {
	f.registry[name] = ctor
}

func (f *Factory) Create(name string) (any, bool) {
	ctor, ok := f.registry[name]
	if !ok {
		return nil, false
	}
	return ctor(), true
}