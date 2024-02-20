package bd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/fmartintbx/gambituser/tools"
	"github.com/fmartintbx/gambituser/models"
)

func SignUp(Sig models.SignUp) error {
	fmt.Println("Start Register")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	Sentence := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES('" + Sig.UserEmail + "', '" + Sig.UserUUID + "', '" + tools.MySQLDate() + "')"
	fmt.Println(Sentence)

	_, err = Db.Exec(Sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

    fmt.Println("SignUp >  Congratulation")
	return nil
}
