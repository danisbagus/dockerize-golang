#!/bin/bash
#

docker run --rm --network="host" --name app-go -v $(pwd)/.env:/go/src/.env danisbagus/app-go

exec "$@"
