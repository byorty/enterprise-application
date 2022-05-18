package collection

import "sort"

type SliceOption func(settings *SliceSettings)

func WithSliceSize(size int) SliceOption {
	return func(settings *SliceSettings) {
		settings.size = size
	}
}

type SliceSettings struct {
	size int
}

type Slice[T any] interface {
	Add(items ...T)
	Set(i int, item T)
	Get(i int) T
	Len() int
	Entries() []T
	Iterator() Iterator[T]
	Remove(i int)
	RemoveAll()
	Sort(f SortFunc[T])
}

type genericSlice[T any] []T

func NewSlice[T any](opts ...SliceOption) Slice[T] {
	settings := new(SliceSettings)
	for _, opt := range opts {
		opt(settings)
	}

	l := make(genericSlice[T], settings.size)
	return &l
}

func ImportList[T any](slice ...T) Slice[T] {
	l := make(genericSlice[T], len(slice))
	copy(l, slice)
	return &l
}

func (l *genericSlice[T]) Add(items ...T) {
	(*l) = append((*l), items...)
}

func (l *genericSlice[T]) Set(i int, item T) {
	(*l)[i] = item
}

func (l *genericSlice[T]) Get(i int) T {
	return (*l)[i]
}

func (l *genericSlice[T]) Len() int {
	return len(*l)
}

func (l *genericSlice[T]) Entries() []T {
	return (*l)
}

func (l *genericSlice[T]) Iterator() Iterator[T] {
	return NewSliceIterator[T](l)
}

func (l *genericSlice[T]) Remove(i int) {
	(*l) = append((*l)[:i], (*l)[i+1:]...)
}

func (l *genericSlice[T]) RemoveAll() {
	(*l) = make(genericSlice[T], 0)
}

func (l *genericSlice[T]) Sort(f SortFunc[T]) {
	sort.Sort(&sorter[T]{
		list: l,
		f:    f,
	})
}

type SortFunc[T any] func(i, j T) bool

type sorter[T any] struct {
	list Slice[T]
	f    SortFunc[T]
}

func (s *sorter[T]) Len() int {
	return s.list.Len()
}

func (s *sorter[T]) Less(i, j int) bool {
	return s.f(s.list.Get(i), s.list.Get(j))
}

func (s *sorter[T]) Swap(i, j int) {
	a := s.list.Get(i)
	b := s.list.Get(j)
	s.list.Set(i, b)
	s.list.Set(j, a)
}
