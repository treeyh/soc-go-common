package copyer

import (
	"context"

	// https://github.com/jinzhu/copier
	"github.com/jinzhu/copier"
	"github.com/treeyh/soc-go-common/core/errors"
	"github.com/treeyh/soc-go-common/core/logger"
)

func Copy(ctx context.Context, source interface{}, target interface{}) errors.AppError {
	if source == nil {
		return nil
	}

	err := copier.Copy(target, source)
	if err != nil {
		logger.Logger().ErrorCtx(ctx, err)
		return errors.NewAppErrorByExistError(errors.ObjectCopyFail, err)
	}
	// jsonStr, err := json.ToJson(src)
	// if err != nil {
	// 	logger.Logger().ErrorCtx(ctx, err)
	// 	return errors.NewAppErrorByExistError(errors.ObjectCopyFail, err)
	// }
	// err = json.FromJson(jsonStr, target)
	// if err != nil {
	// 	logger.Logger().ErrorCtx(ctx, err)
	// 	return errors.NewAppErrorByExistError(errors.ObjectCopyFail, err)
	// }
	return nil
}

func CopyList(ctx context.Context, source interface{}, target interface{}) errors.AppError {
	if source == nil {
		return nil
	}

	err := copier.Copy(target, source)
	if err != nil {
		logger.Logger().ErrorCtx(ctx, err)
		return errors.NewAppErrorByExistError(errors.ObjectCopyFail, err)
	}
	return nil
}
