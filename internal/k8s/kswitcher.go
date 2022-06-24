package k8s

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func GetClusterNames() []Cluster {

	clusters := make([]Cluster, 0)

	f, err := os.Open(os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(fmt.Sprintf("failed to open the kube config file. %s", err))
	}

	scanner := bufio.NewScanner(f)

	var newClusterToken bool

	var cluster Cluster

	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		// Check to see if we are in - context. If so we are in
		// a cluster block.
		if strings.HasPrefix(line, "- context:") {
			newClusterToken = true
			cluster = Cluster{}
		}

		if newClusterToken && strings.HasPrefix(line, "cluster") {

			line = strings.ReplaceAll(line, "cluster:", "")
			if line != "" {
				line = strings.TrimSpace(line)
				cluster.Name = line
				clusters = append(clusters, cluster)
			}
		}
	}

	for _, c := range clusters {
		fmt.Println(c.Name)
	}

	return clusters
}
