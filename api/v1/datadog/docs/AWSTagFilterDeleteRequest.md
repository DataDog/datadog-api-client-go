# AWSTagFilterDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountIdentifier** | Pointer to **string** | The unique identifier of your AWS account. | [optional] 
**Namespace** | Pointer to [**AWSNamespace**](AWSNamespace.md) |  | [optional] 

## Methods

### NewAWSTagFilterDeleteRequest

`func NewAWSTagFilterDeleteRequest() *AWSTagFilterDeleteRequest`

NewAWSTagFilterDeleteRequest instantiates a new AWSTagFilterDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSTagFilterDeleteRequestWithDefaults

`func NewAWSTagFilterDeleteRequestWithDefaults() *AWSTagFilterDeleteRequest`

NewAWSTagFilterDeleteRequestWithDefaults instantiates a new AWSTagFilterDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountIdentifier

`func (o *AWSTagFilterDeleteRequest) GetAwsAccountIdentifier() string`

GetAwsAccountIdentifier returns the AwsAccountIdentifier field if non-nil, zero value otherwise.

### GetAwsAccountIdentifierOk

`func (o *AWSTagFilterDeleteRequest) GetAwsAccountIdentifierOk() (*string, bool)`

GetAwsAccountIdentifierOk returns a tuple with the AwsAccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountIdentifier

`func (o *AWSTagFilterDeleteRequest) SetAwsAccountIdentifier(v string)`

SetAwsAccountIdentifier sets AwsAccountIdentifier field to given value.

### HasAwsAccountIdentifier

`func (o *AWSTagFilterDeleteRequest) HasAwsAccountIdentifier() bool`

HasAwsAccountIdentifier returns a boolean if a field has been set.

### GetNamespace

`func (o *AWSTagFilterDeleteRequest) GetNamespace() AWSNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AWSTagFilterDeleteRequest) GetNamespaceOk() (*AWSNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AWSTagFilterDeleteRequest) SetNamespace(v AWSNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AWSTagFilterDeleteRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


