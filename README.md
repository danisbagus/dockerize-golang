# dockerize-golang

Golang Appplication Using Docker
## Setup

Create golang app container

```bash
docker-compose up
```

## Testing

Hit transactions API

```bash
curl --location --request GET 'http://localhost:8100/api/transactions' --header 'Content-Type: application/json'
```
