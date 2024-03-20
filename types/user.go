package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPassWordLen  = 3
)

type UpdateUserReq struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *UpdateUserReq) ToBson() bson.M {
	bsonValue := bson.M{}
	if len(u.FirstName) > 0 {
		bsonValue["firstName"] = u.FirstName
	}
	if len(u.LastName) > 0 {
		bsonValue["lastName"] = u.LastName
	}
	return bsonValue
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func (p *CreateUserParams) Validate() error {
	if len(p.FirstName) < minFirstNameLen {
		return fmt.Errorf("minimum firstName length should be %d", minFirstNameLen)
	}
	if len(p.LastName) < minLastNameLen {
		return fmt.Errorf("minimum firstName length should be %d", minFirstNameLen)
	}
	if !ValidateEmail(p.Email) {
		return fmt.Errorf("invalid email")
	}
	if err := ValidatePassword(p.Password); err != nil {
		return err
	}
	return nil
}

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")
	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) error {
	minLength := 8
	uppercaseRegex := `[A-Z]+`
	lowercaseRegex := `[a-z]+`
	digitRegex := `[0-9]+`
	specialCharRegex := `[!@#$%^&*]+`

	if len(password) < minLength {
		return fmt.Errorf("Password must be at least %d characters long", minLength)
	}

	if !regexp.MustCompile(uppercaseRegex).MatchString(password) {
		return fmt.Errorf("Password must contain at least one uppercase letter (A-Z)")
	}

	if !regexp.MustCompile(lowercaseRegex).MatchString(password) {
		return fmt.Errorf("Password must contain at least one lowercase letter (a-z)")
	}

	if !regexp.MustCompile(digitRegex).MatchString(password) {
		return fmt.Errorf("Password must contain at least one digit (0-9)")
	}

	if !regexp.MustCompile(specialCharRegex).MatchString(password) {
		return fmt.Errorf("Password must contain at least one special character (!@#$%^&*)")
	}

	return nil // Password is valid
}
func NewUsersFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil

}
