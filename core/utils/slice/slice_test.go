package slice

import (
	"fmt"
	"testing"
)

func TestSliceUniqueString(t *testing.T) {

	ss := []string{"aaa", "bbb", "aacc"}
	ss = SliceUniqueString(ss)
	t.Log(ss)

	ii := []int64{1, 2, 3, 1}
	ii = SliceUniqueInt64(ii)
	t.Log(ii)

	r := ContainString("aaaaa", ss)

	fmt.Println(r)

	r = ContainString("aaaabbbb", ss)

	fmt.Println(r)

}
