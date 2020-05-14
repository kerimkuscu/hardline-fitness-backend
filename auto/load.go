package auto

import (
	"github.com/kerimkuscu/hardline-fitness-backend/api/database"
	"github.com/kerimkuscu/hardline-fitness-backend/api/models"
	"log"
)

func Load() {
	db, err := database.Connect()

	if err != nil {
		log.Fatalf("cannot connect to the database: %v", err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.Program{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Program{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Program{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Program{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}

		// err = db.Debug().Model(&models.Program{}).Where("author_id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
		// if err != nil {
		// 	log.Fatalf("cannot seed posts table: %v", err)
		// }
		// console.Pretty(posts[i])
	}
}
