package doors

import "strings"

type Doors struct {
	states []bool
}

func (ds *Doors) Len() int {
	return len(ds.states)
}

func (ds *Doors) String() string {
	sb := strings.Builder{}
	for _, b := range ds.states {
		if b {
			sb.WriteString("@")
		} else {
			sb.WriteString("#")
		}

	}
	return sb.String()
}

func (ds *Doors) Run(step int) {
	i := -1 + step
	for ; i < len(ds.states); i += step {
		ds.states[i] = !ds.states[i]
	}
}

func New() Doors {
	return Doors{states: make([]bool, 100)}
}
