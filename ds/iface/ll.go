package iface

type NodeInter interface {
	GetValue() interface{}
	GetNext() NodeInter // 返回接口
}

type LL interface {
	GetHead() NodeInter // 返回接口
}
