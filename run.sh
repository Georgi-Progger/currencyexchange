#!/bin/bash
set -a
source .env
set +a

export GOPATH=$HOME/projects

goose -dir db/migrations postgres "user=$DB_USER dbname=$DB_NAME password=$DB_PASSWORD host=$DB_HOST port=$DB_PORT sslmode=disable" up