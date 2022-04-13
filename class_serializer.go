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

func (this *Serializer) Register(class contracts.Class) {
	this.classes.Store(class.ClassName(), class)
}

func (this *Serializer) Serialize(instance interface{}) string {
	return this.serializer.Serialize(Class{
		Class:   class.Make(instance).ClassName(),
		Payload: this.serializer.Serialize(instance),
	})
}

func (this *Serializer) Parse(serialized string) (interface{}, error) {
	var c Class
	if err := this.serializer.Unserialize(serialized, &c); err != nil {
		return nil, err
	}

	classItem, exists := this.classes.Load(c.Class)
	if !exists {
		return nil, errors.New("unregistered class")
	}

	targetClass := classItem.(contracts.Class)

	instance := reflect.New(targetClass.GetType()).Interface()

	if err := this.serializer.Unserialize(c.Payload, instance); err != nil {
		return nil, err
	}

	if component, isComponent := instance.(contracts.Component); isComponent {
		component.Construct(this.container)
		return component, nil
	}

	return instance, nil
}
