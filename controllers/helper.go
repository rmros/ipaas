package controllers

import (
	"fmt"
	"strings"

	"ipaas/models"
	k8s "ipaas/pkg/k8s/client"
	"ipaas/pkg/tools/log"
	"ipaas/pkg/tools/parse"

	"k8s.io/api/apps/v1beta1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TOK8sNamespace(name string) *v1.Namespace {
	return &v1.Namespace{
		TypeMeta:   metav1.TypeMeta{Kind: "Namespace", APIVersion: "v1"},
		ObjectMeta: metav1.ObjectMeta{Name: name},
	}
}

// DelpoyService deploy service of paas
func DelpoyService(clusterID string, svc *v1.Service, deploy *v1beta1.Deployment) (interface{}, error) {
	fake := k8s.GetClientset(clusterID)
	if fake == nil {
		return nil, fmt.Errorf("the k8s cluster %q has no client exist", clusterID)
	}

	var result []interface{}
CREATE_SERVICE:
	ressvc, err := fake.Services().CreateService(svc)
	if err != nil {
		if errors.IsConflict(err) {
			goto CREATE_SERVICE
		}
		return nil, err
	}
CREATE_DEPLOYMENT:
	resdeploy, err := fake.DeploymentsExtensions().CreateDeployment(deploy)
	if err != nil {
		if errors.IsConflict(err) {
			goto CREATE_DEPLOYMENT
		}
		go fake.Services().DeleteService(svc.Name, svc.Namespace)
		return nil, err
	}
	result = append(result, resdeploy, ressvc)
	return result, nil
}

// DeleteService delete service of paas
func DeleteService(name, namespace, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	if fake == nil {
		return fmt.Errorf("the k8s cluster %q has no client exist", clusterID)
	}
DELETE_SERVICE:
	err := fake.Services().DeleteService(name, namespace)
	if err != nil {
		if errors.IsConflict(err) {
			goto DELETE_SERVICE
		}
		return err
	}
DELETE_DEPLOYMENT:
	err = fake.DeploymentsExtensions().DeleteDeployment(name, namespace)
	if err != nil {
		if errors.IsConflict(err) {
			goto DELETE_DEPLOYMENT
		}
		return err
	}
	return nil
}

// DeleteServiceByAppName delete service of paas by app name
func DeleteServiceByAppName(name, namespace, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	labels := fmt.Sprintf("%v=%v", models.MinipaasAppName, name)
	if err := fake.DeploymentsExtensions().DeleteDeploymentByLabels(labels, namespace); err != nil {
		return err
	}
	if err := fake.Services().DeleteServiceByLabels(labels, namespace); err != nil {
		return err
	}
	return nil
}

// ListServiceByAppName list service of paas by app name
func ListServiceByAppName(name, namespace, clusterID string) ([]v1.Service, []v1beta1.Deployment, error) {
	fake := k8s.GetClientset(clusterID)
	labels := fmt.Sprintf("%v=%v", models.MinipaasAppName, name)
	services, err := fake.Services().ListService(labels, namespace)
	if err != nil {
		return []v1.Service{}, []v1beta1.Deployment{}, err
	}
	deployments, err := fake.DeploymentsExtensions().ListDeployment(labels, namespace)
	if err != nil {
		return []v1.Service{}, []v1beta1.Deployment{}, err
	}
	return services, deployments, nil
}

// ServiceExist assert service of paas exsit or not
func ServiceExist(name, namespace, clusterID string) (*v1.Service, *v1beta1.Deployment, bool) {
	fake := k8s.GetClientset(clusterID)
	svc, err := fake.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, false
	}
	deploy, err := fake.AppsV1beta1().Deployments(namespace).Get(name, metav1.GetOptions{})
	if err != nil {
		return nil, nil, false
	}
	return svc, deploy, true
}

func stopService(svc *v1.Service, deploy *v1beta1.Deployment, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	deploy.Spec.Replicas = parse.IntToInt32Pointer(0)
	if _, err := fake.DeploymentsExtensions().UpdateDeployment(deploy); err != nil {
		return err
	}
	return nil
}

func startService(svc *v1.Service, deploy *v1beta1.Deployment, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	deploy.Spec.Replicas = parse.StringToInt32Pointer(deploy.ObjectMeta.Labels["replicas"])
	if _, err := fake.DeploymentsExtensions().UpdateDeployment(deploy); err != nil {
		return err
	}
	return nil
}

