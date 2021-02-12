# UsageBillableSummaryKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApmHostSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**ApmHostTop99p** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**ApmTraceSearchSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**FargateContainerAverage** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**InfraContainerSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**InfraHostSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**InfraHostTop99p** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**IotTop99p** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**LambdaFunctionAverage** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**LogsIndexedSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**LogsIngestedSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**NpmFlowSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**NpmHostTop99p** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**RumSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**ServerlessInvocationSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**SiemSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**SyntheticsApiTestsSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**SyntheticsBrowserChecksSum** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 
**TimeseriesAverage** | Pointer to [**UsageBillableSummaryBody**](UsageBillableSummaryBody.md) |  | [optional] 

## Methods

### NewUsageBillableSummaryKeys

`func NewUsageBillableSummaryKeys() *UsageBillableSummaryKeys`

NewUsageBillableSummaryKeys instantiates a new UsageBillableSummaryKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageBillableSummaryKeysWithDefaults

`func NewUsageBillableSummaryKeysWithDefaults() *UsageBillableSummaryKeys`

NewUsageBillableSummaryKeysWithDefaults instantiates a new UsageBillableSummaryKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApmHostSum

`func (o *UsageBillableSummaryKeys) GetApmHostSum() UsageBillableSummaryBody`

GetApmHostSum returns the ApmHostSum field if non-nil, zero value otherwise.

### GetApmHostSumOk

`func (o *UsageBillableSummaryKeys) GetApmHostSumOk() (*UsageBillableSummaryBody, bool)`

GetApmHostSumOk returns a tuple with the ApmHostSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostSum

`func (o *UsageBillableSummaryKeys) SetApmHostSum(v UsageBillableSummaryBody)`

SetApmHostSum sets ApmHostSum field to given value.

### HasApmHostSum

`func (o *UsageBillableSummaryKeys) HasApmHostSum() bool`

HasApmHostSum returns a boolean if a field has been set.

### GetApmHostTop99p

`func (o *UsageBillableSummaryKeys) GetApmHostTop99p() UsageBillableSummaryBody`

GetApmHostTop99p returns the ApmHostTop99p field if non-nil, zero value otherwise.

### GetApmHostTop99pOk

`func (o *UsageBillableSummaryKeys) GetApmHostTop99pOk() (*UsageBillableSummaryBody, bool)`

GetApmHostTop99pOk returns a tuple with the ApmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmHostTop99p

`func (o *UsageBillableSummaryKeys) SetApmHostTop99p(v UsageBillableSummaryBody)`

SetApmHostTop99p sets ApmHostTop99p field to given value.

### HasApmHostTop99p

`func (o *UsageBillableSummaryKeys) HasApmHostTop99p() bool`

HasApmHostTop99p returns a boolean if a field has been set.

### GetApmTraceSearchSum

`func (o *UsageBillableSummaryKeys) GetApmTraceSearchSum() UsageBillableSummaryBody`

GetApmTraceSearchSum returns the ApmTraceSearchSum field if non-nil, zero value otherwise.

### GetApmTraceSearchSumOk

`func (o *UsageBillableSummaryKeys) GetApmTraceSearchSumOk() (*UsageBillableSummaryBody, bool)`

GetApmTraceSearchSumOk returns a tuple with the ApmTraceSearchSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApmTraceSearchSum

`func (o *UsageBillableSummaryKeys) SetApmTraceSearchSum(v UsageBillableSummaryBody)`

SetApmTraceSearchSum sets ApmTraceSearchSum field to given value.

### HasApmTraceSearchSum

`func (o *UsageBillableSummaryKeys) HasApmTraceSearchSum() bool`

HasApmTraceSearchSum returns a boolean if a field has been set.

### GetFargateContainerAverage

`func (o *UsageBillableSummaryKeys) GetFargateContainerAverage() UsageBillableSummaryBody`

GetFargateContainerAverage returns the FargateContainerAverage field if non-nil, zero value otherwise.

### GetFargateContainerAverageOk

`func (o *UsageBillableSummaryKeys) GetFargateContainerAverageOk() (*UsageBillableSummaryBody, bool)`

GetFargateContainerAverageOk returns a tuple with the FargateContainerAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFargateContainerAverage

`func (o *UsageBillableSummaryKeys) SetFargateContainerAverage(v UsageBillableSummaryBody)`

SetFargateContainerAverage sets FargateContainerAverage field to given value.

