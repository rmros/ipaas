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
	"time"

	"ipaas/pkg/tools/storage/mysql"

	"github.com/jinzhu/gorm"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	mysql.GetDB().SingularTable(true)
	mysql.GetDB().CreateTable(
		new(App),
		new(Audit),
		new(Cluster),
		new(PodLifeCycle),
		new(Storage),
		new(TickScaleTask),
		new(User),
		new(Team),
		new(Space),
	)
}

// //Cluster the k8s cluster info
// type Cluster struct {
// 	ID            string `json:"cluster_id,omitempty" gorm:"primary_key"`
// 	Name          string `json:"name,omitempty"`
// 	Describe      string `json:"describe,omitempty"`
// 	ConfigContent string `json:"config_content,omitempty"`
// }

// Cluster is the structure of a cluster both in mysql table and json output
type Cluster struct {
	ID             string
	ClusterName    string
	Description    string
	PublicIPs      string
	BindingIPs     string
	BindingDomains string
	ConfigDetail   string
	Extention      string
	WebTerminal    string
	StorageID      string
	CreationTime   time.Time
	IsDefault      int8
	ResourcePrice  string
	Type           int8
	Cert           string
	Key            string

	// qing cloud's config
	Zone            string
	QingcloudURL    string
	AccessKeyID     string
	SecretAccessKey string

	// kubernetes cluster config
	APIProtocol string
	APIHost     string
	APIToken    string
	Content     string
	APIVersion  string
}

//App the application that use deploy
type App struct {
	ID            uint       `json:"id,omitempty"`
	CreatedAt     time.Time  `json:"createAt,omitempty"`
	Name          string     `json:"name,omitempty"`
	UserName      string     `json:"nameSpace,omitempty"`
	Description   string     `json:"description,omitempty"`
	AppStatus     int        `json:"appStatus,omitempty"`
	ServiceCount  int        `json:"serviceCount,omitempty"`
	InstanceCount int        `json:"intanceCount,omitempty"`
	External      string     `json:"external,omitempty"`
	Items         []*Service `json:"services,omitempty" gorm:"-"`
}

//Service the app's service
type Service struct {
	AppName       string           `json:"appName,omitempty"`
	Name          string           `json:"name,omitempty"`
	Image         string           `json:"image,omitempty"`
	InstanceCount int              `json:"instanceCount" `
	Status        int              `json:"status,omitempty"`
	Type          int              `json:"type,omitempty"` //0 stateless 1 statefulset
	NodeName      string           `json:"nodeName,omitempty"`
	External      string           `json:"external,omitempty"`
	LoadbalanceIP string           `json:"loadbalanceIP,omitempty"`
	CPU           string           `json:"cpu,omitempty"`
	Memory        string           `json:"memory,omitempty"`
	Ports         []v1.ServicePort `json:"ports,omitempty"`
	Envs          []v1.EnvVar      `json:"envs,omitempty"`
	Cmds          []string         `json:"cmds,omitempty"`
	Volumes       []Volume         `json:"volumes,omitempty"`
	Storage       *Storage         `json:"storage,omitempty"`
}

//Volume service volume
type Volume struct {
	Type          int      `json:"type,omitempty"` //0 挂载整个配置组（目录的方式挂载）1 挂载配置组中的一个key（挂载配置组中的单个文件）
	MountPath     string   `json:"mountPath,omitempty"`
	ConfigMapName string   `json:"configMapName,omitempty"`
	ConfigMapKey  []string `json:"configMapKey,omitempty"`
}

//Storage rdb storage
type Storage struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name"`
	Type        string    `json:"type"` //default is rbd
	Size        string    `json:"size"`
	AccessModes string    `json:"accessModes"`
	Namespace   string    `json:"namespace"`
	CreateAt    time.Time `json:"createAt"`
	Status      string    `json:"status"`
	Used        bool      `json:"used"`
	ServiceName string    `json:"serviceName"`
	MountPath   string    `json:"mountPath"`
}

//Config config group
type Config struct {
	Name      string            `json:"name,omitempty"`
	Namespace string            `json:"namespace,omitempty"`
	Data      map[string]string `json:"data,omitempty"`
}

