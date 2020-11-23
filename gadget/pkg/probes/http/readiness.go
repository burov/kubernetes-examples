package http

import (
	"github.com/burov/kubernetes-examples/gadget/pkg/apis/gadget/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
	"net/http"
	"time"

	"k8s.io/client-go/dynamic"
)

func NewReadiness(client dynamic.Interface, ns string, timeout time.Duration) http.HandlerFunc {
	sCli := client.Resource(v1beta1.SchemeGroupVersion.WithResource(v1beta1.SolverResource))
	cCli := client.Resource(v1beta1.SchemeGroupVersion.WithResource(v1beta1.CaseResource))

	return func(w http.ResponseWriter, r *http.Request) {
		t := int64(timeout.Seconds())
		opts := metav1.ListOptions{
			TimeoutSeconds: &t,
			Limit:          1,
		}

		if _, err := sCli.Namespace(ns).List(opts); err != nil {
			klog.Errorf("unable to get solvers from namespace %s, error - %s\n", ns, err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		if _, err := cCli.Namespace(ns).List(opts); err != nil {
			klog.Errorf("unable to get cases from namespace %s, error - %s\n", ns, err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
