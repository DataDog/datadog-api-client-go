# AwsAccountListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AwsAccount**](AWSAccount.md) |  | [optional] 

## Methods

### NewAwsAccountListResponse

`func NewAwsAccountListResponse() *AwsAccountListResponse`

NewAwsAccountListResponse instantiates a new AwsAccountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsAccountListResponseWithDefaults

`func NewAwsAccountListResponseWithDefaults() *AwsAccountListResponse`

NewAwsAccountListResponseWithDefaults instantiates a new AwsAccountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AwsAccountListResponse) GetAccounts() []AwsAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AwsAccountListResponse) GetAccountsOk() ([]AwsAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccounts

`func (o *AwsAccountListResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccounts

`func (o *AwsAccountListResponse) SetAccounts(v []AwsAccount)`

SetAccounts gets a reference to the given []AwsAccount and assigns it to the Accounts field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


