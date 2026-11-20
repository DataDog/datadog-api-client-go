// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetRequest - Test-only generated schema for the dashboard versioning proof of concept.
type TimeseriesWidgetRequest struct {
	TimeseriesWidgetFormulaRequest_20270101              *TimeseriesWidgetFormulaRequest
	TimeseriesWidgetMetricRequest_20270101               *TimeseriesWidgetMetricRequest
	TimeseriesWidgetLogRequest_20270101                  *TimeseriesWidgetLogRequest
	TimeseriesWidgetApmRequest_20270101                  *TimeseriesWidgetApmRequest
	TimeseriesWidgetNetworkRequest_20270101              *TimeseriesWidgetNetworkRequest
	TimeseriesWidgetRumRequest_20270101                  *TimeseriesWidgetRumRequest
	TimeseriesWidgetSecurityRequest_20270101             *TimeseriesWidgetSecurityRequest
	TimeseriesWidgetComplianceFindingsRequest_20270101   *TimeseriesWidgetComplianceFindingsRequest
	TimeseriesWidgetEventRequest_20270101                *TimeseriesWidgetEventRequest
	TimeseriesWidgetProfilingMetricsRequest_20270101     *TimeseriesWidgetProfilingMetricsRequest
	TimeseriesWidgetCiPipelineRequest_20270101           *TimeseriesWidgetCiPipelineRequest
	TimeseriesWidgetCiTestRequest_20270101               *TimeseriesWidgetCiTestRequest
	TimeseriesWidgetAuditRequest_20270101                *TimeseriesWidgetAuditRequest
	TimeseriesWidgetAppsecSpanRequest_20270101           *TimeseriesWidgetAppsecSpanRequest
	TimeseriesWidgetSyntheticsBatchesRequest_20270101    *TimeseriesWidgetSyntheticsBatchesRequest
	TimeseriesWidgetSyntheticsTestRunsRequest_20270101   *TimeseriesWidgetSyntheticsTestRunsRequest
	TimeseriesWidgetProcessRequest_20270101              *TimeseriesWidgetProcessRequest
	TimeseriesWidgetLegacyProfileMetricsRequest_20270101 *TimeseriesWidgetLegacyProfileMetricsRequest
	TimeseriesWidgetLegacyEventRequest_20270101          *TimeseriesWidgetLegacyEventRequest
	TimeseriesWidgetDdsqlRequest_20270101                *TimeseriesWidgetDdsqlRequest
	TimeseriesWidgetAdvancedQueryRequest_20270101        *TimeseriesWidgetAdvancedQueryRequest
	TimeseriesWidgetLocalDatasetRequest_20270101         *TimeseriesWidgetLocalDatasetRequest

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// TimeseriesWidgetFormulaRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetFormulaRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetFormulaRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetFormulaRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetFormulaRequest_20270101: v}
}

// TimeseriesWidgetMetricRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetMetricRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetMetricRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetMetricRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetMetricRequest_20270101: v}
}

// TimeseriesWidgetLogRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetLogRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetLogRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetLogRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetLogRequest_20270101: v}
}

// TimeseriesWidgetApmRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetApmRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetApmRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetApmRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetApmRequest_20270101: v}
}

// TimeseriesWidgetNetworkRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetNetworkRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetNetworkRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetNetworkRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetNetworkRequest_20270101: v}
}

// TimeseriesWidgetRumRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetRumRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetRumRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetRumRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetRumRequest_20270101: v}
}

// TimeseriesWidgetSecurityRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetSecurityRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetSecurityRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetSecurityRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetSecurityRequest_20270101: v}
}

// TimeseriesWidgetComplianceFindingsRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetComplianceFindingsRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetComplianceFindingsRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetComplianceFindingsRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetComplianceFindingsRequest_20270101: v}
}

// TimeseriesWidgetEventRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetEventRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetEventRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetEventRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetEventRequest_20270101: v}
}

// TimeseriesWidgetProfilingMetricsRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetProfilingMetricsRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetProfilingMetricsRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetProfilingMetricsRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetProfilingMetricsRequest_20270101: v}
}

// TimeseriesWidgetCiPipelineRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetCiPipelineRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetCiPipelineRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetCiPipelineRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetCiPipelineRequest_20270101: v}
}

// TimeseriesWidgetCiTestRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetCiTestRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetCiTestRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetCiTestRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetCiTestRequest_20270101: v}
}

// TimeseriesWidgetAuditRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetAuditRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetAuditRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetAuditRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetAuditRequest_20270101: v}
}

// TimeseriesWidgetAppsecSpanRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetAppsecSpanRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetAppsecSpanRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetAppsecSpanRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetAppsecSpanRequest_20270101: v}
}

// TimeseriesWidgetSyntheticsBatchesRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetSyntheticsBatchesRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetSyntheticsBatchesRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetSyntheticsBatchesRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetSyntheticsBatchesRequest_20270101: v}
}

// TimeseriesWidgetSyntheticsTestRunsRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetSyntheticsTestRunsRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetSyntheticsTestRunsRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetSyntheticsTestRunsRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetSyntheticsTestRunsRequest_20270101: v}
}

// TimeseriesWidgetProcessRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetProcessRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetProcessRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetProcessRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetProcessRequest_20270101: v}
}

// TimeseriesWidgetLegacyProfileMetricsRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetLegacyProfileMetricsRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetLegacyProfileMetricsRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetLegacyProfileMetricsRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetLegacyProfileMetricsRequest_20270101: v}
}

// TimeseriesWidgetLegacyEventRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetLegacyEventRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetLegacyEventRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetLegacyEventRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetLegacyEventRequest_20270101: v}
}

// TimeseriesWidgetDdsqlRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetDdsqlRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetDdsqlRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetDdsqlRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetDdsqlRequest_20270101: v}
}

// TimeseriesWidgetAdvancedQueryRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetAdvancedQueryRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetAdvancedQueryRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetAdvancedQueryRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetAdvancedQueryRequest_20270101: v}
}

