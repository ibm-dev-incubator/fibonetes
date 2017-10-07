package main

import (
	"flag"
	"os"
	"path/filepath"

    apiv1 "k8s.io/api/core/v1"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	batchClient := clientset.BatchV1()
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "fibonetes",
		},
		Spec: batchv1.JobSpec{
            Template: apiv1.PodTemplateSpec{
                Spec: apiv1.PodSpec{
                    RestartPolicy: apiv1.RestartPolicyNever,
                    Containers: []apiv1.Container{
                        {
                            Name:  "fibonetes",
                            Image: "python",
                        },
                    },
                },
            },
        },
	}
	_, err = batchClient.Jobs("default").Create(job)
	if err != nil {
		panic(err)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
