// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type StagedConfigService struct {
	GetDeployedProductCredentialStub        func(api.GetDeployedProductCredentialInput) (api.GetDeployedProductCredentialOutput, error)
	getDeployedProductCredentialMutex       sync.RWMutex
	getDeployedProductCredentialArgsForCall []struct {
		arg1 api.GetDeployedProductCredentialInput
	}
	getDeployedProductCredentialReturns struct {
		result1 api.GetDeployedProductCredentialOutput
		result2 error
	}
	getDeployedProductCredentialReturnsOnCall map[int]struct {
		result1 api.GetDeployedProductCredentialOutput
		result2 error
	}
	GetStagedProductByNameStub        func(string) (api.StagedProductsFindOutput, error)
	getStagedProductByNameMutex       sync.RWMutex
	getStagedProductByNameArgsForCall []struct {
		arg1 string
	}
	getStagedProductByNameReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	getStagedProductByNameReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	GetStagedProductJobResourceConfigStub        func(string, string) (api.JobProperties, error)
	getStagedProductJobResourceConfigMutex       sync.RWMutex
	getStagedProductJobResourceConfigArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getStagedProductJobResourceConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	getStagedProductJobResourceConfigReturnsOnCall map[int]struct {
		result1 api.JobProperties
		result2 error
	}
	GetStagedProductNetworksAndAZsStub        func(string) (map[string]interface{}, error)
	getStagedProductNetworksAndAZsMutex       sync.RWMutex
	getStagedProductNetworksAndAZsArgsForCall []struct {
		arg1 string
	}
	getStagedProductNetworksAndAZsReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	getStagedProductNetworksAndAZsReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	GetStagedProductPropertiesStub        func(string) (map[string]api.ResponseProperty, error)
	getStagedProductPropertiesMutex       sync.RWMutex
	getStagedProductPropertiesArgsForCall []struct {
		arg1 string
	}
	getStagedProductPropertiesReturns struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	getStagedProductPropertiesReturnsOnCall map[int]struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	ListDeployedProductsStub        func() ([]api.DeployedProductOutput, error)
	listDeployedProductsMutex       sync.RWMutex
	listDeployedProductsArgsForCall []struct {
	}
	listDeployedProductsReturns struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	listDeployedProductsReturnsOnCall map[int]struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	ListStagedProductErrandsStub        func(string) (api.ErrandsListOutput, error)
	listStagedProductErrandsMutex       sync.RWMutex
	listStagedProductErrandsArgsForCall []struct {
		arg1 string
	}
	listStagedProductErrandsReturns struct {
		result1 api.ErrandsListOutput
		result2 error
	}
	listStagedProductErrandsReturnsOnCall map[int]struct {
		result1 api.ErrandsListOutput
		result2 error
	}
	ListStagedProductJobsStub        func(string) (map[string]string, error)
	listStagedProductJobsMutex       sync.RWMutex
	listStagedProductJobsArgsForCall []struct {
		arg1 string
	}
	listStagedProductJobsReturns struct {
		result1 map[string]string
		result2 error
	}
	listStagedProductJobsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StagedConfigService) GetDeployedProductCredential(arg1 api.GetDeployedProductCredentialInput) (api.GetDeployedProductCredentialOutput, error) {
	fake.getDeployedProductCredentialMutex.Lock()
	ret, specificReturn := fake.getDeployedProductCredentialReturnsOnCall[len(fake.getDeployedProductCredentialArgsForCall)]
	fake.getDeployedProductCredentialArgsForCall = append(fake.getDeployedProductCredentialArgsForCall, struct {
		arg1 api.GetDeployedProductCredentialInput
	}{arg1})
	fake.recordInvocation("GetDeployedProductCredential", []interface{}{arg1})
	fake.getDeployedProductCredentialMutex.Unlock()
	if fake.GetDeployedProductCredentialStub != nil {
		return fake.GetDeployedProductCredentialStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getDeployedProductCredentialReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedConfigService) GetDeployedProductCredentialCallCount() int {
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	return len(fake.getDeployedProductCredentialArgsForCall)
}

func (fake *StagedConfigService) GetDeployedProductCredentialCalls(stub func(api.GetDeployedProductCredentialInput) (api.GetDeployedProductCredentialOutput, error)) {
	fake.getDeployedProductCredentialMutex.Lock()
	defer fake.getDeployedProductCredentialMutex.Unlock()
	fake.GetDeployedProductCredentialStub = stub
}

func (fake *StagedConfigService) GetDeployedProductCredentialArgsForCall(i int) api.GetDeployedProductCredentialInput {
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	argsForCall := fake.getDeployedProductCredentialArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedConfigService) GetDeployedProductCredentialReturns(result1 api.GetDeployedProductCredentialOutput, result2 error) {
	fake.getDeployedProductCredentialMutex.Lock()
	defer fake.getDeployedProductCredentialMutex.Unlock()
	fake.GetDeployedProductCredentialStub = nil
	fake.getDeployedProductCredentialReturns = struct {
		result1 api.GetDeployedProductCredentialOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetDeployedProductCredentialReturnsOnCall(i int, result1 api.GetDeployedProductCredentialOutput, result2 error) {
	fake.getDeployedProductCredentialMutex.Lock()
	defer fake.getDeployedProductCredentialMutex.Unlock()
	fake.GetDeployedProductCredentialStub = nil
	if fake.getDeployedProductCredentialReturnsOnCall == nil {
		fake.getDeployedProductCredentialReturnsOnCall = make(map[int]struct {
			result1 api.GetDeployedProductCredentialOutput
			result2 error
		})
	}
	fake.getDeployedProductCredentialReturnsOnCall[i] = struct {
		result1 api.GetDeployedProductCredentialOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductByName(arg1 string) (api.StagedProductsFindOutput, error) {
	fake.getStagedProductByNameMutex.Lock()
	ret, specificReturn := fake.getStagedProductByNameReturnsOnCall[len(fake.getStagedProductByNameArgsForCall)]
	fake.getStagedProductByNameArgsForCall = append(fake.getStagedProductByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStagedProductByName", []interface{}{arg1})
	fake.getStagedProductByNameMutex.Unlock()
	if fake.GetStagedProductByNameStub != nil {
		return fake.GetStagedProductByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedConfigService) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *StagedConfigService) GetStagedProductByNameCalls(stub func(string) (api.StagedProductsFindOutput, error)) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = stub
}

func (fake *StagedConfigService) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	argsForCall := fake.getStagedProductByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedConfigService) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.getStagedProductByNameMutex.Lock()
	defer fake.getStagedProductByNameMutex.Unlock()
	fake.GetStagedProductByNameStub = nil
	if fake.getStagedProductByNameReturnsOnCall == nil {
		fake.getStagedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.getStagedProductByNameReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfig(arg1 string, arg2 string) (api.JobProperties, error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	ret, specificReturn := fake.getStagedProductJobResourceConfigReturnsOnCall[len(fake.getStagedProductJobResourceConfigArgsForCall)]
	fake.getStagedProductJobResourceConfigArgsForCall = append(fake.getStagedProductJobResourceConfigArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetStagedProductJobResourceConfig", []interface{}{arg1, arg2})
	fake.getStagedProductJobResourceConfigMutex.Unlock()
	if fake.GetStagedProductJobResourceConfigStub != nil {
		return fake.GetStagedProductJobResourceConfigStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductJobResourceConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigCallCount() int {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	return len(fake.getStagedProductJobResourceConfigArgsForCall)
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigCalls(stub func(string, string) (api.JobProperties, error)) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	defer fake.getStagedProductJobResourceConfigMutex.Unlock()
	fake.GetStagedProductJobResourceConfigStub = stub
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigArgsForCall(i int) (string, string) {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	argsForCall := fake.getStagedProductJobResourceConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigReturns(result1 api.JobProperties, result2 error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	defer fake.getStagedProductJobResourceConfigMutex.Unlock()
	fake.GetStagedProductJobResourceConfigStub = nil
	fake.getStagedProductJobResourceConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigReturnsOnCall(i int, result1 api.JobProperties, result2 error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	defer fake.getStagedProductJobResourceConfigMutex.Unlock()
	fake.GetStagedProductJobResourceConfigStub = nil
	if fake.getStagedProductJobResourceConfigReturnsOnCall == nil {
		fake.getStagedProductJobResourceConfigReturnsOnCall = make(map[int]struct {
			result1 api.JobProperties
			result2 error
		})
	}
	fake.getStagedProductJobResourceConfigReturnsOnCall[i] = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZs(arg1 string) (map[string]interface{}, error) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	ret, specificReturn := fake.getStagedProductNetworksAndAZsReturnsOnCall[len(fake.getStagedProductNetworksAndAZsArgsForCall)]
	fake.getStagedProductNetworksAndAZsArgsForCall = append(fake.getStagedProductNetworksAndAZsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStagedProductNetworksAndAZs", []interface{}{arg1})
	fake.getStagedProductNetworksAndAZsMutex.Unlock()
	if fake.GetStagedProductNetworksAndAZsStub != nil {
		return fake.GetStagedProductNetworksAndAZsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductNetworksAndAZsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsCallCount() int {
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	return len(fake.getStagedProductNetworksAndAZsArgsForCall)
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsCalls(stub func(string) (map[string]interface{}, error)) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	defer fake.getStagedProductNetworksAndAZsMutex.Unlock()
	fake.GetStagedProductNetworksAndAZsStub = stub
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsArgsForCall(i int) string {
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	argsForCall := fake.getStagedProductNetworksAndAZsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsReturns(result1 map[string]interface{}, result2 error) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	defer fake.getStagedProductNetworksAndAZsMutex.Unlock()
	fake.GetStagedProductNetworksAndAZsStub = nil
	fake.getStagedProductNetworksAndAZsReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	defer fake.getStagedProductNetworksAndAZsMutex.Unlock()
	fake.GetStagedProductNetworksAndAZsStub = nil
	if fake.getStagedProductNetworksAndAZsReturnsOnCall == nil {
		fake.getStagedProductNetworksAndAZsReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.getStagedProductNetworksAndAZsReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductProperties(arg1 string) (map[string]api.ResponseProperty, error) {
	fake.getStagedProductPropertiesMutex.Lock()
	ret, specificReturn := fake.getStagedProductPropertiesReturnsOnCall[len(fake.getStagedProductPropertiesArgsForCall)]
	fake.getStagedProductPropertiesArgsForCall = append(fake.getStagedProductPropertiesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStagedProductProperties", []interface{}{arg1})
	fake.getStagedProductPropertiesMutex.Unlock()
	if fake.GetStagedProductPropertiesStub != nil {
		return fake.GetStagedProductPropertiesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStagedProductPropertiesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedConfigService) GetStagedProductPropertiesCallCount() int {
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	return len(fake.getStagedProductPropertiesArgsForCall)
}

func (fake *StagedConfigService) GetStagedProductPropertiesCalls(stub func(string) (map[string]api.ResponseProperty, error)) {
	fake.getStagedProductPropertiesMutex.Lock()
	defer fake.getStagedProductPropertiesMutex.Unlock()
	fake.GetStagedProductPropertiesStub = stub
}

func (fake *StagedConfigService) GetStagedProductPropertiesArgsForCall(i int) string {
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	argsForCall := fake.getStagedProductPropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedConfigService) GetStagedProductPropertiesReturns(result1 map[string]api.ResponseProperty, result2 error) {
	fake.getStagedProductPropertiesMutex.Lock()
	defer fake.getStagedProductPropertiesMutex.Unlock()
	fake.GetStagedProductPropertiesStub = nil
	fake.getStagedProductPropertiesReturns = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductPropertiesReturnsOnCall(i int, result1 map[string]api.ResponseProperty, result2 error) {
	fake.getStagedProductPropertiesMutex.Lock()
	defer fake.getStagedProductPropertiesMutex.Unlock()
	fake.GetStagedProductPropertiesStub = nil
	if fake.getStagedProductPropertiesReturnsOnCall == nil {
		fake.getStagedProductPropertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]api.ResponseProperty
			result2 error
		})
	}
	fake.getStagedProductPropertiesReturnsOnCall[i] = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListDeployedProducts() ([]api.DeployedProductOutput, error) {
	fake.listDeployedProductsMutex.Lock()
	ret, specificReturn := fake.listDeployedProductsReturnsOnCall[len(fake.listDeployedProductsArgsForCall)]
	fake.listDeployedProductsArgsForCall = append(fake.listDeployedProductsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListDeployedProducts", []interface{}{})
	fake.listDeployedProductsMutex.Unlock()
	if fake.ListDeployedProductsStub != nil {
		return fake.ListDeployedProductsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listDeployedProductsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedConfigService) ListDeployedProductsCallCount() int {
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	return len(fake.listDeployedProductsArgsForCall)
}

func (fake *StagedConfigService) ListDeployedProductsCalls(stub func() ([]api.DeployedProductOutput, error)) {
	fake.listDeployedProductsMutex.Lock()
	defer fake.listDeployedProductsMutex.Unlock()
	fake.ListDeployedProductsStub = stub
}

func (fake *StagedConfigService) ListDeployedProductsReturns(result1 []api.DeployedProductOutput, result2 error) {
	fake.listDeployedProductsMutex.Lock()
	defer fake.listDeployedProductsMutex.Unlock()
	fake.ListDeployedProductsStub = nil
	fake.listDeployedProductsReturns = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListDeployedProductsReturnsOnCall(i int, result1 []api.DeployedProductOutput, result2 error) {
	fake.listDeployedProductsMutex.Lock()
	defer fake.listDeployedProductsMutex.Unlock()
	fake.ListDeployedProductsStub = nil
	if fake.listDeployedProductsReturnsOnCall == nil {
		fake.listDeployedProductsReturnsOnCall = make(map[int]struct {
			result1 []api.DeployedProductOutput
			result2 error
		})
	}
	fake.listDeployedProductsReturnsOnCall[i] = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListStagedProductErrands(arg1 string) (api.ErrandsListOutput, error) {
	fake.listStagedProductErrandsMutex.Lock()
	ret, specificReturn := fake.listStagedProductErrandsReturnsOnCall[len(fake.listStagedProductErrandsArgsForCall)]
	fake.listStagedProductErrandsArgsForCall = append(fake.listStagedProductErrandsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListStagedProductErrands", []interface{}{arg1})
	fake.listStagedProductErrandsMutex.Unlock()
	if fake.ListStagedProductErrandsStub != nil {
		return fake.ListStagedProductErrandsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listStagedProductErrandsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedConfigService) ListStagedProductErrandsCallCount() int {
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	return len(fake.listStagedProductErrandsArgsForCall)
}

func (fake *StagedConfigService) ListStagedProductErrandsCalls(stub func(string) (api.ErrandsListOutput, error)) {
	fake.listStagedProductErrandsMutex.Lock()
	defer fake.listStagedProductErrandsMutex.Unlock()
	fake.ListStagedProductErrandsStub = stub
}

func (fake *StagedConfigService) ListStagedProductErrandsArgsForCall(i int) string {
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	argsForCall := fake.listStagedProductErrandsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedConfigService) ListStagedProductErrandsReturns(result1 api.ErrandsListOutput, result2 error) {
	fake.listStagedProductErrandsMutex.Lock()
	defer fake.listStagedProductErrandsMutex.Unlock()
	fake.ListStagedProductErrandsStub = nil
	fake.listStagedProductErrandsReturns = struct {
		result1 api.ErrandsListOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListStagedProductErrandsReturnsOnCall(i int, result1 api.ErrandsListOutput, result2 error) {
	fake.listStagedProductErrandsMutex.Lock()
	defer fake.listStagedProductErrandsMutex.Unlock()
	fake.ListStagedProductErrandsStub = nil
	if fake.listStagedProductErrandsReturnsOnCall == nil {
		fake.listStagedProductErrandsReturnsOnCall = make(map[int]struct {
			result1 api.ErrandsListOutput
			result2 error
		})
	}
	fake.listStagedProductErrandsReturnsOnCall[i] = struct {
		result1 api.ErrandsListOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListStagedProductJobs(arg1 string) (map[string]string, error) {
	fake.listStagedProductJobsMutex.Lock()
	ret, specificReturn := fake.listStagedProductJobsReturnsOnCall[len(fake.listStagedProductJobsArgsForCall)]
	fake.listStagedProductJobsArgsForCall = append(fake.listStagedProductJobsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ListStagedProductJobs", []interface{}{arg1})
	fake.listStagedProductJobsMutex.Unlock()
	if fake.ListStagedProductJobsStub != nil {
		return fake.ListStagedProductJobsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listStagedProductJobsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *StagedConfigService) ListStagedProductJobsCallCount() int {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	return len(fake.listStagedProductJobsArgsForCall)
}

func (fake *StagedConfigService) ListStagedProductJobsCalls(stub func(string) (map[string]string, error)) {
	fake.listStagedProductJobsMutex.Lock()
	defer fake.listStagedProductJobsMutex.Unlock()
	fake.ListStagedProductJobsStub = stub
}

func (fake *StagedConfigService) ListStagedProductJobsArgsForCall(i int) string {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	argsForCall := fake.listStagedProductJobsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StagedConfigService) ListStagedProductJobsReturns(result1 map[string]string, result2 error) {
	fake.listStagedProductJobsMutex.Lock()
	defer fake.listStagedProductJobsMutex.Unlock()
	fake.ListStagedProductJobsStub = nil
	fake.listStagedProductJobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListStagedProductJobsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.listStagedProductJobsMutex.Lock()
	defer fake.listStagedProductJobsMutex.Unlock()
	fake.ListStagedProductJobsStub = nil
	if fake.listStagedProductJobsReturnsOnCall == nil {
		fake.listStagedProductJobsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.listStagedProductJobsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StagedConfigService) recordInvocation(key string, args []interface{}) {
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