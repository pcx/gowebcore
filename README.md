# Core modules for a custom go web application

## Errors

```
# Creating a new error from a standard error
err := errors.New(stdErr)

# Create a new error from scratch, like fmt.Errorf does
err := errors.Format("This is error %s", someVar)

# Return a nil error, for a standard error this is usually just `nil`
return errors.Nil()

# Checking if error exists, for standard error this is usually `if err != nil`
if err.Present() {
        doSomething()
}

# Printing error stacktrace
err.Report()
```

## Custom Errors

Check `errors/custom_errors.go` for more examples. An example of an authorization error is:

```
func Unauthorized(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("Unauthorized: %s", message), a...),
		UnauthorizedCode,
		false,
		401,
	)
}


// Creating an error with when authorization fails
return errors.Unauthorized("Provided API token is invalid because: %v", decodeErr)

```
