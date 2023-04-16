package k8simple

import (
	"os"

	"k8s.io/client-go/tools/clientcmd"
)

func switchNamespace(namespace string) (status bool) {
	kubeconfig := os.Getenv("~/.kube/config")
	// rules := clientcmd.NewDefaultClientConfigLoadingRules()
	// config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{}).RawConfig()
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	config.Contexts[config.CurrentContext].Namespace = namespace

	// clientSet, err := kubernetes.NewForConfig(config)
	// currentContext := config.CurrentContext

	// if err != nil {
	// 	panic(err.Error())
	// }

	// config.CurrentContext = "test"

	// config = &rest.Config{
	// 	Context: clientcmd.ContextOverrideFlags( ,,Namespace: {"Default": namespace})
	// }

	// clientSet.CoreV1().Namespaces()

	return true
}

