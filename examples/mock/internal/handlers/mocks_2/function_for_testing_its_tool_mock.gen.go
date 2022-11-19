// Code generated by Mock annotation processor. DO NOT EDIT.
// versions:
//		go: go1.18.3
//		go-annotation: 0.1.0
//		Mock: 0.0.1
// Code generated by counterfeiter. DO NOT EDIT.
package mocks_2

import (
	"net/http"
	"sync"

	"github.com/YReshetko/go-annotation/examples/mock/internal/handlers"
)

type FunctionForTestingItsToolMock struct {
	Stub        func(handlers.SomeInternalHandler, http.Response) http.Request
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 handlers.SomeInternalHandler
		arg2 http.Response
	}
	returns struct {
		result1 http.Request
	}
	returnsOnCall map[int]struct {
		result1 http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FunctionForTestingItsToolMock) Spy(arg1 handlers.SomeInternalHandler, arg2 http.Response) http.Request {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 handlers.SomeInternalHandler
		arg2 http.Response
	}{arg1, arg2})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("FunctionForTestingItsTool", []interface{}{arg1, arg2})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return returns.result1
}

func (fake *FunctionForTestingItsToolMock) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FunctionForTestingItsToolMock) Calls(stub func(handlers.SomeInternalHandler, http.Response) http.Request) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *FunctionForTestingItsToolMock) ArgsForCall(i int) (handlers.SomeInternalHandler, http.Response) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2
}

func (fake *FunctionForTestingItsToolMock) Returns(result1 http.Request) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 http.Request
	}{result1}
}

func (fake *FunctionForTestingItsToolMock) ReturnsOnCall(i int, result1 http.Request) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 http.Request
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 http.Request
	}{result1}
}

func (fake *FunctionForTestingItsToolMock) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FunctionForTestingItsToolMock) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ handlers.FunctionForTestingItsTool = new(FunctionForTestingItsToolMock).Spy
