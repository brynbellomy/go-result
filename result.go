package result

import "fmt"

type (
	Result struct {
		value interface{}
		err   error
	}
)

func Success(val interface{}) Result {
	return Result{value: val}
}

func Failure(err error) Result {
	return Result{err: err}
}

func (r Result) Value() interface{} {
	return r.value
}

func (r Result) Error() error {
	return r.err
}

func (r Result) IsError() bool {
	return r.err != nil
}

func (r Result) String() string {
	if r.IsError() {
		return fmt.Sprintf(`Result.Error(%v)`, r.Error())
	} else {
		return fmt.Sprintf(`Result.Success(%T: %v)`, r.Value(), r.Value())
	}
}
