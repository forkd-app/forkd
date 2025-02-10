package graph

type Cursor[T any] interface {
	Encode() (string, error)
	Decode(string) error
	Validate(input T) bool
}
