// Package bowsim provides a way to perform string
// similarity score calculation using bag of words model.
package bowsim

import (
	"github.com/juju/utils/set"
	"strings"
)

func bow(s string) set.Strings {
	result := set.NewStrings()
	for _, word := range strings.Fields(s) {
		result.Add(word)
	}
	return result
}

// Get returns a bag of words string similarity score.
func Get(a, b string) float64 {
	return float64(bow(a).Intersection(bow(b)).Size()) / float64(bow(a).Size())
}
