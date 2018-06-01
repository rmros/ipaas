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

type ConfigMapInterface interface {
	CreateConfigMap(cm *v1.ConfigMap) (*v1.ConfigMap, error)
	UpdateConfigMap(cm *v1.ConfigMap) (*v1.ConfigMap, error)
	GetConfigMap(name, namespace string) (*v1.ConfigMap, error)
	ListConfigMap(namespace string) (*v1.ConfigMapList, error)
	DeleteConfigMap(name, namespace string) error
	ExsitConfigMap(name, namespace string) (bool, error)
}

type configMaps struct {
	*kubernetes.Clientset
}

// ConfigMaps return a configMap.
func ConfigMaps(client *kubernetes.Clientset) ConfigMapInterface {
	return &configMaps{Clientset: client}
}

func (client *configMaps) CreateConfigMap(cm *v1.ConfigMap) (*v1.ConfigMap, error) {
	return client.CoreV1().ConfigMaps(cm.Namespace).Create(cm)
}

func (client *configMaps) UpdateConfigMap(cm *v1.ConfigMap) (*v1.ConfigMap, error) {
	return client.CoreV1().ConfigMaps(cm.Namespace).Update(cm)
}

func (client *configMaps) GetConfigMap(name, namespace string) (*v1.ConfigMap, error) {
	return client.CoreV1().ConfigMaps(namespace).Get(name, metav1.GetOptions{})
}

func (client *configMaps) ListConfigMap(namespace string) (*v1.ConfigMapList, error) {
	return client.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
}

func (client *configMaps) DeleteConfigMap(name, namespace string) error {
DELETE_CONFIGMAP:
	if err := client.CoreV1().ConfigMaps(namespace).Delete(name, &metav1.DeleteOptions{}); err != nil {
		if errors.IsConflict(err) {
			goto DELETE_CONFIGMAP
		}
		return err
	}
	return nil
}

func (client *configMaps) ExsitConfigMap(name, namespace string) (bool, error) {
	cm, err := client.CoreV1().ConfigMaps(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	if cm == nil {
		return false, nil
	}
	return true, nil
}
