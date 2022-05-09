package adapter

// Goose 适配类和方法
type Goose interface {
	Quack() string
}

// NewDuck Duck的工厂函数
func NewDuck() Duck {
	return &DuckImp{}
}

// Duck 被适配类和复用的方法
type Duck interface {
	DuckQuack() string
}

// DuckImp 实例
type DuckImp struct{}

// DuckQuack Duck的方法 -> 鸭子叫
func (*DuckImp) DuckQuack() string {
	return "A Duck quack!"
}

// NewAdapter 适配器Adapter的工厂函数
func NewAdapter(duck Duck) Goose {
	return &adapter{
		Duck: duck,
	}
}

// 转换Duck为Goose的适配器
type adapter struct {
	Duck
}

// Quack 实现goose接口
func (a *adapter) Quack() string {
	return a.DuckQuack()
}
