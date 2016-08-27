package lib

import "fmt"
import . "github.com/FourSigma/validate/lib/misc/err"

type Checker interface {
	Check() error
}

type Helper interface {
	Checker
	value() interface{}
	getString() string
	Name(string) Validator
	isRequired() bool
	Required() Validator
	GetHandlers() []Handler
	SetHandlers(...Handler)
}

type Validator interface {
	IsEmpty() bool
	Helper
}

type Handler interface {
	Handle(interface{}) error
}

func DefaultCheck(v Validator) error {
	if v.IsEmpty() && v.isRequired() == true {
		return fmt.Errorf("%s::Value required.", v.getString())
	}

	if v.IsEmpty() {
		return nil
	}

	for _, u := range v.GetHandlers() {
		err := u.Handle(v.value())
		if err != nil {
			if _, ok := err.(TerminateLoop); ok {
				break
			}
			return err
		}
	}
	return nil
}

func NewDefaultHelper(val interface{}, typname string, handlers ...Handler) Helper {
	return &helper{
		val:      val,
		typname:  typname,
		handlers: handlers,
		checker:  DefaultCheck,
	}
}

//Generic implementation of the Helper interface
type helper struct {
	val      interface{}
	required bool
	meta     string
	typname  string

	handlers []Handler
	checker  func(v Validator) error
}

func (s *helper) Check() error {
	return s.checker(s)
}

func (s *helper) value() interface{} {
	return s.val
}

func (s *helper) Name(name string) Validator {
	s.meta = name
	return s
}
func (s *helper) getString() string {
	if s.meta == "" {
		s.meta = "NO_NAME_GIVEN -- Use Name()"
	}
	return fmt.Sprintf("%s::%s", s.typname, s.meta)
}

func (s *helper) Required() Validator {
	s.required = true
	return s
}
func (s *helper) isRequired() bool {
	return s.required
}

func (s *helper) GetHandlers() []Handler {
	return s.handlers
}

func (s *helper) SetHandlers(hn ...Handler) {
	s.handlers = hn
}

//Dummy method to staisfy interface
func (s *helper) IsEmpty() bool {
	return false
}
