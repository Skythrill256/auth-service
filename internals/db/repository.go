package db

import (
	"database/sql"
	"errors"
	"github.com/Skythrill256/auth-service/internals/models"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) CreateUser(user *models.User) error {
	query := `INSERT INTO users(email,password,is_verified) VALUES($1,$2,$3)`
	err := repo.DB.QueryRow(query, user.Email, user.Password, user.IsVerified, user.GoogleID).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id ,email,password is_verified, created_at, updated_at, google_id FROM users WHERE id=$1`
	err := repo.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GoogleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id , email , password , is_verified , created_at, updated_at, google_id FROM users WHERE email=$1`
	err := repo.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GoogleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (repo *Repository) VerifyUserEmail(email string) error {
	query := `UPDATE users SET is_verified = true, updated_at = CURRENT_TIMESTAMP WHERE email = $1`

	_, err := repo.DB.Exec(query, email)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetUserByGoogleID(googleID string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password, is_verified, created_at, updated_at, google_id FROM users WHERE google_id = $1`
	err := repo.DB.QueryRow(query, googleID).Scan(&user.ID, &user.Email, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt, &user.GoogleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
