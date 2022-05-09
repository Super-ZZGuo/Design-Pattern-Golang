package singleton

import "sync"

// Singleton 单例模式接口
// 通过该接口可以避免 GetInstance 返回一个包私有类型的指针
type Singleton interface {
	foo() string
}

type singleton struct{}

func (*singleton) foo() string {
	return "do some thing"
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 单例模式获取对象
func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
