package k8s

import "testing"

func TestGetKubeConfig(t *testing.T) {

	GetClusterNames()
}

func TestGetCurrentCluster(t *testing.T) {

	current := GetCurrentCluster()
	t.Logf("Current: %s", current)
	if current == "" {
		t.Error("expected a cluster name but received an empty string.")
	}
}
