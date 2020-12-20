package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateBaseRoutine Creates a base routine with the given data
func CreateBaseRoutine(data model.BaseRoutine) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetBaseRoutine Get base routine by id
func GetBaseRoutine(id int) *model.BaseRoutine {
	baseRoutine := model.BaseRoutine{}
	storage.DB.Find(&baseRoutine, id)

	baseRoutine.GetBaseRoutineAssociations()

	return &baseRoutine
}

// GetBaseRoutines Get all non deleted base routines and by using a filter
func GetBaseRoutines(filter model.BaseRoutine) *[]model.BaseRoutine {
	baseRoutines := []model.BaseRoutine{}
	storage.DB.Where(&filter).Find(&baseRoutines)

	for i := 0; i < len(baseRoutines); i++ {
		baseRoutines[i].GetBaseRoutineAssociations()
	}

	return &baseRoutines
}

// UpdateBaseRoutine Update specific base routine using id param in URL
func UpdateBaseRoutine(updatedBaseRoutine model.BaseRoutine) error {

	storage.DB.Model(&updatedBaseRoutine).Association("BaseExercises").Replace(updatedBaseRoutine.BaseExercises)

	output := storage.DB.Save(&updatedBaseRoutine)

	return output.Error
}

// DeleteBaseRoutine Delete BaseRoutine by id, logical delete
func DeleteBaseRoutine(id int) {
	deletedBaseRoutine := GetBaseRoutine(id)

	storage.DB.Delete(&deletedBaseRoutine)

}

// PopulateBaseRoutine populates base routine element with basic data
func PopulateBaseRoutine() {

	phulExercises := []*model.BaseExercise{
		GetBaseExercise(2),
		GetBaseExercise(4),
		GetBaseExercise(7),
		GetBaseExercise(9),
		GetBaseExercise(8),
		GetBaseExercise(10),
		GetBaseExercise(13),
	}

	phul := model.BaseRoutine{
		Name:          "Power Hypertrophy Upper Lower",
		Description:   "The PHUL workout is based around the basic principles of strength and size.  This program will allow you to maximize results on both fronts in an easy adaptable routine.",
		BaseExercises: phulExercises,
		CategoryID:    1,
	}

	strengthGenericExercises := []*model.BaseExercise{
		GetBaseExercise(3),
		GetBaseExercise(11),
		GetBaseExercise(2),
		GetBaseExercise(12),
		GetBaseExercise(10),
		GetBaseExercise(13),
		GetBaseExercise(1),
	}

	strengthGeneric := model.BaseRoutine{
		Name:          "Strength Generic",
		Description:   "We know daily exercise is good for optimizing health. But with so many options and limitless information available, it’s easy to get overwhelmed with what works. But not to worry. We’ve got your back (and body)! Check out the 10 exercises you can do for ultimate fitness. Combine them into a routine for a workout that’s simple but powerful and sure to keep you in shape for the rest of your life.",
		BaseExercises: strengthGenericExercises,
		CategoryID:    1,
	}

	weightLossGenericExercises := []*model.BaseExercise{
		GetBaseExercise(14),
		GetBaseExercise(15),
		GetBaseExercise(16),
		GetBaseExercise(6),
		GetBaseExercise(7),
		GetBaseExercise(11),
	}

	weightLossGeneric := model.BaseRoutine{
		Name:          "Weight Loss",
		Description:   "If you're trying to lose weight, a weight loss workout plan can be very helpful. Getting regular exercise can help you meet your goals in a healthy, sustainable way.",
		BaseExercises: weightLossGenericExercises,
		CategoryID:    3,
	}

	maintenanceExercises := []*model.BaseExercise{
		GetBaseExercise(1),
		GetBaseExercise(15),
		GetBaseExercise(13),
		GetBaseExercise(11),
		GetBaseExercise(5),
		GetBaseExercise(2),
		GetBaseExercise(3),
	}

	maintenance := model.BaseRoutine{
		Name:          "Maintenance",
		Description:   "This program is for anyone who wants to stay in shape, without having a primary interest in bodybuilding competitions or a specific sport of competition, only to be in shape and with a pleasant body. In addition, this routine affects the entire body and is planned to work all the muscles of the body, developing both strength and muscular endurance.",
		BaseExercises: maintenanceExercises,
		CategoryID:    4,
	}

	upperSideStrengthExercises := []*model.BaseExercise{
		GetBaseExercise(2),
		GetBaseExercise(3),
		GetBaseExercise(5),
		GetBaseExercise(6),
		GetBaseExercise(10),
		GetBaseExercise(11),
		GetBaseExercise(12),
		GetBaseExercise(13),
	}
	lowerSideStrengthExercises := []*model.BaseExercise{
		GetBaseExercise(2),
		GetBaseExercise(3),
		GetBaseExercise(8),
		GetBaseExercise(9),
		GetBaseExercise(4),
	}

	upperSideStrength := model.BaseRoutine{
		Name:          "Upper Side Strength",
		Description:   "The way this program was designed will allow you to add significant muscle mass to your torso and shed excess body fat. The result? A new body.",
		BaseExercises: upperSideStrengthExercises,
		CategoryID:    1,
	}

	lowerSideStrength := model.BaseRoutine{
		Name:          "Lower Side Strength",
		Description:   "Do you ever find it difficult to bend forward? Have you noticed yourself rubbing your achy, lower back lately? How about wishing you could lift your little one, a suitcase, or even a bag of groceries more easily? If you answered “yes” to any of the above, you might need to work on your core and lower back strength.",
		BaseExercises: lowerSideStrengthExercises,
		CategoryID:    1,
	}

	routines := []model.BaseRoutine{
		phul,
		strengthGeneric,
		weightLossGeneric,
		maintenance,
		upperSideStrength,
		lowerSideStrength,
	}

	storage.DB.Create(&routines)

}
