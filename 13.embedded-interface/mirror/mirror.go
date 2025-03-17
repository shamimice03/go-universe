package mirror

import "fmt"

type Mirror struct {
	text string
}

func (t *Mirror) Display() {
	fmt.Println(t.text)
}

func (t *Mirror) Tester() {
	fmt.Println(t)
}

func (t *Mirror) Save() error {
	fmt.Println(t)
	return nil
}

func CreateNew(mirrorText string) *Mirror {
	return &Mirror{
		text: mirrorText,
	}
}
