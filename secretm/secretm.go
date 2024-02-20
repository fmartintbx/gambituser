package secretm

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/fmartintbx/gambituser/awsgo"
	"github.com/fmartintbx/gambituser/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println(" > Requesting Secret " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	secretValue, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}
	json.Unmarshal([]byte(*secretValue.SecretString), &secretData)
	fmt.Println(" > Secret Read OK " + secretName)
	return secretData, nil
}
