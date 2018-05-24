// Copyright © 2017 huang jia <449264675@qq.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"fmt"

	"ipaas/pkg/tools/uuid"

	k8s "ipaas/pkg/k8s/util"

	"k8s.io/api/apps/v1beta1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

const (
	//MinipaasAppName minipaas.io/appName label key
	MinipaasAppName = "minipaas.io/appName"
	//MinipaasServiceName minipaas.io/serviceName label key
	MinipaasServiceName = "minipaas.io/serviceName"
	//MinipaasName minipaas.io/name label key
	MinipaasName = "minipaas.io/name"
	//MinipaasStorageClassName minipaas.io.storageclass name, the gloabel StorageClass Name
	MinipaasStorageClassName = "minipaas.io.storageclass"
)

//TOK8SDeployment translate app to k8s deployment
func (svc *Service) TOK8SDeployment(namespace string) *v1beta1.Deployment {
	name := svc.Name
	nodeName := svc.NodeName
	labels := map[string]string{MinipaasAppName: svc.AppName, MinipaasServiceName: name, MinipaasName: name, "replicas": fmt.Sprintf("%v", svc.InstanceCount)}
	anotations := map[string]string{}
	replicas := svc.InstanceCount
	var volumes []v1.Volume
	var volumeMounts []v1.VolumeMount
	for _, v := range svc.Volumes {
		var items []v1.KeyToPath
		volume := v1.Volume{Name: v.ConfigMapName}
		if v.Type == 0 {
			volumeMount := v1.VolumeMount{Name: v.ConfigMapName, MountPath: v.MountPath}
			volumeMounts = append(volumeMounts, volumeMount)
			volume.VolumeSource = v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: v.ConfigMapName},
				},
			}
		} else {
			for _, k := range v.ConfigMapKey {
				items = append(items, v1.KeyToPath{Key: k, Path: k})
				volumeMount := v1.VolumeMount{Name: v.ConfigMapName, SubPath: k}
				mountPath := ""
				if v.MountPath[len(v.MountPath)-2:] == "/" {
					mountPath = v.MountPath + k
				} else {
					mountPath = v.MountPath + "/" + k
				}
				volumeMount.MountPath = mountPath
				volumeMounts = append(volumeMounts, volumeMount)
			}
			volume.VolumeSource = v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: v.ConfigMapName},
					Items: items,
				},
			}
		}
		volumes = append(volumes, v1.Volume{})
	}

	resourceRequirements := v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse(svc.CPU),    //TODO 根据前端传入的值做资源限制
			v1.ResourceMemory: resource.MustParse(svc.Memory), //TODO 根据前端传入的值做资源限制
		},
		Requests: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse(svc.CPU),
			v1.ResourceMemory: resource.MustParse(svc.Memory),
		},
	}

	var (
		initContainers []v1.Container
		containers     []v1.Container
		nodeSelector   map[string]string
	)
	ports := []v1.ContainerPort{}
	for _, p := range svc.Ports {
		ports = append(ports, v1.ContainerPort{ContainerPort: int32(p.TargetPort.IntVal)})
	}
	containers = []v1.Container{
		v1.Container{
			Name:            name,
			Image:           svc.Image,
			ImagePullPolicy: v1.PullIfNotPresent,
			Resources:       resourceRequirements,
			Ports:           ports,
			Env:             svc.Envs,
			VolumeMounts:    volumeMounts,
		},
	}
	return k8s.NewDeployment(name, namespace, nodeName, labels, anotations, replicas, volumes, initContainers, containers, nodeSelector)
}

//TOK8SService translate service to k8s service
func (svc *Service) TOK8SService(namespace string) *v1.Service {
	name := svc.Name
	labels := map[string]string{MinipaasAppName: svc.AppName, MinipaasServiceName: name, MinipaasName: name}
	ports := svc.Ports
	return k8s.NewService(svc.Name, namespace, labels, v1.ServiceTypeNodePort, ports)
}

// TOK8SHeadlessService translate service to k8s headlessService
func (svc *Service) TOK8SHeadlessService(namespace string) *v1.Service {
	name := svc.Name + uuid.UU()
	labels := map[string]string{MinipaasAppName: svc.AppName, MinipaasServiceName: name, MinipaasName: name}
	ports := svc.Ports
	return k8s.NewHeadlessService(svc.Name, namespace, labels, v1.ServiceTypeNodePort, ports)
}

//TOK8SConfigMap translate config to k8s service
func (cfg *Config) TOK8SConfigMap(namespace string) *v1.ConfigMap {
	return k8s.NewConfigMap(cfg.Name, namespace, map[string]string{}, cfg.Data)
}

//TOK8SHPA translate HPA to k8s HorizontalPodAutoscaler
func (hpa *HPA) TOK8SHPA(namespace string) *autoscalingv1.HorizontalPodAutoscaler {
	return k8s.NewHPA(hpa.Name, namespace, hpa.RefObjectName, hpa.MinReplicas, hpa.TargetCPUUtilizationPercentage, hpa.MaxReplicas, map[string]string{})
}

