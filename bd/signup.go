package bd

import (
	"fmt"
	"github.com/JulianDavidGamboa/gambitUser/models"
	"github.com/JulianDavidGamboa/gambitUser/tools"
)

func SignUp(sign models.SignUp) error {
	fmt.Println("Start Register")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sqlStatement := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sign.UserEmail + "','" + sign.UserUUID + "','" + tools.DateMySQL() + "')"
	fmt.Println(sqlStatement)

	_, err = Db.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
