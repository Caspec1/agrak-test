# Agrak Test API

## Connect PostgreSQL DB Using Docker

Run this for create a container for postgreSQL

```bash
    docker run --name some-postgres -e POSTGRES_USER=your -e POSTGRES_PASSWORD=yousecretpassword -p 5432:5432 -d postgres
```

## Set ENV
Clone the repo and create a .env file whit the next env:

```env
    PGHOST=localhost
    PGUSER=youruser
    PGPASSWORD=yourpassword
    PGDATABASE=agrak-test
    PGPORT=5432
    PORT=3000
```

## Building

Run this on the root of project for build api

```go
    go run build
```

## Execute

Run this for execute api

MacOs
```go
    ./agrak-test
```
Windows
```go
    agrak-test.exe
```

## Usage
Use the api with Postman or Thunder Client

## Requests

```go
    GET "/api/products" //Get all products
	POST "/api/products" // Create a product
	GET "/api/products/:sku" // Get product by sku
	PUT "/api/products/:sku" // Update product
	DELETE "/api/products/:sku" // Delete product
```

## Testing

Run this on root of the project for run all tests

```go
    go test ./...
```
