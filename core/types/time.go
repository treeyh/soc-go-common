package types

import (
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/logger"
	"io"
	"strconv"
	"strings"
	"time"
)

var (
	// globalTimeZoneMap 全球时区字典，key为offset值
	globalTimeZoneMap = map[int]*time.Location{
		-39600: time.FixedZone("-1100", 3600*(-11)),
		-36000: time.FixedZone("-1000", 3600*(-10)),
		-34200: time.FixedZone("-0930", 3600*(-9.5)),
		-32400: time.FixedZone("-0900", 3600*(-9)),
		-28800: time.FixedZone("-0800", 3600*(-8)),
		-25200: time.FixedZone("-0700", 3600*(-7)),
		-21600: time.FixedZone("-0600", 3600*(-6)),
		-18000: time.FixedZone("-0500", 3600*(-5)),
		-14400: time.FixedZone("-0400", 3600*(-4)),
		-12600: time.FixedZone("-0330", 3600*(-3.5)),
		-10800: time.FixedZone("-0300", 3600*(-3)),
		-7200:  time.FixedZone("-0200", 3600*(-2)),
		-3600:  time.FixedZone("-0100", 3600*(-1)),
		0:      time.FixedZone("+0000", 3600*0),
		3600:   time.FixedZone("+0100", 3600*1),
		7200:   time.FixedZone("+0200", 3600*2),
		10800:  time.FixedZone("+0300", 3600*3),
		12600:  time.FixedZone("+0330", 3600*3.5),
		14400:  time.FixedZone("+0400", 3600*4),
		16200:  time.FixedZone("+0430", 3600*4.5),
		18000:  time.FixedZone("+0500", 3600*5),
		19800:  time.FixedZone("+0530", 3600*5.5),
		20700:  time.FixedZone("+0545", 3600*5.75),
		21600:  time.FixedZone("+0600", 3600*6),
		23400:  time.FixedZone("+0630", 3600*6.5),
		25200:  time.FixedZone("+0700", 3600*7),
		28800:  time.FixedZone("+0800", 3600*8),
		31500:  time.FixedZone("+0845", 3600*8.75),
		32400:  time.FixedZone("+0900", 3600*9),
		34200:  time.FixedZone("+0930", 3600*9.5),
		36000:  time.FixedZone("+1000", 3600*10),
		37800:  time.FixedZone("+1030", 3600*10.5),
		39600:  time.FixedZone("+1100", 3600*11),
		43200:  time.FixedZone("+1200", 3600*12),
		45900:  time.FixedZone("+1245", 3600*12.75),
		46800:  time.FixedZone("+1300", 3600*13),
		50400:  time.FixedZone("+1400", 3600*14),
	}
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

// InByOffset 时间时区转换
func InByOffset(t time.Time, offset int) time.Time {
	if zone, ok := globalTimeZoneMap[offset]; ok {
		return t.In(zone)
	} else {
		return t.In(time.FixedZone(strconv.Itoa(offset), offset))
	}
}

func (t *Time) IsNotNull() bool {
	return time.Time(*t).After(consts.BlankTimeObject)
}

func (t *Time) IsNull() bool {
	return !t.IsNotNull()
}

func (t Time) ToTime() time.Time {
	return time.Time(t)
}

// InByOffset 时间时区转换
func (t Time) InByOffset(offset int) Time {
	if zone, ok := globalTimeZoneMap[offset]; ok {
		return Time(t.ToTime().In(zone))
	} else {
		return Time(t.ToTime().In(time.FixedZone(strconv.Itoa(offset), offset)))
	}
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	//fmt.Println("UnmarshalJSON:", string(data))
	date := strings.ReplaceAll(string(data), "\"", "")
	if date == "" {
		return
	}
	now, err := time.ParseInLocation(consts.AppTimeFormat, date, time.Local)
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
func ConvertTimes(times []Time) []time.Time {
	if times == nil {
		return nil
	}

	ts := make([]time.Time, len(times))
	for k, v := range times {
		ts[k] = time.Time(v)
	}
	return ts
}
