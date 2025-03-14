package persistent

import (
	"ai-seller/internal/entity"
	"ai-seller/pkg/postgres"
	"context"
	"fmt"
)

// AuthRepo -.
type IntegrationRepo struct {
	*postgres.Postgres
}

// NewAuthRepo -.
func NewIntegrationRepo(pg *postgres.Postgres) *IntegrationRepo {
	return &IntegrationRepo{pg}
}

// -------------- Integration --------------

// CreateIntegration -.
func (r *IntegrationRepo) CreateIntegration(ctx context.Context, i entity.Integration) error {
	sql, args, err := r.Builder.
		Insert("integration").
		Columns("name, created_at, updated_at").
		Values(i.Name, i.CreatedAt, i.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("IntegrationRepo - CreateIntegration - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&i.ID)
	if err != nil {
		return fmt.Errorf("IntegrationRepo - CreateIntegration - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetIntegration -.
func (r *IntegrationRepo) GetIntegration(ctx context.Context, id string) (entity.Integration, error) {
	var i entity.Integration

	sql, args, err := r.Builder.
		Select("id, name, created_at, updated_at").
		From("integration").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return i, fmt.Errorf("IntegrationRepo - GetIntegration - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&i.ID, &i.Name, &i.CreatedAt, &i.UpdatedAt)
	if err != nil {
		return i, fmt.Errorf("IntegrationRepo - GetIntegration - r.Pool.QueryRow: %w", err)
	}

	return i, nil
}

// UpdateIntegration -.
func (r *IntegrationRepo) UpdateIntegration(ctx context.Context, i entity.Integration) error {
	sql, args, err := r.Builder.
		Update("integration").
		Set("name", i.Name).
		Set("created_at", i.CreatedAt).
		Set("updated_at", i.UpdatedAt).
		Where("id = ?", i.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("IntegrationRepo - UpdateIntegration - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("IntegrationRepo - UpdateIntegration - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteIntegration -.
func (r *IntegrationRepo) DeleteIntegration(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("integration").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("IntegrationRepo - DeleteIntegration - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("IntegrationRepo - DeleteIntegration - r.Pool.Exec: %w", err)
	}

	return nil
}
