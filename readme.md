# Learn Go with tests

Exercise: https://quii.gitbook.io/learn-go-with-tests/

## Tests

### Running Tests

`cd src/hello && go test`

### Examples

`cd src/adder && go test -v`

### Benchmarks

`cd src/iteration && go test -bench=.`

### Test Coverage

`cd src/sum && go test -cover`

## Notes

Run `go mod init SOMENAME` for each folder

Use `errcheck` linter to find tests that does not check all error conditions.
