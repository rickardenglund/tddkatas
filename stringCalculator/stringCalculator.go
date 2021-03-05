package stringCalculator

import (
	"fmt"
	"regexp"
	"strconv"
)

type ErrNegativeNotAllowedType struct {
	msg             string
	NegativeNumbers []int
}

func (e ErrNegativeNotAllowedType)Error() string {
	return e.msg
}

// Kata described at
// https://osherove.com/tdd-kata-1
func Add(s string) (int, error) {
	delimiters := `[,\n]`
	reHeader := regexp.MustCompile(`^//(.+)\n`)

	r := reHeader.FindStringSubmatch(s)
	if len(r) > 0 {
		delimiters = regexp.QuoteMeta(r[1])
		s = s[len(r[0]):]
	}

	reDelimiter := regexp.MustCompile(fmt.Sprintf(`%s`, delimiters))
	strs := reDelimiter.Split(s, -1)

	total, negatives := sum(strs)

	if len(negatives) > 0 {
		return 0, ErrNegativeNotAllowedType{"negatives not allowed", negatives}
	}

	return total, nil
}

func sum(strs []string) (int, []int) {
	sum := 0
	var negatives []int
	for _, s := range strs {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, nil
		}

		if i > 1000 {
			continue
		}

		if i < 0 {
			negatives = append(negatives, i)
		}

		sum += i
	}
	return sum, negatives
}
