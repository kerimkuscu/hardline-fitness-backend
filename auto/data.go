package auto

import "github.com/kerimkuscu/hardline-fitness-backend/api/models"

var users = []models.User{
	models.User{
		Nickname: "Kerim Kuscu",
		Email:    "kerimkuscu@gmail.com",
		Password: "123456",
	}, {
		Nickname: "Buse Demirbas",
		Email:    "busedemirbas@gmail.com",
		Password: "123456",
	},
}

var posts = []models.Program{
	models.Program{
		Title:   "Leg Training",
		Content: "1)Leg Press\n2)Step-Up\n3)Pistol Squat\n4)Glute-Ham Raise\n5)Walking Lunge",
	}, {
		Title:   "Biceps Training",
		Content: "1)Incline Dumbbell Hammer Curl\n2)Incline Inner-Biceps Curl\n3)EZ-Bar Curl\n4)Dumbbell Biceps Curl\n5)Hammer Curl",
	}, {
		Title:   "Triceps Training",
		Content: "1)Overall Triceps Mass\n2)Close Grip Barbell Bench Press\n3)Parallel Dips\n4)Triceps Pushdown\n5)Skullcrushers",
	}, {
		Title:   "Shoulder Training",
		Content: "1)Barbell Overhead Shoulder Press\n2)Seated Dumbbell Shoulder Press\n3)Front Raise\n4)Reverse Pec Deck Fly\n5)Bent-Over Dumbbell Lateral Raise",
	},
}
