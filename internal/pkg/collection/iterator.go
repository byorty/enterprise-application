package collection

type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Reset()
}

func NewListIterator[T any](iterable List[T]) Iterator[T] {
	return &listIterator[T]{
		iterable: iterable,
	}
}

type listIterator[T any] struct {
	iterable List[T]
	index    int
}

func (i *listIterator[T]) HasNext() bool {
	return i.index < i.iterable.Len()
}

func (i *listIterator[T]) Next() T {
	item := i.iterable.Get(i.index)
	i.index++
	return item
}

func (i *listIterator[T]) Reset() {
	i.index = 0
}

func NewMapIterator[K comparable, V any](iterable Map[K, V]) Iterator[V] {
	keys := make([]K, 0)
	for k, _ := range iterable.Entries() {
		keys = append(keys, k)
	}

	return &mapIterator[K, V]{
		keys:     keys,
		iterable: iterable,
	}
}

type mapIterator[K comparable, V any] struct {
	index    int
	keys     []K
	keysLen  int
	iterable Map[K, V]
}

func (i *mapIterator[K, V]) HasNext() bool {
	return i.index < i.keysLen
}

func (i *mapIterator[K, V]) Next() V {
	value, _ := i.iterable.Get(i.keys[i.index])
	i.index++
	return value
}

func (i *mapIterator[K, V]) Reset() {
	i.index = 0
}
