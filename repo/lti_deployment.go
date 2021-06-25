package repo

import "github.com/atomicjolt/atomiclti/model"

type LtiDeploymentRepo struct {
	*BaseRepo
}

func (r *LtiDeploymentRepo) All() ([]model.LtiDeployment, error) {
	var ltiDeployments []model.LtiDeployment
	err := r.DB.Model(&ltiDeployments).Select()

	return ltiDeployments, err
}

func (r *LtiDeploymentRepo) Find(id int64) (*model.LtiDeployment, error) {
	ltiDeployment := &model.LtiDeployment{ID: id}
	err := r.DB.Model(ltiDeployment).WherePK().Select()

	return ltiDeployment, err
}
