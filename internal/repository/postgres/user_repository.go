package postgres

import (
	"context"
	"database/sql"
	"hexnity/internal/domain"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository links the implementation to the domain interface
func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	query := `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query, u.Username, u.Email).Scan(&u.ID, &u.CreatedAt)
}

// Implement other methods (GetByID, etc.) similarly...
