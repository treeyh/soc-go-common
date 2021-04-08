package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceUniqueString(t *testing.T) {

	ss := []string{"aaa", "bbb", "aacc", "bbb", "ccc"}
	ss = SliceUniqueString(ss)
	assert.Equal(t, ss, []string{"aaa", "bbb", "aacc", "ccc"}, "SliceUniqueString error.")

	ii := []int64{1, 2, 3, 1}
	ii = SliceUniqueInt64(ii)
	assert.Equal(t, ii, []int64{1, 2, 3}, "SliceUniqueString error.")

	r := ContainString("aaaabbbb", ss)

	assert.True(t, r, "ContainString error.")

}
