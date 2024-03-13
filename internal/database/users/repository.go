package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
)

func New(userDB *pgx.Conn, timeout time.Duration) *Repository {
	return &Repository{userDB: userDB, timeout: timeout}
}

type Repository struct {
	userDB  *pgx.Conn
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateUserReq) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	u.ID = req.ID
	u.Username = req.Username
	u.Password = req.Password
	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt

	sqlStr := "INSERT INTO users (id, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.userDB.Exec(ctx, sqlStr, u.ID, u.Username, u.Password, u.CreatedAt, u.UpdatedAt)

	return u, err
}

func (r *Repository) FindByID(ctx context.Context, userID uuid.UUID) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	sqlStr := "SELECT * FROM users WHERE id = $1"

	err := r.userDB.QueryRow(ctx, sqlStr, userID).Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt)

	return u, err
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (database.User, error) {
	var u database.User

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	sqlStr := "SELECT * FROM users WHERE username = $1"

	err := r.userDB.QueryRow(ctx, sqlStr, username).Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.UpdatedAt)

	return u, err
}
