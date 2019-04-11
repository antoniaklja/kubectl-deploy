package main

import (
	"kubectl-deploy/constants"
	"kubectl-deploy/internal"
	"kubectl-deploy/plugin"
	"os"

	"github.com/spf13/cobra"
)

const (
	fileParameterName      = "file"
	configParameterName    = "config"
	contextParameterName   = "context"
	namespaceParameterName = "namespace"
	dryRunParameterName    = "dry-run"
)

func main() {
	var file string
	var config string
	var context string
	var namespace string
	var dryRun bool

	var deployCmd = &cobra.Command{
		Use:   "kubectl deploy",
		Short: "render and apply kubernetes manifests",
		Run: func(cmd *cobra.Command, args []string) {
			err := plugin.Deploy(file, config, context, namespace, dryRun)
			if err != nil {
				os.Exit(1)
			}
		},
	}

	deployCmd.PersistentFlags().StringVarP(&file, fileParameterName, "f", "", "If true, only print the object that would be sent, without sending it")
	if err := deployCmd.MarkPersistentFlagRequired(fileParameterName); err != nil {
		os.Exit(1)
	}
	deployCmd.PersistentFlags().StringVarP(&config, configParameterName, "", "", "config.yaml (required)")
	if err := deployCmd.MarkPersistentFlagRequired(configParameterName); err != nil {
		os.Exit(1)
	}
	deployCmd.PersistentFlags().StringVarP(&context, contextParameterName, "c", "", "k8s context (optional)")
	deployCmd.PersistentFlags().StringVarP(&namespace, namespaceParameterName, "n", "", "k8s namespace (optional)")
	deployCmd.PersistentFlags().BoolVarP(&dryRun, dryRunParameterName, "", false, "If true, only print the object that would be sent, without sending it")

	deployCmd.PersistentFlags().BoolP(constants.TraceParameterName, "t", false, "enable trace logging level output (optional)")
	deployCmd.PersistentFlags().BoolP(constants.DebugParameterName, "d", false, "enable debug logging level output (optional)")

	debug, err := internal.GetDebug(deployCmd)
	if err != nil {
		os.Exit(1)
	}
	trace, err := internal.GetTrace(deployCmd)
	if err != nil {
		os.Exit(1)
	}

	internal.SetupLogging(debug, trace)
	deployCmd.Execute()
}
