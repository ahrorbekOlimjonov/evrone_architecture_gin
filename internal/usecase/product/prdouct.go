package product

import (
	"context"
	"fmt"

	"ai-seller/internal/entity"
	"ai-seller/internal/repo"
)

// UseCase -.
type UseCase struct {
	auth        repo.AuthRepo
	product     repo.ProductRepo
	integration repo.IntegrationRepo
}

// New -.
func New(a repo.AuthRepo, p repo.ProductRepo, i repo.IntegrationRepo) *UseCase {
	return &UseCase{
		auth:        a,
		product:     p,
		integration: i,
	}
}

// -------------- Auth --------------

// CreateUser -.
func (uc *UseCase) CreateUser(ctx context.Context, u entity.User) error {
	err := uc.auth.CreateUser(ctx, u)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateUser - s.auth.CreateUser: %w", err)
	}
	return nil
}

// GetUser -.
func (uc *UseCase) GetUser(ctx context.Context, id string) (entity.User, error) {
	user, err := uc.auth.GetUser(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("ProductUseCase - GetUser - s.auth.GetUser: %w", err)
	}

	return user, nil
}

// UpdateUser -.
func (uc *UseCase) UpdateUser(ctx context.Context, u entity.User) error {
	err := uc.auth.UpdateUser(ctx, u)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateUser - s.auth.UpdateUser: %w", err)
	}

	return nil
}

// DeleteUser -.
func (uc *UseCase) DeleteUser(ctx context.Context, id string) error {
	err := uc.auth.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteUser - s.auth.DeleteUser: %w", err)
	}

	return nil
}

// -------------- Role --------------

// CreateRole -.
func (uc *UseCase) CreateRole(ctx context.Context, r entity.Role) error {
	err := uc.auth.CreateRole(ctx, r)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateRole - s.auth.CreateRole: %w", err)
	}

	return nil
}

// GetRole -.
func (uc *UseCase) GetRole(ctx context.Context, id string) (entity.Role, error) {
	role, err := uc.auth.GetRole(ctx, id)
	if err != nil {
		return entity.Role{}, fmt.Errorf("ProductUseCase - GetRole - s.auth.GetRole: %w", err)
	}

	return role, nil
}

// UpdateRole -.
func (uc *UseCase) UpdateRole(ctx context.Context, r entity.Role) error {
	err := uc.auth.UpdateRole(ctx, r)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateRole - s.auth.UpdateRole: %w", err)
	}

	return nil
}

// DeleteRole -.
func (uc *UseCase) DeleteRole(ctx context.Context, id string) error {
	err := uc.auth.DeleteRole(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteRole - s.auth.DeleteRole: %w", err)
	}

	return nil
}

// -------------- ClientType --------------

// CreateClientType -.
func (uc *UseCase) CreateClientType(ctx context.Context, ct entity.ClientType) error {
	err := uc.auth.CreateClientType(ctx, ct)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateClientType - s.auth.CreateClientType: %w", err)
	}

	return nil
}

// GetClientType -.
func (uc *UseCase) GetClientType(ctx context.Context, id string) (entity.ClientType, error) {
	clientType, err := uc.auth.GetClientType(ctx, id)
	if err != nil {
		return entity.ClientType{}, fmt.Errorf("ProductUseCase - GetClientType - s.auth.GetClientType: %w", err)
	}

	return clientType, nil
}

// UpdateClientType -.
func (uc *UseCase) UpdateClientType(ctx context.Context, ct entity.ClientType) error {
	err := uc.auth.UpdateClientType(ctx, ct)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateClientType - s.auth.UpdateClientType: %w", err)
	}

	return nil
}

// DeleteClientType -.
func (uc *UseCase) DeleteClientType(ctx context.Context, id string) error {
	err := uc.auth.DeleteClientType(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteClientType - s.auth.DeleteClientType: %w", err)
	}

	return nil
}

// -------------- Integration --------------
// CreateIntegration -.
func (uc *UseCase) CreateIntegration(ctx context.Context, i entity.Integration) error {
	err := uc.integration.CreateIntegration(ctx, i)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateIntegration - s.integration.CreateIntegration: %w", err)
	}

	return nil
}

// GetIntegration -.
func (uc *UseCase) GetIntegration(ctx context.Context, id string) (entity.Integration, error) {
	integration, err := uc.integration.GetIntegration(ctx, id)
	if err != nil {
		return entity.Integration{}, fmt.Errorf("ProductUseCase - GetIntegration - s.integration.GetIntegration: %w", err)
	}

	return integration, nil
}

// UpdateIntegration -.
func (uc *UseCase) UpdateIntegration(ctx context.Context, i entity.Integration) error {
	err := uc.integration.UpdateIntegration(ctx, i)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateIntegration - s.integration.UpdateIntegration: %w", err)
	}

	return nil
}

// DeleteIntegration -.
func (uc *UseCase) DeleteIntegration(ctx context.Context, id string) error {
	err := uc.integration.DeleteIntegration(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteIntegration - s.integration.DeleteIntegration: %w", err)
	}

	return nil
}

// -------------- Product --------------

// CreateProduct -.
func (uc *UseCase) CreateProduct(ctx context.Context, p entity.Product) error {
	err := uc.product.CreateProduct(ctx, p)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateProduct - s.product.CreateProduct: %w", err)
	}

	return nil
}

// GetProduct -.
func (uc *UseCase) GetProduct(ctx context.Context, id string) (entity.Product, error) {
	product, err := uc.product.GetProduct(ctx, id)
	if err != nil {
		return entity.Product{}, fmt.Errorf("ProductUseCase - GetProduct - s.product.GetProduct: %w", err)
	}

	return product, nil
}

