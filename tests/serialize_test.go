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
