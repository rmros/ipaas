package account

import (
	"fmt"

	"ipaas/pkg/k8s/client"
	iv1 "ipaas/pkg/k8s/typed/core/v1"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/golang/glog"
)

var (
	deletePropagationForeground = new(metav1.DeletionPropagation)
)

func init() {
	*deletePropagationForeground = metav1.DeletePropagationForeground
}

func toK8sNamespace(name string) *v1.Namespace {
	return &v1.Namespace{
		TypeMeta:   metav1.TypeMeta{Kind: "Namespace", APIVersion: "v1"},
		ObjectMeta: metav1.ObjectMeta{Name: name},
	}
}

func createNamespace(cluster string, ns *v1.Namespace) error {
	cls := client.GetClientset(cluster)
	if cls == nil {
		return fmt.Errorf("cluster [%v] k8s client doesn't exsit", cluster)
	}

	_, err := iv1.Namespaces(cls.Clientset).Create(ns)
	if err != nil {
		glog.Errorf("when add user,create k8s namespace [%v] in cluster [%v] err: %v", ns.Name, cluster, err)
		return err
	}

	return nil
}

func deleteNamespace(cluster, name, labels string) error {
	cls := client.GetClientset(cluster)
	if cls == nil {
		return fmt.Errorf("cluster [%v] k8s client doesn't exsit", cluster)
	}

	if name == "" {
		err := iv1.Namespaces(cls.Clientset).DeleteCollection(
			&metav1.DeleteOptions{
				PropagationPolicy: deletePropagationForeground,
			},
			metav1.ListOptions{
				LabelSelector: labels,
			},
		)
		if err != nil {
			glog.Errorf("delete k8s namespace in cluster [%v] by label [%v] err: %v", cluster, labels, err)
			return err
		}
	} else {
		if err := iv1.Namespaces(cls.Clientset).Delete(name, &metav1.DeleteOptions{}); err != nil {
			glog.Errorf("delete k8s namespace [%v] in cluster [%v] err: %v", name, cluster, err)
			return err
		}
	}

	return nil
}

func listNamespace(cluster string) ([]v1.Namespace, error) {
	cls := client.GetClientset(cluster)
	if cls == nil {
		return nil, fmt.Errorf("cluster [%v] k8s client doesn't exsit", cluster)
	}
	namespaces, err := iv1.Namespaces(cls.Clientset).List(metav1.ListOptions{})
	if err != nil {
		glog.Errorf("list namespace in cluster [%v] err: %v", cluster, err)
	}
	return namespaces, err
}
