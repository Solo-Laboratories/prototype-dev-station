package helpers

import (
    "log"
    // "os"

    // "helm.sh/helm/v3/pkg/action"
    // "helm.sh/helm/v3/pkg/chart/loader"
    // "helm.sh/helm/v3/pkg/cli"
    // "helm.sh/helm/v3/pkg/kube"
    // "helm.sh/helm/v3/pkg/release"
    // "helm.sh/helm/v3/pkg/repo"
)

func DeployGit() string {
    // Set up Helm action configuration
    // settings := cli.New()
    // actionConfig := new(action.Configuration)
    // if err := actionConfig.Init(kube.GetConfig(settings.KubeConfig, "", settings.Namespace, log.Printf), settings.Namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
    //     log.Fatalf("Failed to initialize Helm action configuration: %v", err)
    // }

    // // Add Gitea Helm repository
    // repoEntry := repo.Entry{
    //     Name: "gitea",
    //     URL:  "https://dl.gitea.io/charts",
    // }
    // chartRepo, err := repo.NewChartRepository(&repoEntry, getter.All(settings))
    // if err != nil {
    //     log.Fatalf("Failed to create chart repository: %v", err)
    // }
    // if _, err := chartRepo.DownloadIndexFile(); err != nil {
    //     log.Fatalf("Failed to download index file: %v", err)
    // }

    // // Load the Gitea chart
    // chartPath, err := chartRepo.DownloadChart("gitea", "10.6.0")
    // if err != nil {
    //     log.Fatalf("Failed to download chart: %v", err)
    // }
    // chart, err := loader.Load(chartPath)
    // if err != nil {
    //     log.Fatalf("Failed to load chart: %v", err)
    // }

    // // Set up Helm install action
    // client := action.NewInstall(actionConfig)
    // client.ReleaseName = "gitea"
    // client.Namespace = "dev-station"=
	// client.Atomic = true
    // client.ValuesFiles = []string{"/values-files/gitea.values.yaml"}

    // // Install the Helm chart
    // release, err := client.Run(chart, nil)
    // if err != nil {
    //     log.Fatalf("Failed to install Helm chart: %v", err)
    // }

    // Return the URL
    log.Println("Gitea URL: https://dev-station.cluster.sololab.one/gitea\n")
	return "https://dev-station.cluster.sololab.one/gitea"
}
