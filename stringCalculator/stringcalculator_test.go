package stringCalculator_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"tddkata/stringCalculator"
)

func TestCalculator_AddEmpty(t *testing.T) {
	i, _ := stringCalculator.Add("")
	require.Equal(t, 0, i)
}

func TestCalculator_AddOneNumber(t *testing.T) {
	i, _ := stringCalculator.Add("9")
	require.Equal(t, 9, i)
}

func TestCalculator_AddTwoNumber(t *testing.T) {
	i, _ := stringCalculator.Add("9,2")
	require.Equal(t, 11, i)
}

func TestCalculator_AddSomeNumber(t *testing.T) {
	i, _ := stringCalculator.Add("9,2,10,4")
	require.Equal(t, 25, i)
}

func TestCalculator_AddWithNewline(t *testing.T) {
	i, _ := stringCalculator.Add("1,2\n3")
	require.Equal(t, 6, i)
}

func TestCalculator_AddCustomDelimiter(t *testing.T) {
	i, _ := stringCalculator.Add("//;\n1;2")
	require.Equal(t, 3, i)
}

func TestCalculator_AddLongDelimiters(t *testing.T) {
	i, _ := stringCalculator.Add("//***\n1***2***3")
	require.Equal(t, 6, i)
}

func TestCalculator_AddNegativeError(t *testing.T) {
	_, err := stringCalculator.Add("//;\n1;-2")
	require.Error(t, err)
}

func TestCalculator_AddNegativesError(t *testing.T) {
	_, err := stringCalculator.Add("//;\n1;-2;-5")
	e := stringCalculator.ErrNegativeNotAllowedType{}
	require.True(t, errors.As(err, &e))

	require.Equal(t, []int{-2,-5}, e.NegativeNumbers)
	require.Equal(t, "negatives not allowed", e.Error())
}

func TestCalculator_AddIgnoreLarge(t *testing.T) {
	i, _ := stringCalculator.Add("1002,2")
	require.Equal(t, 2, i)
}

