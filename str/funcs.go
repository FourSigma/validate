package str

import "fmt"

type CheckError struct {
	Name  string
	Err   error
	Value string
}

func (c CheckError) Error() string {
	return fmt.Sprintf("ERROR: %s - %s - %s", c.Name, c.Err, c.Value)
}
