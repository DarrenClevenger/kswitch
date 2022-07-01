package k8s

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"os/exec"
)

type Cluster struct {
	Name string
}

func GetCurrentCluster() string {

	f, err := os.Open(os.Getenv("KUBECONFIG"))
	if err != nil {
		return "Unknown"
	}

	scan := bufio.NewScanner(f)

	for scan.Scan() {

		line := scan.Text()
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "current-context:") {
			line = strings.ReplaceAll(line, "current-context:", "")
			line = strings.TrimSpace(line)
			return line
		}
	}

	return "Unknown"
}

func SetCurrentClusterContext(cluster_name string) error {

	cmd := exec.Command("kubectl", "config", "use-context", cluster_name)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func GetClusterNames() []Cluster {

	var clusters []Cluster

	f, err := os.Open(os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(fmt.Sprintf("failed to open the kube config file. %s", err))
	}

	scanner := bufio.NewScanner(f)

	var cluster Cluster

	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "cluster:") {

			line = strings.ReplaceAll(line, "cluster:", "")
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "\"", "")

			if len(line) > 0 {
				cluster = Cluster{
					Name: line,
				}
				clusters = append(clusters, cluster)
			}
		}
	}

	return clusters
}
