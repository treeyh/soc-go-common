package tests

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
