package user

import (
	"github.com/MungaSoftwiz/org-authenticator-api/types"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// User model methods
func (s *Storage) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password, phone) VALUES (?, ?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password, user.Phone)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetUserByEmail(email string) (*types.User, error) {
	// register user
	u := new(types.User)
	err := s.db.Get(u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Storage) GetUserByID(id int) (*types.User, error) {
	u := new(types.User)
	err := s.db.Get(u, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
