# Golang Appplication Using Docker

## Instalation

Pull base image from [docker-hub](https://hub.docker.com/repository/docker/danisbagus/base-go)

Create `.env` file

```
cp .env.example .env
```

Build new image

```
docker build -t danisbagus/app-go .
```

## Running App

Create and run new container

```
./run.sh
```
