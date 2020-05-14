package crud

import (
	"errors"
	"github.com/kerimkuscu/hardline-fitness-backend/api/models"
	"github.com/kerimkuscu/hardline-fitness-backend/api/utils/channels"
	"time"

	"github.com/jinzhu/gorm"
)

type repositoryProgramsCRUD struct {
	db *gorm.DB
}

func NewRepositoryProgramsCRUD(db *gorm.DB) *repositoryProgramsCRUD {
	return &repositoryProgramsCRUD{db}
}

func (r *repositoryProgramsCRUD) Save(program models.Program) (models.Program, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Program{}).Create(&program).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return program, nil
	}
	return models.Program{}, err
}

func (r *repositoryProgramsCRUD) FindAll() ([]models.Program, error) {
	var err error
	programs := []models.Program{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Program{}).Limit(100).Find(&programs).Error
		if err != nil {
			ch <- false
			return
		}
		if len(programs) > 0 {
			for i, _ := range programs {
				err := r.db.Debug().Model(&models.User{}).Where("id = ?", programs[i].AuthorID).Take(&programs[i].Author).Error
				if err != nil {
					ch <- false
					return
				}
			}
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return programs, nil
	}
	return nil, err
}

func (r *repositoryProgramsCRUD) FindByID(pid uint64) (models.Program, error) {
	var err error
	program := models.Program{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.Program{}).Where("id = ?", pid).Take(&program).Error
		if err != nil {
			ch <- false
			return
		}

		if program.ID != 0 {
			err = r.db.Debug().Model(&models.User{}).Where("id = ?", program.AuthorID).Take(&program.Author).Error

			if err != nil {
				ch <- false
				return
			}
		}
		ch <- true
	}(done)

	if channels.OK(done) {
		return program, nil
	}
	return models.Program{}, err
}

func (r *repositoryProgramsCRUD) Update(pid uint64, program models.Program) (int64, error) {
	// var err error
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Program{}).Where("id = ?", pid).Take(&models.Program{}).UpdateColumns(
			map[string]interface{}{
				"title":      program.Title,
				"content":    program.Content,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)

	if channels.OK(done) {
		// return program, nil
		if rs.Error != nil {
			if gorm.IsRecordNotFoundError(rs.Error) {
				return 0, errors.New("Program not found")
			}
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *repositoryProgramsCRUD) Delete(pid uint64, uid uint32) (int64, error) {
	// var err error
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.Program{}).Where("id = ? and author_id = ?", pid, uid).Take(&models.Program{}).Delete(&models.Program{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			if gorm.IsRecordNotFoundError(rs.Error) {
				return 0, errors.New("Program not found")
			}
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}
