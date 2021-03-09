package doors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	doors "tddkata/100doors"
)

func TestMake_Len(t *testing.T) {
	ds := doors.New()

	require.Equal(t, 100, ds.Len())
}

func TestMake_String(t *testing.T) {
	ds := doors.New()

	require.Equal(t, "#####", ds.String()[0:5])
}

func TestRun(t *testing.T) {
	tt := []struct{
		n int
		expected string
	}{
		{1, "@@@@@"},
		{2, "@#@#@"},
		{3, "@###@"},
		{4, "@##@@"},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("Run - %d", tc.n), func(t *testing.T) {
			ds := doors.New()
			for i := 1; i<=tc.n; i++ {
				ds.Run(i)
			}

			require.Equal(t, tc.expected, ds.String()[0:len(tc.expected)])
		})
	}

}
