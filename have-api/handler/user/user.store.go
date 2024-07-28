package user

// import (
// 	"github.com/ohmtanawin02/go-postgres-basic/have-api/config"
// 	"github.com/ohmtanawin02/go-postgres-basic/have-api/models"
// 	"golang.org/x/crypto/bcrypt"
// )

// func CreateUser(user *models.User) error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = config.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, string(hashedPassword))
// 	return err
// }

// func GetUserByUsername(username string) (*models.User, error) {
// 	var user models.User
// 	err := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username=$1", username).Scan(&user.ID, &user.Username, &user.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
