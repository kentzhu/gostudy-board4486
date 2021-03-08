package admin

import (
	"board4486/service/database"
)

func IsAdminLoginCorrect(username string, password string) bool {
	var passwordDb string
	rows := database.GetDB().QueryRow("SELECT PASSWORD FROM `admin` where username=? LIMIT 1", username)
	err := rows.Scan(&passwordDb)
	if err != nil {
		return false
	}
	if passwordDb != password {
		return false
	}
	return true
}
