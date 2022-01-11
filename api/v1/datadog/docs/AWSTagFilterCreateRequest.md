# AWSTagFilterCreateRequest

## Properties

| Name             | Type                                           | Description                         | Notes      |
| ---------------- | ---------------------------------------------- | ----------------------------------- | ---------- |
| **AccountId**    | Pointer to **string**                          | Your AWS Account ID without dashes. | [optional] |
| **Namespace**    | Pointer to [**AWSNamespace**](AWSNamespace.md) |                                     | [optional] |
| **TagFilterStr** | Pointer to **string**                          | The tag filter string.              | [optional] |

## Methods

### NewAWSTagFilterCreateRequest

`func NewAWSTagFilterCreateRequest() *AWSTagFilterCreateRequest`

NewAWSTagFilterCreateRequest instantiates a new AWSTagFilterCreateRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAWSTagFilterCreateRequestWithDefaults

`func NewAWSTagFilterCreateRequestWithDefaults() *AWSTagFilterCreateRequest`

NewAWSTagFilterCreateRequestWithDefaults instantiates a new AWSTagFilterCreateRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAccountId

`func (o *AWSTagFilterCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AWSTagFilterCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AWSTagFilterCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AWSTagFilterCreateRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetNamespace

`func (o *AWSTagFilterCreateRequest) GetNamespace() AWSNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AWSTagFilterCreateRequest) GetNamespaceOk() (*AWSNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AWSTagFilterCreateRequest) SetNamespace(v AWSNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AWSTagFilterCreateRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTagFilterStr

`func (o *AWSTagFilterCreateRequest) GetTagFilterStr() string`

GetTagFilterStr returns the TagFilterStr field if non-nil, zero value otherwise.

### GetTagFilterStrOk

`func (o *AWSTagFilterCreateRequest) GetTagFilterStrOk() (*string, bool)`

GetTagFilterStrOk returns a tuple with the TagFilterStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterStr

`func (o *AWSTagFilterCreateRequest) SetTagFilterStr(v string)`

SetTagFilterStr sets TagFilterStr field to given value.

### HasTagFilterStr

`func (o *AWSTagFilterCreateRequest) HasTagFilterStr() bool`

HasTagFilterStr returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
