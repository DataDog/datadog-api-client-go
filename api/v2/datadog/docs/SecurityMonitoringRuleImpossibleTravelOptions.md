# SecurityMonitoringRuleImpossibleTravelOptions

## Properties

| Name                      | Type                | Description                                                                                                                                                                                                       | Notes      |
| ------------------------- | ------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **BaselineUserLocations** | Pointer to **bool** | If true, signals are suppressed for the first 24 hours. In that time, Datadog learns the user&#39;s regular access locations. This can be helpful to reduce noise and infer VPN usage or credentialed API access. | [optional] |

## Methods

### NewSecurityMonitoringRuleImpossibleTravelOptions

`func NewSecurityMonitoringRuleImpossibleTravelOptions() *SecurityMonitoringRuleImpossibleTravelOptions`

NewSecurityMonitoringRuleImpossibleTravelOptions instantiates a new SecurityMonitoringRuleImpossibleTravelOptions object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityMonitoringRuleImpossibleTravelOptionsWithDefaults

`func NewSecurityMonitoringRuleImpossibleTravelOptionsWithDefaults() *SecurityMonitoringRuleImpossibleTravelOptions`

NewSecurityMonitoringRuleImpossibleTravelOptionsWithDefaults instantiates a new SecurityMonitoringRuleImpossibleTravelOptions object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBaselineUserLocations

`func (o *SecurityMonitoringRuleImpossibleTravelOptions) GetBaselineUserLocations() bool`

GetBaselineUserLocations returns the BaselineUserLocations field if non-nil, zero value otherwise.

### GetBaselineUserLocationsOk

`func (o *SecurityMonitoringRuleImpossibleTravelOptions) GetBaselineUserLocationsOk() (*bool, bool)`

GetBaselineUserLocationsOk returns a tuple with the BaselineUserLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineUserLocations

`func (o *SecurityMonitoringRuleImpossibleTravelOptions) SetBaselineUserLocations(v bool)`

SetBaselineUserLocations sets BaselineUserLocations field to given value.

### HasBaselineUserLocations

`func (o *SecurityMonitoringRuleImpossibleTravelOptions) HasBaselineUserLocations() bool`

HasBaselineUserLocations returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
