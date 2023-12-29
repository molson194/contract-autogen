# Contract Testing

## Ideas:
* 100% code coverage with contract tests that run during merge and deployment is the best defense against regressions.
* Local development should use contracts when serving dependency responses.

## Getting started:
* In separate terminals, run `go run .` with `useServiceYContractServer = false` from the service_x and service_y folder. Then, navigate to `http://localhost:8080/api?name=Matt` to see the api response.
* Run contract tests for both services with `go test` command in service folders.
* Run local instances of service_x `go run .` with `useServiceYContractServer = true`. Navigate to `http://localhost:8080/api?name=Matt`. See that the contract gets the response from service_y.
