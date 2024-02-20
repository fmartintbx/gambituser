package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/fmartintbx/gambituser/awsgo"
	"github.com/fmartintbx/gambituser/bd"
	"github.com/fmartintbx/gambituser/models"
)

func main() {
	// Iniciar la ejecución de Lambda
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	// Tu lógica aquí
	awsgo.InitializeAWS()

	if !ValidateParameters() {
		fmt.Println("Error in Parameter. Should send 'SecretManager'.")
		err := errors.New("error , should send  SecretName")
		return event, err
	}
	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub =  " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error to read secret" + err.Error())
		return event, err
	}
	err = bd.SignUp(data)
	return event, err
}

func ValidateParameters() bool {
	var hasParameter bool
	_, hasParameter = os.LookupEnv("SecretName")
	return hasParameter
}
