package serialization

import (
	"errors"
	"github.com/goal-web/contracts"
	"github.com/goal-web/serialization/serializers"
	"github.com/goal-web/supports/class"
	"reflect"
)

type Class struct {
	Class   string `json:"c"`
	Payload string `json:"p"`
}

func NewClassSerializer() contracts.ClassSerializer {
	return &Serializer{
		classes:    map[string]contracts.Class{},
		serializer: serializers.Json{},
	}
}

type Serializer struct {
	classes    map[string]contracts.Class
	serializer contracts.Serializer
}

func (this *Serializer) Register(class contracts.Class) {
	this.classes[class.ClassName()] = class
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

	if this.classes[c.Class] == nil {
		return nil, errors.New("unregistered class")
	}

	instance := reflect.New(this.classes[c.Class].GetType()).Interface()

	if err := this.serializer.Unserialize(c.Payload, instance); err != nil {
		return nil, err
	}

	return instance, nil
}