//Event the resource event
type Event struct {
	Reason        string      `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	Message       string      `json:"message,omitempty" protobuf:"bytes,4,opt,name=message"`
	LastTimestamp metav1.Time `json:"lastTimestamp,omitempty" protobuf:"bytes,7,opt,name=lastTimestamp"`
	Type          string      `json:"type,omitempty" protobuf:"bytes,9,opt,name=type"`
}

//Process container's process
type Process struct {
	User        string  `json:"user,omitempty"`
	PID         int64   `json:"pid,omitempty"`
	ParentPID   int64   `json:"parent_pid,omitempty"`
	StartTime   string  `json:"start_time,omitempty"`
	PercentCPU  float64 `json:"percent_cpu,omitempty"`
	PercentMEM  float64 `json:"percent_mem,omitempty"`
	RSS         int64   `json:"rss,omitempty"`
	VirtualSize int64   `json:"virtual_size,omitempty"`
	Status      string  `json:"status,omitempty"`
	RunningTime string  `json:"running_time,omitempty"`
	CgroupPath  string  `json:"cgroup_path,omitempty"`
	Cmd         string  `json:"cmd,omitempty"`
}

//Item Cadvisor api result item
type Item struct {
	Name string `json:"name,omitempty"`
}

//CadvisorResult request Cadvisor api result
type CadvisorResult struct {
	Subcontainers []*Item `json:"subcontainers,omitempty"`
}

//HPA k8s rsource hpa's dto
type HPA struct {
	Name                           string `json:"name,omitempty"`
	RefObjectName                  string `json:"refObject,omitempty"`
	MinReplicas                    *int32 `json:"minReplicas,omitempty"`
	MaxReplicas                    int32  `json:"maxReplicas,omitempty"`
	TargetCPUUtilizationPercentage *int32 `json:"targetCPUUtilizationPercentage,omitempty"`
}

//CephRBD ceph rbd's info
type CephRBD struct {
	Provisioner          string `json:"provisioner,omitempty"`
	Name                 string `json:"name,omitempty"`
	Monitors             string `json:"monitors,omitempty"`
	AdminID              string `json:"adminId,omitempty"`
	AdminSecretName      string `json:"adminSecretName,omitempty"`
	AdminSecretNamespace string `json:"adminSecretNamespace,omitempty"`
	Pool                 string `json:"pool,omitempty"`
	UserID               string `json:"userId,omitempty"`
	UserSecretName       string `json:"userSecretName,omitempty"`
	FsType               string `json:"fsType,omitempty"`
	ImageFormat          string `json:"imageFormat,omitempty"`
	ImageFeatures        string `json:"imageFeatures,omitempty"`
}

// TickScaleTask ticker time scale task
type TickScaleTask struct {
	ID                             uint      `json:"id,omitempty"`
	Name                           string    `json:"name,omitempty"`
	Namespace                      string    `json:"namespace,omitempty"`
	Spec                           string    `json:"spec,omitempty"`
	Desired                        int32     `json:"desired,omitempty"`
	ClusterID                      string    `json:"clusterID,omitempty"`
	Status                         uint      `json:"status,omitempty"` //0 off 1 on
	CreationTime                   time.Time `json:"creationTime,omitempty"`
	RefObjectName                  string    `json:"refObject,omitempty"`
	MinReplicas                    *int32    `json:"minReplicas,omitempty"`
	MaxReplicas                    int32     `json:"maxReplicas,omitempty"`
	TargetCPUUtilizationPercentage *int32    `json:"targetCPUUtilizationPercentage,omitempty"`
}

// Audit is the operation recoder
type Audit struct {
	ID              uint      `json:"id,omitempty"`
	UserName        string    `json:"userName,omitempty"`
	Namespace       string    `json:"namespace,omitempty"`
	ClusterName     string    `json:"clusterName,omitempty"`
	ReferenceObject string    `json:"referenceObject,omitempty"`
	Operation       string    `json:"operation,omitempty"`
	Status          int       `json:"appStatus,omitempty"`
	CreatedAt       time.Time `json:"createAt,omitempty"`
}

// PodLifeCycle record service's pod life cycle
type PodLifeCycle struct {
	ID          uint   `json:"id"`
	ClusterID   string `json:"clusterID"`
	Namespace   string `json:"namespace"`
	ServiceName string `json:"serviceName"`
	PodName     string `json:"podName"`
	DeleteAt    string `json:"delete_at"`
	CreateAt    string `json:"create_at"`
	Status      string `json:"status"`
}

const (
	_ = iota
	// RoleNormal the normal user
	RoleNormal
	// RoleTeam the team manager
	RoleTeam
	// RoleAdmin the system manager
	RoleAdmin
	// RoleSuper the system super manager
	RoleSuper
)

const (
	// SuperUserName default super user name for now
	SuperUserName = "admin"
)

// User user info
type User struct {
	ID             uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt      time.Time  `json:"createdAt,omitempty"`
	UpdatedAt      time.Time  `json:"updatedAt,omitempty"`
	DeletedAt      *time.Time `json:"deletedAt,omitempty" sql:"index"`
	Name           string     `json:"name,omitempty" gorm:"UNIQUE_INDEX"`
	Displayname    string     `json:"displayname,omitempty"`
	Password       string     `json:"password,omitempty"`
	Email          string     `json:"email,omitempty"`
	Phone          string     `json:"phone,omitempty"`
	LoginFrequency int        `json:"loginFrequency,omitempty"`
	Active         int8       `json:"active,omitempty" gorm:"DEFAULT:1"`
	APIToken       string     `json:"apiToken,omitempty"`
	Role           int32      `json:"role,omitempty" gorm:"DEFAULT:1"`
	Type           int        `json:"type,omitempty"`
	Company        string     `json:"company,omitempty"`
	Teams          []*Team    `gorm:"teams,many2many:user_teams;"`
}

// TableName return user model's  table name
func (user *User) TableName() string {
	return "users"
}

// Team team info
type Team struct {
	gorm.Model
	Name        string  `json:"name,omitempty" gorm:"primary_key"`
	Description string  `json:"description,omitempty"`
	CreatorID   int32   `json:"creatorID,omitempty"`
	Users       []*User `json:"users,omitempty" `
}

// TableName return team model's  table name
func (team *Team) TableName() string {
	return "teams"
}

// Space space info
type Space struct {
	ID           string    `json:"id,omitempty" gorm:"primary_key"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	TeamID       string    `json:"teamID"`
	CreationTime time.Time `json:"creationTime,omitempty"`
	Type         int       `json:"type,omitempty"` // 1 personal namespace 2 team's namespace
}

