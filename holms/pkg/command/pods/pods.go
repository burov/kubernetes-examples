package pods

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Command struct {}

const (
	deployment  = "deployment"
	service     = "service"
	daemonset   = "daemonset"
	replicaset  = "replicaset"
	statefulset = "statefulset"

	pluralSuffix = "s"
)

var (
	namespace, resource string
)

func init() {
	pflag.StringVar(&namespace, "namespace", "tagging-v2", "--namespace kube-system")
	pflag.StringVar(&resource, "resource", "replicaset", "--resource deployments/app")
}



func (c Command) Execute(args []string) error {
	pflag.Parse()

	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return fmt.Errorf("error find kuberntes config on path - %s, error - %w", kubeconfig, err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("error init kubernetes client path - %s, error - %w", kubeconfig, err)
	}

	labels, err := getSelector(clientset, namespace, resource)
	if err != nil {
		return fmt.Errorf("unable to query pod selector, namespace - %s, resource - %s, error - %w", namespace, resource, err)
	}

	query := buildQueryString(labels)
	pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: query})
	if err != nil {
		return fmt.Errorf("error to query pods with selector - %s, error - %w", query, err)
	}

	fmt.Printf("Namespace - %s\n", namespace)
	fmt.Printf("Resource - %s\n", resource)
	fmt.Println("Pods:")
	for _, p := range pods.Items {
		fmt.Printf("\t%s\n", p.Name)
	}

	return nil
}

func getSelector(clientset kubernetes.Interface, namespace, resource string) (map[string]string, error) {
	ttype, name := strings.Split(resource, "/")[0], strings.Split(resource, "/")[1]
	switch ttype {
	case deployment, deployment + pluralSuffix:
		d, err := clientset.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get deployment, error - %w", err)
		}
		return d.Spec.Selector.MatchLabels, nil
	case service, service + pluralSuffix:
		s, err := clientset.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get service, error - %w", err)
		}

		return s.Spec.Selector, nil
	case replicaset, replicaset + pluralSuffix:
		s, err := clientset.AppsV1().ReplicaSets(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get replicaset, error - %w", err)
		}
		return s.Spec.Selector.MatchLabels, nil
	case daemonset, daemonset + pluralSuffix:
		d, err := clientset.AppsV1().DaemonSets(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get daemonset, error - %w", err)
		}
		return d.Spec.Selector.MatchLabels, nil
	case statefulset, statefulset + pluralSuffix:
		s, err := clientset.AppsV1().StatefulSets(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get statefulset, error - %w", err)
		}
		return s.Spec.Selector.MatchLabels, nil
	default:
		return nil, fmt.Errorf("unsupported resource type %s", ttype)
	}
}

func buildQueryString(labels map[string]string) string {
	str := ""

	for k, v := range labels {
		str += fmt.Sprintf("%s=%s,", k, v)
	}

	return strings.TrimSuffix(str, ",")
}
