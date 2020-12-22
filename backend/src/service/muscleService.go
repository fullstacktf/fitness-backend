package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateMuscle Create a new muscle with specified data in body
func CreateMuscle(data model.Muscle) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetMuscle Get muscle by id
func GetMuscle(id int) *model.Muscle {
	muscle := model.Muscle{}
	storage.DB.Find(&muscle, id)

	return &muscle
}

// GetMuscles Get all non deleted muscles and by using a filter
func GetMuscles() *[]model.Muscle {
	muscles := []model.Muscle{}
	storage.DB.Find(&muscles)

	return &muscles
}

// UpdateMuscle Update specific muscle using id param in URL
func UpdateMuscle(updatedMuscle model.Muscle) error {

	output := storage.DB.Save(&updatedMuscle)

	return output.Error
}

// DeleteMuscle Delete muscle by id, logical delete
func DeleteMuscle(id int) {
	deletedMuscle := GetMuscle(id)

	storage.DB.Delete(&deletedMuscle)

}

// PopulateMuscles populates muscles element with basic data
func PopulateMuscles() {

	pectoral := model.Muscle{
		Name: "Pectoral",
	}

	dorsal := model.Muscle{
		Name: "Dorsal",
	}

	quadriceps := model.Muscle{
		Name: "Quadriceps",
	}

	biceps := model.Muscle{
		Name: "Bíceps",
	}

	triceps := model.Muscle{
		Name: "Tríceps",
	}

	shoulders := model.Muscle{
		Name: "Shoulders",
	}

	hamstrings := model.Muscle{
		Name: "Hamstrings",
	}

	buttocks := model.Muscle{
		Name: "Buttocks",
	}

	abs := model.Muscle{
		Name: "ABS",
	}

	neck := model.Muscle{
		Name: "Neck",
	}

	calf := model.Muscle{
		Name: "Calf",
	}

	forearm := model.Muscle{
		Name: "Forearm",
	}

	lumbar := model.Muscle{
		Name: "Lumbar",
	}

	calves := model.Muscle{
		Name: "Calves",
	}

	scapular := model.Muscle{
		Name: "Scapular Muscles",
	}

	muscles := []model.Muscle{
		pectoral,
		dorsal,
		quadriceps,
		biceps,
		triceps,
		shoulders,
		hamstrings,
		buttocks,
		abs,
		neck,
		calf,
		forearm,
		lumbar,
		calves,
		scapular,
	}

	storage.DB.Create(&muscles)
}