// TimeseriesWidgetLocalDatasetRequest_20270101AsTimeseriesWidgetRequest is a convenience function that returns TimeseriesWidgetLocalDatasetRequest wrapped in TimeseriesWidgetRequest.
func TimeseriesWidgetLocalDatasetRequest_20270101AsTimeseriesWidgetRequest(v *TimeseriesWidgetLocalDatasetRequest) TimeseriesWidgetRequest {
	return TimeseriesWidgetRequest{TimeseriesWidgetLocalDatasetRequest_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *TimeseriesWidgetRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TimeseriesWidgetFormulaRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetFormulaRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetFormulaRequest_20270101 != nil && obj.TimeseriesWidgetFormulaRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetFormulaRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetFormulaRequest_20270101)
			if string(jsonTimeseriesWidgetFormulaRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetFormulaRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetFormulaRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetFormulaRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetMetricRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetMetricRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetMetricRequest_20270101 != nil && obj.TimeseriesWidgetMetricRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetMetricRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetMetricRequest_20270101)
			if string(jsonTimeseriesWidgetMetricRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetMetricRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetMetricRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetMetricRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetLogRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetLogRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetLogRequest_20270101 != nil && obj.TimeseriesWidgetLogRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetLogRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetLogRequest_20270101)
			if string(jsonTimeseriesWidgetLogRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetLogRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetLogRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetLogRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetApmRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetApmRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetApmRequest_20270101 != nil && obj.TimeseriesWidgetApmRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetApmRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetApmRequest_20270101)
			if string(jsonTimeseriesWidgetApmRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetApmRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetApmRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetApmRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetNetworkRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetNetworkRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetNetworkRequest_20270101 != nil && obj.TimeseriesWidgetNetworkRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetNetworkRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetNetworkRequest_20270101)
			if string(jsonTimeseriesWidgetNetworkRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetNetworkRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetNetworkRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetNetworkRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetRumRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetRumRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetRumRequest_20270101 != nil && obj.TimeseriesWidgetRumRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetRumRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetRumRequest_20270101)
			if string(jsonTimeseriesWidgetRumRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetRumRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetRumRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetRumRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetSecurityRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetSecurityRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetSecurityRequest_20270101 != nil && obj.TimeseriesWidgetSecurityRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetSecurityRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetSecurityRequest_20270101)
			if string(jsonTimeseriesWidgetSecurityRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetSecurityRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetSecurityRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetSecurityRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetComplianceFindingsRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetComplianceFindingsRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetComplianceFindingsRequest_20270101 != nil && obj.TimeseriesWidgetComplianceFindingsRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetComplianceFindingsRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetComplianceFindingsRequest_20270101)
			if string(jsonTimeseriesWidgetComplianceFindingsRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetComplianceFindingsRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetComplianceFindingsRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetComplianceFindingsRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetEventRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetEventRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetEventRequest_20270101 != nil && obj.TimeseriesWidgetEventRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetEventRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetEventRequest_20270101)
			if string(jsonTimeseriesWidgetEventRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetEventRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetEventRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetEventRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetProfilingMetricsRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetProfilingMetricsRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetProfilingMetricsRequest_20270101 != nil && obj.TimeseriesWidgetProfilingMetricsRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetProfilingMetricsRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetProfilingMetricsRequest_20270101)
			if string(jsonTimeseriesWidgetProfilingMetricsRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetProfilingMetricsRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetProfilingMetricsRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetProfilingMetricsRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetCiPipelineRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetCiPipelineRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetCiPipelineRequest_20270101 != nil && obj.TimeseriesWidgetCiPipelineRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetCiPipelineRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetCiPipelineRequest_20270101)
			if string(jsonTimeseriesWidgetCiPipelineRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetCiPipelineRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetCiPipelineRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetCiPipelineRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetCiTestRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetCiTestRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetCiTestRequest_20270101 != nil && obj.TimeseriesWidgetCiTestRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetCiTestRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetCiTestRequest_20270101)
			if string(jsonTimeseriesWidgetCiTestRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetCiTestRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetCiTestRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetCiTestRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetAuditRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetAuditRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetAuditRequest_20270101 != nil && obj.TimeseriesWidgetAuditRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetAuditRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetAuditRequest_20270101)
			if string(jsonTimeseriesWidgetAuditRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetAuditRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetAuditRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetAuditRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetAppsecSpanRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetAppsecSpanRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetAppsecSpanRequest_20270101 != nil && obj.TimeseriesWidgetAppsecSpanRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetAppsecSpanRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetAppsecSpanRequest_20270101)
			if string(jsonTimeseriesWidgetAppsecSpanRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetAppsecSpanRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetAppsecSpanRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetAppsecSpanRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetSyntheticsBatchesRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101 != nil && obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetSyntheticsBatchesRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101)
			if string(jsonTimeseriesWidgetSyntheticsBatchesRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetSyntheticsTestRunsRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101 != nil && obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetSyntheticsTestRunsRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101)
			if string(jsonTimeseriesWidgetSyntheticsTestRunsRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetProcessRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetProcessRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetProcessRequest_20270101 != nil && obj.TimeseriesWidgetProcessRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetProcessRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetProcessRequest_20270101)
			if string(jsonTimeseriesWidgetProcessRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetProcessRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetProcessRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetProcessRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetLegacyProfileMetricsRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101 != nil && obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetLegacyProfileMetricsRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101)
			if string(jsonTimeseriesWidgetLegacyProfileMetricsRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetLegacyEventRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetLegacyEventRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetLegacyEventRequest_20270101 != nil && obj.TimeseriesWidgetLegacyEventRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetLegacyEventRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetLegacyEventRequest_20270101)
			if string(jsonTimeseriesWidgetLegacyEventRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetLegacyEventRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetLegacyEventRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetLegacyEventRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetDdsqlRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetDdsqlRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetDdsqlRequest_20270101 != nil && obj.TimeseriesWidgetDdsqlRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetDdsqlRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetDdsqlRequest_20270101)
			if string(jsonTimeseriesWidgetDdsqlRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetDdsqlRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetDdsqlRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetDdsqlRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetAdvancedQueryRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetAdvancedQueryRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetAdvancedQueryRequest_20270101 != nil && obj.TimeseriesWidgetAdvancedQueryRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetAdvancedQueryRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetAdvancedQueryRequest_20270101)
			if string(jsonTimeseriesWidgetAdvancedQueryRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetAdvancedQueryRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetAdvancedQueryRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetAdvancedQueryRequest_20270101 = nil
	}

	// try to unmarshal data into TimeseriesWidgetLocalDatasetRequest_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesWidgetLocalDatasetRequest_20270101)
	if err == nil {
		if obj.TimeseriesWidgetLocalDatasetRequest_20270101 != nil && obj.TimeseriesWidgetLocalDatasetRequest_20270101.UnparsedObject == nil {
			jsonTimeseriesWidgetLocalDatasetRequest_20270101, _ := datadog.Marshal(obj.TimeseriesWidgetLocalDatasetRequest_20270101)
			if string(jsonTimeseriesWidgetLocalDatasetRequest_20270101) == "{}" { // empty struct
				obj.TimeseriesWidgetLocalDatasetRequest_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesWidgetLocalDatasetRequest_20270101 = nil
		}
	} else {
		obj.TimeseriesWidgetLocalDatasetRequest_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.TimeseriesWidgetFormulaRequest_20270101 = nil
		obj.TimeseriesWidgetMetricRequest_20270101 = nil
		obj.TimeseriesWidgetLogRequest_20270101 = nil
		obj.TimeseriesWidgetApmRequest_20270101 = nil
		obj.TimeseriesWidgetNetworkRequest_20270101 = nil
		obj.TimeseriesWidgetRumRequest_20270101 = nil
		obj.TimeseriesWidgetSecurityRequest_20270101 = nil
		obj.TimeseriesWidgetComplianceFindingsRequest_20270101 = nil
		obj.TimeseriesWidgetEventRequest_20270101 = nil
		obj.TimeseriesWidgetProfilingMetricsRequest_20270101 = nil
		obj.TimeseriesWidgetCiPipelineRequest_20270101 = nil
		obj.TimeseriesWidgetCiTestRequest_20270101 = nil
		obj.TimeseriesWidgetAuditRequest_20270101 = nil
		obj.TimeseriesWidgetAppsecSpanRequest_20270101 = nil
		obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101 = nil
		obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101 = nil
		obj.TimeseriesWidgetProcessRequest_20270101 = nil
		obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101 = nil
		obj.TimeseriesWidgetLegacyEventRequest_20270101 = nil
		obj.TimeseriesWidgetDdsqlRequest_20270101 = nil
		obj.TimeseriesWidgetAdvancedQueryRequest_20270101 = nil
		obj.TimeseriesWidgetLocalDatasetRequest_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj TimeseriesWidgetRequest) MarshalJSON() ([]byte, error) {
	if obj.TimeseriesWidgetFormulaRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetFormulaRequest_20270101)
	}

	if obj.TimeseriesWidgetMetricRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetMetricRequest_20270101)
	}

	if obj.TimeseriesWidgetLogRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetLogRequest_20270101)
	}

	if obj.TimeseriesWidgetApmRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetApmRequest_20270101)
	}

	if obj.TimeseriesWidgetNetworkRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetNetworkRequest_20270101)
	}

	if obj.TimeseriesWidgetRumRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetRumRequest_20270101)
	}

	if obj.TimeseriesWidgetSecurityRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetSecurityRequest_20270101)
	}

	if obj.TimeseriesWidgetComplianceFindingsRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetComplianceFindingsRequest_20270101)
	}

	if obj.TimeseriesWidgetEventRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetEventRequest_20270101)
	}

	if obj.TimeseriesWidgetProfilingMetricsRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetProfilingMetricsRequest_20270101)
	}

	if obj.TimeseriesWidgetCiPipelineRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetCiPipelineRequest_20270101)
	}

	if obj.TimeseriesWidgetCiTestRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetCiTestRequest_20270101)
	}

	if obj.TimeseriesWidgetAuditRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetAuditRequest_20270101)
	}

	if obj.TimeseriesWidgetAppsecSpanRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetAppsecSpanRequest_20270101)
	}

	if obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101)
	}

	if obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101)
	}

	if obj.TimeseriesWidgetProcessRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetProcessRequest_20270101)
	}

	if obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101)
	}

	if obj.TimeseriesWidgetLegacyEventRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetLegacyEventRequest_20270101)
	}

	if obj.TimeseriesWidgetDdsqlRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetDdsqlRequest_20270101)
	}

	if obj.TimeseriesWidgetAdvancedQueryRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetAdvancedQueryRequest_20270101)
	}

	if obj.TimeseriesWidgetLocalDatasetRequest_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesWidgetLocalDatasetRequest_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *TimeseriesWidgetRequest) GetActualInstance() interface{} {
	if obj.TimeseriesWidgetFormulaRequest_20270101 != nil {
		return obj.TimeseriesWidgetFormulaRequest_20270101
	}

	if obj.TimeseriesWidgetMetricRequest_20270101 != nil {
		return obj.TimeseriesWidgetMetricRequest_20270101
	}

	if obj.TimeseriesWidgetLogRequest_20270101 != nil {
		return obj.TimeseriesWidgetLogRequest_20270101
	}

	if obj.TimeseriesWidgetApmRequest_20270101 != nil {
		return obj.TimeseriesWidgetApmRequest_20270101
	}

	if obj.TimeseriesWidgetNetworkRequest_20270101 != nil {
		return obj.TimeseriesWidgetNetworkRequest_20270101
	}

	if obj.TimeseriesWidgetRumRequest_20270101 != nil {
		return obj.TimeseriesWidgetRumRequest_20270101
	}

	if obj.TimeseriesWidgetSecurityRequest_20270101 != nil {
		return obj.TimeseriesWidgetSecurityRequest_20270101
	}

	if obj.TimeseriesWidgetComplianceFindingsRequest_20270101 != nil {
		return obj.TimeseriesWidgetComplianceFindingsRequest_20270101
	}

	if obj.TimeseriesWidgetEventRequest_20270101 != nil {
		return obj.TimeseriesWidgetEventRequest_20270101
	}

	if obj.TimeseriesWidgetProfilingMetricsRequest_20270101 != nil {
		return obj.TimeseriesWidgetProfilingMetricsRequest_20270101
	}

	if obj.TimeseriesWidgetCiPipelineRequest_20270101 != nil {
		return obj.TimeseriesWidgetCiPipelineRequest_20270101
	}

	if obj.TimeseriesWidgetCiTestRequest_20270101 != nil {
		return obj.TimeseriesWidgetCiTestRequest_20270101
	}

	if obj.TimeseriesWidgetAuditRequest_20270101 != nil {
		return obj.TimeseriesWidgetAuditRequest_20270101
	}

	if obj.TimeseriesWidgetAppsecSpanRequest_20270101 != nil {
		return obj.TimeseriesWidgetAppsecSpanRequest_20270101
	}

	if obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101 != nil {
		return obj.TimeseriesWidgetSyntheticsBatchesRequest_20270101
	}

	if obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101 != nil {
		return obj.TimeseriesWidgetSyntheticsTestRunsRequest_20270101
	}

	if obj.TimeseriesWidgetProcessRequest_20270101 != nil {
		return obj.TimeseriesWidgetProcessRequest_20270101
	}

	if obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101 != nil {
		return obj.TimeseriesWidgetLegacyProfileMetricsRequest_20270101
	}

	if obj.TimeseriesWidgetLegacyEventRequest_20270101 != nil {
		return obj.TimeseriesWidgetLegacyEventRequest_20270101
	}

	if obj.TimeseriesWidgetDdsqlRequest_20270101 != nil {
		return obj.TimeseriesWidgetDdsqlRequest_20270101
	}

	if obj.TimeseriesWidgetAdvancedQueryRequest_20270101 != nil {
		return obj.TimeseriesWidgetAdvancedQueryRequest_20270101
	}

	if obj.TimeseriesWidgetLocalDatasetRequest_20270101 != nil {
		return obj.TimeseriesWidgetLocalDatasetRequest_20270101
	}

	// all schemas are nil
	return nil
}
