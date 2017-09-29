package terraform

import (
	"fmt"
	"os"
)

// This file holds feature flags for the next release

var featureOutputErrors bool

func init() {
	featureOutputErrors = os.Getenv("TF_OUTPUT_ERRORS") != ""
	if featureOutputErrors {
		fmt.Println("WARNING: output errors feature enabled")
	}
}
