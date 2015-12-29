package str

import "fmt"

func NewStr(s string, sh []Handler) *Str {

	return &Str{
		s: s,
		h: sh,
	}

}

type Handler func(string) error

type Str struct {
	s string
	h []Handler
}

type TerminateLoop struct {
	Name string
}

func (t TerminateLoop) Error() string {
	return fmt.Sprintf("Loop terminated by %s", t.Name)
}
func (s *Str) Check() error {
	for i, v := range s.h {
		err := v(s.s)
		fmt.Println("Error:", err, i)
		if err != nil {
			if _, ok := err.(TerminateLoop); ok {
				fmt.Println("Terminate Checkout")
				break
			}
			return err
		}
	}
	return nil
}

func (s *Str) Add(b ...Handler) *Str {
	s.h = append(b, s.h...)
	return s
}

func (s *Str) Finally(a ...Handler) *Str {
	s.h = append(s.h, a...)
	return s
}
