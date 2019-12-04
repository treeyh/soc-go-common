package copyer

import (
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"github.com/treeyh/soc-go-common/core/utils/json"
)

func Copy(src interface{}, source interface{}) errors.AppError {
	jsonStr, err := json.ToJson(src)
	if err != nil {
		logger.Logger().Error(err)
		return errors.NewAppError(errors.ObjectCopyFail)
	}
	err = json.FromJson(jsonStr, source)
	if err != nil {
		logger.Logger().Error(err)
		return errors.NewAppError(errors.ObjectCopyFail)
	}
	return nil
}
