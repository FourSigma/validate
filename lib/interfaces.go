package lib

type Validator interface {
	Empty() bool
	Check() error
	Add(...Handler) Validator
	Finally(...Handler) Validator
	Name(string) Validator
	Required() Validator
}

type Handler interface {
	Handle(interface{}) error
}
