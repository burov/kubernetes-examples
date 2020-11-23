package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const SolverResource = "solvers"

const (
	SolverStatusError  = SolverStatusType("error")
	SolverStatusSynced = SolverStatusType("synced")
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Solver `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Solver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SolverSpec   `json:"spec"`
	Status SolverStatus `json:"status,omitempty"`
}

type SolverSpec struct {
	Environment Environment `json:"environment"`
	Retries     int         `json:"retries"`
}

type Environment struct {
	Image   string `json:"image"`
	Command string
}

type SolverStatus struct {
	Message string           `json:"message"`
	Status  SolverStatusType `json:"status"`
}

type SolverStatusType string
