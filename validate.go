package validate

type Checker interface {
	Check() error
}

//type TerminateLoop struct {
//	Name string
//}

//func (t TerminateLoop) Error() string {
//	return fmt.Sprintf("Loop terminated by %s", t.Name)
//}

func Check(c ...Checker) error {
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
