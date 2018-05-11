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
