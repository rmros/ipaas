package v1beta1

import (
	v1beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//StatefulsetInterface  has methods to work with StatefulSet resources.
type StatefulsetInterface interface {
	CreateStatefulSet(statefulset *v1beta1.StatefulSet) (*v1beta1.StatefulSet, error)
	UpdateStatefulSet(statefulset *v1beta1.StatefulSet) (*v1beta1.StatefulSet, error)
	DeleteStatefulSet(name, namespace string) error
	GetStatefulSet(name, namespace string) (*v1beta1.StatefulSet, error)
	ListStatefulSet(namespace string) (*v1beta1.StatefulSetList, error)
	ExsitStatefulSet(name, namespace string) (bool, error)
}

//statefulsets implemenets StatefulsetInterface
type statefulsets struct {
	*kubernetes.Clientset
}

// Statefulsets return statefulsets
func Statefulsets(client *kubernetes.Clientset) StatefulsetInterface {
	return &statefulsets{Clientset: client}
}

func (client *statefulsets) CreateStatefulSet(statefulset *v1beta1.StatefulSet) (*v1beta1.StatefulSet, error) {
	return client.AppsV1beta1().StatefulSets(statefulset.Namespace).Create(statefulset)
}

func (client *statefulsets) UpdateStatefulSet(statefulset *v1beta1.StatefulSet) (*v1beta1.StatefulSet, error) {
	return client.AppsV1beta1().StatefulSets(statefulset.Namespace).Update(statefulset)
}

func (client *statefulsets) DeleteStatefulSet(name, namespace string) error {
DELETE_STATEFULSET:
	if err := client.AppsV1beta1().StatefulSets(namespace).Delete(name, &metav1.DeleteOptions{}); err != nil {
		if errors.IsConflict(err) {
			goto DELETE_STATEFULSET
		}
		return err
	}
	return nil
}

func (client *statefulsets) GetStatefulSet(name, namespace string) (*v1beta1.StatefulSet, error) {
	return client.AppsV1beta1().StatefulSets(namespace).Get(name, metav1.GetOptions{})
}

func (client *statefulsets) ListStatefulSet(namespace string) (*v1beta1.StatefulSetList, error) {
	return client.AppsV1beta1().StatefulSets(namespace).List(metav1.ListOptions{})
}

func (client *statefulsets) ExsitStatefulSet(name, namespace string) (bool, error) {
	statefuset, err := client.AppsV1beta1().StatefulSets(namespace).Get(name, metav1.GetOptions{})
	if err != nil {

		return false, err
	}
	if statefuset == nil {
		return false, nil
	}
	return true, nil
}
