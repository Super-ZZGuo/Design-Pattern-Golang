package facade

import "fmt"

// NEWAPI 返回一个Api实例，包含a、b两个模组的实例
func NEWAPI() API {
	return &ApiImp{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API 定义顶层通用API接口
type API interface {
	DoThing() string
}

// ApiImp 接口实现
type ApiImp struct {
	a AModuleAPI
	b BModuleAPI
}

// DoThing 封装功能
func (api *ApiImp) DoThing() string {
	aRes := api.a.DoSomeThing()
	bRes := api.b.DoSomeThing()
	return fmt.Sprintf("%s and %s", aRes, bRes)
}

// AMoudleAPI A模组接口
type AModuleAPI interface {
	DoSomeThing() string
}

// NewAModuleAPI 返回一个A模组实例
func NewAModuleAPI() AModuleAPI {
	return &AModuleImp{}
}

// AModuleImp A模组实例
type AModuleImp struct{}

// DoSomeThing A模组实现的功能
func (*AModuleImp) DoSomeThing() string {
	return "AModule do some things"
}

// BModuleAPI B模组接口
type BModuleAPI interface {
	DoSomeThing() string
}

// NewBModuleAPI 返回B模组实例
func NewBModuleAPI() BModuleAPI {
	return &BModuleImp{}
}

// BModuleImp B模组实例
type BModuleImp struct{}

// DoSomeThing B模组实现的功能
func (*BModuleImp) DoSomeThing() string {
	return "BModule do some things"
}
