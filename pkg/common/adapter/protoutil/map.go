package protoutil

import (
	"fmt"
	"github.com/byorty/enterprise-application/pkg/common/collection"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewMap[T any]() *Map[T] {
	return &Map[T]{
		m: collection.NewMap[string, T](),
	}
}

type Map[T any] struct {
	m collection.Map[string, T]
}

func (m *Map[T]) Set(key string, value T) {
	m.m.Set(key, value)
}

func (m Map[T]) Get(message protoreflect.Message, field protoreflect.FieldDescriptor) (T, error) {
	messageName := message.Descriptor().Name()
	fieldName := field.Name()
	value, ok := m.m.Get(fmt.Sprintf("%s.%s", messageName, fieldName))
	if ok {
		return value, nil
	}

	value, ok = m.m.Get(string(fieldName))
	if ok {
		return value, nil
	}

	return value, ErrKeyNotFound
}
