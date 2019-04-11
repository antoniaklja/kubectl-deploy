package plugin

import (
	"kubectl-deploy/internal"
	"kubectl-deploy/kubectl"
	"kubectl-deploy/render"
)

// Deploy render and apply Kubernetes manifests
func Deploy(path, config, context, namespace string, dryRun bool) error {
	files, err := internal.GetFiles(path)
	if err != nil {
		return err
	}
	for _, file := range files {

		rendered, err := render.Render(file, config)
		if err != nil {
			return err
		}

		err = kubectl.Apply(rendered, context, namespace, dryRun)
		if err != nil {
			return err
		}
	}
	return nil
}
