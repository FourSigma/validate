Validate 
========
Validate is a functional, type safe, flexible, and simple approach to string validation and transformation for the Go programming language.  


```go

type Person struct{
  FirstName string
  LastName  string
  EMail     string
}

func(p Person) OK() (err error){
  err = Check(
    String(p.Email).Validate(EMail),
    String(p.FirstName).Validate(Name...), 
    String(p.LastName).Validate(Name...),
  )

}

//Satisfies str.Handler type ---> func(string) error
func EMail(s string) error {
	_, err := mail.ParseAddress(s)
	return err
}

// str.Handler --> func(string) error
var Name = []str.Handler{MaxLen(14), MinLen(2)}

func MaxLen(i int) str.Handler {
	return func(s string) error {
		if len(s) > i {
			return fmt.Errorf("String %s excedes maximum length of %d.", s, i)
		}
		return nil
	}
}

func MinLen(i int) str.Handler {
	return func(s string) error {
		if len(s) < i {
			return fmt.Errorf("String %s doesn't meet minimum length of %d.", s, i)
		}
		return nil
	}
}



```

