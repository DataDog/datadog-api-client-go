# APIKeyResponse

## Properties

| Name         | Type                                                                         | Description                              | Notes      |
| ------------ | ---------------------------------------------------------------------------- | ---------------------------------------- | ---------- |
| **Data**     | Pointer to [**FullAPIKey**](FullAPIKey.md)                                   |                                          | [optional] |
| **Included** | Pointer to [**[]APIKeyResponseIncludedItem**](APIKeyResponseIncludedItem.md) | Array of objects related to the API key. | [optional] |

## Methods

### NewAPIKeyResponse

`func NewAPIKeyResponse() *APIKeyResponse`

NewAPIKeyResponse instantiates a new APIKeyResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAPIKeyResponseWithDefaults

`func NewAPIKeyResponseWithDefaults() *APIKeyResponse`

NewAPIKeyResponseWithDefaults instantiates a new APIKeyResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetData

`func (o *APIKeyResponse) GetData() FullAPIKey`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *APIKeyResponse) GetDataOk() (*FullAPIKey, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *APIKeyResponse) SetData(v FullAPIKey)`

SetData sets Data field to given value.

### HasData

`func (o *APIKeyResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIncluded

`func (o *APIKeyResponse) GetIncluded() []APIKeyResponseIncludedItem`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *APIKeyResponse) GetIncludedOk() (*[]APIKeyResponseIncludedItem, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *APIKeyResponse) SetIncluded(v []APIKeyResponseIncludedItem)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *APIKeyResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
