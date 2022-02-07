package serialization

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/serialization/serializers"
)

func New() contracts.Serialization {
	return &Serialization{serializers: map[string]contracts.Serializer{
		"json": serializers.Json{},
		"gob":  serializers.Gob{},
		"xml":  serializers.Xml{},
	}}
}

type Serialization struct {
	serializers map[string]contracts.Serializer
}

func (s *Serialization) Method(name string) contracts.Serializer {
	return s.serializers[name]
}

func (s *Serialization) Extend(name string, serializer contracts.Serializer) {
	s.serializers[name] = serializer
}
