// Code generated by counterfeiter. DO NOT EDIT.
package emitfakes

import (
	"mysql-metrics/emit"
	"sync"
)

type FakeLogger struct {
	ErrorStub        func(string, error)
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
		arg1 string
		arg2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogger) Error(arg1 string, arg2 error) {
	fake.errorMutex.Lock()
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
		arg1 string
		arg2 error
	}{arg1, arg2})
	fake.recordInvocation("Error", []interface{}{arg1, arg2})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		fake.ErrorStub(arg1, arg2)
	}
}

func (fake *FakeLogger) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeLogger) ErrorArgsForCall(i int) (string, error) {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return fake.errorArgsForCall[i].arg1, fake.errorArgsForCall[i].arg2
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogger) recordInvocation(key string, args []interface{}) {
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

var _ emit.Logger = new(FakeLogger)
