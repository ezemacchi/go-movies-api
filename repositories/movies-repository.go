package repositories

import (
	"movies-api/models"
)

type movieRepository struct {
	DbContext *DbContext
}

func (mr *movieRepository) getById(id int) (*models.Movie, error) {

	mr.DbContext.openDB()
}

func (mr *movieRepository) Insert(doc interface{}) (int64, error) { return 0, nil }
func (mr *movieRepository) Update(doc interface{}) (int64, error) { return 0, nil }
func (mr *movieRepository) Delete(doc interface{}) (int64, error) { return 0, nil }
