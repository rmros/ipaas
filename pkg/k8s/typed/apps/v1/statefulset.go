/*
Copyright [huangjia] [name of copyright owner]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//StatefulsetInterface  has methods to work with StatefulSet resources.
type StatefulsetInterface interface {
	CreateStatefulSet(statefulset *appsv1.StatefulSet) (*appsv1.StatefulSet, error)
	UpdateStatefulSet(statefulset *appsv1.StatefulSet) (*appsv1.StatefulSet, error)
	DeleteStatefulSet(name, namespace string) error
	GetStatefulSet(name, namespace string) (*appsv1.StatefulSet, error)
	ListStatefulSet(namespace string) (*appsv1.StatefulSetList, error)
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

func (client *statefulsets) CreateStatefulSet(statefulset *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	return client.AppsV1().StatefulSets(statefulset.Namespace).Create(statefulset)
}

func (client *statefulsets) UpdateStatefulSet(statefulset *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	return client.AppsV1().StatefulSets(statefulset.Namespace).Update(statefulset)
}

func (client *statefulsets) DeleteStatefulSet(name, namespace string) error {
DELETE_STATEFULSET:
	if err := client.AppsV1().StatefulSets(namespace).Delete(name, &metav1.DeleteOptions{}); err != nil {
		if errors.IsConflict(err) {
			goto DELETE_STATEFULSET
		}
		return err
	}
	return nil
}

func (client *statefulsets) GetStatefulSet(name, namespace string) (*appsv1.StatefulSet, error) {
	return client.AppsV1().StatefulSets(namespace).Get(name, metav1.GetOptions{})
}

func (client *statefulsets) ListStatefulSet(namespace string) (*appsv1.StatefulSetList, error) {
	return client.AppsV1().StatefulSets(namespace).List(metav1.ListOptions{})
}

func (client *statefulsets) ExsitStatefulSet(name, namespace string) (bool, error) {
	statefuset, err := client.AppsV1().StatefulSets(namespace).Get(name, metav1.GetOptions{})
	if err != nil {

		return false, err
	}
	if statefuset == nil {
		return false, nil
	}
	return true, nil
}
