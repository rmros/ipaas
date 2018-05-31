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
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NodeInterface interface {
	Create(node *v1.Node) (*v1.Node, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Update(node *v1.Node) (*v1.Node, error)
	List(options metav1.ListOptions) ([]v1.Node, error)
	Get(name string) (*v1.Node, error)
}

type nodes struct {
	*kubernetes.Clientset
}

func Nodes(cl *kubernetes.Clientset) NodeInterface {
	return &nodes{Clientset: cl}
}

func (n *nodes) Create(node *v1.Node) (*v1.Node, error) {
	return nil, nil
}

func (n *nodes) Delete(name string, options *metav1.DeleteOptions) error {
	return nil
}
func (n *nodes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return nil
}
func (n *nodes) Update(node *v1.Node) (*v1.Node, error) {
	return n.CoreV1().Nodes().Update(node)
}
func (n *nodes) List(options metav1.ListOptions) ([]v1.Node, error) {
	nodes, err := n.CoreV1().Nodes().List(options)
	if err != nil || len(nodes.Items) == 0 {
		return nil, err
	}
	return nodes.Items, nil

}
func (n *nodes) Get(name string) (*v1.Node, error) {
	return n.CoreV1().Nodes().Get(name, metav1.GetOptions{})
}
