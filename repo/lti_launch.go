package repo

import "github.com/atomicjolt/atomiclti/model"

type LtiLaunchRepo struct {
	*BaseRepo
}

func (r *LtiLaunchRepo) All() ([]model.LtiLaunch, error) {
	var ltiLaunches []model.LtiLaunch
	err := r.DB.Model(&ltiLaunches).Select()

	return ltiLaunches, err
}

func (r *LtiLaunchRepo) Find(id int64) (*model.LtiLaunch, error) {
	client := &model.LtiLaunch{ID: id}
	err := r.DB.Model(client).WherePK().Select()

	return client, err
}
