// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"ai-seller/internal/entity"
)

type UseCases interface {
	Auth
	Integration
	Product
}

type (

	// Auth -.
	Auth interface {
		CreateUser(context.Context, entity.User) error
		GetUser(context.Context, string) (entity.User, error)
		UpdateUser(context.Context, entity.User) error
		DeleteUser(context.Context, string) error

		CreateRole(context.Context, entity.Role) error
		GetRole(context.Context, string) (entity.Role, error)
		UpdateRole(context.Context, entity.Role) error
		DeleteRole(context.Context, string) error

		CreateClientType(context.Context, entity.ClientType) error
		GetClientType(context.Context, string) (entity.ClientType, error)
		UpdateClientType(context.Context, entity.ClientType) error
		DeleteClientType(context.Context, string) error
	}

	// Integration -.
	Integration interface {
		CreateIntegration(context.Context, entity.Integration) error
		GetIntegration(context.Context, string) (entity.Integration, error)
		UpdateIntegration(context.Context, entity.Integration) error
		DeleteIntegration(context.Context, string) error
	}

	// Product -.
	Product interface {
		CreateProduct(context.Context, entity.Product) error
		GetProduct(context.Context, string) (entity.Product, error)
		UpdateProduct(context.Context, entity.Product) error
		DeleteProduct(context.Context, string) error

		CreateCategory(context.Context, entity.Category) error
		GetCategory(context.Context, string) (entity.Category, error)
		UpdateCategory(context.Context, entity.Category) error
		DeleteCategory(context.Context, string) error

		CreateAttribute(context.Context, entity.Attribute) error
		GetAttribute(context.Context, string) (entity.Attribute, error)
		UpdateAttribute(context.Context, entity.Attribute) error
		DeleteAttribute(context.Context, string) error

		CreateOrder(context.Context, entity.Order) error
		GetOrder(context.Context, string) (entity.Order, error)
		UpdateOrder(context.Context, entity.Order) error
		DeleteOrder(context.Context, string) error

		CreateOrderProducts(context.Context, entity.OrderProducts) error
		GetOrderProducts(context.Context, string) (entity.OrderProducts, error)
		UpdateOrderProducts(context.Context, entity.OrderProducts) error
		DeleteOrderProducts(context.Context, string) error
	}
)
