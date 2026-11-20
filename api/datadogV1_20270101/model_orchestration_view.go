// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1_20270101

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrchestrationView Test-only generated schema for the dashboard versioning proof of concept.
type OrchestrationView string

// List of OrchestrationView.
const (
	ORCHESTRATIONVIEW_SERVICE                 OrchestrationView = "service"
	ORCHESTRATIONVIEW_CLUSTER                 OrchestrationView = "cluster"
	ORCHESTRATIONVIEW_JOB                     OrchestrationView = "job"
	ORCHESTRATIONVIEW_NODE                    OrchestrationView = "node"
	ORCHESTRATIONVIEW_DEPLOYMENT              OrchestrationView = "deployment"
	ORCHESTRATIONVIEW_NAMESPACE               OrchestrationView = "namespace"
	ORCHESTRATIONVIEW_POD                     OrchestrationView = "pod"
	ORCHESTRATIONVIEW_GENERICWORKLOAD         OrchestrationView = "genericWorkload"
	ORCHESTRATIONVIEW_INFERENCEWORKLOAD       OrchestrationView = "inferenceWorkload"
	ORCHESTRATIONVIEW_REPLICASET              OrchestrationView = "replicaSet"
	ORCHESTRATIONVIEW_CRONJOB                 OrchestrationView = "cronJob"
	ORCHESTRATIONVIEW_DAEMONSET               OrchestrationView = "daemonSet"
	ORCHESTRATIONVIEW_STATEFULSET             OrchestrationView = "statefulSet"
	ORCHESTRATIONVIEW_PERSISTENTVOLUME        OrchestrationView = "persistentVolume"
	ORCHESTRATIONVIEW_PERSISTENTVOLUMECLAIM   OrchestrationView = "persistentVolumeClaim"
	ORCHESTRATIONVIEW_ROLE                    OrchestrationView = "role"
	ORCHESTRATIONVIEW_ROLEBINDING             OrchestrationView = "roleBinding"
	ORCHESTRATIONVIEW_CLUSTERROLE             OrchestrationView = "clusterRole"
	ORCHESTRATIONVIEW_CLUSTERROLEBINDING      OrchestrationView = "clusterRoleBinding"
	ORCHESTRATIONVIEW_SERVICEACCOUNT          OrchestrationView = "serviceAccount"
	ORCHESTRATIONVIEW_INGRESS                 OrchestrationView = "ingress"
	ORCHESTRATIONVIEW_CONFIGMAP               OrchestrationView = "configMap"
	ORCHESTRATIONVIEW_VERTICALPODAUTOSCALER   OrchestrationView = "verticalPodAutoscaler"
	ORCHESTRATIONVIEW_HORIZONTALPODAUTOSCALER OrchestrationView = "horizontalPodAutoscaler"
	ORCHESTRATIONVIEW_CRD                     OrchestrationView = "crd"
	ORCHESTRATIONVIEW_CR                      OrchestrationView = "cr"
	ORCHESTRATIONVIEW_NETWORKPOLICY           OrchestrationView = "networkPolicy"
	ORCHESTRATIONVIEW_PODDISRUPTIONBUDGET     OrchestrationView = "podDisruptionBudget"
	ORCHESTRATIONVIEW_ECSTASK                 OrchestrationView = "ecsTask"
	ORCHESTRATIONVIEW_ECSTASKDEFINITION       OrchestrationView = "ecsTaskDefinition"
	ORCHESTRATIONVIEW_ECSDAEMONTASKDEFINITION OrchestrationView = "ecsDaemonTaskDefinition"
	ORCHESTRATIONVIEW_ECSSERVICE              OrchestrationView = "ecsService"
	ORCHESTRATIONVIEW_ECSDAEMON               OrchestrationView = "ecsDaemon"
	ORCHESTRATIONVIEW_ECSCLUSTER              OrchestrationView = "ecsCluster"
	ORCHESTRATIONVIEW_ECSCONTAINERINSTANCE    OrchestrationView = "ecsContainerInstance"
)

var allowedOrchestrationViewEnumValues = []OrchestrationView{
	ORCHESTRATIONVIEW_SERVICE,
	ORCHESTRATIONVIEW_CLUSTER,
	ORCHESTRATIONVIEW_JOB,
	ORCHESTRATIONVIEW_NODE,
	ORCHESTRATIONVIEW_DEPLOYMENT,
	ORCHESTRATIONVIEW_NAMESPACE,
	ORCHESTRATIONVIEW_POD,
	ORCHESTRATIONVIEW_GENERICWORKLOAD,
	ORCHESTRATIONVIEW_INFERENCEWORKLOAD,
	ORCHESTRATIONVIEW_REPLICASET,
	ORCHESTRATIONVIEW_CRONJOB,
	ORCHESTRATIONVIEW_DAEMONSET,
	ORCHESTRATIONVIEW_STATEFULSET,
	ORCHESTRATIONVIEW_PERSISTENTVOLUME,
	ORCHESTRATIONVIEW_PERSISTENTVOLUMECLAIM,
	ORCHESTRATIONVIEW_ROLE,
	ORCHESTRATIONVIEW_ROLEBINDING,
	ORCHESTRATIONVIEW_CLUSTERROLE,
	ORCHESTRATIONVIEW_CLUSTERROLEBINDING,
	ORCHESTRATIONVIEW_SERVICEACCOUNT,
	ORCHESTRATIONVIEW_INGRESS,
	ORCHESTRATIONVIEW_CONFIGMAP,
	ORCHESTRATIONVIEW_VERTICALPODAUTOSCALER,
	ORCHESTRATIONVIEW_HORIZONTALPODAUTOSCALER,
	ORCHESTRATIONVIEW_CRD,
	ORCHESTRATIONVIEW_CR,
	ORCHESTRATIONVIEW_NETWORKPOLICY,
	ORCHESTRATIONVIEW_PODDISRUPTIONBUDGET,
	ORCHESTRATIONVIEW_ECSTASK,
	ORCHESTRATIONVIEW_ECSTASKDEFINITION,
	ORCHESTRATIONVIEW_ECSDAEMONTASKDEFINITION,
	ORCHESTRATIONVIEW_ECSSERVICE,
	ORCHESTRATIONVIEW_ECSDAEMON,
	ORCHESTRATIONVIEW_ECSCLUSTER,
	ORCHESTRATIONVIEW_ECSCONTAINERINSTANCE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *OrchestrationView) GetAllowedValues() []OrchestrationView {
	return allowedOrchestrationViewEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *OrchestrationView) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = OrchestrationView(value)
	return nil
}

// NewOrchestrationViewFromValue returns a pointer to a valid OrchestrationView
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewOrchestrationViewFromValue(v string) (*OrchestrationView, error) {
	ev := OrchestrationView(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for OrchestrationView: valid values are %v", v, allowedOrchestrationViewEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v OrchestrationView) IsValid() bool {
	for _, existing := range allowedOrchestrationViewEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrchestrationView value.
func (v OrchestrationView) Ptr() *OrchestrationView {
	return &v
}
