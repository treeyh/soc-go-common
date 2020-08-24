package types

import (
	"github.com/treeyh/soc-go-common/core/utils/strs"
	"io"
	"strconv"
	"strings"
)

type Int64 int64

func (t *Int64) toInt64() int64 {
	return int64(*t)
}

func (t *Int64) UnmarshalJSON(data []byte) (err error) {
	num := strings.ReplaceAll(string(data), "\"", "")
	if num == "" {
		*t = Int64(0)
		return
	}
	numInt, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		*t = Int64(0)
		return
	}

	*t = Int64(numInt)
	return
}

func (t Int64) MarshalJSON() ([]byte, error) {

	str := strconv.FormatInt(int64(t), 10)

	bts := strs.Str2Bytes(str)
	b := make([]byte, 0, len(bts)+2)
	b = append(b, '"')
	b = append(b, bts...)
	b = append(b, '"')
	return b, nil
}

// UnmarshalGQL implements the graphql.Marshaler interface
func (t *Int64) UnmarshalGQL(v string) error {
	num := strings.ReplaceAll(v, "\"", "")
	if num == "" {
		return nil
	}
	now, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return nil
	}

	*t = Int64(now)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (t Int64) MarshalGQL(w io.Writer) {
	//fmt.Println("MarshalGQL:", t)
	str := strconv.FormatInt(int64(t), 10)
	bts := strs.Str2Bytes(str)
	b := make([]byte, 0, len(bts)+2)
	b = append(b, '"')
	b = append(b, bts...)
	b = append(b, '"')

	w.Write(b)
}

func (t Int64) String() string {
	return strconv.FormatInt(int64(t), 10)
}
