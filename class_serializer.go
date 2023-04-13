package serialization

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/serialization/serializers"
	"github.com/goal-web/supports/class"
	"reflect"
	"sync"
)

type Class struct {
	Class   string `json:"c"`
	Payload string `json:"p"`
}

func NewClassSerializer(container contracts.Container) contracts.ClassSerializer {
	return &Serializer{
		container:  container,
		classes:    sync.Map{},
		serializer: serializers.Json{},
	}
}

type Serializer struct {
	container  contracts.Container
	classes    sync.Map
	serializer contracts.Serializer
}

func (serializer *Serializer) Register(class contracts.Class[any]) {
	serializer.classes.Store(class.ClassName(), class)
}

func (serializer *Serializer) Serialize(instance any) string {
	return serializer.serializer.Serialize(Class{
		Class:   class.Any(instance).ClassName(),
		Payload: serializer.serializer.Serialize(instance),
	})
}

func (serializer *Serializer) Parse(serialized string) (any, error) {
	var c Class
	if err := serializer.serializer.Unserialize(serialized, &c); err != nil {
		return nil, err
	}

	classItem, exists := serializer.classes.Load(c.Class)
	if !exists {
		return nil, errors.New("unregistered class")
	}

	targetClass := classItem.(contracts.Class[any])

	instance := reflect.New(targetClass.GetType()).Interface()

	if err := serializer.serializer.Unserialize(c.Payload, instance); err != nil {
		return nil, err
	}

	if component, isComponent := instance.(contracts.Component); isComponent {
		component.Construct(serializer.container)
		return component, nil
	}

	return instance, nil
}
