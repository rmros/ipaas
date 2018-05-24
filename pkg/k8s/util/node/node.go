package node

import (
	"k8s.io/api/core/v1"
)

// IsNodeReady returns true if a node is ready; false otherwise.
func IsNodeReady(node *v1.Node) bool {
	for _, c := range node.Status.Conditions {
		if c.Type == v1.NodeReady {
			return c.Status == v1.ConditionTrue
		}
	}
	return false
}

// IsNodeNoSchedule returns true if a node is schedulable; false otherwise.
func IsNodeSchedule(node *v1.Node) bool {
	for _, t := range node.Spec.Taints {
		if t.Effect == v1.TaintEffectNoSchedule || node.Spec.Unschedulable {
			return false
		}
	}
	return true
}

// IsDiskPressure returns true if a node is DiskPressure; false otherwise.
func IsDiskPressure(node *v1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == v1.NodeDiskPressure && condition.Status == v1.ConditionTrue {
			return true
		}
	}
	return false
}

// IsMemoryPressure returns true if a node is MemoryPressure; false otherwise.
func IsMemoryPressure(node *v1.Node) bool {
	for _, condition := range node.Status.Conditions {
		if condition.Type == v1.NodeMemoryPressure && condition.Status == v1.ConditionTrue {
			return true
		}
	}
	return false
}

func GetHostName(node *v1.Node) string {
	for _, address := range node.Status.Addresses {
		if address.Type == v1.NodeHostName {
			return address.Address
		}
	}
	return ""
}

func GetInternalIP(node *v1.Node) string {
	for _, addresse := range node.Status.Addresses {
		if addresse.Type == v1.NodeInternalIP {
			return addresse.Address
		}
	}
	return ""
}
