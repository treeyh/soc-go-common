package json

import (
	"github.com/treeyh/soc-go-common/core/errors"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/treeyh/soc-go-common/core/consts"
)

type TimeDecoder struct {
}
type TimeEncoder struct {
	precision time.Duration
}

var loc, _ = time.LoadLocation("Local")

func (td *TimeDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	str := iter.ReadString()

	mayBlank, _ := time.Parse(consts.AppTimeFormat, str)
	now, err := time.ParseInLocation(consts.AppTimeFormat, str, loc)

	if err != nil {
		*((*time.Time)(ptr)) = time.Unix(0, 0)
	} else if mayBlank.IsZero() {
		*((*time.Time)(ptr)) = mayBlank
	} else {
		*((*time.Time)(ptr)) = now
	}
}

func (codec *TimeEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	ts := *((*time.Time)(ptr))
	return ts.IsZero()
}

func (codec *TimeEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ts := *((*time.Time)(ptr))
	if !ts.IsZero() {
		timestamp := ts.Unix()
		tm := time.Unix(timestamp, 0)
		format := tm.Format(consts.AppTimeFormat)
		stream.WriteString(format)
	} else {
		mayBlank, _ := time.Parse(consts.AppTimeFormat, consts.BlankString)
		stream.WriteString(mayBlank.Format(consts.AppTimeFormat))
	}

}

var json = jsoniter.ConfigFastest

func init() {
	jsoniter.RegisterTypeEncoder("time.Time", &TimeEncoder{})
	jsoniter.RegisterTypeDecoder("time.Time", &TimeDecoder{})
}

func ToJson(obj interface{}) (string, errors.AppError) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return "", errors.NewAppErrorByExistError(errors.JsonEncodeFail, err)
	}
	return string(bs), nil
}

func ToJsonIgnoreError(obj interface{}) string {
	if obj == nil {
		return ""
	}
	bs, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bs)
}

func FromJson(jsonStr string, obj interface{}) errors.AppError {
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		return errors.NewAppErrorByExistError(errors.JsonDecodeFail, err)
	}
	return nil
}
