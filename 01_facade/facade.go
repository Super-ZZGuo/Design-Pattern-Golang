package facade

import "fmt"

func NEWAPI() API {
	return &ApiImp{
		a: NewAMoudleAPI(),
		b: NewBMoudleAPI(),
	}
}

type API interface {
	DoThing() string
}

type ApiImp struct {
	a AMoudleAPI
	b BMoudleAPI
}

func (api *ApiImp) DoThing() string {
	aRes := api.a.DoSomeThing()
	bRes := api.b.DoSomeThing()
	return fmt.Sprintf("%s and %s", aRes, bRes)
}

type AMoudleAPI interface {
	DoSomeThing() string
}

func NewAMoudleAPI() AMoudleAPI {
	return &AModuleImp{}
}

type AModuleImp struct{}

func (*AModuleImp) DoSomeThing() string {
	return "AModule do some things"
}

type BMoudleAPI interface {
	DoSomeThing() string
}

func NewBMoudleAPI() BMoudleAPI {
	return &BModuleImp{}
}

type BModuleImp struct{}

func (*BModuleImp) DoSomeThing() string {
	return "BModule do some things"
}
