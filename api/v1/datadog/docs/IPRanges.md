# IPRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**IPPrefixes**](IPPrefixes.md) |  | [optional] 
**Api** | Pointer to [**IPPrefixes**](IPPrefixes.md) |  | [optional] 
**Apm** | Pointer to [**IPPrefixes**](IPPrefixes.md) |  | [optional] 
**Logs** | Pointer to [**IPPrefixes**](IPPrefixes.md) |  | [optional] 
**Modified** | Pointer to **string** | Date when last updated, in the form &#x60;YYYY-MM-DD-hh-mm-ss&#x60;. | [optional] 
**Process** | Pointer to [**IPPrefixes**](IPPrefixes.md) |  | [optional] 
**Synthetics** | Pointer to [**IPPrefixes**](IPPrefixes.md) |  | [optional] 
**Version** | Pointer to **int64** | Version of the IP list. | [optional] 
**Webhooks** | Pointer to [**IPPrefixes**](IPPrefixes.md) |  | [optional] 

## Methods

### NewIPRanges

`func NewIPRanges() *IPRanges`

NewIPRanges instantiates a new IPRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPRangesWithDefaults

`func NewIPRangesWithDefaults() *IPRanges`

NewIPRangesWithDefaults instantiates a new IPRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *IPRanges) GetAgents() IPPrefixes`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *IPRanges) GetAgentsOk() (IPPrefixes, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAgents

`func (o *IPRanges) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgents

`func (o *IPRanges) SetAgents(v IPPrefixes)`

SetAgents gets a reference to the given IPPrefixes and assigns it to the Agents field.

### GetApi

`func (o *IPRanges) GetApi() IPPrefixes`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *IPRanges) GetApiOk() (IPPrefixes, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApi

`func (o *IPRanges) HasApi() bool`

HasApi returns a boolean if a field has been set.

### SetApi

`func (o *IPRanges) SetApi(v IPPrefixes)`

SetApi gets a reference to the given IPPrefixes and assigns it to the Api field.

### GetApm

`func (o *IPRanges) GetApm() IPPrefixes`

GetApm returns the Apm field if non-nil, zero value otherwise.

### GetApmOk

`func (o *IPRanges) GetApmOk() (IPPrefixes, bool)`

GetApmOk returns a tuple with the Apm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApm

`func (o *IPRanges) HasApm() bool`

HasApm returns a boolean if a field has been set.

### SetApm

`func (o *IPRanges) SetApm(v IPPrefixes)`

SetApm gets a reference to the given IPPrefixes and assigns it to the Apm field.

### GetLogs

`func (o *IPRanges) GetLogs() IPPrefixes`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *IPRanges) GetLogsOk() (IPPrefixes, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogs

`func (o *IPRanges) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### SetLogs

`func (o *IPRanges) SetLogs(v IPPrefixes)`

SetLogs gets a reference to the given IPPrefixes and assigns it to the Logs field.

### GetModified

`func (o *IPRanges) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IPRanges) GetModifiedOk() (string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *IPRanges) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *IPRanges) SetModified(v string)`

SetModified gets a reference to the given string and assigns it to the Modified field.

### GetProcess

`func (o *IPRanges) GetProcess() IPPrefixes`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *IPRanges) GetProcessOk() (IPPrefixes, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProcess

`func (o *IPRanges) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### SetProcess

`func (o *IPRanges) SetProcess(v IPPrefixes)`

SetProcess gets a reference to the given IPPrefixes and assigns it to the Process field.

### GetSynthetics

`func (o *IPRanges) GetSynthetics() IPPrefixes`

GetSynthetics returns the Synthetics field if non-nil, zero value otherwise.

### GetSyntheticsOk

`func (o *IPRanges) GetSyntheticsOk() (IPPrefixes, bool)`

GetSyntheticsOk returns a tuple with the Synthetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSynthetics

`func (o *IPRanges) HasSynthetics() bool`

HasSynthetics returns a boolean if a field has been set.

### SetSynthetics

`func (o *IPRanges) SetSynthetics(v IPPrefixes)`

SetSynthetics gets a reference to the given IPPrefixes and assigns it to the Synthetics field.

### GetVersion

`func (o *IPRanges) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IPRanges) GetVersionOk() (int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasVersion

`func (o *IPRanges) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersion

`func (o *IPRanges) SetVersion(v int64)`

SetVersion gets a reference to the given int64 and assigns it to the Version field.

### GetWebhooks

`func (o *IPRanges) GetWebhooks() IPPrefixes`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *IPRanges) GetWebhooksOk() (IPPrefixes, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasWebhooks

`func (o *IPRanges) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

### SetWebhooks

`func (o *IPRanges) SetWebhooks(v IPPrefixes)`

SetWebhooks gets a reference to the given IPPrefixes and assigns it to the Webhooks field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


