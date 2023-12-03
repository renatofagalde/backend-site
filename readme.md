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


### Create S3 bucket to store codebuild artifacts
- br.com.likwi.artifacts.apps.backend-site-dev
- br.com.likwi.artifacts.apps.backend-site-hom
- br.com.likwi.artifacts.apps.backend-site-prd

### Actions necessary before pipeline creation
```shell
GOOS=linux GOARCH=amd64  CGO_ENABLED=0 go build -o backend-site main.go
```

```shell
zip backend-site.zip backend-site
```

```shell
aws s3 cp backend-site.zip s3://br.com.likwi.artifacts.apps.backend-site-dev
```

```shell
aws cloudformation package --template-file template.yaml --output-template-file packaged.yaml --s3-bucket br.com.likwi.artifacts.apps.backend-site-dev
```

[//]: # (aws cloudformation deploy --template-file packaged.yaml --stack-name backend-site-dev-stack --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM)
```shell
aws cloudformation deploy --template-file packaged.yaml --stack-name backend-site-dev-stack --capabilities CAPABILITY_IAM
```
deletar a stack
```shell
aws cloudformation delete-stack --stack-name backend-site-stack-dev
```
local test FindByID
```shell
export CGO_ENABLED=0 && sam build && sam local invoke -e ./_miscellaneous/lambda/apigateway-aws-proxy-get.json BackendSite
```

### BuildSpec
- Use this documentation: [AWS CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/getting-started-create-build-spec-console.html)
- Available runtimes: [runtime](https://docs.aws.amazon.com/codebuild/latest/userguide/available-runtimes.html)

```yaml
version: 0.2

phases:
  install:
    runtime-versions:
      java: corretto11
  pre_build:
    commands:
      - echo Nothing to do in the pre_build phase...
  build:
    commands:
      - echo Build started on `date`
      - mvn install
  post_build:
    commands:
      - echo Build completed on `date`
artifacts:
  files:
    - target/messageUtil-1.0.jar
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


# Testing project localy

### Simple call
```shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
go mod tidy
go build -o main main.go
sam local invoke -e ./_miscellaneous/lambda/simple_event.json BackendSiteAppFunction
```
### GET Content route
```shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
go mod tidy
go build -o main main.go
sam local invoke -e ./_miscellaneous/lambda/apigateway-aws-proxy-get.json BackendSiteAppFunction
```

[Commits](https://github.com/iuricode/padroes-de-commits)