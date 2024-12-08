package helpers

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
	"log"
	"os"
	"sigs.k8s.io/yaml"
)

func DeployGit() string {
	// Create a Kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error creating in-cluster config: %v", err)
	}

	// Initialize Helm action configuration
	actionConfig := new(action.Configuration)
	var kubeConfig *genericclioptions.ConfigFlags
    namespace := "dev-station"
	// Create the ConfigFlags struct instance with initialized values from ServiceAccount
	kubeConfig = genericclioptions.NewConfigFlags(false)
	kubeConfig.APIServer = &config.Host
	kubeConfig.BearerToken = &config.BearerToken
	kubeConfig.CAFile = &config.CAFile
	kubeConfig.Namespace = &namespace
	if err := actionConfig.Init(kubeConfig, namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Fatalf("Error configuring the Helm Runner: %v", err)
	}

	// Add Helm repository
	repoEntry := repo.Entry{
		Name: "gitea",
		URL:  "https://dl.gitea.io/charts",
	}
	chartRepo, err := repo.NewChartRepository(&repoEntry, getter.All(cli.New()))
	if err != nil {
		log.Fatalf("Error creating chart repository: %v", err)
	}
	if _, err := chartRepo.DownloadIndexFile(); err != nil {
		log.Fatalf("Error downloading index file: %v", err)
	}

	// Install Helm chart
	install := action.NewInstall(actionConfig)
	install.ReleaseName = "gitea"
	install.Namespace = "dev-station"
	install.Atomic = true
	install.Wait = true
	chartPath, err := install.LocateChart("gitea/gitea", cli.New())

	// Load values file
	valuesFile := "values-files/gitea.values.yaml"
	values, err := os.ReadFile(valuesFile)
	if err != nil {
		log.Fatalf("Error reading values file: %v", err)
	}
	var valuesMap map[string]interface{}
	if err := yaml.Unmarshal(values, &valuesMap); err != nil {
		log.Fatalf("Error unmarshalling values file: %v", err)
	}

	if err != nil {
		log.Fatalf("Error locating chart: %v", err)
	}
	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Fatalf("Error loading chart: %v", err)
	}
	if _, err := install.Run(chart, valuesMap); err != nil {
		log.Fatalf("Error installing chart: %v", err)
	}

	log.Println("Gitea URL: https://dev-station.cluster.sololab.one/gitea\n")
	return "https://dev-station.cluster.sololab.one/gitea"
}
