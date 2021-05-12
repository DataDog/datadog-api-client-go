# AWSAccountListResponse

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Accounts** | Pointer to [**[]AWSAccount**](AWSAccount.md) | List of enabled AWS accounts. | [optional] 

## Methods

### NewAWSAccountListResponse

`func NewAWSAccountListResponse() *AWSAccountListResponse`

NewAWSAccountListResponse instantiates a new AWSAccountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSAccountListResponseWithDefaults

`func NewAWSAccountListResponseWithDefaults() *AWSAccountListResponse`

NewAWSAccountListResponseWithDefaults instantiates a new AWSAccountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AWSAccountListResponse) GetAccounts() []AWSAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AWSAccountListResponse) GetAccountsOk() (*[]AWSAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AWSAccountListResponse) SetAccounts(v []AWSAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AWSAccountListResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