### HasFargateContainerAverage

`func (o *UsageBillableSummaryKeys) HasFargateContainerAverage() bool`

HasFargateContainerAverage returns a boolean if a field has been set.

### GetInfraContainerSum

`func (o *UsageBillableSummaryKeys) GetInfraContainerSum() UsageBillableSummaryBody`

GetInfraContainerSum returns the InfraContainerSum field if non-nil, zero value otherwise.

### GetInfraContainerSumOk

`func (o *UsageBillableSummaryKeys) GetInfraContainerSumOk() (*UsageBillableSummaryBody, bool)`

GetInfraContainerSumOk returns a tuple with the InfraContainerSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraContainerSum

`func (o *UsageBillableSummaryKeys) SetInfraContainerSum(v UsageBillableSummaryBody)`

SetInfraContainerSum sets InfraContainerSum field to given value.

### HasInfraContainerSum

`func (o *UsageBillableSummaryKeys) HasInfraContainerSum() bool`

HasInfraContainerSum returns a boolean if a field has been set.

### GetInfraHostSum

`func (o *UsageBillableSummaryKeys) GetInfraHostSum() UsageBillableSummaryBody`

GetInfraHostSum returns the InfraHostSum field if non-nil, zero value otherwise.

### GetInfraHostSumOk

`func (o *UsageBillableSummaryKeys) GetInfraHostSumOk() (*UsageBillableSummaryBody, bool)`

GetInfraHostSumOk returns a tuple with the InfraHostSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostSum

`func (o *UsageBillableSummaryKeys) SetInfraHostSum(v UsageBillableSummaryBody)`

SetInfraHostSum sets InfraHostSum field to given value.

### HasInfraHostSum

`func (o *UsageBillableSummaryKeys) HasInfraHostSum() bool`

HasInfraHostSum returns a boolean if a field has been set.

### GetInfraHostTop99p

`func (o *UsageBillableSummaryKeys) GetInfraHostTop99p() UsageBillableSummaryBody`

GetInfraHostTop99p returns the InfraHostTop99p field if non-nil, zero value otherwise.

### GetInfraHostTop99pOk

`func (o *UsageBillableSummaryKeys) GetInfraHostTop99pOk() (*UsageBillableSummaryBody, bool)`

GetInfraHostTop99pOk returns a tuple with the InfraHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraHostTop99p

`func (o *UsageBillableSummaryKeys) SetInfraHostTop99p(v UsageBillableSummaryBody)`

SetInfraHostTop99p sets InfraHostTop99p field to given value.

### HasInfraHostTop99p

`func (o *UsageBillableSummaryKeys) HasInfraHostTop99p() bool`

HasInfraHostTop99p returns a boolean if a field has been set.

### GetIotTop99p

`func (o *UsageBillableSummaryKeys) GetIotTop99p() UsageBillableSummaryBody`

GetIotTop99p returns the IotTop99p field if non-nil, zero value otherwise.

### GetIotTop99pOk

`func (o *UsageBillableSummaryKeys) GetIotTop99pOk() (*UsageBillableSummaryBody, bool)`

GetIotTop99pOk returns a tuple with the IotTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIotTop99p

`func (o *UsageBillableSummaryKeys) SetIotTop99p(v UsageBillableSummaryBody)`

SetIotTop99p sets IotTop99p field to given value.

### HasIotTop99p

`func (o *UsageBillableSummaryKeys) HasIotTop99p() bool`

HasIotTop99p returns a boolean if a field has been set.

### GetLambdaFunctionAverage

`func (o *UsageBillableSummaryKeys) GetLambdaFunctionAverage() UsageBillableSummaryBody`

GetLambdaFunctionAverage returns the LambdaFunctionAverage field if non-nil, zero value otherwise.

### GetLambdaFunctionAverageOk

`func (o *UsageBillableSummaryKeys) GetLambdaFunctionAverageOk() (*UsageBillableSummaryBody, bool)`

GetLambdaFunctionAverageOk returns a tuple with the LambdaFunctionAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLambdaFunctionAverage

`func (o *UsageBillableSummaryKeys) SetLambdaFunctionAverage(v UsageBillableSummaryBody)`

SetLambdaFunctionAverage sets LambdaFunctionAverage field to given value.

### HasLambdaFunctionAverage

`func (o *UsageBillableSummaryKeys) HasLambdaFunctionAverage() bool`

HasLambdaFunctionAverage returns a boolean if a field has been set.

