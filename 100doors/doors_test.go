package doors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	doors "tddkata/100doors"
)

func TestMake_Len(t *testing.T) {
	ds := doors.New(2)

	require.Equal(t, 100, ds.Len())
}

func TestMake_String(t *testing.T) {
	ds := doors.New(2)

	require.Equal(t, "#####", ds.String()[0:5])
}

func TestRun(t *testing.T) {
	tt := []struct{
		n int
		states int
		expected string
	}{
		{1, 2, "@@@@@"},
		{2, 2, "@#@#@"},
		{3, 2, "@###@"},
		{4, 2, "@##@@"},

		{1, 3, "@@@@@"},
		{2, 3, "@H@H@"},
		{3, 3, "@HHH@"},
		{4, 3, "@HH#@"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Run - %d", tc.n), func(t *testing.T) {
			ds := doors.New(tc.states)
			for i := 1; i<=tc.n; i++ {
				ds.Run(i)
			}

			require.Equal(t, tc.expected, ds.String()[0:len(tc.expected)])
		})
	}

}

