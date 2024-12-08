package helpers

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
)

func DeployManifestFile(file string) {
	// Use the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to initialize Configuration: %v", err)
	}

	// Create the dynamicClient
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to initialize ClientSet: %v", err)
	}

	// Read the manifest file
	manifestFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to Read in Manifest File: %v", err)
	}

	// Decode the manifest file
	// Convert []byte to io.Reader
	manifestReader := bytes.NewReader(manifestFile)
	decoder := yaml.NewYAMLOrJSONDecoder(manifestReader, 100)
	var obj unstructured.Unstructured
	if err := decoder.Decode(&obj); err != nil {
		log.Fatalf("Failed to Decode Manifest yaml file to data structure: %v", err)
	}

	// Apply the manifest
	gvr := obj.GroupVersionKind().GroupVersion().WithResource(obj.GetKind())
	namespace := obj.GetNamespace()
	//name := obj.GetName()

	log.Println(gvr)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err := dynamicClient.Resource(gvr).Namespace(namespace).Create(context.TODO(), &obj, metav1.CreateOptions{})
		if errors.IsAlreadyExists(err) {
			_, err = dynamicClient.Resource(gvr).Namespace(namespace).Update(context.TODO(), &obj, metav1.UpdateOptions{})
		}
		return err
	})
	if retryErr != nil {
		log.Fatalf("Failed to apply manifest file: %v", retryErr)
	}

	log.Println("Manifest applied successfully")
}
