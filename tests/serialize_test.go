package tests

import (
	"fmt"
	"github.com/goal-web/serialization/serializers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func TestJsonSerialize(t *testing.T) {

	json := serializers.Json{}

	jsonStr := json.Serialize(User{
		Id:   "10086",
		Name: "goal",
	})
	fmt.Println(jsonStr)
	assert.True(t, "{\"id\":\"10086\",\"name\":\"goal\"}" == jsonStr)

	var user User
	assert.Nil(t, json.Unserialize(jsonStr, &user))
	fmt.Println(user)
}

func TestGobSerialize(t *testing.T) {

	gob := serializers.Gob{}

	jsonStr := gob.Serialize(User{
		Id:   "10086",
		Name: "goal",
	})
	fmt.Println(jsonStr)

	var user User
	assert.Nil(t, gob.Unserialize(jsonStr, &user))
	fmt.Println(user)
}

/**
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkJsonSerialize
BenchmarkJsonSerialize-4   	  894805	      1673 ns/op
*/
func BenchmarkJsonSerialize(b *testing.B) {
	serializer := serializers.Json{}
	for i := 0; i < b.N; i++ {

		jsonStr := serializer.Serialize(User{
			Id:   "10086",
			Name: "goal",
		})

		var user User
		_ = serializer.Unserialize(jsonStr, &user)
	}
}

/**
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkGobSerialize
BenchmarkGobSerialize-4   	   46285	     22629 ns/op
*/
func BenchmarkGobSerialize(b *testing.B) {
	serializer := serializers.Gob{}
	for i := 0; i < b.N; i++ {

		jsonStr := serializer.Serialize(User{
			Id:   "10086",
			Name: "goal",
		})

		var user User
		_ = serializer.Unserialize(jsonStr, &user)
	}
}
