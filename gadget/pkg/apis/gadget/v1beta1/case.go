package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const CaseResource = "cases"

const (
	CaseStatusError  = CaseStatusType("error")
	CaseStatusSynced = CaseStatusType("synced")
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Case `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Case struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CaseSpec   `json:"spec"`
	Status CaseStatus `json:"status,omitempty"`
}

type CaseSpec struct {
	Environment Environment `json:"environment"`
	Retries     int         `json:"retries"`
}

type CaseStatus struct {
	Message string           `json:"message"`
	Status  SolverStatusType `json:"status"`
}

type CaseStatusType string
