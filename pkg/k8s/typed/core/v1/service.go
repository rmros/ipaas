/*
Copyright 2018 huangjia.

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
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//ServiceInterface has methods to work with Service resources.
type ServiceInterface interface {
	CreateService(service *v1.Service) (*v1.Service, error)
	UpdateService(svc *v1.Service) (*v1.Service, error)
	DeleteService(name, namespace string) error
	GetService(name, namespace string) (*v1.Service, error)
	ListService(labels, namespace string) ([]v1.Service, error)
	ExsitService(name, namespace string) (bool, error)
	DeleteServiceByLabels(labels, namespace string) error
}

//hpas implements ServiceInterface.
type services struct {
	*kubernetes.Clientset
}

//Services return a service.
func Services(client *kubernetes.Clientset) ServiceInterface {
	return &services{Clientset: client}
}

func (client *services) CreateService(service *v1.Service) (*v1.Service, error) {
	return client.CoreV1().Services(service.Namespace).Create(service)
}

func (client *services) UpdateService(svc *v1.Service) (*v1.Service, error) {
	return client.CoreV1().Services(svc.Namespace).Update(svc)
}

func (client *services) DeleteService(name, namespace string) error {
DELETE_SERVICE:
	if err := client.CoreV1().Services(namespace).Delete(name, &metav1.DeleteOptions{}); err != nil {
		if errors.IsConflict(err) {
			goto DELETE_SERVICE
		}
		return err
	}
	return nil
}

func (client *services) DeleteServiceByLabels(labels, namespace string) error {
	list, err := client.CoreV1().Services(namespace).List(metav1.ListOptions{LabelSelector: labels})
	if err != nil {
		return err
	}
	for _, item := range list.Items {
	DELETE_SERVICE:
		if err = client.DeleteService(item.Name, namespace); err != nil {
			if errors.IsConflict(err) {
				goto DELETE_SERVICE
			}
			return err
		}
	}
	return nil
}

func (client *services) GetService(name, namespace string) (*v1.Service, error) {
	return client.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
}

func (client *services) ListService(labels, namespace string) ([]v1.Service, error) {
	listOption := metav1.ListOptions{}
	if labels != "" {
		listOption.LabelSelector = labels
	}
	serviceList, err := client.CoreV1().Services(namespace).List(listOption)
	if err != nil || len(serviceList.Items) == 0 {
		return []v1.Service{}, err
	}
	return serviceList.Items, nil
}

func (client *services) ExsitService(name, namespace string) (bool, error) {
	svc, err := client.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	if svc == nil {
		return false, nil
	}
	return true, nil
}
