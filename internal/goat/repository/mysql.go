package repository

import (
	"context"
	"database/sql"

	"github.com/bontusss/goat/internal/goat"
	"github.com/bontusss/goat/internal/goat/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// MySQLUserRepository is a struct for MySQL operations, encapsulating the DB connection.
type MySQLUserRepository struct {
	db *sql.DB // MySQL DB connection.
}

// NewMySQLUserRepository initializes a new MySQLUserRepository with a given DSN (Data Source Name).
func NewMySQLUserRepository(dsn string) (*MySQLUserRepository, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Ensure that an index on the email field is created to enforce uniqueness.
	_, err = db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_email ON users(email)")
	if err != nil {
		return nil, err
	}

	return &MySQLUserRepository{db: db}, nil
}

// Register adds a new user to the MySQL database. It hashes the user's password before saving.
func (r *MySQLUserRepository) Register(user *models.User) error {
	ctx := context.Background()
	user.ID = uint(uuid.New().ID()) // Generate a unique ID for the user.

	// Hash the user's password for secure storage.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert the new user into the database.
	_, err = r.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES (?, ?, ?)", user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// Login checks a user's credentials against the stored values in the MySQL database.
// If the credentials are valid, it returns the user object; otherwise, it returns an error.
func (r *MySQLUserRepository) Login(email, password string) (*models.User, error) {
	ctx := context.Background()
	user := &models.User{}

	// Attempt to find the user by email.
	err := r.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
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

func (r *MySQLUserRepository) DeleteUser(id uint) error {
	ctx := context.Background()
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLUserRepository) ResetPassword(email string, newPassword string) error {
	ctx := context.Background()
	_, err := r.db.ExecContext(ctx, "UPDATE users SET password = ? WHERE email = ?", newPassword, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLUserRepository) UpdateUser(user *models.User) error {
	ctx := context.Background()
	_, err := r.db.ExecContext(ctx, "UPDATE users SET email = ?, password = ? WHERE id = ?", user.Email, user.Password, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLUserRepository) GetUserByID(id uint) (*models.User, error) {
	ctx := context.Background()
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLUserRepository) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()
	user := &models.User{}
	err := r.db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *MySQLUserRepository) GetAllUsers() ([]*models.User, error) {
	ctx := context.Background()
	rows, err := r.db.QueryContext(ctx, "SELECT id, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *MySQLUserRepository) GetUsersByEmail(email string) ([]*models.User, error) {
	ctx := context.Background()
	rows, err := r.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}


