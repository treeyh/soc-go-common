package types

import (
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/logger"
	"io"
	"strings"
	"time"
)

type Time time.Time

func NowTime() Time {
	return Time(time.Now())
}

func Time0() Time {
	return Time(time.Unix(0, 0))
}

func TTime0() time.Time {
	return time.Unix(0, 0)
}

func (t *Time) IsNotNull() bool {
	return time.Time(*t).After(consts.BlankTimeObject)
}

func (t *Time) IsNull() bool {
	return !t.IsNotNull()
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	//fmt.Println("UnmarshalJSON:", string(data))
	date := strings.ReplaceAll(string(data), "\"", "")
	if date == "" {
		return
	}
	now, err := time.Parse(consts.AppTimeFormat, date)
	if err != nil {
		now, err = time.Parse(consts.AppSystemTimeFormat, date)
		if err != nil {
			now, err = time.Parse(consts.AppSystemTimeFormat8, date)
			if err != nil {
				*t = Time(consts.BlankTimeObject)
				return nil
			}
		}
	}
	*t = Time(now)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(consts.AppTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, consts.AppTimeFormat)
	b = append(b, '"')
	return b, nil
}

// UnmarshalGQL implements the graphql.Marshaler interface
func (t *Time) UnmarshalGQL(v string) error {

	date := strings.ReplaceAll(v, "\"", "")
	if date == "" {
		return nil
	}

	now, err := time.Parse(consts.AppTimeFormat, date)
	if err != nil {
		now, err = time.Parse(consts.AppSystemTimeFormat, date)
		if err != nil {
			now, err = time.Parse(consts.AppSystemTimeFormat8, date)
			if err != nil {
				logger.Logger().Error(err)
			}
		}
	}
	*t = Time(now)
	//fmt.Println("UnmarshalGQL:", t)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (t Time) MarshalGQL(w io.Writer) {
	//fmt.Println("MarshalGQL:", t)
	b := make([]byte, 0, len(consts.AppTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, consts.AppTimeFormat)
	b = append(b, '"')
	//fmt.Println("MarshalGQL:", string(b))
	w.Write(b)
}

func (t Time) String() string {
	return time.Time(t).Format(consts.AppTimeFormat)
}

// ConvertTimes 转换数组列表
func ConvertTimes(times *[]Time) *[]time.Time {
	if times == nil {
		return nil
	}

	ts := make([]time.Time, len(*times))
	for k, v := range *times {
		ts[k] = time.Time(v)
	}
	return &ts
}
