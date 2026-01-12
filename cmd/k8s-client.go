package cmd

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getKubeConfig() (clientcmd.ClientConfig, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides), nil
}

func getKubeClient() (*kubernetes.Clientset, string, error) {
	clientConfig, err := getKubeConfig()
	if err != nil {
		return nil, "", err
	}

	config, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, "", err
	}

	rawConfig, err := clientConfig.RawConfig()
	if err != nil {
		return nil, "", err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, "", err
	}

	return clientset, rawConfig.CurrentContext, nil
}
