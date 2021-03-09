package doors

import "strings"

type Doors struct {
	states  []int
	nstates int
}

func (ds *Doors) Len() int {
	return len(ds.states)
}

func (ds *Doors) String() string {
	sb := strings.Builder{}
	for _, b := range ds.states {
		switch b {
		case 0:
			sb.WriteString("#")
		case 1:
			sb.WriteString("@")
		case 2:
			sb.WriteString("H")
		}
	}

	return sb.String()
}

func (ds *Doors) Run(step int) {
	i := -1 + step
	for ; i < len(ds.states); i += step {
		ds.states[i] = (ds.states[i] + 1) % ds.nstates
	}
}

func New(states int) Doors {
	return Doors{states: make([]int, 100), nstates: states}
}
