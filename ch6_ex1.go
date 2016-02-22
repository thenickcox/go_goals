package main

import "fmt"

type font struct {
	family string
	size   int
}

func NewFont(family string, size int) *font {
	return &font{family, size}
}

func (font *font) SetFamily(family string) {
	font.family = family
}

func (font *font) SetSize(size int) {
	if size < 5 || size > 144 {
		return fmt.Printf("invalid size %d", size)
	}
	font.size = size
}

func (font *font) GetFamily() string {
	return font.family
}

func (font *font) GetSize() int {
	return font.size
}

func (font *font) String() string {
	return fmt.Sprintf("{font-family: \"%v\"; font-size: %dpt)}", font.GetFamily(), font.GetSize())
}

func main() {
	titleFont := NewFont("serif", 11)
	titleFont.SetFamily("Helvetica")
	titleFont.SetSize(20)
	fmt.Println(titleFont)
}
