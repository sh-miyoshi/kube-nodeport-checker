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

	fmt.Println("Used NodePort List")
	for _, service := range services.Items {
		if service.Spec.Type == "NodePort" {
			for _, port := range service.Spec.Ports {
				portName := "no name"
				if port.Name != "" {
					portName = port.Name
				}
				fmt.Printf("%d: %s (in %s)\n", port.NodePort, portName, service.Name)
			}
		}
	}
}

func newClient(kubeConfigPath string) (kubernetes.Interface, error) {
	if kubeConfigPath == "" {
		kubeConfigPath = clientcmd.RecommendedHomeFile
	}

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(kubeConfig)
}