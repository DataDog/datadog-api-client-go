// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineConfigDestinationItem - A destination for the pipeline.
type ObservabilityPipelineConfigDestinationItem struct {
	ObservabilityPipelineDatadogLogsDestination   *ObservabilityPipelineDatadogLogsDestination
	ObservabilityPipelineSumoLogicDestination     *ObservabilityPipelineSumoLogicDestination
	ObservabilityPipelineElasticsearchDestination *ObservabilityPipelineElasticsearchDestination
	ObservabilityPipelineRsyslogDestination       *ObservabilityPipelineRsyslogDestination
	ObservabilityPipelineSyslogNgDestination      *ObservabilityPipelineSyslogNgDestination
	AzureStorageDestination                       *AzureStorageDestination
	MicrosoftSentinelDestination                  *MicrosoftSentinelDestination

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineDatadogLogsDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineDatadogLogsDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineDatadogLogsDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineDatadogLogsDestination: v}
}

// ObservabilityPipelineSumoLogicDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSumoLogicDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSumoLogicDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSumoLogicDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSumoLogicDestination: v}
}

// ObservabilityPipelineElasticsearchDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineElasticsearchDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineElasticsearchDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineElasticsearchDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineElasticsearchDestination: v}
}

// ObservabilityPipelineRsyslogDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineRsyslogDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineRsyslogDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineRsyslogDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineRsyslogDestination: v}
}

// ObservabilityPipelineSyslogNgDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns ObservabilityPipelineSyslogNgDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func ObservabilityPipelineSyslogNgDestinationAsObservabilityPipelineConfigDestinationItem(v *ObservabilityPipelineSyslogNgDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{ObservabilityPipelineSyslogNgDestination: v}
}

// AzureStorageDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns AzureStorageDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func AzureStorageDestinationAsObservabilityPipelineConfigDestinationItem(v *AzureStorageDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{AzureStorageDestination: v}
}

