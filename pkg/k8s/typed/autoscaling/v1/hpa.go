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
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//HPAInterface has methods to work with HPA resources.
type HPAInterface interface {
	CreateHPA(hpa *autoscalingv1.HorizontalPodAutoscaler) (*autoscalingv1.HorizontalPodAutoscaler, error)
	UpdateHPA(hpa *autoscalingv1.HorizontalPodAutoscaler) (*autoscalingv1.HorizontalPodAutoscaler, error)
	GetHPA(name, namespace string) (*autoscalingv1.HorizontalPodAutoscaler, error)
	DeleteHPA(name, namespace string) error
}

//hpas implements HPAInterface.
type hpas struct {
	*kubernetes.Clientset
}

//Hpas return a hpas.
func Hpas(client *kubernetes.Clientset) HPAInterface {
	return &hpas{Clientset: client}
}

func (client *hpas) CreateHPA(hpa *autoscalingv1.HorizontalPodAutoscaler) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	return client.AutoscalingV1().HorizontalPodAutoscalers(hpa.Namespace).Create(hpa)
}

func (client *hpas) UpdateHPA(hpa *autoscalingv1.HorizontalPodAutoscaler) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	return client.AutoscalingV1().HorizontalPodAutoscalers(hpa.Namespace).Update(hpa)
}

func (client *hpas) GetHPA(name, namespace string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	return client.AutoscalingV1().HorizontalPodAutoscalers(namespace).Get(name, metav1.GetOptions{})
}

func (client *hpas) DeleteHPA(name, namespace string) error {
	return client.AutoscalingV1().HorizontalPodAutoscalers(namespace).Delete(name, &metav1.DeleteOptions{})
}
