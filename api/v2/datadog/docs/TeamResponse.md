# TeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TeamResponseData**](TeamResponseData.md) |  | 
**Included** | Pointer to [**[]TeamIncludedItems**](TeamIncludedItems.md) | Included objects from relationships. | [optional] [readonly] 

## Methods

### NewTeamResponse

`func NewTeamResponse(data TeamResponseData, ) *TeamResponse`

NewTeamResponse instantiates a new TeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamResponseWithDefaults

`func NewTeamResponseWithDefaults() *TeamResponse`

NewTeamResponseWithDefaults instantiates a new TeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TeamResponse) GetData() TeamResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TeamResponse) GetDataOk() (*TeamResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TeamResponse) SetData(v TeamResponseData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *TeamResponse) GetIncluded() []TeamIncludedItems`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TeamResponse) GetIncludedOk() (*[]TeamIncludedItems, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TeamResponse) SetIncluded(v []TeamIncludedItems)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TeamResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


