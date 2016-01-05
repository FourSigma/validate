Validate 
========
Validate is a **functional**, **type safe**, **flexible**, and **simple** approach to string validation and transformation for the Go programming language.   Unlike
```go

import (
. "github.com/FourSigma/validate"	
"github.com/FourSigma/validate/str"
)

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
```
#### String Handler
Write functions that either satisfy or return (closures)  ```type Str.Handler func(string) error```.  These function can carry out custom validation and return an error.  
``` go
//Write functions that are of type str.Handler ---> func(string) error
func EMail(s string) error {
	_, err := mail.ParseAddress(s)
	return err
}

//Use Go function closures that returns a str.Handler function
//for more flexibiltiy.
func MaxLen(i int) str.Handler {
	return func(s string) error {
		if len(s) > i {
			return fmt.Errorf("String %s excedes maximum length of %d.", s, i)
		}
		return nil
	}
}

```
#### Multiple string validations
Aggreating handlers into slice can make complex validations easier, composable, and maintainable. Functions are evaluated in the order.   
```go
// You can create []str.Handler for multiple validations.
var Name = []str.Handler{MaxLen(14), MinLen(2)}


func MinLen(i int) str.Handler {
	return func(s string) error {
		if len(s) < i {
			return fmt.Errorf("String %s doesn't meet minimum length of %d.", s, i)
		}
		return nil
	}
}



```

