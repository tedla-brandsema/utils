package generics

import (
	"fmt"
	"testing"
)

// Allowed types
type Primitive interface {
	int | float64 | string
}

type valueImpl[T Primitive] struct {
	v T
}

func (x valueImpl[T]) Any() any {
	return any(x.v)
}

func (x valueImpl[T]) TypeName() string {
	return fmt.Sprintf("%T", x.v)
}

// NewValue enforces allowed types at compile time.
// If you try NewValue(true) the compiler will reject it.
func NewValue[T Primitive](v T) Wrapper {
	return valueImpl[T]{v: v}
}

func AddTo[T Primitive](h *HeteroSlice, v T) {
	h.AddValue(NewValue(v))
}

func TestHetroSlcie(t *testing.T) {
	hs := NewHeteroSlice()

	// ergonomic, compile-time-checked calls:
	AddTo(hs, 42)
	AddTo(hs, 3.14)
	AddTo(hs, "hello")

	// direct wrapper construction also possible:
	hs.AddValue(NewValue(100))

	for _, w := range hs.Items() {
		fmt.Printf("%s -> %v\n", w.TypeName(), w.Any())
	}

	// AddTo(hs, true) // <- compile error: bool does not satisfy Primitive

}
