package adapter

import "testing"

var expect = "A Duck quack!"

func TestAdapter(t *testing.T) {
	// 被适配类
	duck := NewDuck()
	// 适配类
	goose := NewAdapter(duck)
	res := goose.Quack()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
