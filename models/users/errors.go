package users

import "errors"

var (
	// ErrNotFound is returned when a resource cannot be found
	// in the database.
	ErrNotFound = errors.New("models: resource not found")

	// ErrIDInvalid is returned when an invalid id is passed to methods
	// like Delete.
	ErrIDInvalid = errors.New("models: ID provided was invalid")

	// ErrPasswordIncorrect is returned when an invalid password
	// is used when attempting to authenticate a user.
	ErrPasswordIncorrect = errors.New("models: incorrect password provided")

	// ErrPasswordTooShort is returned when a user tries to set
	// a password that is less than 8 characters long
	ErrPasswordTooShort = errors.New("models: password must be at least 8 characters long")

	// ErrPasswordRequired is returned when a create is attempted
	// without a user password provided.
	ErrPasswordRequired = errors.New("models: password is required")

	// ErrEmailRequired is returned when an email address is
	// not provided when creating a user
	ErrEmailRequired = errors.New("models: email address is required")

	// ErrEmailInvalid is returned when an email address provided
	// does not match any of our requirements
	ErrEmailInvalid = errors.New("models: email address is not valid")

	// ErrEmailTaken is returned when an update or create is attempted
	// with an email address what is already in use
	ErrEmailTaken = errors.New("models: email address is already taken")

	// ErrRememberRequired is returned when a create or update
	// is attempted without a user remember token hash
	ErrRememberRequired = errors.New("models: remember token is required")

	// ErrRememberTooShort is returned when a remember token is
	// not at least 32 bytes
	ErrRememberTooShort = errors.New("models: remember token is not at least 32 bytes")

	userPwPepper  = "secret-secret-secret"
	hmacSecretKey = "secret-hmac-hmacSecretKey"
)
