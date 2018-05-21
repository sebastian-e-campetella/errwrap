package errwrap

import (
	"fmt"
)

type Any interface{}

type ErrorWrapper struct {
	Error  error
	Result Any
}

type ErrorCatch interface {
	catch() Any
}

func (ew ErrorWrapper) Catch(any interface{}, err error) Any {
	ew.Error = err
	ew.Result = any

	if ew.Error != nil {
		fmt.Println(err)
	}

	return ew.Result
}

func (ew ErrorWrapper) CatchPanic(any interface{}, err error) Any {
	ew.Error = err
	ew.Result = any

	if ew.Error != nil {
		panic(ew.Error)
	}
	return ew.Result
}

func (ew ErrorWrapper) CatchWrapper(e ErrorWrapper, f Any) Any {
	ew = e

	if ew.Error != nil {
		return f
	}
	return ew.Result
}
