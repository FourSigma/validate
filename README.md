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
    String(p.FirstName).Validate(Name...), 
    String(p.LastName).Validate(Name...),
    String(p.Email).Validate(EMail...),
  )

}

```

