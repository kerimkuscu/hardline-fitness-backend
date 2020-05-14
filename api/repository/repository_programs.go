package repository

import "github.com/kerimkuscu/hardline-fitness-backend/api/models"

type ProgramRepository interface {
	Save(models.Program) (models.Program, error)
	FindAll() ([]models.Program, error)
	FindByID(uint64) (models.Program, error)
	Update(uint64, models.Program) (int64, error)
	Delete(post_id uint64, user_id uint32) (int64, error)
}
