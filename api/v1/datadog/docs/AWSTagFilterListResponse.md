# AWSTagFilterListResponse

## Properties

| Name        | Type                                             | Description              | Notes      |
| ----------- | ------------------------------------------------ | ------------------------ | ---------- |
| **Filters** | Pointer to [**[]AWSTagFilter**](AWSTagFilter.md) | An array of tag filters. | [optional] |

## Methods

### NewAWSTagFilterListResponse

`func NewAWSTagFilterListResponse() *AWSTagFilterListResponse`

NewAWSTagFilterListResponse instantiates a new AWSTagFilterListResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAWSTagFilterListResponseWithDefaults

`func NewAWSTagFilterListResponseWithDefaults() *AWSTagFilterListResponse`

NewAWSTagFilterListResponseWithDefaults instantiates a new AWSTagFilterListResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFilters

`func (o *AWSTagFilterListResponse) GetFilters() []AWSTagFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AWSTagFilterListResponse) GetFiltersOk() (*[]AWSTagFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AWSTagFilterListResponse) SetFilters(v []AWSTagFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AWSTagFilterListResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
