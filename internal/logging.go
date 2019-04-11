package internal

import (
	"kubectl-deploy/constants"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// GetDebug returns the global debug flag value
func GetDebug(cmd *cobra.Command) (debug bool, err error) {
	if cmd.HasFlags() {
		debug, err = cmd.Flags().GetBool(constants.DebugParameterName)
	} else {
		debug, err = cmd.PersistentFlags().GetBool(constants.DebugParameterName)
	}
	if err != nil {
		err = errors.Wrap(err, "unexpected error")
	}
	return
}

// GetTrace returns the global tracing flag value
func GetTrace(cmd *cobra.Command) (trace bool, err error) {
	if cmd.HasFlags() {
		trace, err = cmd.Flags().GetBool(constants.TraceParameterName)
	} else {
		trace, err = cmd.PersistentFlags().GetBool(constants.TraceParameterName)
	}
	if err != nil {
		err = errors.Wrap(err, "unexpected error")
	}
	return
}

func SetupLogging(debug bool, trace bool) error {
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	logrus.SetLevel(logrus.InfoLevel)

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		logrus.Debug("Debug logging enabled")
	}

	if trace {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		logrus.Debug("Trace debug logging enabled")
	}
	return nil
}
