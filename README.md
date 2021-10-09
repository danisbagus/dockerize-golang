# dockerize-golang

Golang Appplication Using Docker
## Setup

Prepare necessary environemt by rename config.example to config.yml

Create golang app container

```bash
docker-compose up
```

## Testing

Hit transactions API

```bash
curl --location --request GET 'http://localhost:{appPort}/api/transactions' --header 'Content-Type: application/json'
```
