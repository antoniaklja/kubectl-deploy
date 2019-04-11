package kubectl

import (
	ctx "context"
	"os"

	"github.com/VirtusLab/go-extended/pkg/cli"
	"github.com/sirupsen/logrus"
)

// Apply applies a configuration to a resource by input
func Apply(input, context, namespace string, dryRun bool) error {
	if dryRun {
		logrus.Println(input)
		return nil
	}
	var args []string
	if len(context) > 0 {
		args = append(args, "--context", context)
	}
	if len(namespace) > 0 {
		args = append(args, "--namespace", namespace)
	}
	args = append(args, "apply", "-f", "-")

	env := os.Environ()
	stdout, stderr, err := cli.Sh(ctx.TODO(), logrus.StandardLogger(), env, &input, "kubectl", args...)
	if len(stdout) > 0 {
		logrus.Printf("Command stdout:\n---\n%s---", stdout)
	}
	if len(stderr) > 0 {
		logrus.Printf("Command stderr:\n---\n%s---", stderr)
	}
	if err != nil {
		return err
	}
	return nil
}
