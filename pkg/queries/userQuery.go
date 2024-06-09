package queries

import (
	"github.com/JacobRWebb/InventoryManagement/pkg/database"
	"github.com/JacobRWebb/InventoryManagement/pkg/models"
	"github.com/JacobRWebb/InventoryManagement/pkg/utils"
)

func CreateUser(user *models.User) error {

	existingUser, err := GetUserByEmail(user.Email)

	if existingUser != nil {
		return &utils.CustomError{ErrorMessage: "Email is already taken."}
	}

	if err != nil {
		return err
	}

	query := `INSERT INTO users VALUES ($1, $2, $3)`

	_, err = database.DB.Exec(query, user.Id, user.Email, user.PasswordHash)

	return err
}

func GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT Id, Email, PasswordHash FROM users WHERE Email = $1"

	var user models.User

	err := database.DB.QueryRow(query, email).Scan(&user.Id, &user.Email, &user.PasswordHash)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
