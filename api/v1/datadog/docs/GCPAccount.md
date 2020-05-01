# GCPAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthProviderX509CertUrl** | Pointer to **string** | Should be https://www.googleapis.com/oauth2/v1/certs. | [optional] 
**AuthUri** | Pointer to **string** | Should be https://accounts.google.com/o/oauth2/auth. | [optional] 
**Automute** | Pointer to **bool** | Silence monitors for expected GCE instance shutdowns. | [optional] 
**ClientEmail** | Pointer to **string** | Your email found in your JSON service account key. | [optional] 
**ClientId** | Pointer to **string** | Your ID found in your JSON service account key. | [optional] 
**ClientX509CertUrl** | Pointer to **string** | Should be https://www.googleapis.com/robot/v1/metadata/x509/&lt;CLIENT_EMAIL&gt; where &lt;CLIENT_EMAIL&gt; is the email found in your JSON service account key. | [optional] 
**Errors** | Pointer to **[]string** | An array of errors. | [optional] 
**HostFilters** | Pointer to **string** | Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. | [optional] 
**PrivateKey** | Pointer to **string** | Your private key name found in your JSON service account key. | [optional] 
**PrivateKeyId** | Pointer to **string** | Your private key ID found in your JSON service account key. | [optional] 
**ProjectId** | Pointer to **string** | Your Google Cloud project ID found in your JSON service account key. | [optional] 
**TokenUri** | Pointer to **string** | Should be https://accounts.google.com/o/oauth2/token. | [optional] 
**Type** | Pointer to **string** | The value for service_account found in your JSON service account key. | [optional] 

## Methods

### NewGCPAccount

`func NewGCPAccount() *GCPAccount`

NewGCPAccount instantiates a new GCPAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPAccountWithDefaults

`func NewGCPAccountWithDefaults() *GCPAccount`

NewGCPAccountWithDefaults instantiates a new GCPAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthProviderX509CertUrl

`func (o *GCPAccount) GetAuthProviderX509CertUrl() string`

GetAuthProviderX509CertUrl returns the AuthProviderX509CertUrl field if non-nil, zero value otherwise.

### GetAuthProviderX509CertUrlOk

`func (o *GCPAccount) GetAuthProviderX509CertUrlOk() (*string, bool)`

GetAuthProviderX509CertUrlOk returns a tuple with the AuthProviderX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProviderX509CertUrl

`func (o *GCPAccount) SetAuthProviderX509CertUrl(v string)`

SetAuthProviderX509CertUrl sets AuthProviderX509CertUrl field to given value.

### HasAuthProviderX509CertUrl

`func (o *GCPAccount) HasAuthProviderX509CertUrl() bool`

HasAuthProviderX509CertUrl returns a boolean if a field has been set.

### GetAuthUri

`func (o *GCPAccount) GetAuthUri() string`

GetAuthUri returns the AuthUri field if non-nil, zero value otherwise.

### GetAuthUriOk

`func (o *GCPAccount) GetAuthUriOk() (*string, bool)`

GetAuthUriOk returns a tuple with the AuthUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUri

`func (o *GCPAccount) SetAuthUri(v string)`

SetAuthUri sets AuthUri field to given value.

### HasAuthUri

`func (o *GCPAccount) HasAuthUri() bool`

HasAuthUri returns a boolean if a field has been set.

### GetAutomute

`func (o *GCPAccount) GetAutomute() bool`

GetAutomute returns the Automute field if non-nil, zero value otherwise.

### GetAutomuteOk

`func (o *GCPAccount) GetAutomuteOk() (*bool, bool)`

GetAutomuteOk returns a tuple with the Automute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomute

`func (o *GCPAccount) SetAutomute(v bool)`

SetAutomute sets Automute field to given value.

### HasAutomute

`func (o *GCPAccount) HasAutomute() bool`

HasAutomute returns a boolean if a field has been set.

### GetClientEmail

`func (o *GCPAccount) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *GCPAccount) GetClientEmailOk() (*string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEmail

`func (o *GCPAccount) SetClientEmail(v string)`

SetClientEmail sets ClientEmail field to given value.

### HasClientEmail

`func (o *GCPAccount) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### GetClientId

`func (o *GCPAccount) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GCPAccount) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GCPAccount) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GCPAccount) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientX509CertUrl

`func (o *GCPAccount) GetClientX509CertUrl() string`

GetClientX509CertUrl returns the ClientX509CertUrl field if non-nil, zero value otherwise.

### GetClientX509CertUrlOk

`func (o *GCPAccount) GetClientX509CertUrlOk() (*string, bool)`

GetClientX509CertUrlOk returns a tuple with the ClientX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientX509CertUrl

`func (o *GCPAccount) SetClientX509CertUrl(v string)`

SetClientX509CertUrl sets ClientX509CertUrl field to given value.

### HasClientX509CertUrl

`func (o *GCPAccount) HasClientX509CertUrl() bool`

HasClientX509CertUrl returns a boolean if a field has been set.

### GetErrors

`func (o *GCPAccount) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GCPAccount) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *GCPAccount) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *GCPAccount) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetHostFilters

`func (o *GCPAccount) GetHostFilters() string`

GetHostFilters returns the HostFilters field if non-nil, zero value otherwise.

### GetHostFiltersOk

`func (o *GCPAccount) GetHostFiltersOk() (*string, bool)`

GetHostFiltersOk returns a tuple with the HostFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFilters

`func (o *GCPAccount) SetHostFilters(v string)`

SetHostFilters sets HostFilters field to given value.

### HasHostFilters

`func (o *GCPAccount) HasHostFilters() bool`

HasHostFilters returns a boolean if a field has been set.

### GetPrivateKey

`func (o *GCPAccount) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GCPAccount) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GCPAccount) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GCPAccount) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyId

`func (o *GCPAccount) GetPrivateKeyId() string`

GetPrivateKeyId returns the PrivateKeyId field if non-nil, zero value otherwise.

### GetPrivateKeyIdOk

`func (o *GCPAccount) GetPrivateKeyIdOk() (*string, bool)`

GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyId

`func (o *GCPAccount) SetPrivateKeyId(v string)`

SetPrivateKeyId sets PrivateKeyId field to given value.

### HasPrivateKeyId

`func (o *GCPAccount) HasPrivateKeyId() bool`

HasPrivateKeyId returns a boolean if a field has been set.

### GetProjectId

`func (o *GCPAccount) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GCPAccount) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GCPAccount) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GCPAccount) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetTokenUri

`func (o *GCPAccount) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *GCPAccount) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *GCPAccount) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.

### HasTokenUri

`func (o *GCPAccount) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### GetType

`func (o *GCPAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GCPAccount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GCPAccount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GCPAccount) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


