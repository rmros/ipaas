package validate

import (
	"encoding/json"
	"net/http"

	"ipaas/models"

	"k8s.io/api/core/v1"
)

func ValidateSpace(req *http.Request) (*models.Space, error) {
	space := &models.Space{}
	if err := json.NewDecoder(req.Body).Decode(space); err != nil {
		return nil, err
	}
	return space, nil
}

func ValidateUser(req *http.Request) (*models.User, error) {
	user := &models.User{}
	if err := json.NewDecoder(req.Body).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func ValidateTeam(req *http.Request) (*models.Team, error) {
	team := &models.Team{}
	if err := json.NewDecoder(req.Body).Decode(team); err != nil {
		return nil, err
	}
	return team, nil
}

func ValidateTeamAddUsers(req *http.Request) ([]*models.User, error) {
	users := []*models.User{}
	if err := json.NewDecoder(req.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func ValidateApp(req *http.Request) (*models.App, error) {
	app := &models.App{}
	if err := json.NewDecoder(req.Body).Decode(app); err != nil {
		return nil, err
	}
	return app, nil
}

func ValidateService(req *http.Request) (*models.Service, error) {
	svc := &models.Service{}
	if err := json.NewDecoder(req.Body).Decode(svc); err != nil {
		return nil, err
	}
	return svc, nil
}

func ValidateConfig(req *http.Request) (*models.Config, error) {
	config := &models.Config{}
	if err := json.NewDecoder(req.Body).Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}

func ValidateConfigData(req *http.Request) (map[string]string, error) {
	data := map[string]string{}
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func ValidateHPA(req *http.Request) (*models.HPA, error) {
	hpa := &models.HPA{}
	if err := json.NewDecoder(req.Body).Decode(hpa); err != nil {
		return nil, err
	}
	return hpa, nil
}

func ValidatePorts(req *http.Request) ([]v1.ServicePort, error) {
	ports := []v1.ServicePort{}
	err := json.NewDecoder(req.Body).Decode(&ports)
	return ports, err
}

func ValidateEnvs(req *http.Request) ([]v1.EnvVar, error) {
	envs := []v1.EnvVar{}
	err := json.NewDecoder(req.Body).Decode(&envs)
	return envs, err
}

func ValidateCephRBD(req *http.Request) (*models.CephRBD, error) {
	rbd := &models.CephRBD{}
	err := json.NewDecoder(req.Body).Decode(rbd)
	return rbd, err
}

func ValidateTickScaleTask(req *http.Request) (*models.TickScaleTask, error) {
	task := &models.TickScaleTask{}
	err := json.NewDecoder(req.Body).Decode(task)
	return task, err
}

func ValidateCluster(req *http.Request) (*models.Cluster, error) {
	cluster := &models.Cluster{}
	err := json.NewDecoder(req.Body).Decode(cluster)
	return cluster, err
}

func ValidateStorage(req *http.Request) (*models.Storage, error) {
	storage := &models.Storage{}
	err := json.NewDecoder(req.Body).Decode(storage)
	return storage, err
}

func Array(req *http.Request) ([]string, error) {
	inputs := []string{}
	err := json.NewDecoder(req.Body).Decode(&inputs)
	return inputs, err
}

func Map(req *http.Request) (map[string]string, error) {
	inputs := map[string]string{}
	err := json.NewDecoder(req.Body).Decode(&inputs)
	return inputs, err
}
