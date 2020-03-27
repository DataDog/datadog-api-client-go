# PagerDutyService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceKey** | Pointer to **string** | Your Service name associated service key in Pagerduty. | 
**ServiceName** | Pointer to **string** | Your Service name in PagerDuty. | 

## Methods

### NewPagerDutyService

`func NewPagerDutyService(serviceKey string, serviceName string, ) *PagerDutyService`

NewPagerDutyService instantiates a new PagerDutyService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagerDutyServiceWithDefaults

`func NewPagerDutyServiceWithDefaults() *PagerDutyService`

NewPagerDutyServiceWithDefaults instantiates a new PagerDutyService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceKey

`func (o *PagerDutyService) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *PagerDutyService) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *PagerDutyService) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetServiceName

`func (o *PagerDutyService) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PagerDutyService) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PagerDutyService) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


