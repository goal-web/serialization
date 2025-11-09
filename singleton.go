package serialization

import (
	"sync"

	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
)

var singleton contracts.Serialization
var classSerializer contracts.ClassSerializer
var defaultSerializer contracts.Serializer
var once sync.Once

func Default() contracts.Serialization {
	once.Do(func() {
		singleton = application.Get("serialization").(contracts.Serialization)
		classSerializer = application.Get("class.serializer").(contracts.ClassSerializer)
		defaultSerializer = application.Get("serializer").(contracts.Serializer)
	})

	return singleton
}

func ClassSerializer() contracts.ClassSerializer {
	Default()
	return classSerializer
}

func Serializer() contracts.Serializer {
	Default()
	return defaultSerializer
}
