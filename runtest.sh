#!/bin/bash
export BINDADDRESS=":8080"
export DATABASEURL="host=localhost user=postgres password=1234 dbname=restapi_dev sslmode=disable"

make test