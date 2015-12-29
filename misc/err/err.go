package err

import "fmt"

type TerminateLoop struct {
	Name string
}

func (t TerminateLoop) Error() string {
	return fmt.Sprintf("Loop terminated by %s", t.Name)
}
