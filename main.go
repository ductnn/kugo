package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Kube config path
func GetKubeConfigPath() string {
	var kubeConfigPath string
	homeDir := homedir.HomeDir()

	if _, err := os.Stat(homeDir + "/.kube/config"); err == nil {
		kubeConfigPath = homeDir + "/.kube/config"
	} else {
		fmt.Println("Enter kubernetes config directory: ")
		fmt.Scanf("%s", kubeConfigPath)
	}

	return kubeConfigPath
}

func main() {
	// Set Kube config
	kubeConfigPath := GetKubeConfigPath()
	fmt.Println(kubeConfigPath)

	// Build configuration from config file
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err)
	}

	// Create clientser
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	deploymentName := flag.String("deployment", "", "deployment name")
	imageName := flag.String("image", "", "new image name")
	appName := flag.String("app", "app", "application name")

	flag.Parse()

	if *deploymentName == "" {
		fmt.Println("The deployment name must be specify.")
		os.Exit(0)
	}
	if *imageName == "" {
		fmt.Println("The new image name must be specify.")
		os.Exit(0)
	}

	// Create Deployment
	deployment, err := clientset.AppsV1().Deployments("default").Get(context.TODO(), *deploymentName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	// fmt.Println(deployment)

	if errors.IsNotFound(err) {
		fmt.Printf("Deployment not found\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting deployment%v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("Found deployment\n")

		name := deployment.GetName()
		fmt.Println("Name deployment: ", name)

		containers := &deployment.Spec.Template.Spec.Containers

		found := false

		for i := range *containers {
			c := *containers
			if c[i].Name == *appName {
				fmt.Println("Old Version Image: ", c[i].Image)
				fmt.Println("New Version Image: ", *imageName)
				c[i].Image = *imageName
				found = true
			}
		}

		if !found {
			fmt.Println("Container not exist in the deployment pods.")
			os.Exit(0)
		}

		_, err := clientset.AppsV1().Deployments("default").Update(context.TODO(), deployment, metav1.UpdateOptions{})
		if err != nil {
			panic(err)
		}
	}
}
