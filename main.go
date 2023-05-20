package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/JulianDavidGamboa/gambitUser/bd"
	"github.com/JulianDavidGamboa/gambitUser/models"
	"os"

	"github.com/JulianDavidGamboa/gambitUser/awsgo"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(StartLambda)
}

func StartLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()

	if !ValidateParameters() {
		fmt.Println("Error en los parametros. Debe enviar 'SecretManager'")
		err := errors.New("Error en los parametros. Debe enviar 'SecretManager'")
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
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el secret " + err.Error())
		return event, err
	}

	err = bd.SignUp(data)
	return event, err
}

func ValidateParameters() bool {
	var getParameter bool
	_, getParameter = os.LookupEnv("SecretName")
	return getParameter
}
