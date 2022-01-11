# AWSTagFilter

## Properties

| Name             | Type                                           | Description            | Notes      |
| ---------------- | ---------------------------------------------- | ---------------------- | ---------- |
| **Namespace**    | Pointer to [**AWSNamespace**](AWSNamespace.md) |                        | [optional] |
| **TagFilterStr** | Pointer to **string**                          | The tag filter string. | [optional] |

## Methods

### NewAWSTagFilter

`func NewAWSTagFilter() *AWSTagFilter`

NewAWSTagFilter instantiates a new AWSTagFilter object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAWSTagFilterWithDefaults

`func NewAWSTagFilterWithDefaults() *AWSTagFilter`

NewAWSTagFilterWithDefaults instantiates a new AWSTagFilter object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetNamespace

`func (o *AWSTagFilter) GetNamespace() AWSNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AWSTagFilter) GetNamespaceOk() (*AWSNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AWSTagFilter) SetNamespace(v AWSNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AWSTagFilter) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTagFilterStr

`func (o *AWSTagFilter) GetTagFilterStr() string`

GetTagFilterStr returns the TagFilterStr field if non-nil, zero value otherwise.

### GetTagFilterStrOk

`func (o *AWSTagFilter) GetTagFilterStrOk() (*string, bool)`

GetTagFilterStrOk returns a tuple with the TagFilterStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagFilterStr

`func (o *AWSTagFilter) SetTagFilterStr(v string)`

SetTagFilterStr sets TagFilterStr field to given value.

### HasTagFilterStr

`func (o *AWSTagFilter) HasTagFilterStr() bool`

HasTagFilterStr returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
