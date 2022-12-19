# Testing in Go
- Recap: types of tests
    - Unit testing - a single module/function
    - Integration testing - test functions/modules working together (integrated)
- Go makes no distinction between the two
    - Same process to create both types

# Setup
- Tests are written in separate files, with extra **_test** suffix
    - importantPkg.go -> importantPkg_test.go
- Unit tests should be in the same package
- The testing package is used to create tests and must be imported in each test file
- run with `$ go test <package_path>`

# Available Testing Functions
- `Fail()` - mark the test as failed
    - `Errorf(string)` - fail and add a message
- `FailNow()` - mark the test as failed and abort the current test
    - `Fatalf(string)` - fail, abort, and add a message
- `Logf()` - Equivalent to `Printf(..)` but only when the test fails