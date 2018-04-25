// Code generated by counterfeiter. DO NOT EDIT.
package gatherfakes

import (
	"mysql-metrics/gather"
	"sync"
)

type FakeDatabaseClient struct {
	ShowGlobalStatusStub        func() (map[string]string, error)
	showGlobalStatusMutex       sync.RWMutex
	showGlobalStatusArgsForCall []struct{}
	showGlobalStatusReturns     struct {
		result1 map[string]string
		result2 error
	}
	showGlobalStatusReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	ShowGlobalVariablesStub        func() (map[string]string, error)
	showGlobalVariablesMutex       sync.RWMutex
	showGlobalVariablesArgsForCall []struct{}
	showGlobalVariablesReturns     struct {
		result1 map[string]string
		result2 error
	}
	showGlobalVariablesReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	ShowSlaveStatusStub        func() (map[string]string, error)
	showSlaveStatusMutex       sync.RWMutex
	showSlaveStatusArgsForCall []struct{}
	showSlaveStatusReturns     struct {
		result1 map[string]string
		result2 error
	}
	showSlaveStatusReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	HeartbeatStatusStub        func() (map[string]string, error)
	heartbeatStatusMutex       sync.RWMutex
	heartbeatStatusArgsForCall []struct{}
	heartbeatStatusReturns     struct {
		result1 map[string]string
		result2 error
	}
	heartbeatStatusReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	ServicePlansDiskAllocatedStub        func() (map[string]string, error)
	servicePlansDiskAllocatedMutex       sync.RWMutex
	servicePlansDiskAllocatedArgsForCall []struct{}
	servicePlansDiskAllocatedReturns     struct {
		result1 map[string]string
		result2 error
	}
	servicePlansDiskAllocatedReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	IsAvailableStub        func() bool
	isAvailableMutex       sync.RWMutex
	isAvailableArgsForCall []struct{}
	isAvailableReturns     struct {
		result1 bool
	}
	isAvailableReturnsOnCall map[int]struct {
		result1 bool
	}
	IsFollowerStub        func() (bool, error)
	isFollowerMutex       sync.RWMutex
	isFollowerArgsForCall []struct{}
	isFollowerReturns     struct {
		result1 bool
		result2 error
	}
	isFollowerReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDatabaseClient) ShowGlobalStatus() (map[string]string, error) {
	fake.showGlobalStatusMutex.Lock()
	ret, specificReturn := fake.showGlobalStatusReturnsOnCall[len(fake.showGlobalStatusArgsForCall)]
	fake.showGlobalStatusArgsForCall = append(fake.showGlobalStatusArgsForCall, struct{}{})
	fake.recordInvocation("ShowGlobalStatus", []interface{}{})
	fake.showGlobalStatusMutex.Unlock()
	if fake.ShowGlobalStatusStub != nil {
		return fake.ShowGlobalStatusStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.showGlobalStatusReturns.result1, fake.showGlobalStatusReturns.result2
}

func (fake *FakeDatabaseClient) ShowGlobalStatusCallCount() int {
	fake.showGlobalStatusMutex.RLock()
	defer fake.showGlobalStatusMutex.RUnlock()
	return len(fake.showGlobalStatusArgsForCall)
}

func (fake *FakeDatabaseClient) ShowGlobalStatusReturns(result1 map[string]string, result2 error) {
	fake.ShowGlobalStatusStub = nil
	fake.showGlobalStatusReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) ShowGlobalStatusReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.ShowGlobalStatusStub = nil
	if fake.showGlobalStatusReturnsOnCall == nil {
		fake.showGlobalStatusReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.showGlobalStatusReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) ShowGlobalVariables() (map[string]string, error) {
	fake.showGlobalVariablesMutex.Lock()
	ret, specificReturn := fake.showGlobalVariablesReturnsOnCall[len(fake.showGlobalVariablesArgsForCall)]
	fake.showGlobalVariablesArgsForCall = append(fake.showGlobalVariablesArgsForCall, struct{}{})
	fake.recordInvocation("ShowGlobalVariables", []interface{}{})
	fake.showGlobalVariablesMutex.Unlock()
	if fake.ShowGlobalVariablesStub != nil {
		return fake.ShowGlobalVariablesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.showGlobalVariablesReturns.result1, fake.showGlobalVariablesReturns.result2
}

func (fake *FakeDatabaseClient) ShowGlobalVariablesCallCount() int {
	fake.showGlobalVariablesMutex.RLock()
	defer fake.showGlobalVariablesMutex.RUnlock()
	return len(fake.showGlobalVariablesArgsForCall)
}

func (fake *FakeDatabaseClient) ShowGlobalVariablesReturns(result1 map[string]string, result2 error) {
	fake.ShowGlobalVariablesStub = nil
	fake.showGlobalVariablesReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) ShowGlobalVariablesReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.ShowGlobalVariablesStub = nil
	if fake.showGlobalVariablesReturnsOnCall == nil {
		fake.showGlobalVariablesReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.showGlobalVariablesReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) ShowSlaveStatus() (map[string]string, error) {
	fake.showSlaveStatusMutex.Lock()
	ret, specificReturn := fake.showSlaveStatusReturnsOnCall[len(fake.showSlaveStatusArgsForCall)]
	fake.showSlaveStatusArgsForCall = append(fake.showSlaveStatusArgsForCall, struct{}{})
	fake.recordInvocation("ShowSlaveStatus", []interface{}{})
	fake.showSlaveStatusMutex.Unlock()
	if fake.ShowSlaveStatusStub != nil {
		return fake.ShowSlaveStatusStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.showSlaveStatusReturns.result1, fake.showSlaveStatusReturns.result2
}

func (fake *FakeDatabaseClient) ShowSlaveStatusCallCount() int {
	fake.showSlaveStatusMutex.RLock()
	defer fake.showSlaveStatusMutex.RUnlock()
	return len(fake.showSlaveStatusArgsForCall)
}

func (fake *FakeDatabaseClient) ShowSlaveStatusReturns(result1 map[string]string, result2 error) {
	fake.ShowSlaveStatusStub = nil
	fake.showSlaveStatusReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) ShowSlaveStatusReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.ShowSlaveStatusStub = nil
	if fake.showSlaveStatusReturnsOnCall == nil {
		fake.showSlaveStatusReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.showSlaveStatusReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) HeartbeatStatus() (map[string]string, error) {
	fake.heartbeatStatusMutex.Lock()
	ret, specificReturn := fake.heartbeatStatusReturnsOnCall[len(fake.heartbeatStatusArgsForCall)]
	fake.heartbeatStatusArgsForCall = append(fake.heartbeatStatusArgsForCall, struct{}{})
	fake.recordInvocation("HeartbeatStatus", []interface{}{})
	fake.heartbeatStatusMutex.Unlock()
	if fake.HeartbeatStatusStub != nil {
		return fake.HeartbeatStatusStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.heartbeatStatusReturns.result1, fake.heartbeatStatusReturns.result2
}

func (fake *FakeDatabaseClient) HeartbeatStatusCallCount() int {
	fake.heartbeatStatusMutex.RLock()
	defer fake.heartbeatStatusMutex.RUnlock()
	return len(fake.heartbeatStatusArgsForCall)
}

func (fake *FakeDatabaseClient) HeartbeatStatusReturns(result1 map[string]string, result2 error) {
	fake.HeartbeatStatusStub = nil
	fake.heartbeatStatusReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) HeartbeatStatusReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.HeartbeatStatusStub = nil
	if fake.heartbeatStatusReturnsOnCall == nil {
		fake.heartbeatStatusReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.heartbeatStatusReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) ServicePlansDiskAllocated() (map[string]string, error) {
	fake.servicePlansDiskAllocatedMutex.Lock()
	ret, specificReturn := fake.servicePlansDiskAllocatedReturnsOnCall[len(fake.servicePlansDiskAllocatedArgsForCall)]
	fake.servicePlansDiskAllocatedArgsForCall = append(fake.servicePlansDiskAllocatedArgsForCall, struct{}{})
	fake.recordInvocation("ServicePlansDiskAllocated", []interface{}{})
	fake.servicePlansDiskAllocatedMutex.Unlock()
	if fake.ServicePlansDiskAllocatedStub != nil {
		return fake.ServicePlansDiskAllocatedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.servicePlansDiskAllocatedReturns.result1, fake.servicePlansDiskAllocatedReturns.result2
}

func (fake *FakeDatabaseClient) ServicePlansDiskAllocatedCallCount() int {
	fake.servicePlansDiskAllocatedMutex.RLock()
	defer fake.servicePlansDiskAllocatedMutex.RUnlock()
	return len(fake.servicePlansDiskAllocatedArgsForCall)
}

func (fake *FakeDatabaseClient) ServicePlansDiskAllocatedReturns(result1 map[string]string, result2 error) {
	fake.ServicePlansDiskAllocatedStub = nil
	fake.servicePlansDiskAllocatedReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) ServicePlansDiskAllocatedReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.ServicePlansDiskAllocatedStub = nil
	if fake.servicePlansDiskAllocatedReturnsOnCall == nil {
		fake.servicePlansDiskAllocatedReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.servicePlansDiskAllocatedReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) IsAvailable() bool {
	fake.isAvailableMutex.Lock()
	ret, specificReturn := fake.isAvailableReturnsOnCall[len(fake.isAvailableArgsForCall)]
	fake.isAvailableArgsForCall = append(fake.isAvailableArgsForCall, struct{}{})
	fake.recordInvocation("IsAvailable", []interface{}{})
	fake.isAvailableMutex.Unlock()
	if fake.IsAvailableStub != nil {
		return fake.IsAvailableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isAvailableReturns.result1
}

func (fake *FakeDatabaseClient) IsAvailableCallCount() int {
	fake.isAvailableMutex.RLock()
	defer fake.isAvailableMutex.RUnlock()
	return len(fake.isAvailableArgsForCall)
}

func (fake *FakeDatabaseClient) IsAvailableReturns(result1 bool) {
	fake.IsAvailableStub = nil
	fake.isAvailableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDatabaseClient) IsAvailableReturnsOnCall(i int, result1 bool) {
	fake.IsAvailableStub = nil
	if fake.isAvailableReturnsOnCall == nil {
		fake.isAvailableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isAvailableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDatabaseClient) IsFollower() (bool, error) {
	fake.isFollowerMutex.Lock()
	ret, specificReturn := fake.isFollowerReturnsOnCall[len(fake.isFollowerArgsForCall)]
	fake.isFollowerArgsForCall = append(fake.isFollowerArgsForCall, struct{}{})
	fake.recordInvocation("IsFollower", []interface{}{})
	fake.isFollowerMutex.Unlock()
	if fake.IsFollowerStub != nil {
		return fake.IsFollowerStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isFollowerReturns.result1, fake.isFollowerReturns.result2
}

func (fake *FakeDatabaseClient) IsFollowerCallCount() int {
	fake.isFollowerMutex.RLock()
	defer fake.isFollowerMutex.RUnlock()
	return len(fake.isFollowerArgsForCall)
}

func (fake *FakeDatabaseClient) IsFollowerReturns(result1 bool, result2 error) {
	fake.IsFollowerStub = nil
	fake.isFollowerReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) IsFollowerReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsFollowerStub = nil
	if fake.isFollowerReturnsOnCall == nil {
		fake.isFollowerReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isFollowerReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.showGlobalStatusMutex.RLock()
	defer fake.showGlobalStatusMutex.RUnlock()
	fake.showGlobalVariablesMutex.RLock()
	defer fake.showGlobalVariablesMutex.RUnlock()
	fake.showSlaveStatusMutex.RLock()
	defer fake.showSlaveStatusMutex.RUnlock()
	fake.heartbeatStatusMutex.RLock()
	defer fake.heartbeatStatusMutex.RUnlock()
	fake.servicePlansDiskAllocatedMutex.RLock()
	defer fake.servicePlansDiskAllocatedMutex.RUnlock()
	fake.isAvailableMutex.RLock()
	defer fake.isAvailableMutex.RUnlock()
	fake.isFollowerMutex.RLock()
	defer fake.isFollowerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDatabaseClient) recordInvocation(key string, args []interface{}) {
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

var _ gather.DatabaseClient = new(FakeDatabaseClient)
