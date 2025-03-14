package persistent

import (
	"ai-seller/internal/entity"
	"ai-seller/pkg/postgres"
	"context"
	"fmt"
)

// ProductRepo -.
type ProductRepo struct {
	*postgres.Postgres
}

// NewProductRepo -.
func NewProductRepo(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{pg}
}

// ---------------- Product ----------------

// CreateProduct -.
func (r *ProductRepo) CreateProduct(ctx context.Context, p entity.Product) error {
	sql, args, err := r.Builder.
		Insert("product").
		Columns("name, category_id, short_info, description, cost, count, discount_cost, discount, created_at, updated_at").
		Values(p.Name, p.CategoryID, p.ShortInfo, p.Description, p.Cost, p.Count, p.DiscountCost, p.Discount, p.CreatedAt, p.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateProduct - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&p.ID)
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateProduct - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetProductByID -.
func (r *ProductRepo) GetProduct(ctx context.Context, id string) (entity.Product, error) {
	var p entity.Product

	sql, args, err := r.Builder.
		Select("id, name, category_id, short_info, description, cost, count, discount_cost, discount, created_at, updated_at").
		From("product").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return p, fmt.Errorf("ProductRepo - GetProductByID - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&p.ID, &p.Name, &p.CategoryID, &p.ShortInfo, &p.Description, &p.Cost, &p.Count, &p.DiscountCost, &p.Discount, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return p, fmt.Errorf("ProductRepo - GetProductByID - r.Pool.QueryRow: %w", err)
	}

	return p, nil
}

// UpdateProduct -.
func (r *ProductRepo) UpdateProduct(ctx context.Context, p entity.Product) error {
	sql, args, err := r.Builder.
		Update("product").
		Set("name", p.Name).
		Set("category_id", p.CategoryID).
		Set("short_info", p.ShortInfo).
		Set("description", p.Description).
		Set("cost", p.Cost).
		Set("count", p.Count).
		Set("discount_cost", p.DiscountCost).
		Set("discount", p.Discount).
		Set("created_at", p.CreatedAt).
		Set("updated_at", p.UpdatedAt).
		Where("id = ?", p.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateProduct - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateProduct - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteProduct -.
func (r *ProductRepo) DeleteProduct(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("product").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteProduct - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteProduct - r.Pool.Exec: %w", err)
	}

	return nil
}

// ---------------- Category ----------------

// CreateCategory -.
func (r *ProductRepo) CreateCategory(ctx context.Context, c entity.Category) error {
	sql, args, err := r.Builder.
		Insert("category").
		Columns("name, created_at, updated_at").
		Values(c.Name, c.CreatedAt, c.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateCategory - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&c.ID)
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateCategory - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetCategoryByID -.
func (r *ProductRepo) GetCategory(ctx context.Context, id string) (entity.Category, error) {
	var c entity.Category

	sql, args, err := r.Builder.
		Select("id, name, created_at, updated_at").
		From("category").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return c, fmt.Errorf("ProductRepo - GetCategoryByID - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&c.ID, &c.Name, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return c, fmt.Errorf("ProductRepo - GetCategoryByID - r.Pool.QueryRow: %w", err)
	}

	return c, nil
}

// UpdateCategory -.
func (r *ProductRepo) UpdateCategory(ctx context.Context, c entity.Category) error {
	sql, args, err := r.Builder.
		Update("category").
		Set("name", c.Name).
		Set("created_at", c.CreatedAt).
		Set("updated_at", c.UpdatedAt).
		Where("id = ?", c.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateCategory - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateCategory - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteCategory -.
func (r *ProductRepo) DeleteCategory(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("category").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteCategory - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteCategory - r.Pool.Exec: %w", err)
	}

	return nil
}

// ---------------- Attribute ----------------

// CreateAttribute -.
func (r *ProductRepo) CreateAttribute(ctx context.Context, a entity.Attribute) error {
	sql, args, err := r.Builder.
		Insert("attribute").
		Columns("name, category_id, created_at, updated_at").
		Values(a.Name, a.CategoryID, a.CreatedAt, a.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateAttribute - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&a.ID)
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateAttribute - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetAttributeByID -.
func (r *ProductRepo) GetAttribute(ctx context.Context, id string) (entity.Attribute, error) {
	var a entity.Attribute

	sql, args, err := r.Builder.
		Select("id, name, category_id, created_at, updated_at").
		From("attribute").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return a, fmt.Errorf("ProductRepo - GetAttributeByID - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&a.ID, &a.Name, &a.CategoryID, &a.CreatedAt, &a.UpdatedAt)
	if err != nil {
		return a, fmt.Errorf("ProductRepo - GetAttributeByID - r.Pool.QueryRow: %w", err)
	}

	return a, nil
}

// UpdateAttribute -.
func (r *ProductRepo) UpdateAttribute(ctx context.Context, a entity.Attribute) error {
	sql, args, err := r.Builder.
		Update("attribute").
		Set("name", a.Name).
		Set("category_id", a.CategoryID).
		Set("created_at", a.CreatedAt).
		Set("updated_at", a.UpdatedAt).
		Where("id = ?", a.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateAttribute - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateAttribute - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteAttribute -.
func (r *ProductRepo) DeleteAttribute(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("attribute").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteAttribute - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteAttribute - r.Pool.Exec: %w", err)
	}

	return nil
}

// ---------------- Order ----------------

// CreateOrder -.
func (r *ProductRepo) CreateOrder(ctx context.Context, o entity.Order) error {
	sql, args, err := r.Builder.
		Insert("order").
		Columns("user_id, integration_id, status, status_changed_time, total_cost, created_at, updated_at").
		Values(o.UserID, o.IntegrationID, o.Status, o.StatusChangedTime, o.TotalCost, o.CreatedAt, o.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateOrder - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&o.ID)
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateOrder - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetOrderByID -.
func (r *ProductRepo) GetOrder(ctx context.Context, id string) (entity.Order, error) {
	var o entity.Order

	sql, args, err := r.Builder.
		Select("id, user_id, integration_id, status, status_changed_time, total_cost, created_at, updated_at").
		From("order").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return o, fmt.Errorf("ProductRepo - GetOrderByID - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&o.ID, &o.UserID, &o.IntegrationID, &o.Status, &o.StatusChangedTime, &o.TotalCost, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		return o, fmt.Errorf("ProductRepo - GetOrderByID - r.Pool.QueryRow: %w", err)
	}

	return o, nil
}

// UpdateOrder -.
func (r *ProductRepo) UpdateOrder(ctx context.Context, o entity.Order) error {
	sql, args, err := r.Builder.
		Update("order").
		Set("user_id", o.UserID).
		Set("integration_id", o.IntegrationID).
		Set("status", o.Status).
		Set("status_changed_time", o.StatusChangedTime).
		Set("total_cost", o.TotalCost).
		Set("created_at", o.CreatedAt).
		Set("updated_at", o.UpdatedAt).
		Where("id = ?", o.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateOrder - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateOrder - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteOrder -.
func (r *ProductRepo) DeleteOrder(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("order").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteOrder - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteOrder - r.Pool.Exec: %w", err)
	}

	return nil
}

// ---------------- OrderProducts ----------------

// CreateOrderProduct -.
func (r *ProductRepo) CreateOrderProducts(ctx context.Context, op entity.OrderProducts) error {
	sql, args, err := r.Builder.
		Insert("order_products").
		Columns("order_id, product_id, count, cost, created_at, updated_at").
		Values(op.OrderID, op.ProductID, op.Count, op.Cost, op.CreatedAt, op.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateOrderProduct - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&op.ID)
	if err != nil {
		return fmt.Errorf("ProductRepo - CreateOrderProduct - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetOrderProductByID -.
func (r *ProductRepo) GetOrderProducts(ctx context.Context, id string) (entity.OrderProducts, error) {
	var op entity.OrderProducts

	sql, args, err := r.Builder.
		Select("id, order_id, product_id, count, cost, created_at, updated_at").
		From("order_products").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return op, fmt.Errorf("ProductRepo - GetOrderProductByID - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&op.ID, &op.OrderID, &op.ProductID, &op.Count, &op.Cost, &op.CreatedAt, &op.UpdatedAt)
	if err != nil {
		return op, fmt.Errorf("ProductRepo - GetOrderProductByID - r.Pool.QueryRow: %w", err)
	}

	return op, nil
}

// UpdateOrderProduct -.
func (r *ProductRepo) UpdateOrderProducts(ctx context.Context, op entity.OrderProducts) error {
	sql, args, err := r.Builder.
		Update("order_products").
		Set("order_id", op.OrderID).
		Set("product_id", op.ProductID).
		Set("count", op.Count).
		Set("cost", op.Cost).
		Set("created_at", op.CreatedAt).
		Set("updated_at", op.UpdatedAt).
		Where("id = ?", op.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateOrderProduct - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - UpdateOrderProduct - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteOrderProduct -.
func (r *ProductRepo) DeleteOrderProducts(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("order_products").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteOrderProduct - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("ProductRepo - DeleteOrderProduct - r.Pool.Exec: %w", err)
	}

	return nil
}
