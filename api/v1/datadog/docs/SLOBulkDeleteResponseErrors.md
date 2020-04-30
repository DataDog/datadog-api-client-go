# SLOBulkDeleteResponseErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the service level objective object associated with this error. | 
**Message** | Pointer to **string** | The error message. | 
**Timeframe** | Pointer to [**SLOErrorTimeframe**](SLOErrorTimeframe.md) |  | 

## Methods

### NewSLOBulkDeleteResponseErrors

`func NewSLOBulkDeleteResponseErrors(id string, message string, timeframe SLOErrorTimeframe, ) *SLOBulkDeleteResponseErrors`

NewSLOBulkDeleteResponseErrors instantiates a new SLOBulkDeleteResponseErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOBulkDeleteResponseErrorsWithDefaults

`func NewSLOBulkDeleteResponseErrorsWithDefaults() *SLOBulkDeleteResponseErrors`

NewSLOBulkDeleteResponseErrorsWithDefaults instantiates a new SLOBulkDeleteResponseErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SLOBulkDeleteResponseErrors) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLOBulkDeleteResponseErrors) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLOBulkDeleteResponseErrors) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *SLOBulkDeleteResponseErrors) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SLOBulkDeleteResponseErrors) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SLOBulkDeleteResponseErrors) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTimeframe

`func (o *SLOBulkDeleteResponseErrors) GetTimeframe() SLOErrorTimeframe`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *SLOBulkDeleteResponseErrors) GetTimeframeOk() (*SLOErrorTimeframe, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *SLOBulkDeleteResponseErrors) SetTimeframe(v SLOErrorTimeframe)`

SetTimeframe sets Timeframe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


