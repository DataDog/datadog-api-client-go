# GcpAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthProviderX509CertUrl** | Pointer to **string** | Should be https://www.googleapis.com/oauth2/v1/certs. | [optional] 
**AuthUri** | Pointer to **string** | Should be https://accounts.google.com/o/oauth2/auth. | [optional] 
**Automute** | Pointer to **bool** | Silence monitors for expected GCE instance shutdowns. | [optional] 
**ClientEmail** | Pointer to **string** | Your email found in your JSON service account key. | [optional] 
**ClientId** | Pointer to **string** | Your ID found in your JSON service account key. | [optional] 
**ClientX509CertUrl** | Pointer to **string** | Should be https://www.googleapis.com/robot/v1/metadata/x509/&lt;CLIENT_EMAIL&gt; where &lt;CLIENT_EMAIL&gt; is the email found in your JSON service account key. | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**HostFilters** | Pointer to **string** | Limit the GCE instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. | [optional] 
**PrivateKey** | Pointer to **string** | Your private key name found in your JSON service account key. | [optional] 
**PrivateKeyId** | Pointer to **string** | Your private key ID found in your JSON service account key. | [optional] 
**ProjectId** | Pointer to **string** | Your Google Cloud project ID found in your JSON service account key. | [optional] 
**TokenUri** | Pointer to **string** | Should be https://accounts.google.com/o/oauth2/token. | [optional] 
**Type** | Pointer to **string** | The value for service_account found in your JSON service account key. | [optional] 

## Methods

### NewGcpAccount

`func NewGcpAccount() *GcpAccount`

NewGcpAccount instantiates a new GcpAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpAccountWithDefaults

`func NewGcpAccountWithDefaults() *GcpAccount`

NewGcpAccountWithDefaults instantiates a new GcpAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthProviderX509CertUrl

`func (o *GcpAccount) GetAuthProviderX509CertUrl() string`

GetAuthProviderX509CertUrl returns the AuthProviderX509CertUrl field if non-nil, zero value otherwise.

### GetAuthProviderX509CertUrlOk

`func (o *GcpAccount) GetAuthProviderX509CertUrlOk() (string, bool)`

GetAuthProviderX509CertUrlOk returns a tuple with the AuthProviderX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthProviderX509CertUrl

`func (o *GcpAccount) HasAuthProviderX509CertUrl() bool`

HasAuthProviderX509CertUrl returns a boolean if a field has been set.

### SetAuthProviderX509CertUrl

`func (o *GcpAccount) SetAuthProviderX509CertUrl(v string)`

SetAuthProviderX509CertUrl gets a reference to the given string and assigns it to the AuthProviderX509CertUrl field.

### GetAuthUri

`func (o *GcpAccount) GetAuthUri() string`

GetAuthUri returns the AuthUri field if non-nil, zero value otherwise.

### GetAuthUriOk

`func (o *GcpAccount) GetAuthUriOk() (string, bool)`

GetAuthUriOk returns a tuple with the AuthUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAuthUri

`func (o *GcpAccount) HasAuthUri() bool`

HasAuthUri returns a boolean if a field has been set.

### SetAuthUri

`func (o *GcpAccount) SetAuthUri(v string)`

SetAuthUri gets a reference to the given string and assigns it to the AuthUri field.

### GetAutomute

`func (o *GcpAccount) GetAutomute() bool`

GetAutomute returns the Automute field if non-nil, zero value otherwise.

### GetAutomuteOk

`func (o *GcpAccount) GetAutomuteOk() (bool, bool)`

GetAutomuteOk returns a tuple with the Automute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAutomute

`func (o *GcpAccount) HasAutomute() bool`

HasAutomute returns a boolean if a field has been set.

### SetAutomute

`func (o *GcpAccount) SetAutomute(v bool)`

SetAutomute gets a reference to the given bool and assigns it to the Automute field.

### GetClientEmail

`func (o *GcpAccount) GetClientEmail() string`

GetClientEmail returns the ClientEmail field if non-nil, zero value otherwise.

### GetClientEmailOk

