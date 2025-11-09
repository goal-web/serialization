package serialization

import (
	"errors"
	"reflect"
	"sync"

	"github.com/goal-web/contracts"
	"github.com/goal-web/serialization/serializers"
	"github.com/goal-web/supports/class"
)

type Class struct {
	Class   string `json:"c"`
	Payload string `json:"p"`
}

func NewClassSerializer(container contracts.Container) contracts.ClassSerializer {
	return &serializer{
		container:  container,
		classes:    sync.Map{},
		serializer: serializers.Json{},
	}
}

type serializer struct {
	container  contracts.Container
	classes    sync.Map
	serializer contracts.Serializer
}

func (s *serializer) Register(class contracts.Class[any]) {
	s.classes.Store(class.ClassName(), class)
}

func (s *serializer) Serialize(instance any) string {
	return s.serializer.Serialize(Class{
		Class:   class.Any(instance).ClassName(),
		Payload: s.serializer.Serialize(instance),
	})
}

func (s *serializer) Parse(serialized string) (any, error) {
	var c Class
	if err := s.serializer.Unserialize(serialized, &c); err != nil {
		return nil, err
	}

	classItem, exists := s.classes.Load(c.Class)
	if !exists {
		return nil, errors.New("unregistered class")
	}

	targetClass := classItem.(contracts.Class[any])

	instance := reflect.New(targetClass.GetType()).Interface()

	if err := s.serializer.Unserialize(c.Payload, instance); err != nil {
		return nil, err
	}

	if component, isComponent := instance.(contracts.Component); isComponent {
		component.Construct(s.container)
		return component, nil
	}

	return instance, nil
}
