package test

import (
	"os"
	"path"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

// 
// Deploy Gin App to AWS
func TestExamples(t *testing.T) {
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		awsRegion = "us-west-1"
	}
	cwd, _ := os.Getwd()
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Quick:       true,
		SkipRefresh: true,
		Dir:         path.Join(cwd, "..", "..", "aws-js-s3-folder"),
		Config: map[string]string{
			"aws:region": awsRegion,
		},
	})
}