// TableName return Space model's  table name
func (space *Space) TableName() string {
	return "spaces"
}

// Company company info
type Company struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Description string
}

// TableName return Company model's  table name
func (company *Company) TableName() string {
	return "companys"
}

// Container pod info
type Container struct {
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	AppName   string    `json:"appName"`
	Namespace string    `json:"namespace"`
	Image     string    `json:"image"`
	URL       string    `json:"url"`
	CreateAt  time.Time `json:"createAt"`
}

// TableName return Container model's  table name
func (container *Container) TableName() string {
	return "containers"
}

type Node struct {
	HostName          string            `json:"hostName"`
	Internal          string            `json:"internal"`
	Status            bool              `json:"status"`
	MasterOrSlave     string            `json:"matserOrslave"`
	ContainerCnt      int               `json:"containerCnt"`
	CPUUsage          int               `json:"cpuUsage"`
	CPUAllocatable    int64             `json:"cpuAllocatable"`
	CPUCapacity       int64             `json:"cpuCapacity"`
	MemoryUsage       int               `json:"memoryUsage"`
	MemoryAllocatable int64             `json:"memoryAllocatable"`
	MemoryCapacity    int64             `json:"memoryCapacity"`
	PodCapacity       int64             `json:"podCapacity"`
	Schedulable       bool              `json:"schedulable"`
	DiskPressure      bool              `json:"diskPressure"`
	MemoryPressure    bool              `json:"memoryPressure"`
	CreateTime        metav1.Time       `json:"createT_at"`
	NodeVersion       v1.NodeSystemInfo `json:"version"`
	Containers        []Container       `json:"containers"`
}
