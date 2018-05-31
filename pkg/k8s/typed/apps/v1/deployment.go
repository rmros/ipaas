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
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//DeploymentInterface has methods to work with Deployment resources.
type DeploymentInterface interface {
	CreateDeployment(deployment *appsv1.Deployment) (*appsv1.Deployment, error)
	UpdateDeployment(deployment *appsv1.Deployment) (*appsv1.Deployment, error)
	DeleteDeployment(name, namespace string) error
	DeleteDeploymentByLabels(labels, namespace string) error
	GetDeployment(name, namespace string) (*appsv1.Deployment, error)
	ListDeployment(labels, namespace string) ([]appsv1.Deployment, error)
	ListPodByDeploymentName(name, namespace string) ([]v1.Pod, error)
	ExsitDeployment(name, namespace string) (bool, error)
}

//deployments implements DeploymentInterface.
type deployments struct {
	*kubernetes.Clientset
}

//Deployments return deployments.
func Deployments(client *kubernetes.Clientset) DeploymentInterface {
	return &deployments{Clientset: client}
}

func (client *deployments) CreateDeployment(deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	return client.AppsV1().Deployments(deployment.Namespace).Create(deployment)
}

func (client *deployments) UpdateDeployment(deploy *appsv1.Deployment) (*appsv1.Deployment, error) {
	return client.AppsV1().Deployments(deploy.Namespace).Update(deploy)
}

func (client *deployments) DeleteDeployment(name, namespace string) error {
DELETE_DEPLOYMENT:
	deletePropagationForeground := new(metav1.DeletionPropagation)
	*deletePropagationForeground = metav1.DeletePropagationForeground
	if err := client.ExtensionsV1beta1().Deployments(namespace).Delete(name, &metav1.DeleteOptions{PropagationPolicy: deletePropagationForeground}); err != nil {
		if errors.IsConflict(err) {
			goto DELETE_DEPLOYMENT
		}
		return err
	}
	return nil
}

func (client *deployments) DeleteDeploymentByLabels(labels, namespace string) error {
	list, err := client.ExtensionsV1beta1().Deployments(namespace).List(metav1.ListOptions{LabelSelector: labels})
	if err != nil {
		return err
	}
	for _, item := range list.Items {
	DELETE_DEPLOYMENT:
		if err = client.DeleteDeployment(item.Name, namespace); err != nil {
			if errors.IsConflict(err) {
				goto DELETE_DEPLOYMENT
			}
			return err
		}
	}
	return nil
}

func (client *deployments) GetDeployment(name, namespace string) (*appsv1.Deployment, error) {
	deploy, err := client.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return deploy, nil
}

func (client *deployments) ListDeployment(labels, namespace string) ([]appsv1.Deployment, error) {
	listOption := metav1.ListOptions{}
	if labels != "" {
		listOption.LabelSelector = labels
	}
	deploymentList, err := client.AppsV1().Deployments(namespace).List(listOption)
	if err != nil || len(deploymentList.Items) == 0 {
		return []appsv1.Deployment{}, err
	}
	return deploymentList.Items, nil
}

func (client *deployments) ListPodByDeploymentName(name, namespace string) ([]v1.Pod, error) {
	list, err := client.CoreV1().Pods(namespace).List(metav1.ListOptions{LabelSelector: "minipaas.io/name=" + name})
	if err != nil || len(list.Items) == 0 {
		return []v1.Pod{}, err
	}
	return list.Items, nil
}

func (client *deployments) ExsitDeployment(name, namespace string) (bool, error) {
	deploy, err := client.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return false, err
	}
	if deploy == nil {
		return false, nil
	}
	return true, nil
}
