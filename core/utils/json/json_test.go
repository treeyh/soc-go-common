package json

import (
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/types"
	"testing"
	"time"
)

type TA struct {
	D time.Time
	N string
	I types.Int64
}

func TestFromJson(t *testing.T) {
	str := "{\"d\":\"2019-01-01 11:11:1000\",\"n\":\"1\"}"
	ta := &TA{}

	FromJson(str, ta)

	t.Log(ta.D)
	t.Log(ToJson(ta))
}

func TestToJson(t *testing.T) {
	ta := TA{
		D: time.Now(),
		N: "abc",
		I: 3,
	}
	str, err := ToJson(ta)
	t.Log(err)
	t.Log(str)

	t.Log(consts.BlankTimeObject)
	t.Log(types.Time(consts.BlankTimeObject))

	ta2 := &TA{}

	FromJson(str, ta2)

	t.Log(ta2.I)
}