### GetLogsIndexedSum

`func (o *UsageBillableSummaryKeys) GetLogsIndexedSum() UsageBillableSummaryBody`

GetLogsIndexedSum returns the LogsIndexedSum field if non-nil, zero value otherwise.

### GetLogsIndexedSumOk

`func (o *UsageBillableSummaryKeys) GetLogsIndexedSumOk() (*UsageBillableSummaryBody, bool)`

GetLogsIndexedSumOk returns a tuple with the LogsIndexedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsIndexedSum

`func (o *UsageBillableSummaryKeys) SetLogsIndexedSum(v UsageBillableSummaryBody)`

SetLogsIndexedSum sets LogsIndexedSum field to given value.

### HasLogsIndexedSum

`func (o *UsageBillableSummaryKeys) HasLogsIndexedSum() bool`

HasLogsIndexedSum returns a boolean if a field has been set.

### GetLogsIngestedSum

`func (o *UsageBillableSummaryKeys) GetLogsIngestedSum() UsageBillableSummaryBody`

GetLogsIngestedSum returns the LogsIngestedSum field if non-nil, zero value otherwise.

### GetLogsIngestedSumOk

`func (o *UsageBillableSummaryKeys) GetLogsIngestedSumOk() (*UsageBillableSummaryBody, bool)`

GetLogsIngestedSumOk returns a tuple with the LogsIngestedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsIngestedSum

`func (o *UsageBillableSummaryKeys) SetLogsIngestedSum(v UsageBillableSummaryBody)`

SetLogsIngestedSum sets LogsIngestedSum field to given value.

### HasLogsIngestedSum

`func (o *UsageBillableSummaryKeys) HasLogsIngestedSum() bool`

HasLogsIngestedSum returns a boolean if a field has been set.

### GetNpmFlowSum

`func (o *UsageBillableSummaryKeys) GetNpmFlowSum() UsageBillableSummaryBody`

GetNpmFlowSum returns the NpmFlowSum field if non-nil, zero value otherwise.

### GetNpmFlowSumOk

`func (o *UsageBillableSummaryKeys) GetNpmFlowSumOk() (*UsageBillableSummaryBody, bool)`

GetNpmFlowSumOk returns a tuple with the NpmFlowSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmFlowSum

`func (o *UsageBillableSummaryKeys) SetNpmFlowSum(v UsageBillableSummaryBody)`

SetNpmFlowSum sets NpmFlowSum field to given value.

### HasNpmFlowSum

`func (o *UsageBillableSummaryKeys) HasNpmFlowSum() bool`

HasNpmFlowSum returns a boolean if a field has been set.

### GetNpmHostTop99p

`func (o *UsageBillableSummaryKeys) GetNpmHostTop99p() UsageBillableSummaryBody`

GetNpmHostTop99p returns the NpmHostTop99p field if non-nil, zero value otherwise.

### GetNpmHostTop99pOk

`func (o *UsageBillableSummaryKeys) GetNpmHostTop99pOk() (*UsageBillableSummaryBody, bool)`

GetNpmHostTop99pOk returns a tuple with the NpmHostTop99p field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmHostTop99p

`func (o *UsageBillableSummaryKeys) SetNpmHostTop99p(v UsageBillableSummaryBody)`

SetNpmHostTop99p sets NpmHostTop99p field to given value.

### HasNpmHostTop99p

`func (o *UsageBillableSummaryKeys) HasNpmHostTop99p() bool`

HasNpmHostTop99p returns a boolean if a field has been set.

### GetRumSum

`func (o *UsageBillableSummaryKeys) GetRumSum() UsageBillableSummaryBody`

GetRumSum returns the RumSum field if non-nil, zero value otherwise.

### GetRumSumOk

`func (o *UsageBillableSummaryKeys) GetRumSumOk() (*UsageBillableSummaryBody, bool)`

GetRumSumOk returns a tuple with the RumSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRumSum

`func (o *UsageBillableSummaryKeys) SetRumSum(v UsageBillableSummaryBody)`

SetRumSum sets RumSum field to given value.

### HasRumSum

`func (o *UsageBillableSummaryKeys) HasRumSum() bool`

HasRumSum returns a boolean if a field has been set.

### GetServerlessInvocationSum

`func (o *UsageBillableSummaryKeys) GetServerlessInvocationSum() UsageBillableSummaryBody`

GetServerlessInvocationSum returns the ServerlessInvocationSum field if non-nil, zero value otherwise.

