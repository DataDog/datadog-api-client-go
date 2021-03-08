# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Host aliases collected by Datadog. | [optional] 
**Apps** | Pointer to **[]string** | The Datadog integrations reporting metrics for the host. | [optional] 
**AwsName** | Pointer to **string** | AWS name of your host. | [optional] 
**HostName** | Pointer to **string** | The host name. | [optional] 
**Id** | Pointer to **int64** | The host ID. | [optional] 
**IsMuted** | Pointer to **bool** | If a host is muted or unmuted. | [optional] 
**LastReportedTime** | Pointer to **int64** | Last time the host reported a metric data point. | [optional] 
**Meta** | Pointer to [**HostMeta**](HostMeta.md) |  | [optional] 
**Metrics** | Pointer to [**HostMetrics**](HostMetrics.md) |  | [optional] 
**MuteTimeout** | Pointer to **int64** | Timeout of the mute applied to your host. | [optional] 
**Name** | Pointer to **string** | The host name. | [optional] 
**Sources** | Pointer to **[]string** | Source or cloud provider associated with your host. | [optional] 
**TagsBySource** | Pointer to **map[string][]string** | List of tags for each source (AWS, Datadog Agent, Chef..). | [optional] 
**Up** | Pointer to **bool** | Displays UP when the expected metrics are received and displays &#x60;???&#x60; if no metrics are received. | [optional] 

## Methods

### NewHost

`func NewHost() *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *Host) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Host) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Host) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Host) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetApps

`func (o *Host) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *Host) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *Host) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *Host) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetAwsName

`func (o *Host) GetAwsName() string`

GetAwsName returns the AwsName field if non-nil, zero value otherwise.

### GetAwsNameOk

`func (o *Host) GetAwsNameOk() (*string, bool)`

GetAwsNameOk returns a tuple with the AwsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsName

`func (o *Host) SetAwsName(v string)`

SetAwsName sets AwsName field to given value.

### HasAwsName

`func (o *Host) HasAwsName() bool`

HasAwsName returns a boolean if a field has been set.

### GetHostName

`func (o *Host) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *Host) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *Host) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *Host) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetId

`func (o *Host) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Host) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Host) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMuted

`func (o *Host) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *Host) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *Host) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *Host) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetLastReportedTime

`func (o *Host) GetLastReportedTime() int64`

GetLastReportedTime returns the LastReportedTime field if non-nil, zero value otherwise.

### GetLastReportedTimeOk

`func (o *Host) GetLastReportedTimeOk() (*int64, bool)`

GetLastReportedTimeOk returns a tuple with the LastReportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedTime

`func (o *Host) SetLastReportedTime(v int64)`

SetLastReportedTime sets LastReportedTime field to given value.

### HasLastReportedTime

`func (o *Host) HasLastReportedTime() bool`

HasLastReportedTime returns a boolean if a field has been set.

### GetMeta

`func (o *Host) GetMeta() HostMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Host) GetMetaOk() (*HostMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Host) SetMeta(v HostMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Host) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetMetrics

`func (o *Host) GetMetrics() HostMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Host) GetMetricsOk() (*HostMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Host) SetMetrics(v HostMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Host) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMuteTimeout

`func (o *Host) GetMuteTimeout() int64`

GetMuteTimeout returns the MuteTimeout field if non-nil, zero value otherwise.

### GetMuteTimeoutOk

`func (o *Host) GetMuteTimeoutOk() (*int64, bool)`

GetMuteTimeoutOk returns a tuple with the MuteTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteTimeout

`func (o *Host) SetMuteTimeout(v int64)`

SetMuteTimeout sets MuteTimeout field to given value.

### HasMuteTimeout

`func (o *Host) HasMuteTimeout() bool`

HasMuteTimeout returns a boolean if a field has been set.

### GetName

`func (o *Host) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Host) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Host) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Host) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSources

`func (o *Host) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Host) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Host) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Host) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTagsBySource

`func (o *Host) GetTagsBySource() map[string][]string`

GetTagsBySource returns the TagsBySource field if non-nil, zero value otherwise.

### GetTagsBySourceOk

`func (o *Host) GetTagsBySourceOk() (*map[string][]string, bool)`

GetTagsBySourceOk returns a tuple with the TagsBySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsBySource

`func (o *Host) SetTagsBySource(v map[string][]string)`

SetTagsBySource sets TagsBySource field to given value.

### HasTagsBySource

`func (o *Host) HasTagsBySource() bool`

HasTagsBySource returns a boolean if a field has been set.

### GetUp

`func (o *Host) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *Host) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *Host) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *Host) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


