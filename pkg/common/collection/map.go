package collection

type Map[K comparable, V any] interface {
	Set(key K, value V)
	Get(key K) (V, bool)
	Len() int
	Entries() map[K]V
	Iterator() Iterator[V]
	Remove(key K)
	RemoveAll()
}

func NewMap[K comparable, V any]() Map[K, V] {
	return make(genericMap[K, V])
}

func ImportMap[K comparable, V any](m map[K]V) Map[K, V] {
	return genericMap[K, V](m)
}

type genericMap[K comparable, V any] map[K]V

func (m genericMap[K, V]) Set(key K, value V) {
	m[key] = value
}

func (m genericMap[K, V]) Get(key K) (V, bool) {
	value, ok := m[key]
	return value, ok
}

func (m genericMap[K, V]) Len() int {
	return len(m)
}

func (m genericMap[K, V]) Entries() map[K]V {
	return m
}

func (m genericMap[K, V]) Iterator() Iterator[V] {
	return NewMapIterator[K, V](m)
}

func (m genericMap[K, V]) Remove(key K) {
	delete(m, key)
}

func (m genericMap[K, V]) RemoveAll() {
	for key, _ := range m {
		m.Remove(key)
	}
}
