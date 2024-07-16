package models

import (
	"net/http"
	"regexp"
)

type AuthResponse struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
}

type SessionUser struct {
	Id         string
	Email      string
	ProfilePic string
}

type User struct {
	Id      string
	Email   string
	Profile *UserProfile
}

type UserProfile struct {
	FirstName          string
	ProfilePic         string
	PhoneNumber        string
	OnboardingComplete bool
}

type CreateAccountFormValues struct {
	Email           string
	Password        string
	PasswordConfirm string
}

type LoginAccountFormValues struct {
	Email    string
	Password string
}

func ParseCreateAccountFormValuesAndValidate(r *http.Request) (CreateAccountFormValues, map[string]string) {
	err := r.ParseForm()
	if err != nil {
		return CreateAccountFormValues{}, map[string]string{}
	}

	values := CreateAccountFormValues{
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		PasswordConfirm: r.FormValue("passwordConfirm"),
	}

	errors := ValidateCreateAccountFormValues(values)
	return values, errors
}

func ValidateCreateAccountFormValues(values CreateAccountFormValues) map[string]string {
	errors := map[string]string{}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if values.Email == "" {
		errors["email"] = "Email is required"
	} else if !emailRegex.MatchString(values.Email) {
		errors["email"] = "Invalid email format"
	}

	if len(values.Password) < 8 {
		errors["password"] = "Password must be at least 8 characters long"
	}

	if len(values.PasswordConfirm) <= 0 {
		errors["passwordConfirm"] = "Password Confirmation is required"
	} else if values.Password != values.PasswordConfirm {
		errors["passwordConfirm"] = "Confirmation password does not match"
	}

	return errors
}

func ParseLoginAccountFormValuesAndValidate(r *http.Request) (LoginAccountFormValues, map[string]string) {
	err := r.ParseForm()
	if err != nil {
		return LoginAccountFormValues{}, map[string]string{}
	}

	values := LoginAccountFormValues{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	errors := ValidateLoginAccountFormValues(values)
	return values, errors
}

func ValidateLoginAccountFormValues(values LoginAccountFormValues) map[string]string {
	errors := map[string]string{}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if values.Email == "" {
		errors["email"] = "Email is required"
	} else if !emailRegex.MatchString(values.Email) {
		errors["email"] = "Invalid email format"
	}

	if len(values.Password) < 8 {
		errors["password"] = "Password must be at least 8 characters long"
	}

	return errors
}
