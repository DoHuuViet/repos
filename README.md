## Usage

First run `go build app.go`. Once build completed you can use `/.app` to start server

Example request:

List repositories of specified user

```
curl http://localhost:5000/DoHuuViet/repositories
```

or get repository by name

```
curl http://localhost:5000/DoHuuViet/anni
```

## Test

Run `go test app_test.go app.go` to run unittest
