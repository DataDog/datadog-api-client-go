# SecurityFilterExclusionFilter

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Name** | **string** | Exclusion filter name. | 
**Query** | **string** | Exclusion filter query. Logs that match this query are excluded from the security filter. | 

## Methods

### NewSecurityFilterExclusionFilter

`func NewSecurityFilterExclusionFilter(name string, query string) *SecurityFilterExclusionFilter`

NewSecurityFilterExclusionFilter instantiates a new SecurityFilterExclusionFilter object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityFilterExclusionFilterWithDefaults

`func NewSecurityFilterExclusionFilterWithDefaults() *SecurityFilterExclusionFilter`

NewSecurityFilterExclusionFilterWithDefaults instantiates a new SecurityFilterExclusionFilter object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetName

`func (o *SecurityFilterExclusionFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityFilterExclusionFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityFilterExclusionFilter) SetName(v string)`

SetName sets Name field to given value.


### GetQuery

`func (o *SecurityFilterExclusionFilter) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SecurityFilterExclusionFilter) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SecurityFilterExclusionFilter) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


