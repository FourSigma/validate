Validate 
========
Validate is a **functional**, **type safe**, **flexible**, and **simple** approach to string validation and transformation for the Go programming language.   Unlike other validation libraries that rely on struct field tags, Validate doesn't use reflection.  I have kept the small and simple.
```go

import (
. "github.com/FourSigma/validate"	
"github.com/FourSigma/validate/str"
)

type Person struct{
  FirstName  string
  LastName   string
  MiddleName string
  EMail      string
  State	     string
  ID	     string
}

func(p Person) OK() (err error){
 // API Overview
 
  err = Check(
    String(p.Email).Validate(EMail),         //Validate takes variadic paramerters Validate(fn ...str.Handler)
    String(p.StateName).Validate(MaxLen(2), MinLen(2)) 
    String(p.FirstName).Validate(Name...),  //More complex validations can be aggregated into a slice
    String(p.LastName).Validate(Name...).Add(RunFirst...),       //  Add  - Runs before Name
    String(p.MiddleName).Validate(Name...).Finally(RunLast...), //Finally - Runs after Name
    String(p.ID).Validate(Id...).Required(), //ID cannot be blank
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
Aggreating handlers into slice can make complex validations easier, composable, and maintainable. Functions are evaluated in the order (0...len(slice)-1).   
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

