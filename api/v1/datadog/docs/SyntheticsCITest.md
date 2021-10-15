# SyntheticsCITest

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AllowInsecureCertificates** | Pointer to **bool** | Disable certificate checks in API tests. | [optional] 
**BasicAuth** | Pointer to [**SyntheticsBasicAuth**](SyntheticsBasicAuth.md) |  | [optional] 
**Body** | Pointer to **string** | Body to include in the test. | [optional] 
**BodyType** | Pointer to **string** | Type of the data sent in a synthetics API test. | [optional] 
**Cookies** | Pointer to **string** | Cookies for the request. | [optional] 
**DeviceIds** | Pointer to [**[]SyntheticsDeviceID**](SyntheticsDeviceID.md) | For browser test, array with the different device IDs used to run the test. | [optional] 
**FollowRedirects** | Pointer to **bool** | For API HTTP test, whether or not the test should follow redirects. | [optional] 
**Headers** | Pointer to **map[string]string** | Headers to include when performing the test. | [optional] 
**Locations** | Pointer to **[]string** | Array of locations used to run the test. | [optional] 
**Metadata** | Pointer to [**SyntheticsCIBatchMetadata**](SyntheticsCIBatchMetadata.md) |  | [optional] 
**PublicId** | **string** | The public ID of the Synthetics test to trigger. | 
**Retry** | Pointer to [**SyntheticsTestOptionsRetry**](SyntheticsTestOptionsRetry.md) |  | [optional] 
**StartUrl** | Pointer to **string** | Starting URL for the browser test. | [optional] 
**Variables** | Pointer to **map[string]string** | Variables to replace in the test. | [optional] 

## Methods

### NewSyntheticsCITest

`func NewSyntheticsCITest(publicId string) *SyntheticsCITest`

NewSyntheticsCITest instantiates a new SyntheticsCITest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsCITestWithDefaults

`func NewSyntheticsCITestWithDefaults() *SyntheticsCITest`

NewSyntheticsCITestWithDefaults instantiates a new SyntheticsCITest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAllowInsecureCertificates

`func (o *SyntheticsCITest) GetAllowInsecureCertificates() bool`

GetAllowInsecureCertificates returns the AllowInsecureCertificates field if non-nil, zero value otherwise.

### GetAllowInsecureCertificatesOk

`func (o *SyntheticsCITest) GetAllowInsecureCertificatesOk() (*bool, bool)`

GetAllowInsecureCertificatesOk returns a tuple with the AllowInsecureCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecureCertificates

`func (o *SyntheticsCITest) SetAllowInsecureCertificates(v bool)`

SetAllowInsecureCertificates sets AllowInsecureCertificates field to given value.

### HasAllowInsecureCertificates

`func (o *SyntheticsCITest) HasAllowInsecureCertificates() bool`

HasAllowInsecureCertificates returns a boolean if a field has been set.

### GetBasicAuth

`func (o *SyntheticsCITest) GetBasicAuth() SyntheticsBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *SyntheticsCITest) GetBasicAuthOk() (*SyntheticsBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *SyntheticsCITest) SetBasicAuth(v SyntheticsBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *SyntheticsCITest) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBody

`func (o *SyntheticsCITest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SyntheticsCITest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SyntheticsCITest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SyntheticsCITest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyType

`func (o *SyntheticsCITest) GetBodyType() string`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *SyntheticsCITest) GetBodyTypeOk() (*string, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *SyntheticsCITest) SetBodyType(v string)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *SyntheticsCITest) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### GetCookies

`func (o *SyntheticsCITest) GetCookies() string`

GetCookies returns the Cookies field if non-nil, zero value otherwise.

### GetCookiesOk

`func (o *SyntheticsCITest) GetCookiesOk() (*string, bool)`

GetCookiesOk returns a tuple with the Cookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookies

`func (o *SyntheticsCITest) SetCookies(v string)`

SetCookies sets Cookies field to given value.

### HasCookies

`func (o *SyntheticsCITest) HasCookies() bool`

HasCookies returns a boolean if a field has been set.

### GetDeviceIds

`func (o *SyntheticsCITest) GetDeviceIds() []SyntheticsDeviceID`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *SyntheticsCITest) GetDeviceIdsOk() (*[]SyntheticsDeviceID, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *SyntheticsCITest) SetDeviceIds(v []SyntheticsDeviceID)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *SyntheticsCITest) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetFollowRedirects

`func (o *SyntheticsCITest) GetFollowRedirects() bool`

GetFollowRedirects returns the FollowRedirects field if non-nil, zero value otherwise.

### GetFollowRedirectsOk

`func (o *SyntheticsCITest) GetFollowRedirectsOk() (*bool, bool)`

GetFollowRedirectsOk returns a tuple with the FollowRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirects

`func (o *SyntheticsCITest) SetFollowRedirects(v bool)`

SetFollowRedirects sets FollowRedirects field to given value.

### HasFollowRedirects

`func (o *SyntheticsCITest) HasFollowRedirects() bool`

HasFollowRedirects returns a boolean if a field has been set.

### GetHeaders

`func (o *SyntheticsCITest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SyntheticsCITest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SyntheticsCITest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SyntheticsCITest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetLocations

`func (o *SyntheticsCITest) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticsCITest) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticsCITest) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SyntheticsCITest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetMetadata

`func (o *SyntheticsCITest) GetMetadata() SyntheticsCIBatchMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SyntheticsCITest) GetMetadataOk() (*SyntheticsCIBatchMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SyntheticsCITest) SetMetadata(v SyntheticsCIBatchMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SyntheticsCITest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPublicId

`func (o *SyntheticsCITest) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *SyntheticsCITest) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *SyntheticsCITest) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.


### GetRetry

`func (o *SyntheticsCITest) GetRetry() SyntheticsTestOptionsRetry`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *SyntheticsCITest) GetRetryOk() (*SyntheticsTestOptionsRetry, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *SyntheticsCITest) SetRetry(v SyntheticsTestOptionsRetry)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *SyntheticsCITest) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetStartUrl

`func (o *SyntheticsCITest) GetStartUrl() string`

GetStartUrl returns the StartUrl field if non-nil, zero value otherwise.

### GetStartUrlOk

`func (o *SyntheticsCITest) GetStartUrlOk() (*string, bool)`

GetStartUrlOk returns a tuple with the StartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartUrl

`func (o *SyntheticsCITest) SetStartUrl(v string)`

SetStartUrl sets StartUrl field to given value.

### HasStartUrl

`func (o *SyntheticsCITest) HasStartUrl() bool`

HasStartUrl returns a boolean if a field has been set.

### GetVariables

`func (o *SyntheticsCITest) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *SyntheticsCITest) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *SyntheticsCITest) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *SyntheticsCITest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


