package lib

import "fmt"
import . "github.com/FourSigma/validate/lib/misc/err"

type Validator interface {
	IsEmpty() bool
	IsRequired() bool
	Required() Validator
	Name(string) Validator
	String() string
	GetHandlers() []Handler
	Value() interface{}
	Checker
}

type Checker interface {
	Check() error
}

type Handler interface {
	Handle(interface{}) error
}

func DefaultCheck(v Validator) error {
	if v.IsEmpty() && v.IsRequired() == true {
		return fmt.Errorf("%s::Value required.", v.String())
	}

	if v.IsEmpty() {
		return nil
	}

	for _, u := range v.GetHandlers() {
		err := u.Handle(v.Value())
		if err != nil {
			if _, ok := err.(TerminateLoop); ok {
				break
			}
			return err
		}
	}
	return nil
}
