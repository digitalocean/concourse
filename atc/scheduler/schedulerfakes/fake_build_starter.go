// Code generated by counterfeiter. DO NOT EDIT.
package schedulerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/scheduler"
)

type FakeBuildStarter struct {
	TryStartPendingBuildsForJobStub        func(lager.Logger, db.Job, db.Resources, atc.VersionedResourceTypes, []db.Build) error
	tryStartPendingBuildsForJobMutex       sync.RWMutex
	tryStartPendingBuildsForJobArgsForCall []struct {
		arg1 lager.Logger
		arg2 db.Job
		arg3 db.Resources
		arg4 atc.VersionedResourceTypes
		arg5 []db.Build
	}
	tryStartPendingBuildsForJobReturns struct {
		result1 error
	}
	tryStartPendingBuildsForJobReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJob(arg1 lager.Logger, arg2 db.Job, arg3 db.Resources, arg4 atc.VersionedResourceTypes, arg5 []db.Build) error {
	var arg5Copy []db.Build
	if arg5 != nil {
		arg5Copy = make([]db.Build, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.tryStartPendingBuildsForJobMutex.Lock()
	ret, specificReturn := fake.tryStartPendingBuildsForJobReturnsOnCall[len(fake.tryStartPendingBuildsForJobArgsForCall)]
	fake.tryStartPendingBuildsForJobArgsForCall = append(fake.tryStartPendingBuildsForJobArgsForCall, struct {
		arg1 lager.Logger
		arg2 db.Job
		arg3 db.Resources
		arg4 atc.VersionedResourceTypes
		arg5 []db.Build
	}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.recordInvocation("TryStartPendingBuildsForJob", []interface{}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.tryStartPendingBuildsForJobMutex.Unlock()
	if fake.TryStartPendingBuildsForJobStub != nil {
		return fake.TryStartPendingBuildsForJobStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tryStartPendingBuildsForJobReturns
	return fakeReturns.result1
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobCallCount() int {
	fake.tryStartPendingBuildsForJobMutex.RLock()
	defer fake.tryStartPendingBuildsForJobMutex.RUnlock()
	return len(fake.tryStartPendingBuildsForJobArgsForCall)
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobCalls(stub func(lager.Logger, db.Job, db.Resources, atc.VersionedResourceTypes, []db.Build) error) {
	fake.tryStartPendingBuildsForJobMutex.Lock()
	defer fake.tryStartPendingBuildsForJobMutex.Unlock()
	fake.TryStartPendingBuildsForJobStub = stub
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobArgsForCall(i int) (lager.Logger, db.Job, db.Resources, atc.VersionedResourceTypes, []db.Build) {
	fake.tryStartPendingBuildsForJobMutex.RLock()
	defer fake.tryStartPendingBuildsForJobMutex.RUnlock()
	argsForCall := fake.tryStartPendingBuildsForJobArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobReturns(result1 error) {
	fake.tryStartPendingBuildsForJobMutex.Lock()
	defer fake.tryStartPendingBuildsForJobMutex.Unlock()
	fake.TryStartPendingBuildsForJobStub = nil
	fake.tryStartPendingBuildsForJobReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildStarter) TryStartPendingBuildsForJobReturnsOnCall(i int, result1 error) {
	fake.tryStartPendingBuildsForJobMutex.Lock()
	defer fake.tryStartPendingBuildsForJobMutex.Unlock()
	fake.TryStartPendingBuildsForJobStub = nil
	if fake.tryStartPendingBuildsForJobReturnsOnCall == nil {
		fake.tryStartPendingBuildsForJobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.tryStartPendingBuildsForJobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
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
