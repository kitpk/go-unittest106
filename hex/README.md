Guide run test coverage

1. run test coverage
```
go test ./... -cover
```
2. build file coverage.out
```
go test ./... -coverprofile=coverage.out
```
3. convert file coverage.out to coverage.html
```
go tool cover -html=coverage.out -o coverage.html
```