`func (o *GcpAccount) GetClientEmailOk() (string, bool)`

GetClientEmailOk returns a tuple with the ClientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientEmail

`func (o *GcpAccount) HasClientEmail() bool`

HasClientEmail returns a boolean if a field has been set.

### SetClientEmail

`func (o *GcpAccount) SetClientEmail(v string)`

SetClientEmail gets a reference to the given string and assigns it to the ClientEmail field.

### GetClientId

`func (o *GcpAccount) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GcpAccount) GetClientIdOk() (string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientId

`func (o *GcpAccount) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientId

`func (o *GcpAccount) SetClientId(v string)`

SetClientId gets a reference to the given string and assigns it to the ClientId field.

### GetClientX509CertUrl

`func (o *GcpAccount) GetClientX509CertUrl() string`

GetClientX509CertUrl returns the ClientX509CertUrl field if non-nil, zero value otherwise.

### GetClientX509CertUrlOk

`func (o *GcpAccount) GetClientX509CertUrlOk() (string, bool)`

GetClientX509CertUrlOk returns a tuple with the ClientX509CertUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasClientX509CertUrl

`func (o *GcpAccount) HasClientX509CertUrl() bool`

HasClientX509CertUrl returns a boolean if a field has been set.

### SetClientX509CertUrl

`func (o *GcpAccount) SetClientX509CertUrl(v string)`

SetClientX509CertUrl gets a reference to the given string and assigns it to the ClientX509CertUrl field.

### GetErrors

`func (o *GcpAccount) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *GcpAccount) GetErrorsOk() ([]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *GcpAccount) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *GcpAccount) SetErrors(v []string)`

SetErrors gets a reference to the given []string and assigns it to the Errors field.

### GetHostFilters

`func (o *GcpAccount) GetHostFilters() string`

GetHostFilters returns the HostFilters field if non-nil, zero value otherwise.

### GetHostFiltersOk

`func (o *GcpAccount) GetHostFiltersOk() (string, bool)`

GetHostFiltersOk returns a tuple with the HostFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHostFilters

`func (o *GcpAccount) HasHostFilters() bool`

HasHostFilters returns a boolean if a field has been set.

### SetHostFilters

`func (o *GcpAccount) SetHostFilters(v string)`

SetHostFilters gets a reference to the given string and assigns it to the HostFilters field.

### GetPrivateKey

`func (o *GcpAccount) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GcpAccount) GetPrivateKeyOk() (string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateKey

`func (o *GcpAccount) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKey

`func (o *GcpAccount) SetPrivateKey(v string)`

SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.

### GetPrivateKeyId

`func (o *GcpAccount) GetPrivateKeyId() string`

GetPrivateKeyId returns the PrivateKeyId field if non-nil, zero value otherwise.

### GetPrivateKeyIdOk

`func (o *GcpAccount) GetPrivateKeyIdOk() (string, bool)`

GetPrivateKeyIdOk returns a tuple with the PrivateKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrivateKeyId

`func (o *GcpAccount) HasPrivateKeyId() bool`

HasPrivateKeyId returns a boolean if a field has been set.

### SetPrivateKeyId

`func (o *GcpAccount) SetPrivateKeyId(v string)`

SetPrivateKeyId gets a reference to the given string and assigns it to the PrivateKeyId field.

### GetProjectId

`func (o *GcpAccount) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GcpAccount) GetProjectIdOk() (string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProjectId

`func (o *GcpAccount) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectId

`func (o *GcpAccount) SetProjectId(v string)`

SetProjectId gets a reference to the given string and assigns it to the ProjectId field.

### GetTokenUri

`func (o *GcpAccount) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *GcpAccount) GetTokenUriOk() (string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTokenUri

`func (o *GcpAccount) HasTokenUri() bool`

HasTokenUri returns a boolean if a field has been set.

### SetTokenUri

`func (o *GcpAccount) SetTokenUri(v string)`

SetTokenUri gets a reference to the given string and assigns it to the TokenUri field.

### GetType

`func (o *GcpAccount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GcpAccount) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *GcpAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *GcpAccount) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


