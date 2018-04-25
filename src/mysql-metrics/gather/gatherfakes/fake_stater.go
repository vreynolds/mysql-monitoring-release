// Code generated by counterfeiter. DO NOT EDIT.
package gatherfakes

import (
	"mysql-metrics/gather"
	"sync"
)

type FakeStater struct {
	StatsStub        func(path string) (bytesFree, bytesTotal, inodesFree, inodesTotal uint64, err error)
	statsMutex       sync.RWMutex
	statsArgsForCall []struct {
		path string
	}
	statsReturns struct {
		result1 uint64
		result2 uint64
		result3 uint64
		result4 uint64
		result5 error
	}
	statsReturnsOnCall map[int]struct {
		result1 uint64
		result2 uint64
		result3 uint64
		result4 uint64
		result5 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStater) Stats(path string) (bytesFree, bytesTotal, inodesFree, inodesTotal uint64, err error) {
	fake.statsMutex.Lock()
	ret, specificReturn := fake.statsReturnsOnCall[len(fake.statsArgsForCall)]
	fake.statsArgsForCall = append(fake.statsArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("Stats", []interface{}{path})
	fake.statsMutex.Unlock()
	if fake.StatsStub != nil {
		return fake.StatsStub(path)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4, ret.result5
	}
	return fake.statsReturns.result1, fake.statsReturns.result2, fake.statsReturns.result3, fake.statsReturns.result4, fake.statsReturns.result5
}

func (fake *FakeStater) StatsCallCount() int {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return len(fake.statsArgsForCall)
}

func (fake *FakeStater) StatsArgsForCall(i int) string {
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	return fake.statsArgsForCall[i].path
}

func (fake *FakeStater) StatsReturns(result1 uint64, result2 uint64, result3 uint64, result4 uint64, result5 error) {
	fake.StatsStub = nil
	fake.statsReturns = struct {
		result1 uint64
		result2 uint64
		result3 uint64
		result4 uint64
		result5 error
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeStater) StatsReturnsOnCall(i int, result1 uint64, result2 uint64, result3 uint64, result4 uint64, result5 error) {
	fake.StatsStub = nil
	if fake.statsReturnsOnCall == nil {
		fake.statsReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 uint64
			result3 uint64
			result4 uint64
			result5 error
		})
	}
	fake.statsReturnsOnCall[i] = struct {
		result1 uint64
		result2 uint64
		result3 uint64
		result4 uint64
		result5 error
	}{result1, result2, result3, result4, result5}
}

func (fake *FakeStater) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.statsMutex.RLock()
	defer fake.statsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStater) recordInvocation(key string, args []interface{}) {
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

var _ gather.Stater = new(FakeStater)
