package examples

import (
	"context"

	"github.com/FourSigma/validate"
)

type Person struct {
	FirstName    string
	LastName     string
	Age          int
	BirthYear    int
	EmailAddress string
}

func (p *Person) OK(ctx context.Context) (err error) {

	err = validate.Check(
		ctx,
		validate.String(p.FirstName).Validate(),
		validate.String(p.LastName).Validate(),
		validate.String(p.EmailAddress).Validate(),
		validate.Int(p.Age).Validate(),
		validate.String(p.EmailAddress).Validate(),
	)

	return err
}
