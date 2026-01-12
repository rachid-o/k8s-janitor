package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for unused resources",
	Run:   func(cmd *cobra.Command, args []string) { runScan(cmd, args) },
}

func init() {
	rootImg.AddCommand(scanCmd)
}

func runScan(cmd *cobra.Command, args []string) {
	//return func(cmd *cobra.Command, args []string) {
	client, context, err := getKubeClient()
	if err != nil {
		fmt.Printf("Error building k8s client: %v\n", err)
		return
	}

	fmt.Printf("🌐 Current Context: %s\n", context)
	fmt.Printf("🧹 Scanning namespace: %s...\n", namespace)

	// Example: List Pods in the namespace
	pods, err := client.CoreV1().Pods(namespace).List(cmd.Context(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing pods: %v\n", err)
		return
	}

	fmt.Printf("Found %d pods in namespace %s\n", len(pods.Items), namespace)
	//for _, pod := range pods.Items {
	//	fmt.Printf("- %s [%s]\n", pod.Name, pod.Status.Phase)
	//}
	for _, pod := range pods.Items {
		for _, containerStatus := range pod.Status.ContainerStatuses {
			// Check if the container is waiting and has an "error" reason
			if containerStatus.State.Waiting != nil {
				reason := containerStatus.State.Waiting.Reason
				if reason == "CrashLoopBackOff" || reason == "ImagePullBackOff" || reason == "ErrImagePull" {
					fmt.Printf("⚠️  Found Issue: Pod %s is in %s\n", pod.Name, reason)
				}
			}
		}
	}
	//}
}
