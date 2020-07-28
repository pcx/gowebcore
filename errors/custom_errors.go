package errors

import "fmt"

const (
	RecordNotFoundCode    ErrorCode = 1000
	InvalidTaskArgsCode   ErrorCode = 1001
	PanicRecoveredCode    ErrorCode = 1002
	ValidationFailedCode  ErrorCode = 1003
	UnauthorizedCode      ErrorCode = 1004
	ParamsInvalidCode     ErrorCode = 1005
	TransactionFailedCode ErrorCode = 1006
	InvalidJobArgsCode    ErrorCode = 1007
)

func RecordNotFound(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("RecordNotFound: %s", message), a...),
		RecordNotFoundCode,
		false,
		404,
	)
}

func ParamsInvalid(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("ParamsInvalid: %s", message), a...),
		ParamsInvalidCode,
		false,
		400,
	)
}

func ValidationFailed(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("ValidationFailed: %s", message), a...),
		ValidationFailedCode,
		false,
		400,
	)
}

func BadRequest(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("BadRequest: %s", message), a...),
		ValidationFailedCode,
		false,
		400,
	)
}

func InvalidTaskArgs(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("InvalidTaskArgs: %s", message), a...),
		InvalidTaskArgsCode,
		true,
		400,
	)
}

func PanicRecovered(err interface{}) Error {
	return build(
		fmt.Errorf("PanicRecovered: %v", err),
		PanicRecoveredCode,
		true,
		500,
	)
}

func Unauthorized(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("Unauthorized: %s", message), a...),
		UnauthorizedCode,
		false,
		401,
	)
}

func TransactionFailed(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("TransactionFailed: %s", message), a...),
		TransactionFailedCode,
		false,
		400,
	)
}

func InvalidJobArgs(message string, a ...interface{}) Error {
	return build(
		fmt.Errorf(fmt.Sprintf("InvalidJobArgs: %s", message), a...),
		InvalidTaskArgsCode,
		false,
		400,
	)
}
