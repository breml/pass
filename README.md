# github.com/breml/pass

[![Test Status](https://github.com/breml/pass/workflows/Main/badge.svg)](https://github.com/breml/pass/actions?query=workflow%3AMain)
 [![Go Report Card](https://goreportcard.com/badge/github.com/breml/pass)](https://goreportcard.com/report/github.com/breml/pass) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Package pass allows you to early end Go unit tests successfully, which also
works from called functions outside of the test function.

## Reasoning

While the testing package provides several methods on `*testing.T` to fail a
test early, there is no method provided to end a test early successfully.
Obviously, you can use `return` to end a test early, but this only works from
inside of the test function and it strictly requires the use of conditions to
verify the test result.

Me personally, I aim to keep my test functions as simple as possible and I try
to avoid conditions in test functions as much as possible. Sometimes, in more
complext test scenarios find myself in a situation, where it is not possible to
continue a test case, because I expect a function to fail (those the test is
successful) but it is not safe to continue the test, since the necessary pre-
conditions for the subsequent logic are not met. In such a case, I would like
to end the test case early, but still mark it as successful.

This is where this package comes into play. It allows you to end a test case
early, but still mark it as successful.

## Distinction to (*testing.T).SkipNow

The `testing` package provides the `(*testing.T).SkipNow` function, which allows
you to skip the current test case. The effect of the two functions is very
similar. The only difference is, that `(*testing.T).SkipNow` will print a message
indicating that the test was skipped (optionally with a reason, if
`(*testing.T).Skip` or `(*testing.T).Skipf` was used) in the following form:

```none
--- SKIP: TestSkipNow (0.00s)
```

The `pass` package does mark the test as successfully completed and those the
message on the console will have the follwing form:

```none
--- PASS: TestNow (0.00s)
```

## Example

See [`example_test.go`](example_test.go) for a descriptive example.

## Installation

```bash
go get github.com/breml/pass
```

## Author

Copyright 2024 by Lucas Bremgartner ([breml](https://github.com/breml))

## License

[MIT License](LICENSE)
