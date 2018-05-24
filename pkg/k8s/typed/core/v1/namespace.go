/*
Copyright [yyyy] [name of copyright owner]

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
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceInterface interface {
	Create(namespace *v1.Namespace) (*v1.Namespace, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Update(namespace *v1.Namespace) (*v1.Namespace, error)
	List(options metav1.ListOptions) ([]v1.Namespace, error)
}

type namespaces struct {
	*kubernetes.Clientset
}

func Namespaces(cl *kubernetes.Clientset) NamespaceInterface {
	return &namespaces{Clientset: cl}
}

func (c *namespaces) Create(namespace *v1.Namespace) (*v1.Namespace, error) {
	return c.CoreV1().Namespaces().Create(namespace)
}

func (c *namespaces) Delete(name string, options *metav1.DeleteOptions) error {
	return c.CoreV1().Namespaces().Delete(name, options)
}

func (c *namespaces) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.CoreV1().Namespaces().DeleteCollection(options, listOptions)
}

func (c *namespaces) Update(namespace *v1.Namespace) (*v1.Namespace, error) {
	return c.CoreV1().Namespaces().Update(namespace)
}

func (c *namespaces) List(options metav1.ListOptions) ([]v1.Namespace, error) {
	list, err := c.CoreV1().Namespaces().List(options)
	if err != nil || len(list.Items) == 0 {
		return []v1.Namespace{}, err
	}
	return list.Items, nil
}
