# IpRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**IpPrefixes**](IPPrefixes.md) |  | [optional] 
**Api** | Pointer to [**IpPrefixes**](IPPrefixes.md) |  | [optional] 
**Apm** | Pointer to [**IpPrefixes**](IPPrefixes.md) |  | [optional] 
**Logs** | Pointer to [**IpPrefixes**](IPPrefixes.md) |  | [optional] 
**Modified** | Pointer to **string** | Date when last updated, in the form &#x60;YYYY-MM-DD-hh-mm-ss&#x60; | [optional] 
**Process** | Pointer to [**IpPrefixes**](IPPrefixes.md) |  | [optional] 
**Synthetics** | Pointer to [**IpPrefixes**](IPPrefixes.md) |  | [optional] 
**Version** | Pointer to **int64** | Version of the IP list | [optional] 
**Webhooks** | Pointer to [**IpPrefixes**](IPPrefixes.md) |  | [optional] 

## Methods

### GetAgents

`func (o *IpRanges) GetAgents() IpPrefixes`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *IpRanges) GetAgentsOk() (IpPrefixes, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgents

`func (o *IpRanges) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgents

`func (o *IpRanges) SetAgents(v IpPrefixes)`

SetAgents gets a reference to the given IpPrefixes and assigns it to the Agents field.

### GetApi

`func (o *IpRanges) GetApi() IpPrefixes`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *IpRanges) GetApiOk() (IpPrefixes, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApi

`func (o *IpRanges) HasApi() bool`

HasApi returns a boolean if a field has been set.

### SetApi

`func (o *IpRanges) SetApi(v IpPrefixes)`

SetApi gets a reference to the given IpPrefixes and assigns it to the Api field.

### GetApm

`func (o *IpRanges) GetApm() IpPrefixes`

GetApm returns the Apm field if non-nil, zero value otherwise.

### GetApmOk

`func (o *IpRanges) GetApmOk() (IpPrefixes, bool)`

GetApmOk returns a tuple with the Apm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApm

`func (o *IpRanges) HasApm() bool`

HasApm returns a boolean if a field has been set.

### SetApm

`func (o *IpRanges) SetApm(v IpPrefixes)`

SetApm gets a reference to the given IpPrefixes and assigns it to the Apm field.

### GetLogs

`func (o *IpRanges) GetLogs() IpPrefixes`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *IpRanges) GetLogsOk() (IpPrefixes, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogs

`func (o *IpRanges) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogs

`func (o *IpRanges) SetLogs(v IpPrefixes)`

SetLogs gets a reference to the given IpPrefixes and assigns it to the Logs field.

### GetModified

`func (o *IpRanges) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IpRanges) GetModifiedOk() (string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *IpRanges) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *IpRanges) SetModified(v string)`

SetModified gets a reference to the given string and assigns it to the Modified field.

### GetProcess

`func (o *IpRanges) GetProcess() IpPrefixes`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *IpRanges) GetProcessOk() (IpPrefixes, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcess

`func (o *IpRanges) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### SetProcess

`func (o *IpRanges) SetProcess(v IpPrefixes)`

SetProcess gets a reference to the given IpPrefixes and assigns it to the Process field.

### GetSynthetics

`func (o *IpRanges) GetSynthetics() IpPrefixes`

GetSynthetics returns the Synthetics field if non-nil, zero value otherwise.

### GetSyntheticsOk

`func (o *IpRanges) GetSyntheticsOk() (IpPrefixes, bool)`

GetSyntheticsOk returns a tuple with the Synthetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSynthetics

`func (o *IpRanges) HasSynthetics() bool`

HasSynthetics returns a boolean if a field has been set.

### SetSynthetics

`func (o *IpRanges) SetSynthetics(v IpPrefixes)`

SetSynthetics gets a reference to the given IpPrefixes and assigns it to the Synthetics field.

### GetVersion

`func (o *IpRanges) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IpRanges) GetVersionOk() (int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *IpRanges) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *IpRanges) SetVersion(v int64)`

SetVersion gets a reference to the given int64 and assigns it to the Version field.

### GetWebhooks

`func (o *IpRanges) GetWebhooks() IpPrefixes`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *IpRanges) GetWebhooksOk() (IpPrefixes, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhooks

`func (o *IpRanges) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### SetWebhooks

`func (o *IpRanges) SetWebhooks(v IpPrefixes)`

SetWebhooks gets a reference to the given IpPrefixes and assigns it to the Webhooks field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