func redeployService(svc *v1.Service, deploy *v1beta1.Deployment, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	pods, err := fake.Pods().ListPodByDeploymentName(deploy.Name, deploy.Namespace)
	if err != nil {
		return nil
	}
	for i := range pods {
		if err := fake.Pods().DeletePod(pods[i]); err != nil {
			return err
		}
	}
	return nil
}

// StartOrStopOrRedployService start or stop or redeploy service of paas
func StartOrStopOrRedployService(svc *v1.Service, deploy *v1beta1.Deployment, verb, clusterID string) error {
	if verb == "stop" {
		if err := stopService(svc, deploy, clusterID); err != nil {
			return fmt.Errorf("stop service %v err: %v", deploy.Name, err)
		}
	}
	if verb == "start" {
		if err := startService(svc, deploy, clusterID); err != nil {
			return fmt.Errorf("start service %v err: %v", deploy.Name, err)
		}
	}
	if verb == "redeploy" {
		if err := redeployService(svc, deploy, clusterID); err != nil {
			return fmt.Errorf("redeploy service %v err: %v", deploy.Name, err)
		}
	}
	return nil
}

// StartOrStopOrRedployApp  start or stop or redeploy app of paas
func StartOrStopOrRedployApp(svcs []v1.Service, deploys []v1beta1.Deployment, verb, clusterID string) (errs []error) {
	for _, deploy := range deploys {
		if err := StartOrStopOrRedployService(nil, &deploy, verb, clusterID); err != nil {
			errs = append(errs, fmt.Errorf("%v service %v err: %v", verb, deploy.Name, err))
			continue
		}
	}
	return
}

// ListService list service of paas
func ListService(labels, namespace, clusterID string) ([]v1.Service, []v1beta1.Deployment, error) {
	fake := k8s.GetClientset(clusterID)
	services, err := fake.Services().ListService(labels, namespace)
	if err != nil {
		return []v1.Service{}, []v1beta1.Deployment{}, err
	}
	deployments, err := fake.DeploymentsExtensions().ListDeployment("", namespace)
	if err != nil {
		return []v1.Service{}, []v1beta1.Deployment{}, err
	}
	return services, deployments, nil
}

// CreateConfigMap create configMap of k8s
func CreateConfigMap(clusterID string, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.ConfigMaps().CreateConfigMap(configMap)
}

// GetConfigMapByName get configMap by name and namespace
func GetConfigMapByName(name, namespace, clusterID string) (*v1.ConfigMap, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.ConfigMaps().GetConfigMap(name, namespace)
}

// UpdateConfigMap update configMap of k8s
func UpdateConfigMap(clusterID string, configMap *v1.ConfigMap) (*v1.ConfigMap, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.ConfigMaps().UpdateConfigMap(configMap)
}

// DeleteConfigMap delete configMap of k8s
func DeleteConfigMap(name, namespace, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	return fake.ConfigMaps().DeleteConfigMap(name, namespace)
}

// ListConfigMap list configMap of k8s
func ListConfigMap(namespace, clusterID string) ([]v1.ConfigMap, error) {
	fake := k8s.GetClientset(clusterID)
	configMapList, err := fake.ConfigMaps().ListConfigMap(namespace)
	if err != nil {
		return []v1.ConfigMap{}, err
	}
	return configMapList.Items, nil
}

// GetPodEvents get pod events
func GetPodEvents(name, namespace, clusterID string) ([]models.Event, error) {
	fake := k8s.GetClientset(clusterID)
	events, err := fake.Events().GetEvents(namespace)
	if err != nil {
		return []models.Event{}, err
	}
	var list []models.Event
	for _, event := range events {
		if strings.Contains(event.Name, name) {
			list = append(
				list,
				models.Event{
					Reason:        event.Reason,
					Type:          event.Type,
					LastTimestamp: event.LastTimestamp,
					Message:       event.Message,
				},
			)
		}
	}
	return list, nil
}

// GetPodLogs get pod logs
func GetPodLogs(name, namespace, clusterID string, logOptions *v1.PodLogOptions) (string, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.Pods().GetPodLogs(name, namespace, logOptions)
}

// GetPod get pod of k8s
func GetPod(name, namespace, clusterID string) (*v1.Pod, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.Pods().GetPod(name, namespace)
}

//GetServiceEvents get service event
func GetServiceEvents(name, namespace, clusterID string) ([]models.Event, error) {
	fake := k8s.GetClientset(clusterID)
	events, err := fake.Events().GetEvents(namespace)
	if err != nil {
		return []models.Event{}, err
	}
	var list []models.Event
	for _, event := range events {
		if strings.Contains(event.Name, name) {
			list = append(
				list,
				models.Event{
					Reason:        event.Reason,
					Type:          event.Type,
					LastTimestamp: event.LastTimestamp,
					Message:       event.Message,
				},
			)
		}
	}
	return list, nil
}

