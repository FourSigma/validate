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

func NewTransStr(s []byte, sh []TransHandler) *TransStr {
	t := make([]TransHandler, len(sh))
	copy(t, sh)

	tb := make([]byte, len(s))
	copy(tb, s)
	return &TransStr{
		b:  s,
		cp: tb,
		h:  t,
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

type TransStr struct {
	b  []byte
	cp []byte
	h  []TransHandler
}

func (s *TransStr) Transform() (err error) {
	for _, v := range s.h {
		s.b, err = v(s.b)
		if err != nil {
			s.b = s.cp
			if _, ok := err.(TerminateLoop); ok {
				fmt.Println("Loop Terminated")
				break
			}
			return err
		}
	}
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
