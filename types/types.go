package types

type User struct {
	ID        int    `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

type UserStorage interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

// RegisterUserPayload represents the payload used for registering a user in the system.
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=32"`
	Phone     string `json:"phone" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

// Organisation Model
type Organisation struct {
	OrgID       string `json:"orgId" db:"orgId"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type OrganisationStorage interface {
	GetAllOrganisations() ([]Organisation, error)
	GetOrganisationByID(id string) (*Organisation, error)
	CreateOrganisation(Organisation) error
	AddUserToOrganisation(orgID, userID string) error
}
