package decoding

import (
	"encoding/json"
	"errors"
	"net/http"

	"gopkg.in/matryer/respond.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ok interface {
	OK() error
}

type ErrMissingField string

func (e ErrMissingField) Error() string {
	return string(e) + " is required"
}

type User struct {
	ID           bson.ObjectId
	Name         string
	Email        string
	PasswordHash string
	State        string
}

type NewUser struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
}

func (u *NewUser) OK() error {
	if len(u.Email) == 0 {
		return ErrMissingField("email")
	}

	if len(u.Password) == 0 {
		return ErrMissingField("password")
	}

	if u.Password != u.PasswordConfirm {
		return errors.New("passwords don't match")
	}

	return nil
}

func (u *NewUser) CreateUser(c *mgo.Collection) (*User, error) {
	user := &User{
		ID:           bson.NewObjectId(),
		State:        "new",
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: md5(u.Email + u.Password),
	}

	if err := c.Insert(user); err != nil {
		return nil, err
	}

	return user, nil
}

func decode(r *http.Request, v ok) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	return v.OK()
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	if err := decode(r, &u); err != nil {
		respond.With(w, r, http.StatusBadRequest, err)
		return
	}

	user, err := u.CreateUser(nil)
	if err != nil {
		respond.With(w, r, http.StatusInternalServerError, err)
		return
	}
	respond.With(w, r, http.StatusCreated, user)
}

func md5(s string) string {
	return s
}
