package serialization

import "github.com/goal-web/contracts"

type ServiceProvider struct {
}

func (s ServiceProvider) Register(application contracts.Application) {
	application.Singleton("serialization", func() contracts.Serialization {
		return New()
	})
	application.Singleton("serializer", func(serialization contracts.Serialization) contracts.Serializer {
		return serialization.Method("json")
	})
	application.Singleton("class_serializer", func() contracts.ClassSerializer {
		return NewClassSerializer()
	})
}

func (s ServiceProvider) Start() error {
	return nil
}

func (s ServiceProvider) Stop() {
}