### GetServerlessInvocationSumOk

`func (o *UsageBillableSummaryKeys) GetServerlessInvocationSumOk() (*UsageBillableSummaryBody, bool)`

GetServerlessInvocationSumOk returns a tuple with the ServerlessInvocationSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessInvocationSum

`func (o *UsageBillableSummaryKeys) SetServerlessInvocationSum(v UsageBillableSummaryBody)`

SetServerlessInvocationSum sets ServerlessInvocationSum field to given value.

### HasServerlessInvocationSum

`func (o *UsageBillableSummaryKeys) HasServerlessInvocationSum() bool`

HasServerlessInvocationSum returns a boolean if a field has been set.

### GetSiemSum

`func (o *UsageBillableSummaryKeys) GetSiemSum() UsageBillableSummaryBody`

GetSiemSum returns the SiemSum field if non-nil, zero value otherwise.

### GetSiemSumOk

`func (o *UsageBillableSummaryKeys) GetSiemSumOk() (*UsageBillableSummaryBody, bool)`

GetSiemSumOk returns a tuple with the SiemSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiemSum

`func (o *UsageBillableSummaryKeys) SetSiemSum(v UsageBillableSummaryBody)`

SetSiemSum sets SiemSum field to given value.

### HasSiemSum

`func (o *UsageBillableSummaryKeys) HasSiemSum() bool`

HasSiemSum returns a boolean if a field has been set.

### GetSyntheticsApiTestsSum

`func (o *UsageBillableSummaryKeys) GetSyntheticsApiTestsSum() UsageBillableSummaryBody`

GetSyntheticsApiTestsSum returns the SyntheticsApiTestsSum field if non-nil, zero value otherwise.

### GetSyntheticsApiTestsSumOk

`func (o *UsageBillableSummaryKeys) GetSyntheticsApiTestsSumOk() (*UsageBillableSummaryBody, bool)`

GetSyntheticsApiTestsSumOk returns a tuple with the SyntheticsApiTestsSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsApiTestsSum

`func (o *UsageBillableSummaryKeys) SetSyntheticsApiTestsSum(v UsageBillableSummaryBody)`

SetSyntheticsApiTestsSum sets SyntheticsApiTestsSum field to given value.

### HasSyntheticsApiTestsSum

`func (o *UsageBillableSummaryKeys) HasSyntheticsApiTestsSum() bool`

HasSyntheticsApiTestsSum returns a boolean if a field has been set.

### GetSyntheticsBrowserChecksSum

`func (o *UsageBillableSummaryKeys) GetSyntheticsBrowserChecksSum() UsageBillableSummaryBody`

GetSyntheticsBrowserChecksSum returns the SyntheticsBrowserChecksSum field if non-nil, zero value otherwise.

### GetSyntheticsBrowserChecksSumOk

`func (o *UsageBillableSummaryKeys) GetSyntheticsBrowserChecksSumOk() (*UsageBillableSummaryBody, bool)`

GetSyntheticsBrowserChecksSumOk returns a tuple with the SyntheticsBrowserChecksSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticsBrowserChecksSum

`func (o *UsageBillableSummaryKeys) SetSyntheticsBrowserChecksSum(v UsageBillableSummaryBody)`

SetSyntheticsBrowserChecksSum sets SyntheticsBrowserChecksSum field to given value.

### HasSyntheticsBrowserChecksSum

`func (o *UsageBillableSummaryKeys) HasSyntheticsBrowserChecksSum() bool`

HasSyntheticsBrowserChecksSum returns a boolean if a field has been set.

### GetTimeseriesAverage

`func (o *UsageBillableSummaryKeys) GetTimeseriesAverage() UsageBillableSummaryBody`

GetTimeseriesAverage returns the TimeseriesAverage field if non-nil, zero value otherwise.

### GetTimeseriesAverageOk

`func (o *UsageBillableSummaryKeys) GetTimeseriesAverageOk() (*UsageBillableSummaryBody, bool)`

GetTimeseriesAverageOk returns a tuple with the TimeseriesAverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesAverage

`func (o *UsageBillableSummaryKeys) SetTimeseriesAverage(v UsageBillableSummaryBody)`

SetTimeseriesAverage sets TimeseriesAverage field to given value.

### HasTimeseriesAverage

`func (o *UsageBillableSummaryKeys) HasTimeseriesAverage() bool`

HasTimeseriesAverage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


