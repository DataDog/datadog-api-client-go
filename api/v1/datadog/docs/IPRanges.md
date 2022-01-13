# IPRanges

## Properties

| Name           | Type                                                           | Description                                                          | Notes      |
| -------------- | -------------------------------------------------------------- | -------------------------------------------------------------------- | ---------- |
| **Agents**     | Pointer to [**IPPrefixesAgents**](IPPrefixesAgents.md)         |                                                                      | [optional] |
| **Api**        | Pointer to [**IPPrefixesAPI**](IPPrefixesAPI.md)               |                                                                      | [optional] |
| **Apm**        | Pointer to [**IPPrefixesAPM**](IPPrefixesAPM.md)               |                                                                      | [optional] |
| **Logs**       | Pointer to [**IPPrefixesLogs**](IPPrefixesLogs.md)             |                                                                      | [optional] |
| **Modified**   | Pointer to **string**                                          | Date when last updated, in the form &#x60;YYYY-MM-DD-hh-mm-ss&#x60;. | [optional] |
| **Process**    | Pointer to [**IPPrefixesProcess**](IPPrefixesProcess.md)       |                                                                      | [optional] |
| **Synthetics** | Pointer to [**IPPrefixesSynthetics**](IPPrefixesSynthetics.md) |                                                                      | [optional] |
| **Version**    | Pointer to **int64**                                           | Version of the IP list.                                              | [optional] |
| **Webhooks**   | Pointer to [**IPPrefixesWebhooks**](IPPrefixesWebhooks.md)     |                                                                      | [optional] |

## Methods

### NewIPRanges

`func NewIPRanges() *IPRanges`

NewIPRanges instantiates a new IPRanges object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewIPRangesWithDefaults

`func NewIPRangesWithDefaults() *IPRanges`

NewIPRangesWithDefaults instantiates a new IPRanges object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAgents

`func (o *IPRanges) GetAgents() IPPrefixesAgents`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *IPRanges) GetAgentsOk() (*IPPrefixesAgents, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *IPRanges) SetAgents(v IPPrefixesAgents)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *IPRanges) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetApi

`func (o *IPRanges) GetApi() IPPrefixesAPI`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *IPRanges) GetApiOk() (*IPPrefixesAPI, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *IPRanges) SetApi(v IPPrefixesAPI)`

SetApi sets Api field to given value.

### HasApi

`func (o *IPRanges) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetApm

`func (o *IPRanges) GetApm() IPPrefixesAPM`

GetApm returns the Apm field if non-nil, zero value otherwise.

### GetApmOk

`func (o *IPRanges) GetApmOk() (*IPPrefixesAPM, bool)`

GetApmOk returns a tuple with the Apm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApm

`func (o *IPRanges) SetApm(v IPPrefixesAPM)`

SetApm sets Apm field to given value.

### HasApm

`func (o *IPRanges) HasApm() bool`

HasApm returns a boolean if a field has been set.

### GetLogs

`func (o *IPRanges) GetLogs() IPPrefixesLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *IPRanges) GetLogsOk() (*IPPrefixesLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *IPRanges) SetLogs(v IPPrefixesLogs)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *IPRanges) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetModified

`func (o *IPRanges) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *IPRanges) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *IPRanges) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *IPRanges) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetProcess

`func (o *IPRanges) GetProcess() IPPrefixesProcess`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *IPRanges) GetProcessOk() (*IPPrefixesProcess, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *IPRanges) SetProcess(v IPPrefixesProcess)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *IPRanges) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetSynthetics

`func (o *IPRanges) GetSynthetics() IPPrefixesSynthetics`

GetSynthetics returns the Synthetics field if non-nil, zero value otherwise.

### GetSyntheticsOk

`func (o *IPRanges) GetSyntheticsOk() (*IPPrefixesSynthetics, bool)`

GetSyntheticsOk returns a tuple with the Synthetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynthetics

`func (o *IPRanges) SetSynthetics(v IPPrefixesSynthetics)`

SetSynthetics sets Synthetics field to given value.

### HasSynthetics

`func (o *IPRanges) HasSynthetics() bool`

HasSynthetics returns a boolean if a field has been set.

### GetVersion

`func (o *IPRanges) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IPRanges) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IPRanges) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IPRanges) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWebhooks

`func (o *IPRanges) GetWebhooks() IPPrefixesWebhooks`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *IPRanges) GetWebhooksOk() (*IPPrefixesWebhooks, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *IPRanges) SetWebhooks(v IPPrefixesWebhooks)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *IPRanges) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
