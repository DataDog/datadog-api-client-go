# SyntheticsTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicAuth** | Pointer to [**SyntheticsTestRequestBasicAuth**](SyntheticsTestRequest_basicAuth.md) |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Method** | Pointer to [**HTTPMethod**](HTTPMethod.md) |  | 
**Port** | Pointer to **int64** |  | [optional] 
**Query** | Pointer to [**interface{}**](.md) |  | [optional] 
**Timeout** | Pointer to **float64** |  | [optional] 
**Url** | Pointer to **string** |  | 

## Methods

### NewSyntheticsTestRequest

`func NewSyntheticsTestRequest(method HTTPMethod, url string, ) *SyntheticsTestRequest`

NewSyntheticsTestRequest instantiates a new SyntheticsTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticsTestRequestWithDefaults

`func NewSyntheticsTestRequestWithDefaults() *SyntheticsTestRequest`

NewSyntheticsTestRequestWithDefaults instantiates a new SyntheticsTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicAuth

`func (o *SyntheticsTestRequest) GetBasicAuth() SyntheticsTestRequestBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *SyntheticsTestRequest) GetBasicAuthOk() (SyntheticsTestRequestBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBasicAuth

`func (o *SyntheticsTestRequest) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### SetBasicAuth

`func (o *SyntheticsTestRequest) SetBasicAuth(v SyntheticsTestRequestBasicAuth)`

SetBasicAuth gets a reference to the given SyntheticsTestRequestBasicAuth and assigns it to the BasicAuth field.

### GetBody

`func (o *SyntheticsTestRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SyntheticsTestRequest) GetBodyOk() (string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBody

`func (o *SyntheticsTestRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBody

`func (o *SyntheticsTestRequest) SetBody(v string)`

SetBody gets a reference to the given string and assigns it to the Body field.

### GetHeaders

`func (o *SyntheticsTestRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SyntheticsTestRequest) GetHeadersOk() (map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHeaders

`func (o *SyntheticsTestRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeaders

`func (o *SyntheticsTestRequest) SetHeaders(v map[string]string)`

SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.

### GetHost

`func (o *SyntheticsTestRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SyntheticsTestRequest) GetHostOk() (string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHost

`func (o *SyntheticsTestRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHost

`func (o *SyntheticsTestRequest) SetHost(v string)`

SetHost gets a reference to the given string and assigns it to the Host field.

### GetMethod

`func (o *SyntheticsTestRequest) GetMethod() HTTPMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SyntheticsTestRequest) GetMethodOk() (HTTPMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMethod

`func (o *SyntheticsTestRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethod

`func (o *SyntheticsTestRequest) SetMethod(v HTTPMethod)`

SetMethod gets a reference to the given HTTPMethod and assigns it to the Method field.

### GetPort

`func (o *SyntheticsTestRequest) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyntheticsTestRequest) GetPortOk() (int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPort

`func (o *SyntheticsTestRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPort

`func (o *SyntheticsTestRequest) SetPort(v int64)`

SetPort gets a reference to the given int64 and assigns it to the Port field.

### GetQuery

`func (o *SyntheticsTestRequest) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SyntheticsTestRequest) GetQueryOk() (interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasQuery

`func (o *SyntheticsTestRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQuery

`func (o *SyntheticsTestRequest) SetQuery(v interface{})`

SetQuery gets a reference to the given interface{} and assigns it to the Query field.

### GetTimeout

`func (o *SyntheticsTestRequest) GetTimeout() float64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SyntheticsTestRequest) GetTimeoutOk() (float64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeout

`func (o *SyntheticsTestRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeout

`func (o *SyntheticsTestRequest) SetTimeout(v float64)`

SetTimeout gets a reference to the given float64 and assigns it to the Timeout field.

### GetUrl

`func (o *SyntheticsTestRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyntheticsTestRequest) GetUrlOk() (string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUrl

`func (o *SyntheticsTestRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrl

`func (o *SyntheticsTestRequest) SetUrl(v string)`

SetUrl gets a reference to the given string and assigns it to the Url field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


