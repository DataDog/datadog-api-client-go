# SyntheticsTestRequestProxy

## Properties

| Name        | Type                             | Description                                  | Notes      |
| ----------- | -------------------------------- | -------------------------------------------- | ---------- |
| **Headers** | Pointer to **map[string]string** | Headers to include when performing the test. | [optional] |
| **Url**     | **string**                       | URL of the proxy to perform the test.        |

## Methods

### NewSyntheticsTestRequestProxy

`func NewSyntheticsTestRequestProxy(url string) *SyntheticsTestRequestProxy`

NewSyntheticsTestRequestProxy instantiates a new SyntheticsTestRequestProxy object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsTestRequestProxyWithDefaults

`func NewSyntheticsTestRequestProxyWithDefaults() *SyntheticsTestRequestProxy`

NewSyntheticsTestRequestProxyWithDefaults instantiates a new SyntheticsTestRequestProxy object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetHeaders

`func (o *SyntheticsTestRequestProxy) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SyntheticsTestRequestProxy) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SyntheticsTestRequestProxy) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SyntheticsTestRequestProxy) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetUrl

`func (o *SyntheticsTestRequestProxy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyntheticsTestRequestProxy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyntheticsTestRequestProxy) SetUrl(v string)`

SetUrl sets Url field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
