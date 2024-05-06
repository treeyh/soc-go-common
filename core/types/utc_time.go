package types

import (
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/logger"
	"io"
	"strconv"
	"strings"
	"time"
)

type UtcTime time.Time

func NowUtcTime() UtcTime {
	return UtcTime(time.Now())
}

func UtcTime0() UtcTime {
	return UtcTime(time.Unix(0, 0))
}

func (t *UtcTime) IsNotNull() bool {
	return time.Time(*t).After(consts.BlankTimeObject)
}

func (t *UtcTime) IsNull() bool {
	return !t.IsNotNull()
}

func (t UtcTime) ToTime() time.Time {
	return time.Time(t)
}

// InByOffset 时间时区转换
func (t UtcTime) InByOffset(offset int) UtcTime {
	if zone, ok := globalTimeZoneMap[offset]; ok {
		return UtcTime(t.ToTime().In(zone))
	} else {
		return UtcTime(t.ToTime().In(time.FixedZone(strconv.Itoa(offset), offset)))
	}
}

func (t *UtcTime) UnmarshalJSON(data []byte) (err error) {
	//fmt.Println("UnmarshalJSON:", string(data))
	date := strings.ReplaceAll(string(data), "\"", "")
	if date == "" {
		return
	}

	now, err := time.Parse(consts.AppSystemTimeFormat8, date)
	if err != nil {
		now, err = time.ParseInLocation(consts.AppSystemTimeFormat, date, time.Local)
		if err != nil {
			now, err = time.Parse(consts.AppTimeFormat, date)
			if err != nil {
				now, err = time.Parse(consts.AppTimeFormatMillisecond, date)
				if err != nil {
					*t = UtcTime(consts.BlankTimeObject)
					return nil
				}
			}
		}
	}
	*t = UtcTime(now)
	return nil
}

func (t UtcTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(consts.AppSystemTimeFormat8)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, consts.AppSystemTimeFormat8)
	b = append(b, '"')
	return b, nil
}

// UnmarshalGQL implements the graphql.Marshaler interface
func (t *UtcTime) UnmarshalGQL(v string) error {

	date := strings.ReplaceAll(v, "\"", "")
	if date == "" {
		return nil
	}

	now, err := time.Parse(consts.AppSystemTimeFormat8, date)
	if err != nil {
		now, err = time.ParseInLocation(consts.AppSystemTimeFormat, date, time.Local)
		if err != nil {
			now, err = time.Parse(consts.AppTimeFormat, date)
			if err != nil {
				now, err = time.Parse(consts.AppTimeFormatMillisecond, date)
				if err != nil {
					logger.Logger().Error(err)
				}
			}
		}
	}
	*t = UtcTime(now)
	//fmt.Println("UnmarshalGQL:", t)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (t UtcTime) MarshalGQL(w io.Writer) {
	//fmt.Println("MarshalGQL:", t)
	b := make([]byte, 0, len(consts.AppSystemTimeFormat8)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, consts.AppSystemTimeFormat8)
	b = append(b, '"')
	//fmt.Println("MarshalGQL:", string(b))
	w.Write(b)
}

func (t UtcTime) String() string {
	return time.Time(t).Format(consts.AppSystemTimeFormat8)
}

// ConvertUtcTimes 转换数组列表
func ConvertUtcTimes(times []UtcTime) []time.Time {
	if times == nil {
		return nil
	}

	ts := make([]time.Time, len(times))
	for k, v := range times {
		ts[k] = time.Time(v)
	}
	return ts
}
