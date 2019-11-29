package copyer

import (
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"github.com/treeyh/soc-go-common/core/utils/json"
)

func Copy(src interface{}, source interface{}) errors.SystemError {
	jsonStr, err := json.ToJson(src)
	if err != nil {
		logger.Logger().Error(err)
		return errors.NewSystemError(errors.ObjectCopyFail)
	}
	err = json.FromJson(jsonStr, source)
	if err != nil {
		logger.Logger().Error(err)
		return errors.NewSystemError(errors.ObjectCopyFail)
	}
	return nil
}