// UpdateProduct -.
func (uc *UseCase) UpdateProduct(ctx context.Context, p entity.Product) error {
	err := uc.product.UpdateProduct(ctx, p)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateProduct - s.product.UpdateProduct: %w", err)
	}

	return nil
}

// DeleteProduct -.
func (uc *UseCase) DeleteProduct(ctx context.Context, id string) error {
	err := uc.product.DeleteProduct(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteProduct - s.product.DeleteProduct: %w", err)
	}

	return nil
}

// CreateCategory -.
func (uc *UseCase) CreateCategory(ctx context.Context, c entity.Category) error {
	err := uc.product.CreateCategory(ctx, c)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateCategory - s.product.CreateCategory: %w", err)
	}

	return nil
}

// GetCategory -.
func (uc *UseCase) GetCategory(ctx context.Context, id string) (entity.Category, error) {
	category, err := uc.product.GetCategory(ctx, id)
	if err != nil {
		return entity.Category{}, fmt.Errorf("ProductUseCase - GetCategory - s.product.GetCategory: %w", err)
	}

	return category, nil
}

// UpdateCategory -.
func (uc *UseCase) UpdateCategory(ctx context.Context, c entity.Category) error {
	err := uc.product.UpdateCategory(ctx, c)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateCategory - s.product.UpdateCategory: %w", err)
	}

	return nil
}

// DeleteCategory -.
func (uc *UseCase) DeleteCategory(ctx context.Context, id string) error {
	err := uc.product.DeleteCategory(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteCategory - s.product.DeleteCategory: %w", err)
	}

	return nil
}

// CreateAttribute -.
func (uc *UseCase) CreateAttribute(ctx context.Context, a entity.Attribute) error {
	err := uc.product.CreateAttribute(ctx, a)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateAttribute - s.product.CreateAttribute: %w", err)
	}

	return nil
}

// GetAttribute -.
func (uc *UseCase) GetAttribute(ctx context.Context, id string) (entity.Attribute, error) {
	attribute, err := uc.product.GetAttribute(ctx, id)
	if err != nil {
		return entity.Attribute{}, fmt.Errorf("ProductUseCase - GetAttribute - s.product.GetAttribute: %w", err)
	}

	return attribute, nil
}

// UpdateAttribute -.
func (uc *UseCase) UpdateAttribute(ctx context.Context, a entity.Attribute) error {
	err := uc.product.UpdateAttribute(ctx, a)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateAttribute - s.product.UpdateAttribute: %w", err)
	}

	return nil
}

// DeleteAttribute -.
func (uc *UseCase) DeleteAttribute(ctx context.Context, id string) error {
	err := uc.product.DeleteAttribute(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteAttribute - s.product.DeleteAttribute: %w", err)
	}

	return nil
}

// CreateOrder -.
func (uc *UseCase) CreateOrder(ctx context.Context, o entity.Order) error {
	err := uc.product.CreateOrder(ctx, o)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateOrder - s.product.CreateOrder: %w", err)
	}

	return nil
}

// GetOrder -.
func (uc *UseCase) GetOrder(ctx context.Context, id string) (entity.Order, error) {
	order, err := uc.product.GetOrder(ctx, id)
	if err != nil {
		return entity.Order{}, fmt.Errorf("ProductUseCase - GetOrder - s.product.GetOrder: %w", err)
	}

	return order, nil
}

// UpdateOrder -.
func (uc *UseCase) UpdateOrder(ctx context.Context, o entity.Order) error {
	err := uc.product.UpdateOrder(ctx, o)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateOrder - s.product.UpdateOrder: %w", err)
	}

	return nil
}

// DeleteOrder -.
func (uc *UseCase) DeleteOrder(ctx context.Context, id string) error {
	err := uc.product.DeleteOrder(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteOrder - s.product.DeleteOrder: %w", err)
	}

	return nil
}

// CreateOrderProduct -.
func (uc *UseCase) CreateOrderProducts(ctx context.Context, op entity.OrderProducts) error {
	err := uc.product.CreateOrderProducts(ctx, op)
	if err != nil {
		return fmt.Errorf("ProductUseCase - CreateOrderProducts - s.product.CreateOrderProducts: %w", err)
	}

	return nil
}

// GetOrderProducts -.
func (uc *UseCase) GetOrderProducts(ctx context.Context, id string) (entity.OrderProducts, error) {
	orderProducts, err := uc.product.GetOrderProducts(ctx, id)
	if err != nil {
		return entity.OrderProducts{}, fmt.Errorf("ProductUseCase - GetOrderProducts - s.product.GetOrderProducts: %w", err)
	}

	return orderProducts, nil
}

// UpdateOrderProducts -.
func (uc *UseCase) UpdateOrderProducts(ctx context.Context, op entity.OrderProducts) error {
	err := uc.product.UpdateOrderProducts(ctx, op)
	if err != nil {
		return fmt.Errorf("ProductUseCase - UpdateOrderProducts - s.product.UpdateOrderProducts: %w", err)
	}

	return nil
}

// DeleteOrderProducts -.
func (uc *UseCase) DeleteOrderProducts(ctx context.Context, id string) error {
	err := uc.product.DeleteOrderProducts(ctx, id)
	if err != nil {
		return fmt.Errorf("ProductUseCase - DeleteOrderProducts - s.product.DeleteOrderProducts: %w", err)
	}

	return nil
}
