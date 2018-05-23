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

package infrastructure

import (
	"ipaas/pkg/tools/storage/mysql"
	"time"
)

func init() {

}

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

func init() {
	mysql.GetDB().SingularTable(true)
	mysql.GetDB().CreateTable(
		new(Cluster),
	)
}
