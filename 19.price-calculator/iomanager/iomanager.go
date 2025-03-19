package iomanager

type IOManager interface {
	ReadFileContents() ([]string, error)
	WriteJSON(data any) error
}
