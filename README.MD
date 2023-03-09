# Simple server

## Server allows to make CRUD operations for User entity.

See [schema.json](./doc/schema.json) for endpoints documentation.

Server can be run using `go` CLI or `Docker`.

### Docker compose
```shell
make up
```
`http://localhost:1234` - server entrypoint

`http://localhost:81` - hosted swagger ui

To stop containers
```shell
make down
```

### Go CLI
From main project directory
```
go run cmd/user_crud/main.go cmd/user_crud/di.go
```