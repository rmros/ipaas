package util

import (
	"ipaas/pkg/tools/parse"

	"k8s.io/api/apps/v1beta1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

//kubernetes resource
const (
	ResourceKindConfigMap               = "configmap"
	ResourceKindDaemonSet               = "daemonset"
	ResourceKindDeployment              = "deployment"
	ResourceKindEvent                   = "event"
	ResourceKindHorizontalPodAutoscaler = "horizontalpodautoscaler"
	ResourceKindIngress                 = "ingress"
	ResourceKindJob                     = "job"
	ResourceKindLimitRange              = "limitrange"
	ResourceKindNamespace               = "namespace"
	ResourceKindNode                    = "node"
	ResourceKindPersistentVolumeClaim   = "persistentvolumeclaim"
	ResourceKindPersistentVolume        = "persistentvolume"
	ResourceKindPod                     = "pod"
	ResourceKindReplicaSet              = "replicaset"
	ResourceKindReplicationController   = "replicationcontroller"
	ResourceKindResourceQuota           = "resourcequota"
	ResourceKindSecret                  = "secret"
	ResourceKindService                 = "service"
	ResourceKindStatefulSet             = "statefulset"
	ResourceKindThirdPartyResource      = "thirdpartyresource"
	ResourceKindStorageClass            = "storageclass"
	ResourceKindRbacRole                = "role"
	ResourceKindRbacClusterRole         = "clusterrole"
	ResourceKindRbacRoleBinding         = "rolebinding"
	ResourceKindRbacClusterRoleBinding  = "clusterrolebinding"
)

// type AppStatus int32
// type UpdateStatus int32

const (
	AppBuilding  = 0
	AppSuccessed = 1
	AppFailed    = 2
	AppRunning   = 3
	AppStop      = 4
	AppDelete    = 5
	AppUnknow    = 6

	StartFailed    = 10
	StartSuccessed = 11

	StopFailed    = 20
	StopSuccessed = 21

	ScaleFailed    = 30
	ScaleSuccessed = 31

	UpdateConfigFailed    = 40
	UpdateConfigSuccessed = 41

	RedeploymentFailed    = 50
	RedeploymentSuccessed = 51
)

var (
	//Status app status
	Status = map[int]string{ //
		0: "AppBuilding",
		1: "AppSuccessed",
		2: "AppFailed",
		3: "AppRunning",
		4: "AppStop",
		5: "AppDelete",
		6: "AppUnknow",
	}

	//ListEverything list all namespace
	ListEverything = metav1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
)

//kubernetes StoragecClass
const (
	//ProvisionerCephRbd ceph rbd provisioner
	ProvisionerCephRbd = "kubernetes.io/rbd"
)

//Event the k8s event
type Event struct {
	Reason        string      `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	Message       string      `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`
	LastTimestamp metav1.Time `json:"lastTimestamp,omitempty" protobuf:"bytes,7,opt,name=lastTimestamp"`
	Type          string      `json:"type,omitempty" protobuf:"bytes,9,opt,name=type"`
}

//NewObjectMeta return metav1.ObjectMeta
func NewObjectMeta(name, namespace string, labels map[string]string) metav1.ObjectMeta {
	return metav1.ObjectMeta{Name: name, Namespace: namespace, Labels: labels}
}

//NewTypeMeta return metav1.TypeMeta
func NewTypeMeta(kind, vereion string) metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       kind,
		APIVersion: vereion,
	}
}

//NewDeployment return v1beta1.Deployment
func NewDeployment(name, namespace, nodeName string, labels, annotations map[string]string, replicas int, volumes []v1.Volume, initContainers, containers []v1.Container, nodeSelector map[string]string) *v1beta1.Deployment {
	return &v1beta1.Deployment{
		TypeMeta:   NewTypeMeta("Deployment", "apps/v1beta1"),
		ObjectMeta: NewObjectMeta(name, namespace, labels),
		Spec:       newDeploymentSpec(replicas, labels, annotations, volumes, initContainers, containers, nodeSelector, nodeName),
	}
}

func newDeploymentSpec(replicas int, labels, annotations map[string]string, volumes []v1.Volume, initContainers, containers []v1.Container, nodeSelector map[string]string, nodeName string) v1beta1.DeploymentSpec {
	return v1beta1.DeploymentSpec{
		Replicas: parse.IntToInt32Pointer(replicas),
		Selector: &metav1.LabelSelector{MatchLabels: labels},
		Template: v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels:      labels,
				Annotations: annotations,
			},
			Spec: v1.PodSpec{
				Volumes:        volumes,
				InitContainers: initContainers,
				Containers:     containers,
				RestartPolicy:  v1.RestartPolicyAlways,
				// NodeSelector:   nodeSelector,
				NodeName: nodeName,
			},
		},
	}
}

//NewService return v1.Service
func NewService(name, namespace string, labels map[string]string, serviceType v1.ServiceType, ports []v1.ServicePort) *v1.Service {
	return &v1.Service{
		TypeMeta:   NewTypeMeta("Service", "v1"),
		ObjectMeta: NewObjectMeta(name, namespace, labels),
		Spec: v1.ServiceSpec{
			Selector: labels,
			Type:     serviceType,
			Ports:    ports,
		},
	}
}

//NewHeadlessService return v1.Service with no clusterIP
func NewHeadlessService(name, namespace string, labels map[string]string, serviceType v1.ServiceType, ports []v1.ServicePort) *v1.Service {
	return &v1.Service{
		TypeMeta:   NewTypeMeta("Service", "v1"),
		ObjectMeta: NewObjectMeta(name, namespace, labels),
		Spec: v1.ServiceSpec{
			ClusterIP: "None",
			Selector:  labels,
			Type:      serviceType,
			Ports:     ports,
		},
	}
}

//NewConfigMap return v1.ConfigMap
func NewConfigMap(name, namespace string, labels, data map[string]string) *v1.ConfigMap {
	return &v1.ConfigMap{
		TypeMeta:   NewTypeMeta("ConfigMap", "v1"),
		ObjectMeta: NewObjectMeta(name, namespace, labels),
		Data:       data,
	}
}

//NewHPA return autoscalingv1.HorizontalPodAutoscaler
func NewHPA(name, namespace, refobject string, minReplicas, targetCPUUtilizationPercentage *int32, maxReplicas int32, labels map[string]string) *autoscalingv1.HorizontalPodAutoscaler {
	return &autoscalingv1.HorizontalPodAutoscaler{
		TypeMeta:   NewTypeMeta("HorizontalPodAutoscaler", "autoscaling/v1"),
		ObjectMeta: NewObjectMeta(name, namespace, labels),
		Spec: autoscalingv1.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: autoscalingv1.CrossVersionObjectReference{
				Kind: "Deployment",
				Name: refobject,
			},
			MaxReplicas:                    maxReplicas,
			MinReplicas:                    minReplicas,
			TargetCPUUtilizationPercentage: targetCPUUtilizationPercentage,
		},
	}
}

//NewPersistenVolumeClaim return v1.PersistentVolumeClaim
func NewPersistenVolumeClaim(name, namespace, accessMode, storageClassName string, labels map[string]string, resources v1.ResourceRequirements) *v1.PersistentVolumeClaim {
	return &v1.PersistentVolumeClaim{
		TypeMeta:   NewTypeMeta("PersistentVolumeClaim", "v1"),
		ObjectMeta: NewObjectMeta(name, namespace, labels),
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{v1.PersistentVolumeAccessMode(accessMode)},
			// Selector:         &metav1.LabelSelector{MatchLabels: labels},
			Resources:        resources,
			StorageClassName: parse.StringToPointer(storageClassName),
		},
	}
}

//NewStorageClass return storagev1.StorageClass
func NewStorageClass(name, provisioner string, parameters map[string]string) *storagev1.StorageClass {
	persistentVolumeReclaimRetain := new(v1.PersistentVolumeReclaimPolicy)
	*persistentVolumeReclaimRetain = v1.PersistentVolumeReclaimRetain
	return &storagev1.StorageClass{
		TypeMeta:      NewTypeMeta("StorageClass", "storage.k8s.io/v1"),
		ObjectMeta:    metav1.ObjectMeta{Name: name},
		Provisioner:   provisioner,
		Parameters:    parameters,
		ReclaimPolicy: persistentVolumeReclaimRetain,
	}
}

// NewStatefulSet return statefulset
func NewStatefulSet(name, namespace, nodeName, headlessServiceName string, labels, annotations, nodeSelector map[string]string, replicas int, volumes []v1.Volume, initContainers, containers []v1.Container, pvc v1.PersistentVolumeClaim) *v1beta1.StatefulSet {
	return &v1beta1.StatefulSet{
		TypeMeta:   NewTypeMeta("StatefulSet", "apps/v1beta1"),
		ObjectMeta: NewObjectMeta(name, namespace, labels),
		Spec: v1beta1.StatefulSetSpec{
			Replicas:             parse.IntToInt32Pointer(replicas),
			Selector:             &metav1.LabelSelector{MatchLabels: labels},
			VolumeClaimTemplates: []v1.PersistentVolumeClaim{pvc},
			ServiceName:          headlessServiceName,
			PodManagementPolicy:  v1beta1.ParallelPodManagement,
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      labels,
					Annotations: annotations,
				},
				Spec: v1.PodSpec{
					Volumes:        volumes,
					InitContainers: initContainers,
					Containers:     containers,
					RestartPolicy:  v1.RestartPolicyAlways,
					NodeSelector:   nodeSelector,
					NodeName:       nodeName,
				},
			},
		},
	}
}
