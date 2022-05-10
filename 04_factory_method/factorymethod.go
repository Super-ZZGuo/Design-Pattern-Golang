package factory_method

// Operator 操作接口类
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 抽象工厂接口
type OperatorFactory interface {
	Creat() Operator
}

// OperatorImp 操作的实现类
type OperatorImp struct {
	a, b int
}

//SetA 设置 A
func (o *OperatorImp) SetA(a int) {
	o.a = a
}

//SetB 设置 B
func (o *OperatorImp) SetB(b int) {
	o.b = b
}

// PlusOperatorFactory PlusOperator的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Creat() Operator {
	return &PlusOperator{
		OperatorImp: &OperatorImp{},
	}
}

// PlusOperator 加法操作的工厂类
type PlusOperator struct {
	*OperatorImp
}

// Result PlusOperatorFactory 实际减法操作
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// MinusOperatorFactory MinOperator的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Creat() Operator {
	return &MinusOperator{
		OperatorImp: &OperatorImp{},
	}
}

// MinusOperator 减法操作的工厂类
type MinusOperator struct {
	*OperatorImp
}

// Result MinusOperatorFactory 实际减法操作
func (o MinusOperator) Result() int {
	return o.a - o.b
}
