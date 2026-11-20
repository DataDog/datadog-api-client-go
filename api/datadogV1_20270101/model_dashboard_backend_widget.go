// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardBackendWidget - Test-only generated schema for the dashboard versioning proof of concept.
type DashboardBackendWidget struct {
	AlertGraphBackendWidget_20270101     *AlertGraphBackendWidget
	AlertValueBackendWidget_20270101     *AlertValueBackendWidget
	BarChartBackendWidget_20270101       *BarChartBackendWidget
	ChangeBackendWidget_20270101         *ChangeBackendWidget
	CheckStatusBackendWidget_20270101    *CheckStatusBackendWidget
	CohortBackendWidget_20270101         *CohortBackendWidget
	DistributionBackendWidget_20270101   *DistributionBackendWidget
	EmbeddedAppBackendWidget_20270101    *EmbeddedAppBackendWidget
	EventStream_20270101                 *EventStream
	EventTimelineBackendWidget_20270101  *EventTimelineBackendWidget
	FreeTextBackendWidget_20270101       *FreeTextBackendWidget
	FunnelBackendWidget_20270101         *FunnelBackendWidget
	GeomapBackendWidget_20270101         *GeomapBackendWidget
	GroupBackendWidget_20270101          *GroupBackendWidget
	HeatgridBackendWidget_20270101       *HeatgridBackendWidget
	HeatmapBackendWidget_20270101        *HeatmapBackendWidget
	HostmapBackendWidget_20270101        *HostmapBackendWidget
	IframeBackendWidget_20270101         *IframeBackendWidget
	ImageBackendWidget_20270101          *ImageBackendWidget
	JourneyMapBackendWidget_20270101     *JourneyMapBackendWidget
	ListStreamBackendWidget_20270101     *ListStreamBackendWidget
	LogStream_20270101                   *LogStream
	ManageStatusBackendWidget_20270101   *ManageStatusBackendWidget
	NoteBackendWidget_20270101           *NoteBackendWidget
	PointPlotBackendWidget_20270101      *PointPlotBackendWidget
	PowerpackBackendWidget_20270101      *PowerpackBackendWidget
	QueryTableBackendWidget_20270101     *QueryTableBackendWidget
	QueryValueBackendWidget_20270101     *QueryValueBackendWidget
	RetentionCurveBackendWidget_20270101 *RetentionCurveBackendWidget
	RunWorkflowBackendWidget_20270101    *RunWorkflowBackendWidget
	SankeyBackendWidget_20270101         *SankeyBackendWidget
	ScatterplotBackendWidget_20270101    *ScatterplotBackendWidget
	ServiceMapBackendWidget_20270101     *ServiceMapBackendWidget
	SloBackendWidget_20270101            *SloBackendWidget
	SloListBackendWidget_20270101        *SloListBackendWidget
	SplitGroupBackendWidget_20270101     *SplitGroupBackendWidget
	SunburstBackendWidget_20270101       *SunburstBackendWidget
	TimeseriesBackendWidget_20270101     *TimeseriesBackendWidget
	ToplistBackendWidget_20270101        *ToplistBackendWidget
	TopologyMapBackendWidget_20270101    *TopologyMapBackendWidget
	ServiceSummaryBackendWidget_20270101 *ServiceSummaryBackendWidget
	TreemapBackendWidget_20270101        *TreemapBackendWidget
	WildcardBackendWidget_20270101       *WildcardBackendWidget

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// AlertGraphBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns AlertGraphBackendWidget wrapped in DashboardBackendWidget.
func AlertGraphBackendWidget_20270101AsDashboardBackendWidget(v *AlertGraphBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{AlertGraphBackendWidget_20270101: v}
}

// AlertValueBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns AlertValueBackendWidget wrapped in DashboardBackendWidget.
func AlertValueBackendWidget_20270101AsDashboardBackendWidget(v *AlertValueBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{AlertValueBackendWidget_20270101: v}
}

// BarChartBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns BarChartBackendWidget wrapped in DashboardBackendWidget.
func BarChartBackendWidget_20270101AsDashboardBackendWidget(v *BarChartBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{BarChartBackendWidget_20270101: v}
}

// ChangeBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns ChangeBackendWidget wrapped in DashboardBackendWidget.
func ChangeBackendWidget_20270101AsDashboardBackendWidget(v *ChangeBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{ChangeBackendWidget_20270101: v}
}

// CheckStatusBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns CheckStatusBackendWidget wrapped in DashboardBackendWidget.
func CheckStatusBackendWidget_20270101AsDashboardBackendWidget(v *CheckStatusBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{CheckStatusBackendWidget_20270101: v}
}

// CohortBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns CohortBackendWidget wrapped in DashboardBackendWidget.
func CohortBackendWidget_20270101AsDashboardBackendWidget(v *CohortBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{CohortBackendWidget_20270101: v}
}

// DistributionBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns DistributionBackendWidget wrapped in DashboardBackendWidget.
func DistributionBackendWidget_20270101AsDashboardBackendWidget(v *DistributionBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{DistributionBackendWidget_20270101: v}
}

// EmbeddedAppBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns EmbeddedAppBackendWidget wrapped in DashboardBackendWidget.
func EmbeddedAppBackendWidget_20270101AsDashboardBackendWidget(v *EmbeddedAppBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{EmbeddedAppBackendWidget_20270101: v}
}

// EventStream_20270101AsDashboardBackendWidget is a convenience function that returns EventStream wrapped in DashboardBackendWidget.
func EventStream_20270101AsDashboardBackendWidget(v *EventStream) DashboardBackendWidget {
	return DashboardBackendWidget{EventStream_20270101: v}
}

// EventTimelineBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns EventTimelineBackendWidget wrapped in DashboardBackendWidget.
func EventTimelineBackendWidget_20270101AsDashboardBackendWidget(v *EventTimelineBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{EventTimelineBackendWidget_20270101: v}
}

// FreeTextBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns FreeTextBackendWidget wrapped in DashboardBackendWidget.
func FreeTextBackendWidget_20270101AsDashboardBackendWidget(v *FreeTextBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{FreeTextBackendWidget_20270101: v}
}

// FunnelBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns FunnelBackendWidget wrapped in DashboardBackendWidget.
func FunnelBackendWidget_20270101AsDashboardBackendWidget(v *FunnelBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{FunnelBackendWidget_20270101: v}
}

// GeomapBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns GeomapBackendWidget wrapped in DashboardBackendWidget.
func GeomapBackendWidget_20270101AsDashboardBackendWidget(v *GeomapBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{GeomapBackendWidget_20270101: v}
}

// GroupBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns GroupBackendWidget wrapped in DashboardBackendWidget.
func GroupBackendWidget_20270101AsDashboardBackendWidget(v *GroupBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{GroupBackendWidget_20270101: v}
}

// HeatgridBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns HeatgridBackendWidget wrapped in DashboardBackendWidget.
func HeatgridBackendWidget_20270101AsDashboardBackendWidget(v *HeatgridBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{HeatgridBackendWidget_20270101: v}
}

// HeatmapBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns HeatmapBackendWidget wrapped in DashboardBackendWidget.
func HeatmapBackendWidget_20270101AsDashboardBackendWidget(v *HeatmapBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{HeatmapBackendWidget_20270101: v}
}

// HostmapBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns HostmapBackendWidget wrapped in DashboardBackendWidget.
func HostmapBackendWidget_20270101AsDashboardBackendWidget(v *HostmapBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{HostmapBackendWidget_20270101: v}
}

// IframeBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns IframeBackendWidget wrapped in DashboardBackendWidget.
func IframeBackendWidget_20270101AsDashboardBackendWidget(v *IframeBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{IframeBackendWidget_20270101: v}
}

// ImageBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns ImageBackendWidget wrapped in DashboardBackendWidget.
func ImageBackendWidget_20270101AsDashboardBackendWidget(v *ImageBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{ImageBackendWidget_20270101: v}
}

// JourneyMapBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns JourneyMapBackendWidget wrapped in DashboardBackendWidget.
func JourneyMapBackendWidget_20270101AsDashboardBackendWidget(v *JourneyMapBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{JourneyMapBackendWidget_20270101: v}
}

// ListStreamBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns ListStreamBackendWidget wrapped in DashboardBackendWidget.
func ListStreamBackendWidget_20270101AsDashboardBackendWidget(v *ListStreamBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{ListStreamBackendWidget_20270101: v}
}

// LogStream_20270101AsDashboardBackendWidget is a convenience function that returns LogStream wrapped in DashboardBackendWidget.
func LogStream_20270101AsDashboardBackendWidget(v *LogStream) DashboardBackendWidget {
	return DashboardBackendWidget{LogStream_20270101: v}
}

// ManageStatusBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns ManageStatusBackendWidget wrapped in DashboardBackendWidget.
func ManageStatusBackendWidget_20270101AsDashboardBackendWidget(v *ManageStatusBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{ManageStatusBackendWidget_20270101: v}
}

// NoteBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns NoteBackendWidget wrapped in DashboardBackendWidget.
func NoteBackendWidget_20270101AsDashboardBackendWidget(v *NoteBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{NoteBackendWidget_20270101: v}
}

// PointPlotBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns PointPlotBackendWidget wrapped in DashboardBackendWidget.
func PointPlotBackendWidget_20270101AsDashboardBackendWidget(v *PointPlotBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{PointPlotBackendWidget_20270101: v}
}

// PowerpackBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns PowerpackBackendWidget wrapped in DashboardBackendWidget.
func PowerpackBackendWidget_20270101AsDashboardBackendWidget(v *PowerpackBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{PowerpackBackendWidget_20270101: v}
}

// QueryTableBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns QueryTableBackendWidget wrapped in DashboardBackendWidget.
func QueryTableBackendWidget_20270101AsDashboardBackendWidget(v *QueryTableBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{QueryTableBackendWidget_20270101: v}
}

// QueryValueBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns QueryValueBackendWidget wrapped in DashboardBackendWidget.
func QueryValueBackendWidget_20270101AsDashboardBackendWidget(v *QueryValueBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{QueryValueBackendWidget_20270101: v}
}

// RetentionCurveBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns RetentionCurveBackendWidget wrapped in DashboardBackendWidget.
func RetentionCurveBackendWidget_20270101AsDashboardBackendWidget(v *RetentionCurveBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{RetentionCurveBackendWidget_20270101: v}
}

// RunWorkflowBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns RunWorkflowBackendWidget wrapped in DashboardBackendWidget.
func RunWorkflowBackendWidget_20270101AsDashboardBackendWidget(v *RunWorkflowBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{RunWorkflowBackendWidget_20270101: v}
}

// SankeyBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns SankeyBackendWidget wrapped in DashboardBackendWidget.
func SankeyBackendWidget_20270101AsDashboardBackendWidget(v *SankeyBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{SankeyBackendWidget_20270101: v}
}

// ScatterplotBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns ScatterplotBackendWidget wrapped in DashboardBackendWidget.
func ScatterplotBackendWidget_20270101AsDashboardBackendWidget(v *ScatterplotBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{ScatterplotBackendWidget_20270101: v}
}

// ServiceMapBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns ServiceMapBackendWidget wrapped in DashboardBackendWidget.
func ServiceMapBackendWidget_20270101AsDashboardBackendWidget(v *ServiceMapBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{ServiceMapBackendWidget_20270101: v}
}

// SloBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns SloBackendWidget wrapped in DashboardBackendWidget.
func SloBackendWidget_20270101AsDashboardBackendWidget(v *SloBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{SloBackendWidget_20270101: v}
}

// SloListBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns SloListBackendWidget wrapped in DashboardBackendWidget.
func SloListBackendWidget_20270101AsDashboardBackendWidget(v *SloListBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{SloListBackendWidget_20270101: v}
}

// SplitGroupBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns SplitGroupBackendWidget wrapped in DashboardBackendWidget.
func SplitGroupBackendWidget_20270101AsDashboardBackendWidget(v *SplitGroupBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{SplitGroupBackendWidget_20270101: v}
}

// SunburstBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns SunburstBackendWidget wrapped in DashboardBackendWidget.
func SunburstBackendWidget_20270101AsDashboardBackendWidget(v *SunburstBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{SunburstBackendWidget_20270101: v}
}

// TimeseriesBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns TimeseriesBackendWidget wrapped in DashboardBackendWidget.
func TimeseriesBackendWidget_20270101AsDashboardBackendWidget(v *TimeseriesBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{TimeseriesBackendWidget_20270101: v}
}

// ToplistBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns ToplistBackendWidget wrapped in DashboardBackendWidget.
func ToplistBackendWidget_20270101AsDashboardBackendWidget(v *ToplistBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{ToplistBackendWidget_20270101: v}
}

// TopologyMapBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns TopologyMapBackendWidget wrapped in DashboardBackendWidget.
func TopologyMapBackendWidget_20270101AsDashboardBackendWidget(v *TopologyMapBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{TopologyMapBackendWidget_20270101: v}
}

// ServiceSummaryBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns ServiceSummaryBackendWidget wrapped in DashboardBackendWidget.
func ServiceSummaryBackendWidget_20270101AsDashboardBackendWidget(v *ServiceSummaryBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{ServiceSummaryBackendWidget_20270101: v}
}

// TreemapBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns TreemapBackendWidget wrapped in DashboardBackendWidget.
func TreemapBackendWidget_20270101AsDashboardBackendWidget(v *TreemapBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{TreemapBackendWidget_20270101: v}
}

// WildcardBackendWidget_20270101AsDashboardBackendWidget is a convenience function that returns WildcardBackendWidget wrapped in DashboardBackendWidget.
func WildcardBackendWidget_20270101AsDashboardBackendWidget(v *WildcardBackendWidget) DashboardBackendWidget {
	return DashboardBackendWidget{WildcardBackendWidget_20270101: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *DashboardBackendWidget) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AlertGraphBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.AlertGraphBackendWidget_20270101)
	if err == nil {
		if obj.AlertGraphBackendWidget_20270101 != nil && obj.AlertGraphBackendWidget_20270101.UnparsedObject == nil {
			jsonAlertGraphBackendWidget_20270101, _ := datadog.Marshal(obj.AlertGraphBackendWidget_20270101)
			if string(jsonAlertGraphBackendWidget_20270101) == "{}" { // empty struct
				obj.AlertGraphBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AlertGraphBackendWidget_20270101 = nil
		}
	} else {
		obj.AlertGraphBackendWidget_20270101 = nil
	}

	// try to unmarshal data into AlertValueBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.AlertValueBackendWidget_20270101)
	if err == nil {
		if obj.AlertValueBackendWidget_20270101 != nil && obj.AlertValueBackendWidget_20270101.UnparsedObject == nil {
			jsonAlertValueBackendWidget_20270101, _ := datadog.Marshal(obj.AlertValueBackendWidget_20270101)
			if string(jsonAlertValueBackendWidget_20270101) == "{}" { // empty struct
				obj.AlertValueBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.AlertValueBackendWidget_20270101 = nil
		}
	} else {
		obj.AlertValueBackendWidget_20270101 = nil
	}

	// try to unmarshal data into BarChartBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.BarChartBackendWidget_20270101)
	if err == nil {
		if obj.BarChartBackendWidget_20270101 != nil && obj.BarChartBackendWidget_20270101.UnparsedObject == nil {
			jsonBarChartBackendWidget_20270101, _ := datadog.Marshal(obj.BarChartBackendWidget_20270101)
			if string(jsonBarChartBackendWidget_20270101) == "{}" { // empty struct
				obj.BarChartBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.BarChartBackendWidget_20270101 = nil
		}
	} else {
		obj.BarChartBackendWidget_20270101 = nil
	}

	// try to unmarshal data into ChangeBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.ChangeBackendWidget_20270101)
	if err == nil {
		if obj.ChangeBackendWidget_20270101 != nil && obj.ChangeBackendWidget_20270101.UnparsedObject == nil {
			jsonChangeBackendWidget_20270101, _ := datadog.Marshal(obj.ChangeBackendWidget_20270101)
			if string(jsonChangeBackendWidget_20270101) == "{}" { // empty struct
				obj.ChangeBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ChangeBackendWidget_20270101 = nil
		}
	} else {
		obj.ChangeBackendWidget_20270101 = nil
	}

	// try to unmarshal data into CheckStatusBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.CheckStatusBackendWidget_20270101)
	if err == nil {
		if obj.CheckStatusBackendWidget_20270101 != nil && obj.CheckStatusBackendWidget_20270101.UnparsedObject == nil {
			jsonCheckStatusBackendWidget_20270101, _ := datadog.Marshal(obj.CheckStatusBackendWidget_20270101)
			if string(jsonCheckStatusBackendWidget_20270101) == "{}" { // empty struct
				obj.CheckStatusBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CheckStatusBackendWidget_20270101 = nil
		}
	} else {
		obj.CheckStatusBackendWidget_20270101 = nil
	}

	// try to unmarshal data into CohortBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.CohortBackendWidget_20270101)
	if err == nil {
		if obj.CohortBackendWidget_20270101 != nil && obj.CohortBackendWidget_20270101.UnparsedObject == nil {
			jsonCohortBackendWidget_20270101, _ := datadog.Marshal(obj.CohortBackendWidget_20270101)
			if string(jsonCohortBackendWidget_20270101) == "{}" { // empty struct
				obj.CohortBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.CohortBackendWidget_20270101 = nil
		}
	} else {
		obj.CohortBackendWidget_20270101 = nil
	}

	// try to unmarshal data into DistributionBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.DistributionBackendWidget_20270101)
	if err == nil {
		if obj.DistributionBackendWidget_20270101 != nil && obj.DistributionBackendWidget_20270101.UnparsedObject == nil {
			jsonDistributionBackendWidget_20270101, _ := datadog.Marshal(obj.DistributionBackendWidget_20270101)
			if string(jsonDistributionBackendWidget_20270101) == "{}" { // empty struct
				obj.DistributionBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.DistributionBackendWidget_20270101 = nil
		}
	} else {
		obj.DistributionBackendWidget_20270101 = nil
	}

	// try to unmarshal data into EmbeddedAppBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.EmbeddedAppBackendWidget_20270101)
	if err == nil {
		if obj.EmbeddedAppBackendWidget_20270101 != nil && obj.EmbeddedAppBackendWidget_20270101.UnparsedObject == nil {
			jsonEmbeddedAppBackendWidget_20270101, _ := datadog.Marshal(obj.EmbeddedAppBackendWidget_20270101)
			if string(jsonEmbeddedAppBackendWidget_20270101) == "{}" { // empty struct
				obj.EmbeddedAppBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EmbeddedAppBackendWidget_20270101 = nil
		}
	} else {
		obj.EmbeddedAppBackendWidget_20270101 = nil
	}

	// try to unmarshal data into EventStream_20270101
	err = datadog.Unmarshal(data, &obj.EventStream_20270101)
	if err == nil {
		if obj.EventStream_20270101 != nil && obj.EventStream_20270101.UnparsedObject == nil {
			jsonEventStream_20270101, _ := datadog.Marshal(obj.EventStream_20270101)
			if string(jsonEventStream_20270101) == "{}" { // empty struct
				obj.EventStream_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventStream_20270101 = nil
		}
	} else {
		obj.EventStream_20270101 = nil
	}

	// try to unmarshal data into EventTimelineBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.EventTimelineBackendWidget_20270101)
	if err == nil {
		if obj.EventTimelineBackendWidget_20270101 != nil && obj.EventTimelineBackendWidget_20270101.UnparsedObject == nil {
			jsonEventTimelineBackendWidget_20270101, _ := datadog.Marshal(obj.EventTimelineBackendWidget_20270101)
			if string(jsonEventTimelineBackendWidget_20270101) == "{}" { // empty struct
				obj.EventTimelineBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.EventTimelineBackendWidget_20270101 = nil
		}
	} else {
		obj.EventTimelineBackendWidget_20270101 = nil
	}

	// try to unmarshal data into FreeTextBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.FreeTextBackendWidget_20270101)
	if err == nil {
		if obj.FreeTextBackendWidget_20270101 != nil && obj.FreeTextBackendWidget_20270101.UnparsedObject == nil {
			jsonFreeTextBackendWidget_20270101, _ := datadog.Marshal(obj.FreeTextBackendWidget_20270101)
			if string(jsonFreeTextBackendWidget_20270101) == "{}" { // empty struct
				obj.FreeTextBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FreeTextBackendWidget_20270101 = nil
		}
	} else {
		obj.FreeTextBackendWidget_20270101 = nil
	}

	// try to unmarshal data into FunnelBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.FunnelBackendWidget_20270101)
	if err == nil {
		if obj.FunnelBackendWidget_20270101 != nil && obj.FunnelBackendWidget_20270101.UnparsedObject == nil {
			jsonFunnelBackendWidget_20270101, _ := datadog.Marshal(obj.FunnelBackendWidget_20270101)
			if string(jsonFunnelBackendWidget_20270101) == "{}" { // empty struct
				obj.FunnelBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.FunnelBackendWidget_20270101 = nil
		}
	} else {
		obj.FunnelBackendWidget_20270101 = nil
	}

	// try to unmarshal data into GeomapBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.GeomapBackendWidget_20270101)
	if err == nil {
		if obj.GeomapBackendWidget_20270101 != nil && obj.GeomapBackendWidget_20270101.UnparsedObject == nil {
			jsonGeomapBackendWidget_20270101, _ := datadog.Marshal(obj.GeomapBackendWidget_20270101)
			if string(jsonGeomapBackendWidget_20270101) == "{}" { // empty struct
				obj.GeomapBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.GeomapBackendWidget_20270101 = nil
		}
	} else {
		obj.GeomapBackendWidget_20270101 = nil
	}

	// try to unmarshal data into GroupBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.GroupBackendWidget_20270101)
	if err == nil {
		if obj.GroupBackendWidget_20270101 != nil && obj.GroupBackendWidget_20270101.UnparsedObject == nil {
			jsonGroupBackendWidget_20270101, _ := datadog.Marshal(obj.GroupBackendWidget_20270101)
			if string(jsonGroupBackendWidget_20270101) == "{}" { // empty struct
				obj.GroupBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.GroupBackendWidget_20270101 = nil
		}
	} else {
		obj.GroupBackendWidget_20270101 = nil
	}

	// try to unmarshal data into HeatgridBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.HeatgridBackendWidget_20270101)
	if err == nil {
		if obj.HeatgridBackendWidget_20270101 != nil && obj.HeatgridBackendWidget_20270101.UnparsedObject == nil {
			jsonHeatgridBackendWidget_20270101, _ := datadog.Marshal(obj.HeatgridBackendWidget_20270101)
			if string(jsonHeatgridBackendWidget_20270101) == "{}" { // empty struct
				obj.HeatgridBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HeatgridBackendWidget_20270101 = nil
		}
	} else {
		obj.HeatgridBackendWidget_20270101 = nil
	}

	// try to unmarshal data into HeatmapBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.HeatmapBackendWidget_20270101)
	if err == nil {
		if obj.HeatmapBackendWidget_20270101 != nil && obj.HeatmapBackendWidget_20270101.UnparsedObject == nil {
			jsonHeatmapBackendWidget_20270101, _ := datadog.Marshal(obj.HeatmapBackendWidget_20270101)
			if string(jsonHeatmapBackendWidget_20270101) == "{}" { // empty struct
				obj.HeatmapBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HeatmapBackendWidget_20270101 = nil
		}
	} else {
		obj.HeatmapBackendWidget_20270101 = nil
	}

	// try to unmarshal data into HostmapBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.HostmapBackendWidget_20270101)
	if err == nil {
		if obj.HostmapBackendWidget_20270101 != nil && obj.HostmapBackendWidget_20270101.UnparsedObject == nil {
			jsonHostmapBackendWidget_20270101, _ := datadog.Marshal(obj.HostmapBackendWidget_20270101)
			if string(jsonHostmapBackendWidget_20270101) == "{}" { // empty struct
				obj.HostmapBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.HostmapBackendWidget_20270101 = nil
		}
	} else {
		obj.HostmapBackendWidget_20270101 = nil
	}

	// try to unmarshal data into IframeBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.IframeBackendWidget_20270101)
	if err == nil {
		if obj.IframeBackendWidget_20270101 != nil && obj.IframeBackendWidget_20270101.UnparsedObject == nil {
			jsonIframeBackendWidget_20270101, _ := datadog.Marshal(obj.IframeBackendWidget_20270101)
			if string(jsonIframeBackendWidget_20270101) == "{}" { // empty struct
				obj.IframeBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.IframeBackendWidget_20270101 = nil
		}
	} else {
		obj.IframeBackendWidget_20270101 = nil
	}

	// try to unmarshal data into ImageBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.ImageBackendWidget_20270101)
	if err == nil {
		if obj.ImageBackendWidget_20270101 != nil && obj.ImageBackendWidget_20270101.UnparsedObject == nil {
			jsonImageBackendWidget_20270101, _ := datadog.Marshal(obj.ImageBackendWidget_20270101)
			if string(jsonImageBackendWidget_20270101) == "{}" { // empty struct
				obj.ImageBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ImageBackendWidget_20270101 = nil
		}
	} else {
		obj.ImageBackendWidget_20270101 = nil
	}

	// try to unmarshal data into JourneyMapBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.JourneyMapBackendWidget_20270101)
	if err == nil {
		if obj.JourneyMapBackendWidget_20270101 != nil && obj.JourneyMapBackendWidget_20270101.UnparsedObject == nil {
			jsonJourneyMapBackendWidget_20270101, _ := datadog.Marshal(obj.JourneyMapBackendWidget_20270101)
			if string(jsonJourneyMapBackendWidget_20270101) == "{}" { // empty struct
				obj.JourneyMapBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.JourneyMapBackendWidget_20270101 = nil
		}
	} else {
		obj.JourneyMapBackendWidget_20270101 = nil
	}

	// try to unmarshal data into ListStreamBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.ListStreamBackendWidget_20270101)
	if err == nil {
		if obj.ListStreamBackendWidget_20270101 != nil && obj.ListStreamBackendWidget_20270101.UnparsedObject == nil {
			jsonListStreamBackendWidget_20270101, _ := datadog.Marshal(obj.ListStreamBackendWidget_20270101)
			if string(jsonListStreamBackendWidget_20270101) == "{}" { // empty struct
				obj.ListStreamBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ListStreamBackendWidget_20270101 = nil
		}
	} else {
		obj.ListStreamBackendWidget_20270101 = nil
	}

	// try to unmarshal data into LogStream_20270101
	err = datadog.Unmarshal(data, &obj.LogStream_20270101)
	if err == nil {
		if obj.LogStream_20270101 != nil && obj.LogStream_20270101.UnparsedObject == nil {
			jsonLogStream_20270101, _ := datadog.Marshal(obj.LogStream_20270101)
			if string(jsonLogStream_20270101) == "{}" { // empty struct
				obj.LogStream_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.LogStream_20270101 = nil
		}
	} else {
		obj.LogStream_20270101 = nil
	}

	// try to unmarshal data into ManageStatusBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.ManageStatusBackendWidget_20270101)
	if err == nil {
		if obj.ManageStatusBackendWidget_20270101 != nil && obj.ManageStatusBackendWidget_20270101.UnparsedObject == nil {
			jsonManageStatusBackendWidget_20270101, _ := datadog.Marshal(obj.ManageStatusBackendWidget_20270101)
			if string(jsonManageStatusBackendWidget_20270101) == "{}" { // empty struct
				obj.ManageStatusBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ManageStatusBackendWidget_20270101 = nil
		}
	} else {
		obj.ManageStatusBackendWidget_20270101 = nil
	}

	// try to unmarshal data into NoteBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.NoteBackendWidget_20270101)
	if err == nil {
		if obj.NoteBackendWidget_20270101 != nil && obj.NoteBackendWidget_20270101.UnparsedObject == nil {
			jsonNoteBackendWidget_20270101, _ := datadog.Marshal(obj.NoteBackendWidget_20270101)
			if string(jsonNoteBackendWidget_20270101) == "{}" { // empty struct
				obj.NoteBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.NoteBackendWidget_20270101 = nil
		}
	} else {
		obj.NoteBackendWidget_20270101 = nil
	}

	// try to unmarshal data into PointPlotBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.PointPlotBackendWidget_20270101)
	if err == nil {
		if obj.PointPlotBackendWidget_20270101 != nil && obj.PointPlotBackendWidget_20270101.UnparsedObject == nil {
			jsonPointPlotBackendWidget_20270101, _ := datadog.Marshal(obj.PointPlotBackendWidget_20270101)
			if string(jsonPointPlotBackendWidget_20270101) == "{}" { // empty struct
				obj.PointPlotBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.PointPlotBackendWidget_20270101 = nil
		}
	} else {
		obj.PointPlotBackendWidget_20270101 = nil
	}

	// try to unmarshal data into PowerpackBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.PowerpackBackendWidget_20270101)
	if err == nil {
		if obj.PowerpackBackendWidget_20270101 != nil && obj.PowerpackBackendWidget_20270101.UnparsedObject == nil {
			jsonPowerpackBackendWidget_20270101, _ := datadog.Marshal(obj.PowerpackBackendWidget_20270101)
			if string(jsonPowerpackBackendWidget_20270101) == "{}" { // empty struct
				obj.PowerpackBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.PowerpackBackendWidget_20270101 = nil
		}
	} else {
		obj.PowerpackBackendWidget_20270101 = nil
	}

	// try to unmarshal data into QueryTableBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.QueryTableBackendWidget_20270101)
	if err == nil {
		if obj.QueryTableBackendWidget_20270101 != nil && obj.QueryTableBackendWidget_20270101.UnparsedObject == nil {
			jsonQueryTableBackendWidget_20270101, _ := datadog.Marshal(obj.QueryTableBackendWidget_20270101)
			if string(jsonQueryTableBackendWidget_20270101) == "{}" { // empty struct
				obj.QueryTableBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.QueryTableBackendWidget_20270101 = nil
		}
	} else {
		obj.QueryTableBackendWidget_20270101 = nil
	}

	// try to unmarshal data into QueryValueBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.QueryValueBackendWidget_20270101)
	if err == nil {
		if obj.QueryValueBackendWidget_20270101 != nil && obj.QueryValueBackendWidget_20270101.UnparsedObject == nil {
			jsonQueryValueBackendWidget_20270101, _ := datadog.Marshal(obj.QueryValueBackendWidget_20270101)
			if string(jsonQueryValueBackendWidget_20270101) == "{}" { // empty struct
				obj.QueryValueBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.QueryValueBackendWidget_20270101 = nil
		}
	} else {
		obj.QueryValueBackendWidget_20270101 = nil
	}

	// try to unmarshal data into RetentionCurveBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.RetentionCurveBackendWidget_20270101)
	if err == nil {
		if obj.RetentionCurveBackendWidget_20270101 != nil && obj.RetentionCurveBackendWidget_20270101.UnparsedObject == nil {
			jsonRetentionCurveBackendWidget_20270101, _ := datadog.Marshal(obj.RetentionCurveBackendWidget_20270101)
			if string(jsonRetentionCurveBackendWidget_20270101) == "{}" { // empty struct
				obj.RetentionCurveBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RetentionCurveBackendWidget_20270101 = nil
		}
	} else {
		obj.RetentionCurveBackendWidget_20270101 = nil
	}

	// try to unmarshal data into RunWorkflowBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.RunWorkflowBackendWidget_20270101)
	if err == nil {
		if obj.RunWorkflowBackendWidget_20270101 != nil && obj.RunWorkflowBackendWidget_20270101.UnparsedObject == nil {
			jsonRunWorkflowBackendWidget_20270101, _ := datadog.Marshal(obj.RunWorkflowBackendWidget_20270101)
			if string(jsonRunWorkflowBackendWidget_20270101) == "{}" { // empty struct
				obj.RunWorkflowBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.RunWorkflowBackendWidget_20270101 = nil
		}
	} else {
		obj.RunWorkflowBackendWidget_20270101 = nil
	}

	// try to unmarshal data into SankeyBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.SankeyBackendWidget_20270101)
	if err == nil {
		if obj.SankeyBackendWidget_20270101 != nil && obj.SankeyBackendWidget_20270101.UnparsedObject == nil {
			jsonSankeyBackendWidget_20270101, _ := datadog.Marshal(obj.SankeyBackendWidget_20270101)
			if string(jsonSankeyBackendWidget_20270101) == "{}" { // empty struct
				obj.SankeyBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SankeyBackendWidget_20270101 = nil
		}
	} else {
		obj.SankeyBackendWidget_20270101 = nil
	}

	// try to unmarshal data into ScatterplotBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.ScatterplotBackendWidget_20270101)
	if err == nil {
		if obj.ScatterplotBackendWidget_20270101 != nil && obj.ScatterplotBackendWidget_20270101.UnparsedObject == nil {
			jsonScatterplotBackendWidget_20270101, _ := datadog.Marshal(obj.ScatterplotBackendWidget_20270101)
			if string(jsonScatterplotBackendWidget_20270101) == "{}" { // empty struct
				obj.ScatterplotBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ScatterplotBackendWidget_20270101 = nil
		}
	} else {
		obj.ScatterplotBackendWidget_20270101 = nil
	}

	// try to unmarshal data into ServiceMapBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.ServiceMapBackendWidget_20270101)
	if err == nil {
		if obj.ServiceMapBackendWidget_20270101 != nil && obj.ServiceMapBackendWidget_20270101.UnparsedObject == nil {
			jsonServiceMapBackendWidget_20270101, _ := datadog.Marshal(obj.ServiceMapBackendWidget_20270101)
			if string(jsonServiceMapBackendWidget_20270101) == "{}" { // empty struct
				obj.ServiceMapBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ServiceMapBackendWidget_20270101 = nil
		}
	} else {
		obj.ServiceMapBackendWidget_20270101 = nil
	}

	// try to unmarshal data into SloBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.SloBackendWidget_20270101)
	if err == nil {
		if obj.SloBackendWidget_20270101 != nil && obj.SloBackendWidget_20270101.UnparsedObject == nil {
			jsonSloBackendWidget_20270101, _ := datadog.Marshal(obj.SloBackendWidget_20270101)
			if string(jsonSloBackendWidget_20270101) == "{}" { // empty struct
				obj.SloBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SloBackendWidget_20270101 = nil
		}
	} else {
		obj.SloBackendWidget_20270101 = nil
	}

	// try to unmarshal data into SloListBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.SloListBackendWidget_20270101)
	if err == nil {
		if obj.SloListBackendWidget_20270101 != nil && obj.SloListBackendWidget_20270101.UnparsedObject == nil {
			jsonSloListBackendWidget_20270101, _ := datadog.Marshal(obj.SloListBackendWidget_20270101)
			if string(jsonSloListBackendWidget_20270101) == "{}" { // empty struct
				obj.SloListBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SloListBackendWidget_20270101 = nil
		}
	} else {
		obj.SloListBackendWidget_20270101 = nil
	}

	// try to unmarshal data into SplitGroupBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.SplitGroupBackendWidget_20270101)
	if err == nil {
		if obj.SplitGroupBackendWidget_20270101 != nil && obj.SplitGroupBackendWidget_20270101.UnparsedObject == nil {
			jsonSplitGroupBackendWidget_20270101, _ := datadog.Marshal(obj.SplitGroupBackendWidget_20270101)
			if string(jsonSplitGroupBackendWidget_20270101) == "{}" { // empty struct
				obj.SplitGroupBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SplitGroupBackendWidget_20270101 = nil
		}
	} else {
		obj.SplitGroupBackendWidget_20270101 = nil
	}

	// try to unmarshal data into SunburstBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.SunburstBackendWidget_20270101)
	if err == nil {
		if obj.SunburstBackendWidget_20270101 != nil && obj.SunburstBackendWidget_20270101.UnparsedObject == nil {
			jsonSunburstBackendWidget_20270101, _ := datadog.Marshal(obj.SunburstBackendWidget_20270101)
			if string(jsonSunburstBackendWidget_20270101) == "{}" { // empty struct
				obj.SunburstBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.SunburstBackendWidget_20270101 = nil
		}
	} else {
		obj.SunburstBackendWidget_20270101 = nil
	}

	// try to unmarshal data into TimeseriesBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.TimeseriesBackendWidget_20270101)
	if err == nil {
		if obj.TimeseriesBackendWidget_20270101 != nil && obj.TimeseriesBackendWidget_20270101.UnparsedObject == nil {
			jsonTimeseriesBackendWidget_20270101, _ := datadog.Marshal(obj.TimeseriesBackendWidget_20270101)
			if string(jsonTimeseriesBackendWidget_20270101) == "{}" { // empty struct
				obj.TimeseriesBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TimeseriesBackendWidget_20270101 = nil
		}
	} else {
		obj.TimeseriesBackendWidget_20270101 = nil
	}

	// try to unmarshal data into ToplistBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.ToplistBackendWidget_20270101)
	if err == nil {
		if obj.ToplistBackendWidget_20270101 != nil && obj.ToplistBackendWidget_20270101.UnparsedObject == nil {
			jsonToplistBackendWidget_20270101, _ := datadog.Marshal(obj.ToplistBackendWidget_20270101)
			if string(jsonToplistBackendWidget_20270101) == "{}" { // empty struct
				obj.ToplistBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ToplistBackendWidget_20270101 = nil
		}
	} else {
		obj.ToplistBackendWidget_20270101 = nil
	}

	// try to unmarshal data into TopologyMapBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.TopologyMapBackendWidget_20270101)
	if err == nil {
		if obj.TopologyMapBackendWidget_20270101 != nil && obj.TopologyMapBackendWidget_20270101.UnparsedObject == nil {
			jsonTopologyMapBackendWidget_20270101, _ := datadog.Marshal(obj.TopologyMapBackendWidget_20270101)
			if string(jsonTopologyMapBackendWidget_20270101) == "{}" { // empty struct
				obj.TopologyMapBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TopologyMapBackendWidget_20270101 = nil
		}
	} else {
		obj.TopologyMapBackendWidget_20270101 = nil
	}

	// try to unmarshal data into ServiceSummaryBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.ServiceSummaryBackendWidget_20270101)
	if err == nil {
		if obj.ServiceSummaryBackendWidget_20270101 != nil && obj.ServiceSummaryBackendWidget_20270101.UnparsedObject == nil {
			jsonServiceSummaryBackendWidget_20270101, _ := datadog.Marshal(obj.ServiceSummaryBackendWidget_20270101)
			if string(jsonServiceSummaryBackendWidget_20270101) == "{}" { // empty struct
				obj.ServiceSummaryBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.ServiceSummaryBackendWidget_20270101 = nil
		}
	} else {
		obj.ServiceSummaryBackendWidget_20270101 = nil
	}

	// try to unmarshal data into TreemapBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.TreemapBackendWidget_20270101)
	if err == nil {
		if obj.TreemapBackendWidget_20270101 != nil && obj.TreemapBackendWidget_20270101.UnparsedObject == nil {
			jsonTreemapBackendWidget_20270101, _ := datadog.Marshal(obj.TreemapBackendWidget_20270101)
			if string(jsonTreemapBackendWidget_20270101) == "{}" { // empty struct
				obj.TreemapBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.TreemapBackendWidget_20270101 = nil
		}
	} else {
		obj.TreemapBackendWidget_20270101 = nil
	}

	// try to unmarshal data into WildcardBackendWidget_20270101
	err = datadog.Unmarshal(data, &obj.WildcardBackendWidget_20270101)
	if err == nil {
		if obj.WildcardBackendWidget_20270101 != nil && obj.WildcardBackendWidget_20270101.UnparsedObject == nil {
			jsonWildcardBackendWidget_20270101, _ := datadog.Marshal(obj.WildcardBackendWidget_20270101)
			if string(jsonWildcardBackendWidget_20270101) == "{}" { // empty struct
				obj.WildcardBackendWidget_20270101 = nil
			} else {
				match++
			}
		} else {
			obj.WildcardBackendWidget_20270101 = nil
		}
	} else {
		obj.WildcardBackendWidget_20270101 = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.AlertGraphBackendWidget_20270101 = nil
		obj.AlertValueBackendWidget_20270101 = nil
		obj.BarChartBackendWidget_20270101 = nil
		obj.ChangeBackendWidget_20270101 = nil
		obj.CheckStatusBackendWidget_20270101 = nil
		obj.CohortBackendWidget_20270101 = nil
		obj.DistributionBackendWidget_20270101 = nil
		obj.EmbeddedAppBackendWidget_20270101 = nil
		obj.EventStream_20270101 = nil
		obj.EventTimelineBackendWidget_20270101 = nil
		obj.FreeTextBackendWidget_20270101 = nil
		obj.FunnelBackendWidget_20270101 = nil
		obj.GeomapBackendWidget_20270101 = nil
		obj.GroupBackendWidget_20270101 = nil
		obj.HeatgridBackendWidget_20270101 = nil
		obj.HeatmapBackendWidget_20270101 = nil
		obj.HostmapBackendWidget_20270101 = nil
		obj.IframeBackendWidget_20270101 = nil
		obj.ImageBackendWidget_20270101 = nil
		obj.JourneyMapBackendWidget_20270101 = nil
		obj.ListStreamBackendWidget_20270101 = nil
		obj.LogStream_20270101 = nil
		obj.ManageStatusBackendWidget_20270101 = nil
		obj.NoteBackendWidget_20270101 = nil
		obj.PointPlotBackendWidget_20270101 = nil
		obj.PowerpackBackendWidget_20270101 = nil
		obj.QueryTableBackendWidget_20270101 = nil
		obj.QueryValueBackendWidget_20270101 = nil
		obj.RetentionCurveBackendWidget_20270101 = nil
		obj.RunWorkflowBackendWidget_20270101 = nil
		obj.SankeyBackendWidget_20270101 = nil
		obj.ScatterplotBackendWidget_20270101 = nil
		obj.ServiceMapBackendWidget_20270101 = nil
		obj.SloBackendWidget_20270101 = nil
		obj.SloListBackendWidget_20270101 = nil
		obj.SplitGroupBackendWidget_20270101 = nil
		obj.SunburstBackendWidget_20270101 = nil
		obj.TimeseriesBackendWidget_20270101 = nil
		obj.ToplistBackendWidget_20270101 = nil
		obj.TopologyMapBackendWidget_20270101 = nil
		obj.ServiceSummaryBackendWidget_20270101 = nil
		obj.TreemapBackendWidget_20270101 = nil
		obj.WildcardBackendWidget_20270101 = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj DashboardBackendWidget) MarshalJSON() ([]byte, error) {
	if obj.AlertGraphBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.AlertGraphBackendWidget_20270101)
	}

	if obj.AlertValueBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.AlertValueBackendWidget_20270101)
	}

	if obj.BarChartBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.BarChartBackendWidget_20270101)
	}

	if obj.ChangeBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.ChangeBackendWidget_20270101)
	}

	if obj.CheckStatusBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.CheckStatusBackendWidget_20270101)
	}

	if obj.CohortBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.CohortBackendWidget_20270101)
	}

	if obj.DistributionBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.DistributionBackendWidget_20270101)
	}

	if obj.EmbeddedAppBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.EmbeddedAppBackendWidget_20270101)
	}

	if obj.EventStream_20270101 != nil {
		return datadog.Marshal(&obj.EventStream_20270101)
	}

	if obj.EventTimelineBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.EventTimelineBackendWidget_20270101)
	}

	if obj.FreeTextBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.FreeTextBackendWidget_20270101)
	}

	if obj.FunnelBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.FunnelBackendWidget_20270101)
	}

	if obj.GeomapBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.GeomapBackendWidget_20270101)
	}

	if obj.GroupBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.GroupBackendWidget_20270101)
	}

	if obj.HeatgridBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.HeatgridBackendWidget_20270101)
	}

	if obj.HeatmapBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.HeatmapBackendWidget_20270101)
	}

	if obj.HostmapBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.HostmapBackendWidget_20270101)
	}

	if obj.IframeBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.IframeBackendWidget_20270101)
	}

	if obj.ImageBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.ImageBackendWidget_20270101)
	}

	if obj.JourneyMapBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.JourneyMapBackendWidget_20270101)
	}

	if obj.ListStreamBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.ListStreamBackendWidget_20270101)
	}

	if obj.LogStream_20270101 != nil {
		return datadog.Marshal(&obj.LogStream_20270101)
	}

	if obj.ManageStatusBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.ManageStatusBackendWidget_20270101)
	}

	if obj.NoteBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.NoteBackendWidget_20270101)
	}

	if obj.PointPlotBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.PointPlotBackendWidget_20270101)
	}

	if obj.PowerpackBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.PowerpackBackendWidget_20270101)
	}

	if obj.QueryTableBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.QueryTableBackendWidget_20270101)
	}

	if obj.QueryValueBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.QueryValueBackendWidget_20270101)
	}

	if obj.RetentionCurveBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.RetentionCurveBackendWidget_20270101)
	}

	if obj.RunWorkflowBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.RunWorkflowBackendWidget_20270101)
	}

	if obj.SankeyBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.SankeyBackendWidget_20270101)
	}

	if obj.ScatterplotBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.ScatterplotBackendWidget_20270101)
	}

	if obj.ServiceMapBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.ServiceMapBackendWidget_20270101)
	}

	if obj.SloBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.SloBackendWidget_20270101)
	}

	if obj.SloListBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.SloListBackendWidget_20270101)
	}

	if obj.SplitGroupBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.SplitGroupBackendWidget_20270101)
	}

	if obj.SunburstBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.SunburstBackendWidget_20270101)
	}

	if obj.TimeseriesBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.TimeseriesBackendWidget_20270101)
	}

	if obj.ToplistBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.ToplistBackendWidget_20270101)
	}

	if obj.TopologyMapBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.TopologyMapBackendWidget_20270101)
	}

	if obj.ServiceSummaryBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.ServiceSummaryBackendWidget_20270101)
	}

	if obj.TreemapBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.TreemapBackendWidget_20270101)
	}

	if obj.WildcardBackendWidget_20270101 != nil {
		return datadog.Marshal(&obj.WildcardBackendWidget_20270101)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *DashboardBackendWidget) GetActualInstance() interface{} {
	if obj.AlertGraphBackendWidget_20270101 != nil {
		return obj.AlertGraphBackendWidget_20270101
	}

	if obj.AlertValueBackendWidget_20270101 != nil {
		return obj.AlertValueBackendWidget_20270101
	}

	if obj.BarChartBackendWidget_20270101 != nil {
		return obj.BarChartBackendWidget_20270101
	}

	if obj.ChangeBackendWidget_20270101 != nil {
		return obj.ChangeBackendWidget_20270101
	}

	if obj.CheckStatusBackendWidget_20270101 != nil {
		return obj.CheckStatusBackendWidget_20270101
	}

	if obj.CohortBackendWidget_20270101 != nil {
		return obj.CohortBackendWidget_20270101
	}

	if obj.DistributionBackendWidget_20270101 != nil {
		return obj.DistributionBackendWidget_20270101
	}

	if obj.EmbeddedAppBackendWidget_20270101 != nil {
		return obj.EmbeddedAppBackendWidget_20270101
	}

	if obj.EventStream_20270101 != nil {
		return obj.EventStream_20270101
	}

	if obj.EventTimelineBackendWidget_20270101 != nil {
		return obj.EventTimelineBackendWidget_20270101
	}

	if obj.FreeTextBackendWidget_20270101 != nil {
		return obj.FreeTextBackendWidget_20270101
	}

	if obj.FunnelBackendWidget_20270101 != nil {
		return obj.FunnelBackendWidget_20270101
	}

	if obj.GeomapBackendWidget_20270101 != nil {
		return obj.GeomapBackendWidget_20270101
	}

	if obj.GroupBackendWidget_20270101 != nil {
		return obj.GroupBackendWidget_20270101
	}

	if obj.HeatgridBackendWidget_20270101 != nil {
		return obj.HeatgridBackendWidget_20270101
	}

	if obj.HeatmapBackendWidget_20270101 != nil {
		return obj.HeatmapBackendWidget_20270101
	}

	if obj.HostmapBackendWidget_20270101 != nil {
		return obj.HostmapBackendWidget_20270101
	}

	if obj.IframeBackendWidget_20270101 != nil {
		return obj.IframeBackendWidget_20270101
	}

	if obj.ImageBackendWidget_20270101 != nil {
		return obj.ImageBackendWidget_20270101
	}

	if obj.JourneyMapBackendWidget_20270101 != nil {
		return obj.JourneyMapBackendWidget_20270101
	}

	if obj.ListStreamBackendWidget_20270101 != nil {
		return obj.ListStreamBackendWidget_20270101
	}

	if obj.LogStream_20270101 != nil {
		return obj.LogStream_20270101
	}

	if obj.ManageStatusBackendWidget_20270101 != nil {
		return obj.ManageStatusBackendWidget_20270101
	}

	if obj.NoteBackendWidget_20270101 != nil {
		return obj.NoteBackendWidget_20270101
	}

	if obj.PointPlotBackendWidget_20270101 != nil {
		return obj.PointPlotBackendWidget_20270101
	}

	if obj.PowerpackBackendWidget_20270101 != nil {
		return obj.PowerpackBackendWidget_20270101
	}

	if obj.QueryTableBackendWidget_20270101 != nil {
		return obj.QueryTableBackendWidget_20270101
	}

	if obj.QueryValueBackendWidget_20270101 != nil {
		return obj.QueryValueBackendWidget_20270101
	}

	if obj.RetentionCurveBackendWidget_20270101 != nil {
		return obj.RetentionCurveBackendWidget_20270101
	}

	if obj.RunWorkflowBackendWidget_20270101 != nil {
		return obj.RunWorkflowBackendWidget_20270101
	}

	if obj.SankeyBackendWidget_20270101 != nil {
		return obj.SankeyBackendWidget_20270101
	}

	if obj.ScatterplotBackendWidget_20270101 != nil {
		return obj.ScatterplotBackendWidget_20270101
	}

	if obj.ServiceMapBackendWidget_20270101 != nil {
		return obj.ServiceMapBackendWidget_20270101
	}

	if obj.SloBackendWidget_20270101 != nil {
		return obj.SloBackendWidget_20270101
	}

	if obj.SloListBackendWidget_20270101 != nil {
		return obj.SloListBackendWidget_20270101
	}

	if obj.SplitGroupBackendWidget_20270101 != nil {
		return obj.SplitGroupBackendWidget_20270101
	}

	if obj.SunburstBackendWidget_20270101 != nil {
		return obj.SunburstBackendWidget_20270101
	}

	if obj.TimeseriesBackendWidget_20270101 != nil {
		return obj.TimeseriesBackendWidget_20270101
	}

	if obj.ToplistBackendWidget_20270101 != nil {
		return obj.ToplistBackendWidget_20270101
	}

	if obj.TopologyMapBackendWidget_20270101 != nil {
		return obj.TopologyMapBackendWidget_20270101
	}

	if obj.ServiceSummaryBackendWidget_20270101 != nil {
		return obj.ServiceSummaryBackendWidget_20270101
	}

	if obj.TreemapBackendWidget_20270101 != nil {
		return obj.TreemapBackendWidget_20270101
	}

	if obj.WildcardBackendWidget_20270101 != nil {
		return obj.WildcardBackendWidget_20270101
	}

	// all schemas are nil
	return nil
}
