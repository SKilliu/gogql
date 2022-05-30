package graph

import "github.com/SKilliu/gogql/storage"

// This file will not be regenerated automatically.

//go:generate go run github.com/99designs/gqlgen generate

// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB storage.QInterface
}
