package fizzbuzz_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"tddkata/fizzbuzz"
)

type printer struct {
	lines []string
}

func (p *printer) Println(s string) {
	p.lines = append(p.lines, s)
}

func TestPrint_noCheck(t *testing.T) {
	printer := printer{}
	fizzbuzz.Print(&printer)
}

func TestPrint_lines(t *testing.T) {
	printer := printer{}
	fizzbuzz.Print(&printer)

	assert.Equal(t, "1", printer.lines[0])
	assert.Len(t, printer.lines, 100)

	assert.Equal(t, "4", printer.lines[3])
	assert.Equal(t, "fizz", printer.lines[2])
	assert.Equal(t, "buzz", printer.lines[4])
	assert.Equal(t, "fizzbuzz", printer.lines[14])
}

func TestStdOutPrinter_Println(t *testing.T) {
	io.ReadAll(os.Stdout)
	r,w, err := os.Pipe()
	require.NoError(t, err)

	os.Stdout = w


	p := fizzbuzz.StdOutPrinter{}

	p.Println("hej")

	require.NoError(t, w.Close())
	s, err := io.ReadAll(r)

	require.Equal(t, "hej\n", string(s))
}
