package validate

type checker interface {
	Check() error
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

//func Terminate(s string) error {
//	return TerminateLoop{Name: "Terminate"}
//}
