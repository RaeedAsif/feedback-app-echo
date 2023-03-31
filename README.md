# feedback-app-echo
# code: feedback-app-echo

## Preface
This repository is the code repo of web application API's for Feedback app as for dogoodpoints assignment.
This sample uses [Echo](https://echo.labstack.com/) as web application framework, [Gorm](https://gorm.io/) as OR mapper and [Zap logger](https://pkg.go.dev/go.uber.org/zap) as logger.
This sample application provides only several functions as Web APIs.

## Technologies/Stack
1. Golang v1.19
2. Postgres (optional)

## Install
Perform the following steps:
1. Download and install [Visual Studio Code(VS Code)](https://code.visualstudio.com/).
2. Download and install [Golang](https://golang.org/).

## Init DB
Since this application on default use in memory storage you can also run queries and store data on a postgres db, 
if so then perform the following steps:
1. Download and install docker. (https://www.docker.com/)
2. Download and intsall docker-compose (https://docs.docker.com/compose/install/)
3. Run "sudo docker-compose up -d"
4. Run "sudo docker ps" to check if contianer's up and running.
5. P.S define your enviroment variables correctly, template is available in ".env.sample"
6. make migrate-db

## Run backend service
perform the following steps:
1. cd ./backend/ 
2. make run
*Please make sure to set your enviroment variables in .env file , template is in ".env.sample"

## Run Swagger docs
perform the following steps:
1. browse to https://localhost:8080/swagger/index.html
