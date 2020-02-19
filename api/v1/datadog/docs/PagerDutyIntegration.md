# PagerDutyIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | Pointer to **string** |  | [optional] 
**Schedules** | Pointer to **[]string** | The array of your schedule URLs. | [optional] 
**Services** | Pointer to [**[]PagerDutyService**](PagerDutyService.md) | The array of PagerDuty service objects. | [optional] 
**Subdomain** | Pointer to **string** |  | [optional] 

## Methods

### NewPagerDutyIntegration

`func NewPagerDutyIntegration() *PagerDutyIntegration`

NewPagerDutyIntegration instantiates a new PagerDutyIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagerDutyIntegrationWithDefaults

`func NewPagerDutyIntegrationWithDefaults() *PagerDutyIntegration`

NewPagerDutyIntegrationWithDefaults instantiates a new PagerDutyIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *PagerDutyIntegration) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *PagerDutyIntegration) GetApiTokenOk() (string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApiToken

`func (o *PagerDutyIntegration) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### SetApiToken

`func (o *PagerDutyIntegration) SetApiToken(v string)`

SetApiToken gets a reference to the given string and assigns it to the ApiToken field.

### GetSchedules

`func (o *PagerDutyIntegration) GetSchedules() []string`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *PagerDutyIntegration) GetSchedulesOk() ([]string, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSchedules

`func (o *PagerDutyIntegration) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### SetSchedules

`func (o *PagerDutyIntegration) SetSchedules(v []string)`

SetSchedules gets a reference to the given []string and assigns it to the Schedules field.

### GetServices

`func (o *PagerDutyIntegration) GetServices() []PagerDutyService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *PagerDutyIntegration) GetServicesOk() ([]PagerDutyService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServices

`func (o *PagerDutyIntegration) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServices

`func (o *PagerDutyIntegration) SetServices(v []PagerDutyService)`

SetServices gets a reference to the given []PagerDutyService and assigns it to the Services field.

### GetSubdomain

`func (o *PagerDutyIntegration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *PagerDutyIntegration) GetSubdomainOk() (string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubdomain

`func (o *PagerDutyIntegration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### SetSubdomain

`func (o *PagerDutyIntegration) SetSubdomain(v string)`

SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


