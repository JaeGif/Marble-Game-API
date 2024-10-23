package repositories

import (
	storage "marble-game-api/cmd/database"
	"marble-game-api/cmd/models"
)

func CreateUser(user models.User) (models.User, error) {
  db := storage.GetDB()
  sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
  err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&user.Id)
  if err != nil {
    return user, err
  }
  return user, nil
}
