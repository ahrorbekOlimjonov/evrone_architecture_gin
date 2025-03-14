package persistent

import (
	"context"
	"fmt"

	"ai-seller/internal/entity"
	"ai-seller/pkg/postgres"
)

// AuthRepo -.
type AuthRepo struct {
	*postgres.Postgres
}

// NewAuthRepo -.
func NewAuthRepo(pg *postgres.Postgres) *AuthRepo {
	return &AuthRepo{pg}
}

// -------------- User --------------

// CreateUser -.
func (r *AuthRepo) CreateUser(ctx context.Context, u entity.User) error {
	sql, args, err := r.Builder.
		Insert("user").
		Columns("name, surname, username, password, birth_date, tg_user_name, phone, instagram, client_from, role_id").
		Values(u.Name, u.Surname, u.Username, u.Password, u.BirthDate, u.TgUserName, u.Phone, u.Instagram, u.ClientFrom, u.RoleID).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - CreateUser - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("AuthRepo - CreateUser - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetUser -.
func (r *AuthRepo) GetUser(ctx context.Context, id string) (entity.User, error) {
	var user entity.User
	sql, args, err := r.Builder.
		Select("id, name, surname, username, password, birth_date, tg_user_name, phone, instagram, client_from, role_id, created_at, updated_at").
		From("user").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return user, fmt.Errorf("AuthRepo - GetUser - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&user.ID, &user.Name, &user.Surname, &user.Username, &user.Password, &user.BirthDate, &user.TgUserName, &user.Phone, &user.Instagram, &user.ClientFrom, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, fmt.Errorf("AuthRepo - GetUser - row.Scan: %w", err)
	}

	return user, nil
}

// UpdateUser -.
func (r *AuthRepo) UpdateUser(ctx context.Context, u entity.User) error {
	sql, args, err := r.Builder.
		Update("user").
		Set("name", u.Name).
		Set("surname", u.Surname).
		Set("username", u.Username).
		Set("password", u.Password).
		Set("birth_date", u.BirthDate).
		Set("tg_user_name", u.TgUserName).
		Set("phone", u.Phone).
		Set("instagram", u.Instagram).
		Set("client_from", u.ClientFrom).
		Set("role_id", u.RoleID).
		Set("updated_at", "CURRENT_TIMESTAMP").
		Where("id = ?", u.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - UpdateUser - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - UpdateUser - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteUser -.
func (r *AuthRepo) DeleteUser(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("user").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - DeleteUser - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - DeleteUser - r.Pool.Exec: %w", err)
	}

	return nil
}

// -------------- Role --------------

// CreateRole -.
func (r *AuthRepo) CreateRole(ctx context.Context, role entity.Role) error {
	sql, args, err := r.Builder.
		Insert("role").
		Columns("name").
		Values(role.Name).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - CreateRole - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&role.ID)
	if err != nil {
		return fmt.Errorf("AuthRepo - CreateRole - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetRole -.
func (r *AuthRepo) GetRole(ctx context.Context, id string) (entity.Role, error) {
	var role entity.Role
	sql, args, err := r.Builder.
		Select("id, name, created_at, updated_at").
		From("role").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return role, fmt.Errorf("AuthRepo - GetRole - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return role, fmt.Errorf("AuthRepo - GetRole - row.Scan: %w", err)
	}

	return role, nil
}

// UpdateRole -.
func (r *AuthRepo) UpdateRole(ctx context.Context, role entity.Role) error {
	sql, args, err := r.Builder.
		Update("role").
		Set("name", role.Name).
		Set("updated_at", "CURRENT_TIMESTAMP").
		Where("id = ?", role.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - UpdateRole - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - UpdateRole - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteRole -.
func (r *AuthRepo) DeleteRole(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("role").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - DeleteRole - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - DeleteRole - r.Pool.Exec: %w", err)
	}

	return nil
}

// -------------- ClientType --------------

// CreateClientType -.
func (r *AuthRepo) CreateClientType(ctx context.Context, clientType entity.ClientType) error {
	sql, args, err := r.Builder.
		Insert("client_type").
		Columns("name").
		Values(clientType.Name).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - CreateClientType - r.Builder: %w", err)
	}

	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&clientType.ID)
	if err != nil {
		return fmt.Errorf("AuthRepo - CreateClientType - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// GetClientType -.
func (r *AuthRepo) GetClientType(ctx context.Context, id string) (entity.ClientType, error) {
	var clientType entity.ClientType
	sql, args, err := r.Builder.
		Select("id, name, created_at, updated_at").
		From("client_type").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return clientType, fmt.Errorf("AuthRepo - GetClientType - r.Builder: %w", err)
	}

	row := r.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&clientType.ID, &clientType.Name, &clientType.CreatedAt, &clientType.UpdatedAt)
	if err != nil {
		return clientType, fmt.Errorf("AuthRepo - GetClientType - row.Scan: %w", err)
	}

	return clientType, nil
}

// UpdateClientType -.
func (r *AuthRepo) UpdateClientType(ctx context.Context, clientType entity.ClientType) error {
	sql, args, err := r.Builder.
		Update("client_type").
		Set("name", clientType.Name).
		Set("updated_at", "CURRENT_TIMESTAMP").
		Where("id = ?", clientType.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - UpdateClientType - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - UpdateClientType - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteClientType -.
func (r *AuthRepo) DeleteClientType(ctx context.Context, id string) error {
	sql, args, err := r.Builder.
		Delete("client_type").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - DeleteClientType - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - DeleteClientType - r.Pool.Exec: %w", err)
	}

	return nil
}
