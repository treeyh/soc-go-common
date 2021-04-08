package pinyin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertCode(t *testing.T) {


	str := "这是One IssueOfFirst任务"
	assert.Equal(t, ConvertCode(str), "ZSOIOFRW")

	str = "One这是OfIssue First1123456"

	assert.Equal(t, ConvertCode(str), "OZSOIF1123456")

	str = "OneOneTwoTwoThirdThirdFourFourFive"
	assert.Equal(t, ConvertCode(str), "OOTTTTFFF")
	assert.NotEqual(t, ConvertCode(str), "OOTTTTFF")

}
