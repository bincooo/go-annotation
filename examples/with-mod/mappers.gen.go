// Code generated by Mapper annotation processor. DO NOT EDIT.
// versions:
//		go: go1.18.3
//		go-annotation: 0.1.0
//		Mapper: 0.0.1-alpha

package with_mod

import (
	_imp_4 "github.com/eapache/go-resiliency/breaker"
	_imp_1 "github.com/eapache/queue"
)

var _ SomeInterface = (*SomeInterfaceImpl)(nil)

type SomeInterfaceImpl struct{}

func (_this_ SomeInterfaceImpl) mapper1(in0 _imp_1.Queue) _imp_1.Queue {
	out0 := _imp_1.Queue{}

	return out0
}

func (_this_ SomeInterfaceImpl) mapper(in0 _imp_4.Breaker) _imp_4.Breaker {
	out0 := _imp_4.Breaker{}

	return out0
}
