package simpleFactory

import "fmt"

// Fruits 接口
type Fruits interface {
	Information(name string) string
}

// NewFruits 根据输入的 name 返回对应的接口实例
func NewFruits(name string) Fruits {
	if name == "apple" {
		return &apple{}
	} else if name == "banana" {
		return &banana{}
	}
	return nil
}

// apple 接口实现类
type apple struct {
}

func (*apple) Information(name string) string {
	return fmt.Sprintf("This is an %s", name)
}

// banana 接口实现类
type banana struct {
}

func (*banana) Information(name string) string {
	return fmt.Sprintf("This is a %s", name)
}
