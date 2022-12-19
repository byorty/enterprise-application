package collection

type Prototyped[T any] interface {
	ToProto() T
}

type PrototypedSlice[T any] interface {
	Slice[Prototyped[T]]
	ToProto() []T
}

func NewPrototypedSlice[T any]() PrototypedSlice[T] {
	sl := make(genericSlice[Prototyped[T]], 0)
	return &genericPrototypedSlice[T]{
		genericSlice: &sl,
	}
}

type genericPrototypedSlice[T any] struct {
	*genericSlice[Prototyped[T]]
}

func (s *genericPrototypedSlice[T]) ToProto() []T {
	var prototypedSlice = make([]T, s.Len())
	for i, item := range s.Entries() {
		prototypedSlice[i] = item.ToProto()
	}

	return prototypedSlice
}
