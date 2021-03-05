package fizzbuzz

import (
	"fmt"
)

type Printer interface {
	Println(s string)
}

func Print(p Printer) {
	for i := 1; i<= 100; i++ {
		p.Println(getText(i))
	}
}

func getText(i int) string {
	text := ""

	if i%3 == 0 {
		text = "fizz"
	}

	if i%5 == 0 {
		text += "buzz"
	}

	if text == "" {
		return fmt.Sprintf("%d", i)
	}

	return text
}


type StdOutPrinter struct {}

func (p *StdOutPrinter) Println(s string) {
	fmt.Printf("%s\n", s)
}

