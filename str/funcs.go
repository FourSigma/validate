package str

import (
	"fmt"
	"regexp"
	"strings"

	e "github.com/FourSigma/validate/misc/err"
)

type CheckError struct {
	Name  string
	Err   error
	Value string
}

func (c CheckError) Error() string {
	return fmt.Sprintf("ERROR: %s - %s - %s", c.Name, c.Err, c.Value)
}

func And(sh ...Handler) Handler {
	return func(s string) error {
		for _, v := range sh {
			err := v(s)
			if err != nil {
				return fmt.Errorf("AND error for string %s", s)
			}
		}
		return nil
	}
}

func Or(sh ...Handler) Handler {
	return func(s string) (err error) {
		for _, v := range sh {
			err = v(s)
			if err == nil {
				return nil
			}
		}
		return fmt.Errorf("OR error for string %s", s)
	}
}

func Required(s string) error {
	if s == "" {
		return fmt.Errorf("Value cannot be blank.")
	}
	return nil
}

func BlankOK(s string) error {
	if s == "" {
		return e.TerminateLoop{}
	}
	return nil
}
func MaxLen(i int) Handler {
	return func(s string) error {
		if len(s) > i {
			return fmt.Errorf("String %s excedes maximum length of %d.", s, i)
		}
		return nil
	}
}

func MinLen(i int) Handler {
	return func(s string) error {
		if len(s) < i {
			return fmt.Errorf("String %s doesn't meet minimum length of %d.", s, i)
		}
		return nil
	}
}

func ContainsAll(substrs ...string) Handler {
	return func(s string) error {
		for _, v := range substrs {
			if ok := strings.Contains(s, v); !ok {
				return fmt.Errorf("SubString %s not found.", v)
			}
		}
		return nil
	}
}

func ContainsAny(substrs ...string) Handler {
	return func(s string) error {
		for _, v := range substrs {
			if ok := strings.Contains(s, v); ok {
				return nil
			}
		}
		return fmt.Errorf("Substrings not found.")
	}

}

func HasSuffix(suffix string) Handler {
	return func(s string) error {
		if ok := strings.HasSuffix(s, suffix); ok {
			return nil
		}
		return fmt.Errorf("No suffix %s found.", suffix)
	}
}

func HasPrefix(prefix string) Handler {
	return func(s string) error {
		if ok := strings.HasPrefix(s, prefix); ok {
			return nil
		}
		return fmt.Errorf("No prefix %s found.", prefix)
	}
}

func Regexp(r string) Handler {

	return func(s string) error {
		m, err := regexp.MatchString(r, s)
		switch {
		case err != nil:
			return err
		case m == false:
			return fmt.Errorf("String %s did not match regexp %s.", s, r)
		default:
			return nil
		}

	}
}
