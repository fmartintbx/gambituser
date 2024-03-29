package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/fmartintbx/gambituser/secretm"
	"github.com/fmartintbx/gambituser/models"
)

var SecretModel models.SecretRDSJson
var err error

var Db *sql.DB

// ReadSecret retrieves secret from environment variables
func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv(os.Getenv("SecretName")))
	return err
}

// DbConnect establishes connection to the database
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Successful connection to the database")
	return nil
}

// ConnStr constructs connection string using secret information
func ConnStr(clave models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = clave.Username
	authToken = clave.Password
	dbEndpoint = clave.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
