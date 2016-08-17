package validate

type checker interface {
	Check() error
}

type Validator interface {
	OK() error
}

type transformer interface {
	Transform() error
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
