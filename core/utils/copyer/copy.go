package copyer

import (
	"context"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
	"github.com/treeyh/soc-go-common/core/utils/json"
	"reflect"
)

func Copy(ctx context.Context, src interface{}, target interface{}) errors.AppError {
	if src == nil {
		return nil
	}

	jsonStr, err := json.ToJson(src)
	if err != nil {
		logger.Logger().ErrorCtx(ctx, err)
		return errors.NewAppError(errors.ObjectCopyFail)
	}
	err = json.FromJson(jsonStr, target)
	if err != nil {
		logger.Logger().ErrorCtx(ctx, err)
		return errors.NewAppError(errors.ObjectCopyFail)
	}
	return nil
}

func CopyList(ctx context.Context, src interface{}, target interface{}) errors.AppError {
	if reflect.TypeOf(src).Kind() == reflect.Slice || reflect.TypeOf(src).Kind() == reflect.Array {
		if reflect.ValueOf(src).Len() <= 0 {
			return nil
		}
	}
	jsonStr, err := json.ToJson(src)
	if err != nil {
		logger.Logger().ErrorCtx(ctx, err)
		return errors.NewAppError(errors.ObjectCopyFail)
	}
	print(jsonStr)
	err = json.FromJson(jsonStr, target)
	if err != nil {
		logger.Logger().ErrorCtx(ctx, err)
		return errors.NewAppError(errors.ObjectCopyFail)
	}
	return nil
}
