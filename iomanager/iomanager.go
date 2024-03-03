package iomanager

type IO interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}
