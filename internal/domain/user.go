package domain

import (
	"context"
	"time"
)

// User represents the core business entity.
// It uses standard Go types to remain framework-agnostic.
type User struct {
	ID            int64      `json:"id"`
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	EmailVerified bool       `json:"email_verified"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

// UserRepository is the "Contract" that the Infrastructure layer must fulfill.
// We use context.Context for production-grade cancellation and tracing.
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	//GetByID(ctx context.Context, id int64) (*User, error)
	//GetByEmail(ctx context.Context, email string) (*User, error)
	//Update(ctx context.Context, user *User) error
	//Delete(ctx context.Context, id int64) error // Soft delete logic happens in implementation
}

// UserUseCase defines the business logic operations available for a User.
type UserUseCase interface {
	Register(ctx context.Context, user *User) error
	//Profile(ctx context.Context, id int64) (*User, error)
}

// Additional domain logic, validations, and methods can be added here as needed.
