# ApiKey

## Properties

| Name          | Type                  | Description                                   | Notes                 |
| ------------- | --------------------- | --------------------------------------------- | --------------------- |
| **Created**   | Pointer to **string** | Date of creation of the API key.              | [optional] [readonly] |
| **CreatedBy** | Pointer to **string** | Datadog user handle that created the API key. | [optional] [readonly] |
| **Key**       | Pointer to **string** | API key.                                      | [optional] [readonly] |
| **Name**      | Pointer to **string** | Name of your API key.                         | [optional]            |

## Methods

### NewApiKey

`func NewApiKey() *ApiKey`

NewApiKey instantiates a new ApiKey object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreated

`func (o *ApiKey) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ApiKey) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ApiKey) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ApiKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ApiKey) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApiKey) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApiKey) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ApiKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetKey

`func (o *ApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *ApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
