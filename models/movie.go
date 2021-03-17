package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MovieModel struct {
	DB *gorm.DB
}

type Movie struct {
	ID         int64     `gorm:"id"`
	Name       string    `gorm:"name"`
	Year       int       `gorm:"year"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
}

func (m *Movie) TableName() string {
	return "movie"
}

func (movieModel *MovieModel) Insert(m *Movie) error {
	if err := movieModel.DB.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (movieModel *MovieModel) GetMovieByName(name string) (*Movie, error) {
	m := Movie{
		Name: name,
	}
	movie := movieModel.DB.Take(&m)
	if movie != nil && movie.Error != nil {
		return nil, movie.Error
	}
	return &m, nil
}
