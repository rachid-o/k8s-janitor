package cmd

import (
	"github.com/spf13/cobra"
)

var namespace string

var rootImg = &cobra.Command{
	Use:   "k8s-janitor",
	Short: "k8s Janitor cleans up messy Kubernetes resources",
	Long:  `A CLI tool to identify and clean up unused ConfigMaps, failing Pods, and more.`,
}

func Execute() error {
	return rootImg.Execute()
}

func init() {
	// Global flag: accessible to all subcommands
	rootImg.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "Kubernetes namespace to scan")
}
