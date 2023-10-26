package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestSpec contains the test spec.
type TestSpec struct {
	// Timeout for the test. Overrides the global timeout set in the Configuration.
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`

	// Skip determines whether the test should skipped.
	// +optional
	Skip *bool `json:"skip,omitempty"`

	// Concurrent determines whether the test should run concurrently with other tests.
	// +optional
	Concurrent *bool `json:"concurrent,omitempty"`

	// SkipDelete determines whether the resources created by the test should be deleted after the test is executed.
	// +optional
	SkipDelete *bool `json:"skipDelete,omitempty"`

	// Namespace determines whether the test should run in a random ephemeral namespace or not.
	// +optional
	Namespace string `json:"namespace,omitempty"`

	// Steps defining the test.
	Steps []TestSpecStep `json:"steps"`
}
