#!/bin/bash
docker-compose up -d
sleep 3
migrate -path migrations -database "postgres://postgres:1234@localhost/restapi_dev?sslmode=disable" up