//TOPersistentVolumeClaim translate service to k8s PersistentVolumeClaim
func (svc *Service) TOPersistentVolumeClaim(namespace string) *v1.PersistentVolumeClaim {
	labels := map[string]string{}
	resources := v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceMemory: resource.MustParse(svc.Storage.Size), //TODO 根据前端传入的值做资源限制
		},
		Requests: v1.ResourceList{
			v1.ResourceMemory: resource.MustParse(svc.Storage.Size),
		},
	}
	return k8s.NewPersistenVolumeClaim(svc.Name, namespace, MinipaasStorageClassName, svc.Storage.AccessModes, labels, resources)
}

//TOStorageClass translate CephRBD to k8s StorageClass
func (rbd *CephRBD) TOStorageClass() *storagev1.StorageClass {
	parameters := map[string]string{
		"monitors":             rbd.Monitors,
		"adminId":              rbd.AdminID,
		"adminSecretName":      rbd.AdminSecretName,
		"adminSecretNamespace": rbd.AdminSecretNamespace,
		"pool":                 rbd.Pool,
		"userId":               rbd.UserID,
		"userSecretName":       rbd.UserSecretName,
	}
	if rbd.FsType != "" {
		parameters["fsType"] = rbd.FsType
	}
	if rbd.ImageFormat != "" {
		parameters["imageFormat"] = rbd.ImageFormat
	}
	if rbd.ImageFeatures != "" {
		parameters["imageFeatures"] = rbd.ImageFeatures
	}
	return k8s.NewStorageClass(MinipaasStorageClassName, rbd.Provisioner, parameters)
}

// TOK8SStatefulset translate service to k8s StatefulSet
func (svc *Service) TOK8SStatefulset(namespace, headlessServiceName string, pvc v1.PersistentVolumeClaim) *v1beta1.StatefulSet {
	name := svc.Name
	nodeName := svc.NodeName
	labels := map[string]string{MinipaasAppName: svc.AppName, MinipaasServiceName: name, MinipaasName: name, "replicas": fmt.Sprintf("%v", svc.InstanceCount)}
	anotations := map[string]string{}
	replicas := svc.InstanceCount
	var volumes []v1.Volume
	var volumeMounts []v1.VolumeMount
	for _, v := range svc.Volumes {
		var items []v1.KeyToPath
		volume := v1.Volume{Name: v.ConfigMapName}
		if v.Type == 0 {
			volumeMount := v1.VolumeMount{Name: v.ConfigMapName, MountPath: v.MountPath}
			volumeMounts = append(volumeMounts, volumeMount)
			volume.VolumeSource = v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: v.ConfigMapName},
				},
			}
		} else {
			for _, k := range v.ConfigMapKey {
				items = append(items, v1.KeyToPath{Key: k, Path: k})
				volumeMount := v1.VolumeMount{Name: v.ConfigMapName, SubPath: k}
				mountPath := ""
				if v.MountPath[len(v.MountPath)-2:] == "/" {
					mountPath = v.MountPath + k
				} else {
					mountPath = v.MountPath + "/" + k
				}
				volumeMount.MountPath = mountPath
				volumeMounts = append(volumeMounts, volumeMount)
			}
			volume.VolumeSource = v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: v.ConfigMapName},
					Items: items,
				},
			}
		}
		volumes = append(volumes, v1.Volume{})
	}

	resourceRequirements := v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse(svc.CPU),    //TODO 根据前端传入的值做资源限制
			v1.ResourceMemory: resource.MustParse(svc.Memory), //TODO 根据前端传入的值做资源限制
		},
		Requests: v1.ResourceList{
			v1.ResourceCPU:    resource.MustParse(svc.CPU),
			v1.ResourceMemory: resource.MustParse(svc.Memory),
		},
	}

	var (
		initContainers []v1.Container
		containers     []v1.Container
		nodeSelector   map[string]string
	)
	ports := []v1.ContainerPort{}
	for _, p := range svc.Ports {
		ports = append(ports, v1.ContainerPort{ContainerPort: int32(p.TargetPort.IntVal)})
	}
	containers = []v1.Container{
		v1.Container{
			Name:            name,
			Image:           svc.Image,
			ImagePullPolicy: v1.PullIfNotPresent,
			Resources:       resourceRequirements,
			Ports:           ports,
			Env:             svc.Envs,
			VolumeMounts:    volumeMounts,
		},
	}
	return k8s.NewStatefulSet(name, namespace, nodeName, headlessServiceName, labels, anotations, nodeSelector, replicas, volumes, initContainers, containers, pvc)
}

// TOK8SPersistentVolumeClaim translate storage to k8s PersistentVolumeClaim
func (storage *Storage) TOK8SPersistentVolumeClaim(namespace string) *v1.PersistentVolumeClaim {
	resourceRequirements := v1.ResourceRequirements{
		Limits: v1.ResourceList{
			v1.ResourceStorage: resource.MustParse(storage.Size), //TODO 根据前端传入的值做资源限制
		},
		Requests: v1.ResourceList{
			v1.ResourceStorage: resource.MustParse(storage.Size), //TODO 根据前端传入的值做资源限制
		},
	}
	return k8s.NewPersistenVolumeClaim(storage.Name, namespace, storage.AccessModes, MinipaasStorageClassName, nil, resourceRequirements)
}