// GetPods get pods of k8s
func GetPods(namespace, clusterID string) ([]v1.Pod, error) {
	fake := k8s.GetClientset(clusterID)
	pods, err := fake.Pods().ListPods(namespace)
	if err != nil {
		return []v1.Pod{}, err
	}
	return pods, nil
}

// CreateHPA create hpa of k8s
func CreateHPA(clusterID string, hpa *autoscalingv1.HorizontalPodAutoscaler) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.HPAs().CreateHPA(hpa)
}

// UpdateHPA update hpa of k8s
func UpdateHPA(clusterID string, hpa *autoscalingv1.HorizontalPodAutoscaler) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.HPAs().UpdateHPA(hpa)
}

// DeleteHPA delete hpa of k8s
func DeleteHPA(name, namespace, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	return fake.HPAs().DeleteHPA(name, namespace)
}

// GetHPA get hpa of k8s
func GetHPA(name, namespace, clusterID string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.HPAs().GetHPA(name, namespace)
}

// GetDeployment get deployment of k8s
func GetDeployment(name, namespace, clusterID string) (*v1beta1.Deployment, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.DeploymentsExtensions().GetDeployment(name, namespace)
}

// GetK8SService get k8s service
func GetK8SService(name, namespace, clusterID string) (*v1.Service, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.Services().GetService(name, namespace)
}

// UpdateDeployment update deployment
func UpdateDeployment(deploy *v1beta1.Deployment, clusterID string) (*v1beta1.Deployment, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.DeploymentsExtensions().UpdateDeployment(deploy)
}

// UpdateService update service of paas
func UpdateService(svc *v1.Service, deploy *v1beta1.Deployment, clusterID string) (*v1.Service, *v1beta1.Deployment, error) {
	fake := k8s.GetClientset(clusterID)
UPDATE_DEPLOYMENT:
	deployment, err := fake.DeploymentsExtensions().UpdateDeployment(deploy)
	if err != nil {
		if errors.IsConflict(err) {
			goto UPDATE_DEPLOYMENT
		}
		return nil, nil, err
	}
UPDATE_SERVICE:
	service, err := fake.Services().UpdateService(svc)
	if err != nil {
		if errors.IsConflict(err) {
			goto UPDATE_SERVICE
		}
		return nil, deployment, err
	}
	return service, deployment, nil
}

// CreateStorageClass create storageclass
func CreateStorageClass(storageclass *storagev1.StorageClass, clusterID string) (*storagev1.StorageClass, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.StorageClasses().CreateStorageClass(storageclass)
}

// DeleteStorageClass delete storageclass by name
func DeleteStorageClass(name, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	return fake.StorageClasses().DeleteStorageClass(name)
}

// DeploySatefulService deploy stateful service
func DeploySatefulService(service, headlessService *v1.Service, statefulset *v1beta1.StatefulSet, clusterID string) ([]interface{}, error) {
	var result []interface{}
	fake := k8s.GetClientset(clusterID)
	services, err := createK8SServices(clusterID, service, headlessService)
	if err != nil {
		go deleteK8SServices(service.Namespace, clusterID, service.Name, headlessService.Name)
		return result, err
	}
	stateful, err := fake.StatefulsetsV1beta1().CreateStatefulSet(statefulset)
	if err != nil {
		go deleteK8SServices(service.Namespace, clusterID, service.Name, headlessService.Name)
		return result, err
	}
	result = append(result, services, stateful)
	return result, nil
}

// DeleteStatefulService delete stateful service
func DeleteStatefulService(serviceName, namespace, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	statefulset, err := getStatefulSet(serviceName, namespace, clusterID)
	if err != nil {
		log.Error("when delete statefulservice %v,get it's statefuleset err: %v", serviceName, err)
		return err
	}
	services, err := getStatefulSetServices(clusterID, statefulset)
	if err != nil {
		log.Error("when delete statefulservice %v,get it's services  err: %v", serviceName, err)
		return err
	}
	serviceNames := []string{}
	for i := range services {
		serviceNames = append(serviceNames, services[i].Name)
	}
	pvcnames := getStatefulSetPVCNames(statefulset)
	// delete statefulset
	if err = fake.Statefulsets().DeleteStatefulSet(statefulset.Name, namespace); err != nil {
		log.Error("when delete statefulservice %v,delete it's statefuleset  err: %v", serviceName, err)
		return err
	}
	// delete statefulset's services
	if err = deleteK8SServices(namespace, clusterID, serviceNames...); err != nil {
		log.Error("when delete statefulservice %v,delete it's serivces  err: %v", serviceName, err)
		return err
	}
	// delete statefulset's pvcs
	if err = deletePVC(namespace, clusterID, pvcnames...); err != nil {
		log.Error("when delete statefulservice %v,delete it's persistentVolumeClaims  err: %v", serviceName, err)
		return err
	}
	return nil
}

