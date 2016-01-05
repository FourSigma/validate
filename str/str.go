package str

import "fmt"

import . "github.com/FourSigma/validate/misc/err"

func NewChkStr(s string, sh []Handler) *Str {
	t := make([]Handler, len(sh))
	copy(t, sh)
	return &Str{
		s: s,
		h: t,
	}

}

type Handler func(string) error

type Str struct {
	s string
	h []Handler
}

func (s *Str) Check() error {
	for i, v := range s.h {
		err := v(s.s)
		fmt.Println("Error:", err, i)
		if err != nil {
			if _, ok := err.(TerminateLoop); ok {
				fmt.Println("Loop Terminated")
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

type TransHandler func([]byte) ([]byte, error) //Not cocurrent safe

func NewTransStr(s *string, sh []TransHandler) *TransStr {

	t := make([]TransHandler, len(sh))
	copy(t, sh)

	return &TransStr{
		o: s,
		b: []byte(*s),
		h: t,
	}

}

type TransStr struct {
	o *string
	b []byte
	h []TransHandler
}

func (s *TransStr) Transform() (err error) {
	for _, v := range s.h {
		s.b, err = v(s.b)
		fmt.Println(string(s.b))
		if err != nil {
			if _, ok := err.(TerminateLoop); ok {
				fmt.Println("Loop Terminated")
				break
			}
			return err
		}
	}
	*s.o = string(s.b)
	return nil
}

func (s *TransStr) Add(b ...TransHandler) *TransStr {
	s.h = append(b, s.h...)
	return s
}

func (s *TransStr) Finally(a ...TransHandler) *TransStr {
	s.h = append(s.h, a...)
	return s
}
