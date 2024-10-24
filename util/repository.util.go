package util

import "gorm.io/gorm"

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func (r *BaseRepository[T]) Save(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *BaseRepository[T]) FindById(id uint) (*T, error) {
	var entity T
	result := r.DB.First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(id uint) error {
	var entity T
	result := r.DB.Delete(&entity, id)
	return result.Error
}