func createK8SServices(clusterID string, services ...*v1.Service) ([]*v1.Service, error) {
	svcs := []*v1.Service{}
	fake := k8s.GetClientset(clusterID)
	for i := range services {
		svc, err := fake.Services().CreateService(services[i])
		if err != nil {
			return svcs, err
		}
		svcs = append(svcs, svc)
	}
	return svcs, nil
}

func deleteK8SServices(namespace, clusterID string, services ...string) error {
	fake := k8s.GetClientset(clusterID)
	for i := range services {
		if err := fake.Services().DeleteService(services[i], namespace); err != nil {
			log.Error("delete service %v err: %v", services[i], err)
			continue
		}
		log.Info("delete service %v success", services[i])
	}
	return nil
}

func deletePVC(namespace, clusterID string, names ...string) error {
	fake := k8s.GetClientset(clusterID)
	for i := range names {
		if err := fake.PersistentVolumeClaims().DeletePersistentVolumeClaim(names[i], namespace); err != nil {
			log.Error("delete PersistentVolumeClaim %v err: %v", names[i], err)
			continue
		}
		log.Info("delete PersistentVolumeClaim %v success", names[i])
	}
	return nil
}

func getStatefulSetPVCNames(statefulset *v1beta1.StatefulSet) []string {
	replicas := statefulset.Spec.Replicas
	pvcsName := []string{}
	for i := 0; i < int(*replicas); i++ {
		pvcname := fmt.Sprintf("%v-%v-%v", statefulset.Spec.VolumeClaimTemplates[0].Name, statefulset.Name, i)
		pvcsName = append(pvcsName, pvcname)
	}
	return pvcsName
}

func getStatefulSet(name, namespace, clusterID string) (*v1beta1.StatefulSet, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.StatefulsetsV1beta1().GetStatefulSet(name, namespace)
}

func getStatefulSetServices(clusterID string, statefulset *v1beta1.StatefulSet) ([]v1.Service, error) {
	fake := k8s.GetClientset(clusterID)
	var labels []string
	for k, v := range statefulset.Labels {
		if k != "replicas" {
			label := fmt.Sprintf("%v=%v", k, v)
			labels = append(labels, label)
		}
	}
	services, err := fake.Services().ListService(strings.Join(labels, ","), statefulset.Namespace)
	if err != nil {
		return []v1.Service{}, err
	}
	return services, nil
}

func getStatefulSetPVCs(clusterID string, statefulset *v1beta1.StatefulSet) ([]v1.PersistentVolumeClaim, error) {
	fake := k8s.GetClientset(clusterID)
	var labels []string
	for k, v := range statefulset.Labels {
		if k != "replicas" {
			label := fmt.Sprintf("%v=%v", k, v)
			labels = append(labels, label)
		}
	}
	return fake.PersistentVolumeClaims().ListPersistentVolumeClaim(strings.Join(labels, ","), statefulset.Namespace)
}

// CreatePersistentVolumeClaim create PersistentVolumeClaim
func CreatePersistentVolumeClaim(pvc *v1.PersistentVolumeClaim, clusterID string) (*v1.PersistentVolumeClaim, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.PersistentVolumeClaims().CreatePersistentVolumeClaim(pvc)
}

// DeletePersistentVolumeClaim delete PersistentVolumeClaim
func DeletePersistentVolumeClaim(name, namespace, clusterID string) error {
	fake := k8s.GetClientset(clusterID)
	return fake.PersistentVolumeClaims().DeletePersistentVolumeClaim(name, namespace)
}

// ListPersistentVolumeClaim list PersistentVolumeClaim by namespace
func ListPersistentVolumeClaim(namespace, clusterID string) ([]v1.PersistentVolumeClaim, error) {
	fake := k8s.GetClientset(clusterID)
	return fake.PersistentVolumeClaims().ListPersistentVolumeClaim("", namespace)
}
