package validate

import "reflect"

type checker interface {
	Check() error
}

type Validator interface {
	OK() error
}

type transformer interface {
	Transform() error
}

type okcheck struct {
	vs []Validator
}

func (t okcheck) Check() error {

	for _, v := range t.vs {
		if val := reflect.ValueOf(v); val.IsNil() {
			continue
		}

		err := v.OK()
		if err != nil {
			return err
		}

	}

	return nil
}

func OKCheck(c ...Validator) *okcheck {
	return &okcheck{
		vs: c,
	}
}

func Check(c ...checker) error {
	for _, v := range c {
		err := v.Check()
		if err != nil {
			return err
		}
	}
	return nil
}

func Transform(c ...transformer) error {
	for _, v := range c {
		err := v.Transform()
		if err != nil {
			return err
		}
	}
	return nil
}
