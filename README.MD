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

To run only server container
```shell
make build
make run
```

After finished work
```shell
make rm
```

### Go CLI
From main project directory
```
go run cmd/user-crud/main.go cmd/user-crud/di.go
```