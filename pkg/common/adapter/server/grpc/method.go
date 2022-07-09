package grpc

import (
	pbv1 "github.com/byorty/enterprise-application/pkg/common/gen/api/proto/v1"
	"reflect"
	"runtime"
	"strings"
)

type MethodDescriptor struct {
	Method     interface{}
	Role       pbv1.Role
	Permission pbv1.Permission
}

func (m *MethodDescriptor) GetName() (string, error) {
	if m.Method == nil {
		return "", ErrMethodDescriptorNotFound
	}

	methodPointer := reflect.ValueOf(m.Method).Pointer()
	fullName := runtime.FuncForPC(methodPointer).Name()
	methodNameParts := strings.Split(fullName, ".")
	return methodNameParts[len(methodNameParts)-1], nil
}

func NewFxMethodDescriptorMap(
	serverDescriptor Descriptor,
) (MethodDescriptorMap, error) {
	m := make(MethodDescriptorMap)
	for _, methodDescriptor := range serverDescriptor.MethodDescriptors {
		methodName, err := methodDescriptor.GetName()
		if err != nil {
			return nil, err
		}

		m[methodName] = methodDescriptor
	}

	return m, nil
}

type MethodDescriptorMap map[string]MethodDescriptor
