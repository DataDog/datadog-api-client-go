# SyntheticsDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int64** |  | 
**Id** | Pointer to [**SyntheticsDeviceID**](SyntheticsDeviceID.md) |  | 
**IsMobile** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | 
**Width** | Pointer to **int64** |  | 

## Methods

### NewSyntheticsDevice

`func NewSyntheticsDevice(height int64, id SyntheticsDeviceID, name string, width int64, ) *SyntheticsDevice`

NewSyntheticsDevice instantiates a new SyntheticsDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsDeviceWithDefaults

`func NewSyntheticsDeviceWithDefaults() *SyntheticsDevice`

NewSyntheticsDeviceWithDefaults instantiates a new SyntheticsDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *SyntheticsDevice) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SyntheticsDevice) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SyntheticsDevice) SetHeight(v int64)`

SetHeight sets Height field to given value.


### GetId

`func (o *SyntheticsDevice) GetId() SyntheticsDeviceID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyntheticsDevice) GetIdOk() (*SyntheticsDeviceID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyntheticsDevice) SetId(v SyntheticsDeviceID)`

SetId sets Id field to given value.


### GetIsMobile

`func (o *SyntheticsDevice) GetIsMobile() bool`

GetIsMobile returns the IsMobile field if non-nil, zero value otherwise.

### GetIsMobileOk

`func (o *SyntheticsDevice) GetIsMobileOk() (*bool, bool)`

GetIsMobileOk returns a tuple with the IsMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMobile

`func (o *SyntheticsDevice) SetIsMobile(v bool)`

SetIsMobile sets IsMobile field to given value.

### HasIsMobile

`func (o *SyntheticsDevice) HasIsMobile() bool`

HasIsMobile returns a boolean if a field has been set.

### GetName

`func (o *SyntheticsDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticsDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticsDevice) SetName(v string)`

SetName sets Name field to given value.


### GetWidth

`func (o *SyntheticsDevice) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SyntheticsDevice) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SyntheticsDevice) SetWidth(v int64)`

SetWidth sets Width field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


