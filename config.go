package serialization

import "github.com/goal-web/contracts"

type Config struct {
	Default string
	Class   []contracts.Class[any]
}
