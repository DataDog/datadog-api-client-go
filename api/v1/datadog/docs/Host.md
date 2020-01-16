# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** |  | [optional] 
**Apps** | Pointer to **[]string** |  | [optional] 
**AwsName** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsMuted** | Pointer to **bool** |  | [optional] 
**LastReportedTime** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**HostMeta**](Host_meta.md) |  | [optional] 
**Metrics** | Pointer to [**HostMetrics**](Host_metrics.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Sources** | Pointer to **[]string** |  | [optional] 
**TagsBySource** | Pointer to [**map[string][]string**](array.md) |  | [optional] 
**Up** | Pointer to **bool** |  | [optional] 

## Methods

### GetAliases

`func (o *Host) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Host) GetAliasesOk() ([]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAliases

`func (o *Host) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### SetAliases

`func (o *Host) SetAliases(v []string)`

SetAliases gets a reference to the given []string and assigns it to the Aliases field.

### GetApps

`func (o *Host) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *Host) GetAppsOk() ([]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApps

`func (o *Host) HasApps() bool`

HasApps returns a boolean if a field has been set.

### SetApps

`func (o *Host) SetApps(v []string)`

SetApps gets a reference to the given []string and assigns it to the Apps field.

### GetAwsName

`func (o *Host) GetAwsName() string`

GetAwsName returns the AwsName field if non-nil, zero value otherwise.

### GetAwsNameOk

`func (o *Host) GetAwsNameOk() (string, bool)`

GetAwsNameOk returns a tuple with the AwsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAwsName

`func (o *Host) HasAwsName() bool`

HasAwsName returns a boolean if a field has been set.

### SetAwsName

`func (o *Host) SetAwsName(v string)`

SetAwsName gets a reference to the given string and assigns it to the AwsName field.

### GetHostName

`func (o *Host) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *Host) GetHostNameOk() (string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostName

`func (o *Host) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostName

`func (o *Host) SetHostName(v string)`

SetHostName gets a reference to the given string and assigns it to the HostName field.

### GetId

`func (o *Host) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Host) GetIdOk() (int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Host) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Host) SetId(v int64)`

SetId gets a reference to the given int64 and assigns it to the Id field.

### GetIsMuted

`func (o *Host) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *Host) GetIsMutedOk() (bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsMuted

`func (o *Host) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### SetIsMuted

`func (o *Host) SetIsMuted(v bool)`

SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.

### GetLastReportedTime

`func (o *Host) GetLastReportedTime() int64`

GetLastReportedTime returns the LastReportedTime field if non-nil, zero value otherwise.

### GetLastReportedTimeOk

`func (o *Host) GetLastReportedTimeOk() (int64, bool)`

GetLastReportedTimeOk returns a tuple with the LastReportedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLastReportedTime

`func (o *Host) HasLastReportedTime() bool`

HasLastReportedTime returns a boolean if a field has been set.

### SetLastReportedTime

`func (o *Host) SetLastReportedTime(v int64)`

SetLastReportedTime gets a reference to the given int64 and assigns it to the LastReportedTime field.

### GetMeta

`func (o *Host) GetMeta() HostMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Host) GetMetaOk() (HostMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMeta

`func (o *Host) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMeta

`func (o *Host) SetMeta(v HostMeta)`

SetMeta gets a reference to the given HostMeta and assigns it to the Meta field.

### GetMetrics

`func (o *Host) GetMetrics() HostMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Host) GetMetricsOk() (HostMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMetrics

`func (o *Host) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetrics

`func (o *Host) SetMetrics(v HostMetrics)`

SetMetrics gets a reference to the given HostMetrics and assigns it to the Metrics field.

### GetName

`func (o *Host) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Host) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Host) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Host) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetSources

`func (o *Host) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Host) GetSourcesOk() ([]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSources

`func (o *Host) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSources

`func (o *Host) SetSources(v []string)`

SetSources gets a reference to the given []string and assigns it to the Sources field.

### GetTagsBySource

`func (o *Host) GetTagsBySource() map[string][]string`

GetTagsBySource returns the TagsBySource field if non-nil, zero value otherwise.

### GetTagsBySourceOk

`func (o *Host) GetTagsBySourceOk() (map[string][]string, bool)`

GetTagsBySourceOk returns a tuple with the TagsBySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTagsBySource

`func (o *Host) HasTagsBySource() bool`

HasTagsBySource returns a boolean if a field has been set.

### SetTagsBySource

`func (o *Host) SetTagsBySource(v map[string][]string)`

SetTagsBySource gets a reference to the given map[string][]string and assigns it to the TagsBySource field.

### GetUp

`func (o *Host) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *Host) GetUpOk() (bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUp

`func (o *Host) HasUp() bool`

HasUp returns a boolean if a field has been set.

### SetUp

`func (o *Host) SetUp(v bool)`

SetUp gets a reference to the given bool and assigns it to the Up field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


