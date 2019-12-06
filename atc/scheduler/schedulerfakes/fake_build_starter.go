// Code generated by counterfeiter. DO NOT EDIT.
package schedulerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/scheduler"
	"github.com/concourse/concourse/atc/scheduler/algorithm"
)

type FakeBuildStarter struct {
	TryStartPendingBuildsForJobStub        func(lager.Logger, db.Pipeline, db.Job, db.Resources, algorithm.NameToIDMap) (bool, error)
	tryStartPendingBuildsForJobMutex       sync.RWMutex
	tryStartPendingBuildsForJobArgsForCall []struct {
		arg1 lager.Logger
		arg2 db.Pipeline
		arg3 db.Job
		arg4 db.Resources
		arg5 algorithm.NameToIDMap
	}
	tryStartPendingBuildsForJobReturns struct {
		result1 bool
		result2 error
	}
	tryStartPendingBuildsForJobReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJob(arg1 lager.Logger, arg2 db.Pipeline, arg3 db.Job, arg4 db.Resources, arg5 algorithm.NameToIDMap) (bool, error) {
	fake.tryStartPendingBuildsForJobMutex.Lock()
	ret, specificReturn := fake.tryStartPendingBuildsForJobReturnsOnCall[len(fake.tryStartPendingBuildsForJobArgsForCall)]
	fake.tryStartPendingBuildsForJobArgsForCall = append(fake.tryStartPendingBuildsForJobArgsForCall, struct {
		arg1 lager.Logger
		arg2 db.Pipeline
		arg3 db.Job
		arg4 db.Resources
		arg5 algorithm.NameToIDMap
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("TryStartPendingBuildsForJob", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.tryStartPendingBuildsForJobMutex.Unlock()
	if fake.TryStartPendingBuildsForJobStub != nil {
		return fake.TryStartPendingBuildsForJobStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.tryStartPendingBuildsForJobReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobCallCount() int {
	fake.tryStartPendingBuildsForJobMutex.RLock()
	defer fake.tryStartPendingBuildsForJobMutex.RUnlock()
	return len(fake.tryStartPendingBuildsForJobArgsForCall)
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobCalls(stub func(lager.Logger, db.Pipeline, db.Job, db.Resources, algorithm.NameToIDMap) (bool, error)) {
	fake.tryStartPendingBuildsForJobMutex.Lock()
	defer fake.tryStartPendingBuildsForJobMutex.Unlock()
	fake.TryStartPendingBuildsForJobStub = stub
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobArgsForCall(i int) (lager.Logger, db.Pipeline, db.Job, db.Resources, algorithm.NameToIDMap) {
	fake.tryStartPendingBuildsForJobMutex.RLock()
	defer fake.tryStartPendingBuildsForJobMutex.RUnlock()
	argsForCall := fake.tryStartPendingBuildsForJobArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobReturns(result1 bool, result2 error) {
	fake.tryStartPendingBuildsForJobMutex.Lock()
	defer fake.tryStartPendingBuildsForJobMutex.Unlock()
	fake.TryStartPendingBuildsForJobStub = nil
	fake.tryStartPendingBuildsForJobReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobReturnsOnCall(i int, result1 bool, result2 error) {
	fake.tryStartPendingBuildsForJobMutex.Lock()
	defer fake.tryStartPendingBuildsForJobMutex.Unlock()
	fake.TryStartPendingBuildsForJobStub = nil
	if fake.tryStartPendingBuildsForJobReturnsOnCall == nil {
		fake.tryStartPendingBuildsForJobReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.tryStartPendingBuildsForJobReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildStarter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tryStartPendingBuildsForJobMutex.RLock()
	defer fake.tryStartPendingBuildsForJobMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBuildStarter) recordInvocation(key string, args []interface{}) {
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

var _ scheduler.BuildStarter = new(FakeBuildStarter)
