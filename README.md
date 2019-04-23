# errwrap
Errwrap is a Go (golang) library for wrapping errors.

base hashicorp/errwrap

## Installation and Docs

Install using `go get github.com/jiankunking/errwrap`.

#### Basic Usage

Below is a very basic example of its usage:

```go
// A function that always returns an error, but wraps it, like a real
// function might.
func tryOpen() error {
	_, err := os.Open("/i/dont/exist")
	if err != nil {
		return errwrap.Wrapf("Doesn't exist: {{err}}", err)
	}

	return nil
}
```
