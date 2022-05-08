package simpleFactory

import "testing"

func TestApple_Information(t *testing.T) {
	// 构造一个苹果类
	fruit := NewFruits("apple")
	// 调用类函数
	s := fruit.Information("apple")
	// 判断是否构造成功
	if s != "This is an apple" {
		t.Fatal("TestApple_Information test fail")
	}
}

func TestBanana_Information(t *testing.T) {
	fruit := NewFruits("banana")
	s := fruit.Information("banana")
	if s != "This is a banana" {
		t.Fatal("TestBanana_Information test fail")
	}
}
