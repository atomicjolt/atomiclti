package repo

import "github.com/atomicjolt/atomiclti/model"

type ApplicationInstanceRepo struct {
	*BaseRepo
}

func (r *ApplicationInstanceRepo) All() ([]model.ApplicationInstance, error) {
	var applicationInstances []model.ApplicationInstance
	err := r.DB.Model(&applicationInstances).Select()

	return applicationInstances, err
}

func (r *ApplicationInstanceRepo) Find(id int64) (*model.ApplicationInstance, error) {
	applicationInstance := &model.ApplicationInstance{ID: id}
	err := r.DB.Model(applicationInstance).WherePK().Select()

	return applicationInstance, err
}
