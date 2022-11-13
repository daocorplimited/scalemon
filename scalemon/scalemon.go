package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func main() {
	http.HandleFunc("/", httpHandler)

	err := http.ListenAndServe(":8085", nil)

	if err != nil {
		fmt.Printf("error happened: %s\n", err)
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	// Get the config
	config, err := clientcmd.BuildConfigFromFlags("", "~/.kube/config")
	if err != nil {
		panic(err.Error())
	}

	// Get the metrics client
	metricsClientset, err := metricsv.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace := "" // leave empty to get data from all namespaces
	podMetricsList, err := metricsClientset.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, v := range podMetricsList.Items {
		fmt.Printf("%s\n", v.GetName())
		fmt.Printf("%s\n", v.GetNamespace())
		fmt.Printf("%vm\n", v.Containers[0].Usage.Cpu().MilliValue())
		fmt.Printf("%vMi\n", v.Containers[0].Usage.Memory().Value()/(1024*1024))
	}
}
