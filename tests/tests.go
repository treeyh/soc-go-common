package tests

import (
	"context"
)

func TestStartUp(testFunc func(), initFunc func()) func() {

	return func() {
		//加载测试配置
		if initFunc != nil {
			initFunc()
		}

		//丢进来的方法立刻执行
		testFunc()
	}
}

// GetNewContext 获取一个新的ctx
func GetNewContext() context.Context {
	ctx := context.Background()
	return ctx
}
