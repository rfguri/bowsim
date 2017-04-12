package bowsim

import (
	"testing"
)

var cases = []struct {
	in  []string
	out float64
}{
	{[]string{"Jane loves me more than Julie loves me", "Jane loves me more than Julie loves me"}, 1.0},
	{[]string{"Jane loves me more than Julie loves me", "Julie likes me more than Jake loves me"}, 0.8333333333333334},
	{[]string{"Jane loves me", "Julie likes me more than Jake loves me"}, 0.6666666666666666},
	{[]string{"Jane loves me more than Julie loves me", "This can't be useful at all in this case"}, 0.0},
}

func TestGet(t *testing.T) {

	for i, c := range cases {
		score := Get(c.in[0], c.in[1])
		if c.out != score {
			t.Errorf("Case #%v: Expected: %v, got: %v", i, c.out, score)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			Get(c.in[0], c.in[1])
		}
	}
}
