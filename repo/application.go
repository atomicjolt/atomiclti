package repo

import "github.com/atomicjolt/atomiclti/model"

type ApplicationRepo struct {
	*BaseRepo
}

func (r *ApplicationRepo) All() ([]model.Application, error) {
	var applications []model.Application
	err := r.DB.Model(&applications).Select()

	return applications, err
}

func (r *ApplicationRepo) Find(id int64) (*model.Application, error) {
	application := &model.Application{ID: id}
	err := r.DB.Model(application).WherePK().Select()

	return application, err
}
