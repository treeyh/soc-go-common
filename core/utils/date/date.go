package date

import (
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/types"
	"strconv"
	"time"
)

func FormatDate(t time.Time) string {
	return t.Format(consts.AppDateFormat)
}
func FormatDateByTypes(t types.Time) string {
	return time.Time(t).Format(consts.AppDateFormat)
}

func FormatTime(t time.Time) string {
	return t.Format(consts.AppTimeFormat)
}
func FormatTimeByTypes(t types.Time) string {
	return time.Time(t).Format(consts.AppTimeFormat)
}

func ParseTime(str string) (types.Time, errors.AppError) {
	t, err := time.Parse(consts.AppTimeFormat, str)
	if err != nil {
		return types.Time0(), errors.NewAppErrorByExistError(errors.ParseTimeFail, err)
	}
	return types.Time(t), nil
}

func Parse(str string) (time.Time, errors.AppError) {
	t, err := time.Parse(consts.AppTimeFormat, str)
	if err != nil {
		return time.Unix(0, 0), errors.NewAppErrorByExistError(errors.ParseTimeFail, err)
	}
	return t, nil
}

// GetWeekStart 获取某一天当前周的周一0点
func GetWeekStart(d time.Time) time.Time {
	offset := int(time.Monday - d.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

// GetMonthStart 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetMonthStart(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetDateInt(t types.Time) int {
	i, _ := strconv.Atoi(time.Time(t).Format(consts.AppDateFormat2))
	return i
}

func GetDateTimeLong(t types.Time) int64 {
	i, _ := strconv.ParseInt(time.Time(t).Format(consts.AppTimeFormat2), 10, 64)
	return i
}
func GetDateIntByTime(t time.Time) int {
	i, _ := strconv.Atoi(t.Format(consts.AppDateFormat2))
	return i
}

func GetDateTimeLongByTime(t time.Time) int64 {
	i, _ := strconv.ParseInt(t.Format(consts.AppTimeFormat2), 10, 64)
	return i
}