// MicrosoftSentinelDestinationAsObservabilityPipelineConfigDestinationItem is a convenience function that returns MicrosoftSentinelDestination wrapped in ObservabilityPipelineConfigDestinationItem.
func MicrosoftSentinelDestinationAsObservabilityPipelineConfigDestinationItem(v *MicrosoftSentinelDestination) ObservabilityPipelineConfigDestinationItem {
	return ObservabilityPipelineConfigDestinationItem{MicrosoftSentinelDestination: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineConfigDestinationItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineDatadogLogsDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineDatadogLogsDestination)
	if err == nil {
		if obj.ObservabilityPipelineDatadogLogsDestination != nil && obj.ObservabilityPipelineDatadogLogsDestination.UnparsedObject == nil {
			jsonObservabilityPipelineDatadogLogsDestination, _ := datadog.Marshal(obj.ObservabilityPipelineDatadogLogsDestination)
			if string(jsonObservabilityPipelineDatadogLogsDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineDatadogLogsDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineDatadogLogsDestination = nil
		}
	} else {
		obj.ObservabilityPipelineDatadogLogsDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSumoLogicDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSumoLogicDestination)
	if err == nil {
		if obj.ObservabilityPipelineSumoLogicDestination != nil && obj.ObservabilityPipelineSumoLogicDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSumoLogicDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSumoLogicDestination)
			if string(jsonObservabilityPipelineSumoLogicDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSumoLogicDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSumoLogicDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSumoLogicDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineElasticsearchDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineElasticsearchDestination)
	if err == nil {
		if obj.ObservabilityPipelineElasticsearchDestination != nil && obj.ObservabilityPipelineElasticsearchDestination.UnparsedObject == nil {
			jsonObservabilityPipelineElasticsearchDestination, _ := datadog.Marshal(obj.ObservabilityPipelineElasticsearchDestination)
			if string(jsonObservabilityPipelineElasticsearchDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineElasticsearchDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineElasticsearchDestination = nil
		}
	} else {
		obj.ObservabilityPipelineElasticsearchDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineRsyslogDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineRsyslogDestination)
	if err == nil {
		if obj.ObservabilityPipelineRsyslogDestination != nil && obj.ObservabilityPipelineRsyslogDestination.UnparsedObject == nil {
			jsonObservabilityPipelineRsyslogDestination, _ := datadog.Marshal(obj.ObservabilityPipelineRsyslogDestination)
			if string(jsonObservabilityPipelineRsyslogDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineRsyslogDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineRsyslogDestination = nil
		}
	} else {
		obj.ObservabilityPipelineRsyslogDestination = nil
	}

	// try to unmarshal data into ObservabilityPipelineSyslogNgDestination
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSyslogNgDestination)
	if err == nil {
		if obj.ObservabilityPipelineSyslogNgDestination != nil && obj.ObservabilityPipelineSyslogNgDestination.UnparsedObject == nil {
			jsonObservabilityPipelineSyslogNgDestination, _ := datadog.Marshal(obj.ObservabilityPipelineSyslogNgDestination)
			if string(jsonObservabilityPipelineSyslogNgDestination) == "{}" { // empty struct
				obj.ObservabilityPipelineSyslogNgDestination = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSyslogNgDestination = nil
		}
	} else {
		obj.ObservabilityPipelineSyslogNgDestination = nil
	}

	// try to unmarshal data into AzureStorageDestination
	err = datadog.Unmarshal(data, &obj.AzureStorageDestination)
	if err == nil {
		if obj.AzureStorageDestination != nil && obj.AzureStorageDestination.UnparsedObject == nil {
			jsonAzureStorageDestination, _ := datadog.Marshal(obj.AzureStorageDestination)
			if string(jsonAzureStorageDestination) == "{}" { // empty struct
				obj.AzureStorageDestination = nil
			} else {
				match++
			}
		} else {
			obj.AzureStorageDestination = nil
		}
	} else {
		obj.AzureStorageDestination = nil
	}

	// try to unmarshal data into MicrosoftSentinelDestination
	err = datadog.Unmarshal(data, &obj.MicrosoftSentinelDestination)
	if err == nil {
		if obj.MicrosoftSentinelDestination != nil && obj.MicrosoftSentinelDestination.UnparsedObject == nil {
			jsonMicrosoftSentinelDestination, _ := datadog.Marshal(obj.MicrosoftSentinelDestination)
			if string(jsonMicrosoftSentinelDestination) == "{}" { // empty struct
				obj.MicrosoftSentinelDestination = nil
			} else {
				match++
			}
		} else {
			obj.MicrosoftSentinelDestination = nil
		}
	} else {
		obj.MicrosoftSentinelDestination = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineDatadogLogsDestination = nil
		obj.ObservabilityPipelineSumoLogicDestination = nil
		obj.ObservabilityPipelineElasticsearchDestination = nil
		obj.ObservabilityPipelineRsyslogDestination = nil
		obj.ObservabilityPipelineSyslogNgDestination = nil
		obj.AzureStorageDestination = nil
		obj.MicrosoftSentinelDestination = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineConfigDestinationItem) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineDatadogLogsDestination)
	}

	if obj.ObservabilityPipelineSumoLogicDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSumoLogicDestination)
	}

	if obj.ObservabilityPipelineElasticsearchDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineElasticsearchDestination)
	}

	if obj.ObservabilityPipelineRsyslogDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineRsyslogDestination)
	}

	if obj.ObservabilityPipelineSyslogNgDestination != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSyslogNgDestination)
	}

	if obj.AzureStorageDestination != nil {
		return datadog.Marshal(&obj.AzureStorageDestination)
	}

	if obj.MicrosoftSentinelDestination != nil {
		return datadog.Marshal(&obj.MicrosoftSentinelDestination)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineConfigDestinationItem) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineDatadogLogsDestination != nil {
		return obj.ObservabilityPipelineDatadogLogsDestination
	}

	if obj.ObservabilityPipelineSumoLogicDestination != nil {
		return obj.ObservabilityPipelineSumoLogicDestination
	}

	if obj.ObservabilityPipelineElasticsearchDestination != nil {
		return obj.ObservabilityPipelineElasticsearchDestination
	}

	if obj.ObservabilityPipelineRsyslogDestination != nil {
		return obj.ObservabilityPipelineRsyslogDestination
	}

	if obj.ObservabilityPipelineSyslogNgDestination != nil {
		return obj.ObservabilityPipelineSyslogNgDestination
	}

	if obj.AzureStorageDestination != nil {
		return obj.AzureStorageDestination
	}

	if obj.MicrosoftSentinelDestination != nil {
		return obj.MicrosoftSentinelDestination
	}

	// all schemas are nil
	return nil
}
