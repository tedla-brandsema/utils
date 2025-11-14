package generics

// runtime holder interface (non-generic)
type Wrapper interface {
	Any() any
	TypeName() string
}

// HeteroSlice stores Wrapper values so different wrappers can coexist.
type HeteroSlice struct {
	items []Wrapper
}

func NewHeteroSlice() *HeteroSlice {
	return &HeteroSlice{
		items: make([]Wrapper, 0),
	}
}

// AddValue accepts a Wrapper (already validated by NewValue).
func (h *HeteroSlice) AddValue(w Wrapper) {
	h.items = append(h.items, w)
}

func (h *HeteroSlice) Items() []Wrapper { 
	return h.items 
}
