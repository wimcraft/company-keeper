This codebase is a fully-fledged full-stack application built with Golang/Echo. It includes CRUD operations, authentication, routing, and more.

# Getting started

Install latest [Golang](https://golang.org/doc/install) (go1.19+), docker.

1. Install dependencies
```bash
$ go mod download
```

2. Start the app and database using Docker Compose:
```bash
$ docker-compose up app db -d
```

The app listens and serves on the port specified in the .env file.

## Environment Config

Create a .env file to configure the server environment and a .test.env file to configure the test environment.

## Tests

Before running tests, create a separate database for testing purposes and specify credentials in `.test.env`. Then, run the tests with the following command:
```bash
$ make test
```

## Linter 

Before running linter install [golangci-lint](https://golangci-lint.run/usage/install/)
```bash
$ make lint
```


## Documentation

To generate documentation, install the necessary dependencies:

```
go install github.com/swaggo/swag/cmd/swag@latest
```

Then, generate the OpenAPI documentation into the `/docs` folder:
```
swag init
```

After starting the server, the documentation will be available in the browser at http://localhost:8080/swagger/index.html. The documentation includes information about the API endpoints and their parameters.