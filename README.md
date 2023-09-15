## Run tests
```
go test ./... -v
```

## Run tests with coverage
```
go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out
```