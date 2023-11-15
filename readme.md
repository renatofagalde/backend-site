# Tarefas para iniciar o projeto
 
### Initialize project module
#### module:
```shell
  go mod init backend-site
  ```
#### dotenv:
```shell
  go get github.com/joho/godotenv
  ```

#### aws lambda:
```shell
  go get github.com/aws/aws-lambda-go
  ```



### This is the explanation of our file structure:
```tree
├── .env
├── docker-compose.yaml
├── go.mod
├── go.sum
├── init_dependencies.go
├── main.go
├── README.md
├── src
│   ├── config
│   │   ├── database
│   │   │   ├── mongodb
│   │   │   │   └── mongodb.go
│   │   │   └── mysqldb
│   │   │       └── mysql-gorm.go
│   │   ├── logger
│   │   │   └── logger.go
│   │   ├── rest_err
│   │   │   └── rest_err.go
│   │   └── validation
│   │       └── user-validate.go
│   ├── controller
│   │   ├── model
│   │   │   ├── request
│   │   │   │   ├── login_request.go
│   │   │   │   └── user_request.go
│   │   │   └── response
│   │   │       └── user_response.go
│   │   ├── routes
│   │   │   └── routes.go
│   │   └── user
│   │       ├── create_controller.go
│   │       ├── create_controller_test.go
│   │       ├── delete_controller.go
│   │       ├── delete_controller_test.go
│   │       ├── find_controller.go
│   │       ├── find_controller_test.go
│   │       ├── login_controller.go
│   │       ├── login_controller_test.go
│   │       ├── update_controller.go
│   │       ├── update_controller_test.go
│   │       └── user_controller.go
│   ├── _miscellaneous
│   │   ├── client-http
│   │   │   ├── http-client.env.json
│   │   │   └── user.http
│   │   ├── idea
│   │   │   └── run
│   │   │       └── golang-basic.run.xml
│   │   └── mongodb
│   │       ├── docker-entrypoint-initdb.d
│   │       │   └── mongo-init.js
│   │       └── mongo-init.js
│   ├── model
│   │   ├── repository
│   │   │   ├── create_user_repository.go
│   │   │   ├── create_user_repository_test.go
│   │   │   ├── delete_repository.go
│   │   │   ├── delete_repository_test.go
│   │   │   ├── entity
│   │   │   │   ├── convert
│   │   │   │   │   ├── convert_domain_to_entity.go
│   │   │   │   │   ├── convert_entity_to_bson.go
│   │   │   │   │   └── convert_entity_to_domain.go
│   │   │   │   └── user_entity.go
│   │   │   ├── find_user_repository_byEmail_test.go
│   │   │   ├── find_user_repository_byID_test.go
│   │   │   ├── find_user_repository.go
│   │   │   ├── update_repository.go
│   │   │   ├── update_repository_test.go
│   │   │   └── user_repository.go
│   │   ├── service
│   │   │   ├── create_service.go
│   │   │   ├── create_service_test.go
│   │   │   ├── delete_service.go
│   │   │   ├── delete_service_test.go
│   │   │   ├── find_user_byEmailAndPassword_service.go
│   │   │   ├── find_user_byEmailAndPassword_service_test.go
│   │   │   ├── find_user_byEmail_service.go
│   │   │   ├── find_user_byEmail_service_test.go
│   │   │   ├── find_user_byID_service.go
│   │   │   ├── find_user_byID_service_test.go
│   │   │   ├── login_service.go
│   │   │   ├── login_service_test.go
│   │   │   ├── update_service.go
│   │   │   ├── update_service_test.go
│   │   │   └── user_interface.go
│   │   ├── user_domain.go
│   │   ├── user_domain_interface.go
│   │   ├── user_domain_password.go
│   │   ├── user_domain_tojson.go
│   │   └── user_token_domain.go
│   └── view
│       └── convert_domain_to_response.go
└── test
    ├── integration
    │   ├── create_user_test.go
    │   ├── delete_user_test.go
    │   ├── find_user_test.go
    │   └── setup.go
    ├── mocks
    │   ├── user_domain_interface_mock.go
    │   ├── user_repository_mock.go
    │   └── user_service_mock.go
    └── mongodb
        └── openconnection.go
```