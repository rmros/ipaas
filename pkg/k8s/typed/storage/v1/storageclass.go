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
	"k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type StorageClassInterface interface {
	CreateStorageClass(storageClass *v1.StorageClass) (*v1.StorageClass, error)
	UpdateStorageClass(storageClass *v1.StorageClass) (*v1.StorageClass, error)
	DeleteStorageClass(name string) error
	GetStorageClass(name string) (*v1.StorageClass, error)
	ListStorageClass() ([]v1.StorageClass, error)
}

type storageclasses struct {
	*kubernetes.Clientset
}

//StorageClasses return a storages.
func StorageClasses(client *kubernetes.Clientset) StorageClassInterface {
	return &storageclasses{Clientset: client}
}

func (client *storageclasses) CreateStorageClass(storageClass *v1.StorageClass) (*v1.StorageClass, error) {
	return client.StorageV1().StorageClasses().Create(storageClass)
}

func (client *storageclasses) UpdateStorageClass(storageClass *v1.StorageClass) (*v1.StorageClass, error) {
	return client.StorageV1().StorageClasses().Update(storageClass)
}

func (client *storageclasses) DeleteStorageClass(name string) error {
	return client.StorageV1().StorageClasses().Delete(name, &metav1.DeleteOptions{})
}

func (client *storageclasses) GetStorageClass(name string) (*v1.StorageClass, error) {
	return client.StorageV1().StorageClasses().Get(name, metav1.GetOptions{})
}

func (client *storageclasses) ListStorageClass() ([]v1.StorageClass, error) {
	list, err := client.StorageV1().StorageClasses().List(metav1.ListOptions{})
	if err != nil {
		return []v1.StorageClass{}, err
	}
	return list.Items, err
}
