package main

import(
	"fmt"
	"log"
	"flag"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	KubeConfig = flag.String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
)

func main(){
	flag.Parse()

	// create kubernetes client
	client, err := newClient(*KubeConfig)
	if err != nil {
		log.Fatal(err)
	}

	// get all service in all namespace
	services, err := client.CoreV1().Services("").List(meta_v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	isAnyService := false
	for _, service := range services.Items {
		if service.Spec.Type == "NodePort" {
			// a service may have multiple ports
			for _, port := range service.Spec.Ports {
				portName := "no name"
				if port.Name != "" {
					portName = port.Name
				}
				isAnyService = true
				fmt.Printf("%d: %s (in %s)\n", port.NodePort, portName, service.Name)
			}
		}
	}
	if(!isAnyService) {
		fmt.Println("*) all NodePort are not used")
	}
}

func newClient(kubeConfigPath string) (kubernetes.Interface, error) {
	if kubeConfigPath == "" {
		// use default path(.kube/config) when kubeconfig path is not set
		kubeConfigPath = clientcmd.RecommendedHomeFile
	}

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}
