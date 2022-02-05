package tests

import (
	"fmt"
	"github.com/goal-web/serialization"
	"github.com/goal-web/serialization/serializers"
	"github.com/goal-web/supports/class"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Detail struct {
	Title     string `json:"title"`
	User      User   `json:"user"`
	Followers []User `json:"followers"`
}

var (
	UserClass = class.Make(new(User))
)

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

func TestComplexJsonSerialize(t *testing.T) {

	json := serializers.Json{}

	jsonStr := json.Serialize(Detail{
		Title: "goal",
		User: User{
			Id:   "10086",
			Name: "goal",
		},
		Followers: []User{
			{
				Id:   "1",
				Name: "goal",
			},
			{
				Id:   "2",
				Name: "goal",
			},
		},
	})
	fmt.Println(jsonStr)
	assert.True(t, "{\"title\":\"goal\",\"user\":{\"id\":\"10086\",\"name\":\"goal\"},\"followers\":[{\"id\":\"1\",\"name\":\"goal\"},{\"id\":\"2\",\"name\":\"goal\"}]}" == jsonStr)

	var detail Detail
	assert.Nil(t, json.Unserialize(jsonStr, &detail))
	fmt.Println(detail)
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

func TestClassSerializer(t *testing.T) {
	serializer := serialization.NewClassSerializer()

	serializer.Register(UserClass)
	serializer.Register(class.Make(new(Detail)))

	serialized := serializer.Serialize(Detail{
		Title: "goal",
		User: User{
			Id:   "10086",
			Name: "goal",
		},
		Followers: []User{
			{
				Id:   "1",
				Name: "goal",
			},
			{
				Id:   "2",
				Name: "goal",
			},
		},
	})

	fmt.Println("serialized", serialized)
	detail, err := serializer.Parse(serialized)
	assert.Nil(t, err)
	assert.True(t, detail.(*Detail).Title == "goal")
	fmt.Println(detail)
}

/**
goos: darwin
goarch: amd64
pkg: github.com/goal-web/serialization/tests
cpu: Intel(R) Core(TM) i7-7660U CPU @ 2.50GHz
BenchmarkClassSerialize
BenchmarkClassSerialize-4   	   95265	     11996 ns/op
*/
func BenchmarkClassSerialize(b *testing.B) {
	serializer := serialization.NewClassSerializer()

	serializer.Register(class.Make(new(Detail)))
	detail := Detail{
		Title: "goal",
		User: User{
			Id:   "10086",
			Name: "goal",
		},
		Followers: []User{
			{
				Id:   "1",
				Name: "goal",
			},
			{
				Id:   "2",
				Name: "goal",
			},
		},
	}

	for i := 0; i < b.N; i++ {
		serialized := serializer.Serialize(detail)
		_, _ = serializer.Parse(serialized)
	}
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
