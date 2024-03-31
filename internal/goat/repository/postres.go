package repository

import (
	"context"

	"github.com/bontusss/goat/internal/goat"
	"github.com/bontusss/goat/internal/goat/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

// PostgreSQLUserRepository is a struct for PostgreSQL operations, encapsulating connection information.
type PostgreSQLUserRepository struct {
	conn *pgx.Conn // PostgreSQL connection for database access.
}

// NewPostgreSQLUserRepository initializes a new PostgreSQLUserRepository with a given connection string.
func NewPostgreSQLUserRepository(connString string) (*PostgreSQLUserRepository, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	// Ensure that an index on the email field is created to enforce uniqueness.
	_, err = conn.Exec(ctx, "CREATE UNIQUE INDEX IF NOT EXISTS users_email_idx ON users (email)")
	if err != nil {
		return nil, err
	}

	return &PostgreSQLUserRepository{conn: conn}, nil
}

// Register adds a new user to the PostgreSQL database. It hashes the user's password before saving.
func (r *PostgreSQLUserRepository) Register(user *models.User) error {
	ctx := context.Background()
	user.ID = uint(uuid.New().ID()) // Generate a unique ID for the user.

	// Hash the user's password for secure storage.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert the new user into the database.
	_, err = r.conn.Exec(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// Login checks a user's credentials against the stored values in the PostgreSQL database.
// If the credentials are valid, it returns the user object; otherwise, it returns an error.
func (r *PostgreSQLUserRepository) Login(email, password string) (*models.User, error) {
	ctx := context.Background()
	user := &models.User{}

	// Attempt to find the user by email.
	err := r.conn.QueryRow(ctx, "SELECT id, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			// If no row is found, return an invalid credentials error.
			return nil, goat.ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify the password against the hashed password stored in the database.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// If the password does not match, return an invalid credentials error.
			return nil, goat.ErrInvalidCredentials
		}
		return nil, err
	}
	return user, nil
}

func (r *PostgreSQLUserRepository) DeleteUser(id uint) error {
	ctx := context.Background()
	_, err := r.conn.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLUserRepository) ResetPassword(email string, newPassword string) error {
	ctx := context.Background()
	_, err := r.conn.Exec(ctx, "UPDATE users SET password = $1 WHERE email = $2", newPassword, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLUserRepository) UpdateUser(user *models.User) error {
	ctx := context.Background()
	_, err := r.conn.Exec(ctx, "UPDATE users SET email = $1, password = $2 WHERE id = $3", user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgreSQLUserRepository) GetUserByID(id uint) (*models.User, error) {
	ctx := context.Background()
	user := &models.User{}
	err := r.conn.QueryRow(ctx, "SELECT id, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

