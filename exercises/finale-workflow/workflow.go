package finale

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func CertificateGeneratorWorkflow(ctx workflow.Context, name string) (string, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started", "recipientName", name)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var outFilePath string
	err := workflow.ExecuteActivity(ctx, "io.temporal.training.GenerateCertificateActivity", name).Get(ctx, &outFilePath)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("CertificateGeneratorWorkflow created file", outFilePath)

	return outFilePath, nil
}
