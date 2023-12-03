package secrets

import (
	"backend-site/src/config/aws.secrets/model"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"strconv"
)

func GetSecrets() (*model.MySQLPropertiesInterface, error) {
	// Configurar a sessão AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		fmt.Println("Erro ao criar a sessão AWS:", err)
		return nil, err
	}

	// Criar um cliente Secrets Manager
	client := secretsmanager.New(sess)

	// Nome da secret no Secrets Manager
	secretName := "arn:aws:secretsmanager:us-east-1:012821307542:secret:dev/mysql/api-jNAmkR"

	// Configurar o input para obter o valor da secret
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	// Recuperar o valor da secret
	result, err := client.GetSecretValue(input)
	if err != nil {
		fmt.Printf("Erro ao recuperar a secret %s: %v\n", secretName, err)
		return nil, err
	}

	// Se a secret tiver um JSON como SecretString, você pode analisá-lo para um mapa de chave/valor
	// Certifique-se de tratar possíveis erros durante a análise JSON
	secretMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(*result.SecretString), &secretMap)
	if err != nil {
		fmt.Printf("Erro ao analisar a secret JSON: %v\n", err)
		return nil, err
	}
	port, _ := strconv.ParseUint(secretMap["port"].(string), 10, 64)

	var mysqlProperties = model.NewMySQLProperties(
		secretMap["host"].(string),
		port,
		secretMap["dbname"].(string),
		secretMap["username"].(string),
		secretMap["password"].(string))
	return &mysqlProperties, nil
}
