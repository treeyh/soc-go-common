package excel

import (
	"fmt"
	"github.com/treeyh/soc-go-common/core/utils/slice"
	"testing"
)

func TestGenerateCSVFromXLSXFile(t *testing.T) {
	t.Log(GenerateCSVFromXLSXFile("test.xlsx", "", 0, 0, []int64{1, 2}))
}

func TestGenerateCSVFromXLSXFile2(t *testing.T) {
	a := []int64{1, 2}
	fmt.Println(slice.Contain(a, 1))
}
