package k8simple

import (
	"strconv",
	"os",

	"k8s.io/client-go/kubernetes",
	"k8s.io/client-go/tools/clientcmd"
	"github.com/SeethalakshmiB/k8simple/pkg/k8simple"
)

func switchNamespace(namespace string) (status bool) {
	kubeconfig := os.GetEnv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	clientset.CoreV1().RESTClient().SetNamespace(namespace)

	return true
}