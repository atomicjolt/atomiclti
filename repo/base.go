package repo

import "github.com/go-pg/pg/v10/orm"

// BaseRepo is intended to be used as an embedded struct in other Repo types. It
// provides shared methods
type BaseRepo struct {
	DB orm.DB
}

// Insert inserts a record into any table, based on the type of struct passed in
func (r *BaseRepo) Insert(record interface{}) error {
	_, err := r.DB.Model(record).Insert()
	return err
}

// Update an existing record
func (r *BaseRepo) Update(record interface{}) error {
	_, err := r.DB.Model(record).WherePK().Update()
	return err
}

// Delete a record (actually a soft delete)
func (r *BaseRepo) Delete(record interface{}) {
	r.DB.Model(record).WherePK().Delete()
}
