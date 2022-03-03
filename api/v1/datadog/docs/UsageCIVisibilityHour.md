# UsageCIVisibilityHour

## Properties

| Name                               | Type                  | Description                                                                                                                                                    | Notes      |
| ---------------------------------- | --------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **CiPipelineIndexedSpans**         | Pointer to **int32**  | The number of spans for pipelines in the queried hour.                                                                                                         | [optional] |
| **CiTestIndexedSpans**             | Pointer to **int32**  | The number of spans for tests in the queried hour.                                                                                                             | [optional] |
| **CiVisibilityPipelineCommitters** | Pointer to **int32**  | Shows the total count of all active Git committers for Pipelines in the current month. A committer is active if they commit at least 3 times in a given month. | [optional] |
| **CiVisibilityTestCommitters**     | Pointer to **int32**  | The total count of all active Git committers for tests in the current month. A committer is active if they commit at least 3 times in a given month.           | [optional] |
| **OrgName**                        | Pointer to **string** | The organization name.                                                                                                                                         | [optional] |
| **PublicId**                       | Pointer to **string** | The organization public ID.                                                                                                                                    | [optional] |

## Methods

### NewUsageCIVisibilityHour

`func NewUsageCIVisibilityHour() *UsageCIVisibilityHour`

NewUsageCIVisibilityHour instantiates a new UsageCIVisibilityHour object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageCIVisibilityHourWithDefaults

`func NewUsageCIVisibilityHourWithDefaults() *UsageCIVisibilityHour`

NewUsageCIVisibilityHourWithDefaults instantiates a new UsageCIVisibilityHour object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCiPipelineIndexedSpans

`func (o *UsageCIVisibilityHour) GetCiPipelineIndexedSpans() int32`

GetCiPipelineIndexedSpans returns the CiPipelineIndexedSpans field if non-nil, zero value otherwise.

### GetCiPipelineIndexedSpansOk

`func (o *UsageCIVisibilityHour) GetCiPipelineIndexedSpansOk() (*int32, bool)`

GetCiPipelineIndexedSpansOk returns a tuple with the CiPipelineIndexedSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiPipelineIndexedSpans

`func (o *UsageCIVisibilityHour) SetCiPipelineIndexedSpans(v int32)`

SetCiPipelineIndexedSpans sets CiPipelineIndexedSpans field to given value.

### HasCiPipelineIndexedSpans

`func (o *UsageCIVisibilityHour) HasCiPipelineIndexedSpans() bool`

HasCiPipelineIndexedSpans returns a boolean if a field has been set.

### GetCiTestIndexedSpans

`func (o *UsageCIVisibilityHour) GetCiTestIndexedSpans() int32`

GetCiTestIndexedSpans returns the CiTestIndexedSpans field if non-nil, zero value otherwise.

### GetCiTestIndexedSpansOk

`func (o *UsageCIVisibilityHour) GetCiTestIndexedSpansOk() (*int32, bool)`

GetCiTestIndexedSpansOk returns a tuple with the CiTestIndexedSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiTestIndexedSpans

`func (o *UsageCIVisibilityHour) SetCiTestIndexedSpans(v int32)`

SetCiTestIndexedSpans sets CiTestIndexedSpans field to given value.

### HasCiTestIndexedSpans

`func (o *UsageCIVisibilityHour) HasCiTestIndexedSpans() bool`

HasCiTestIndexedSpans returns a boolean if a field has been set.

### GetCiVisibilityPipelineCommitters

`func (o *UsageCIVisibilityHour) GetCiVisibilityPipelineCommitters() int32`

GetCiVisibilityPipelineCommitters returns the CiVisibilityPipelineCommitters field if non-nil, zero value otherwise.

### GetCiVisibilityPipelineCommittersOk

`func (o *UsageCIVisibilityHour) GetCiVisibilityPipelineCommittersOk() (*int32, bool)`

GetCiVisibilityPipelineCommittersOk returns a tuple with the CiVisibilityPipelineCommitters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiVisibilityPipelineCommitters

`func (o *UsageCIVisibilityHour) SetCiVisibilityPipelineCommitters(v int32)`

SetCiVisibilityPipelineCommitters sets CiVisibilityPipelineCommitters field to given value.

### HasCiVisibilityPipelineCommitters

`func (o *UsageCIVisibilityHour) HasCiVisibilityPipelineCommitters() bool`

HasCiVisibilityPipelineCommitters returns a boolean if a field has been set.

### GetCiVisibilityTestCommitters

`func (o *UsageCIVisibilityHour) GetCiVisibilityTestCommitters() int32`

GetCiVisibilityTestCommitters returns the CiVisibilityTestCommitters field if non-nil, zero value otherwise.

### GetCiVisibilityTestCommittersOk

`func (o *UsageCIVisibilityHour) GetCiVisibilityTestCommittersOk() (*int32, bool)`

GetCiVisibilityTestCommittersOk returns a tuple with the CiVisibilityTestCommitters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiVisibilityTestCommitters

`func (o *UsageCIVisibilityHour) SetCiVisibilityTestCommitters(v int32)`

SetCiVisibilityTestCommitters sets CiVisibilityTestCommitters field to given value.

### HasCiVisibilityTestCommitters

`func (o *UsageCIVisibilityHour) HasCiVisibilityTestCommitters() bool`

HasCiVisibilityTestCommitters returns a boolean if a field has been set.

### GetOrgName

`func (o *UsageCIVisibilityHour) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *UsageCIVisibilityHour) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *UsageCIVisibilityHour) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *UsageCIVisibilityHour) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPublicId

`func (o *UsageCIVisibilityHour) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *UsageCIVisibilityHour) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *UsageCIVisibilityHour) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *UsageCIVisibilityHour) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
