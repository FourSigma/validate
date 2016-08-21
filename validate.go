package validate

import "fmt"

type Checker interface {
	Check() error
}

type Transformer interface {
	Transform() error
}

func Check(c ...Checker) error {
	for _, v := range c {
		err := v.Check()
		if err != nil {
			return err
		}
	}
	return nil
}

func Transform(c ...Transformer) error {
	for _, v := range c {
		err := v.Transform()
		if err != nil {
			return err
		}
	}
	return nil
}

type Validator interface {
	Empty() bool
	Check() error
	Add(...Handler) Validator
	Finally(...Handler) Validator
	Name(string) Validator
	Required() Validator
}

type Handler interface {
	Handle(interface{}) error
}

func And(sh ...Handler) Handler {
	return and(sh)
}

type and []Handler

func (a and) Handle(i interface{}) error {
	for _, v := range a {
		err := v.Handle(i)
		if err != nil {
			return fmt.Errorf("AND error for type %v", i)
		}
	}

	return nil
}

func Or(sh ...Handler) Handler {
	return or(sh)
}

type or []Handler

func (a or) Handle(i interface{}) error {
	for _, v := range a {
		err := v.Handle(i)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("OR error for type %v", i)
}
