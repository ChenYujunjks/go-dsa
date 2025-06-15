package iface

type NodeInter interface {
	GetValue() interface{}
	GetNext() NodeInter
}

type LL interface {
	Append(value interface{})
	GetHead() NodeInter // 返回类型你自己控制
}
