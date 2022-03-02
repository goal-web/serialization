package serialization

import (
	"github.com/goal-web/contracts"
)

type ServiceProvider struct {
	app contracts.Application
}

func (this *ServiceProvider) Register(application contracts.Application) {
	this.app = application
	application.Singleton("serialization", func() contracts.Serialization {
		return New()
	})
	application.Singleton("serializer", func(config contracts.Config, serialization contracts.Serialization) contracts.Serializer {
		return serialization.Method(config.Get("serialization").(Config).Default)
	})
	application.Singleton("class.serializer", func(config contracts.Config) contracts.ClassSerializer {
		var serializer = NewClassSerializer(application)
		for _, class := range config.Get("serialization").(Config).Class {
			serializer.Register(class)
		}
		return serializer
	})
}

func (this *ServiceProvider) Start() error {
	return nil
}

func (this *ServiceProvider) Stop() {
}
