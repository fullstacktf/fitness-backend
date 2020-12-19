package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateBaseExercise Creates a BaseExercise with the given data
func CreateBaseExercise(data model.BaseExercise) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetBaseExercise Get BaseExercise by id
func GetBaseExercise(id int) *model.BaseExercise {
	baseEx := model.BaseExercise{}
	storage.DB.Find(&baseEx, id)

	return &baseEx
}

// GetBaseExercises Get all non deleted BaseExercises and by using a filter
func GetBaseExercises(filter model.BaseExercise) *[]model.BaseExercise {
	baseExs := []model.BaseExercise{}
	storage.DB.Where(&filter).Find(&baseExs)

	for i := 0; i < len(baseExs); i++ {
		baseExs[i].ExerciseCategory = *GetExerciseCategory(int(baseExs[i].CategoryID))
	}

	return &baseExs
}

// UpdateBaseExercise Update specific BaseExercise using id param in URL
func UpdateBaseExercise(updatedBaseExercise model.BaseExercise) error {

	output := storage.DB.Save(&updatedBaseExercise)

	return output.Error
}

// DeleteBaseExercise Delete BaseExercise by id, logical delete
func DeleteBaseExercise(id int) {
	deletedBaseExercise := GetBaseExercise(id)

	storage.DB.Delete(&deletedBaseExercise)

}

// PopulateBaseExercise populates base exercise element with basic data
func PopulateBaseExercise() {

	muscles := make([]*model.Muscle, 1)
	muscles[0] = GetMuscle(9)

	fuerza := model.BaseExercise{
		Name:               "Plancha",
		Description:        "Ejercicio clásico de plancha",
		VideoURL:           "https://www.youtube.com/watch?v=2wRv2J3sufM",
		DefaultSeries:      1,
		DefaultRepetitions: 1,
		CategoryID:         1,
		Muscles:            muscles,
	}

	storage.DB.Create(&fuerza)

}
