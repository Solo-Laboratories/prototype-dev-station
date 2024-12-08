package helpers

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
    "k8s.io/client-go/restmapper"
	"k8s.io/client-go/discovery/cached/memory"
)

// KubeClient is a dynamic type of client for kubernetes
type KubeClient struct {
    client          *dynamic.DynamicClient
    discoveryMapper *restmapper.DeferredDiscoveryRESTMapper
}
// NewKubeClient creates an instance of KubeClient
func NewKubeClient() *KubeClient {
    config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to initialize Configuration: %v", err)
	}
    // create the dynamic client
    client, err := dynamic.NewForConfig(config)
    if err != nil {
        log.Fatalf("failed to create dynamic client: %w", err)
    }
    // create a discovery client to map dynamic API resources
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        log.Fatalf("failed to create discovery client: %w", err)
    }
    discoveryClient := memory.NewMemCacheClient(clientset.Discovery())
    discoveryMapper := restmapper.NewDeferredDiscoveryRESTMapper(discoveryClient)
    return &KubeClient{client: client, discoveryMapper: discoveryMapper}
}

func (k *KubeClient) Apply(file string) {
	// Read the manifest file
	manifestFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to Read in Manifest File: %v", err)
	}

	// Decode the manifest file
	// Convert []byte to io.Reader
	manifestReader := bytes.NewReader(manifestFile)
	decoder := yaml.NewYAMLOrJSONDecoder(manifestReader, 100)
	obj := &unstructured.Unstructured{Object: map[string]interface{}{}}
	if err := decoder.Decode(&obj); err != nil {
		log.Fatalf("Failed to Decode Manifest yaml file to data structure: %v", err)
	}

	// get GroupVersionResource to invoke the dynamic client
	gvk := obj.GroupVersionKind()
	restMapping, err := k.discoveryMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		log.Fatalf("Failed REST mapping discovery step: %v", err)
	}
	gvr := restMapping.Resource
	// apply the YAML doc
	namespace := obj.GetNamespace()
	if len(namespace) == 0 {
		namespace = "dev-station"
	}
	applyOpts := metav1.ApplyOptions{FieldManager: "kube-apply"}
	_, err = k.client.Resource(gvr).Namespace(namespace).Apply(context.TODO(), obj.GetName(), obj, applyOpts)
	if err != nil {
		log.Fatalf("apply error: %v", err)
	}
	log.Printf("applied YAML for %s %q", obj.GetKind(), obj.GetName())

	// // get GroupVersionResource to invoke the dynamic client
	// gvk := obj.GroupVersionKind()
	// restMapping, err := k.discoveryMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	// if err != nil {
	// 	return err
	// }
	// gvr := restMapping.Resource


	// // gvr := obj.GroupVersionKind().GroupVersion().WithResource(obj.GetKind())
	// // namespace := obj.GetNamespace()
	// //name := obj.GetName()

	// log.Println(gvr)
	// retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
	// 	_, err := dynamicClient.Resource(gvr).Namespace(namespace).Create(context.TODO(), &obj, metav1.CreateOptions{})
	// 	if errors.IsAlreadyExists(err) {
	// 		_, err = dynamicClient.Resource(gvr).Namespace(namespace).Update(context.TODO(), &obj, metav1.UpdateOptions{})
	// 	}
	// 	return err
	// })
	// if retryErr != nil {
	// 	log.Fatalf("Failed to apply manifest file: %v", retryErr)
	// }

	log.Println("Manifest applied successfully")
}
