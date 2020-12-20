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

	baseEx.GetBaseExerciseAssociations()

	return &baseEx
}

// GetBaseExercises Get all non deleted BaseExercises and by using a filter
func GetBaseExercises(filter model.BaseExercise) *[]model.BaseExercise {
	baseExs := []model.BaseExercise{}
	storage.DB.Where(&filter).Find(&baseExs)

	for i := 0; i < len(baseExs); i++ {
		baseExs[i].GetBaseExerciseAssociations()
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

	kettlebellGobletSquatMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(9),
	}

	kettlebellGlobletSquat := model.BaseExercise{
		Name:               "Kettlebell Goblet Squat",
		Description:        "The Kettlebell Goblet Squat is a great way to develop full lower body strength. The positioning of this squat is easy on the back when performed correctly, making it a great squat for advanced and beginner strength trainers alike.",
		VideoURL:           "https://www.youtube.com/watch?v=mF5tnEBrdkc",
		DefaultSeries:      3,
		DefaultRepetitions: 10,
		CategoryID:         1,
		Muscles:            kettlebellGobletSquatMuscles,
	}

	backSquatMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(15),
	}

	backSquat := model.BaseExercise{
		Name:               "Back Squat",
		Description:        "The back squat is the most basic strength exercise in weightlifting, and one of the most commonly used exercises other than the competition lifts.",
		VideoURL:           "https://www.youtube.com/watch?v=ultWZbUMPL8",
		DefaultSeries:      5,
		DefaultRepetitions: 5,
		CategoryID:         1,
		Muscles:            backSquatMuscles,
	}

	bulgarianSplitSquatMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(14),
	}

	bulgarianSplitSquat := model.BaseExercise{
		Name:               "Bulgarian Split Squat",
		Description:        "The Bulgarian split squat is a version of a single-leg squat where the back leg is elevated on a bench or a sturdy chair. As a single-leg, unilateral squat, the exercise places a greater focus on the quadriceps than other, similar lower-body compound movements.",
		VideoURL:           "https://www.youtube.com/watch?v=Sx3qaXu8XTY",
		DefaultSeries:      2,
		DefaultRepetitions: 8,
		CategoryID:         1,
		Muscles:            bulgarianSplitSquatMuscles,
	}

	snatchGripDeadliftMuscles := []*model.Muscle{
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(13),
	}

	snatchGripDeadlift := model.BaseExercise{
		Name:               "Snatch-Grip Deadlift",
		Description:        "A snatch grip deadlift is an advanced variation of the traditional deadlift. The snatch grip is done with a wider grip on the barbell. Some weight lifters prefer a wider snatch grip because it’s more comfortable for the lower back.",
		VideoURL:           "https://www.youtube.com/watch?v=NMDPXmGJI88",
		DefaultSeries:      3,
		DefaultRepetitions: 10,
		CategoryID:         1,
		Muscles:            snatchGripDeadliftMuscles,
	}

	kettlebellSwingMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(13),
	}

	kettlebellSwing := model.BaseExercise{
		Name:               "Kettlebell swing",
		Description:        "Kettlebell swing (AKA Russian swing, double-arm swing, or conventional kettlebell swing) is a basic ballistic exercise used to train the posterior chain in a manner similar to broad jumping. It involves moving the bell in a pendulum motion from between the knees to anywhere between eye level to fully overhead and can be performed either two-handed or using one hand.",
		VideoURL:           "https://www.youtube.com/watch?v=mKDIuUbH94Q",
		DefaultSeries:      3,
		DefaultRepetitions: 10,
		CategoryID:         1,
		Muscles:            kettlebellSwingMuscles,
	}

	kettlebellSumoDeadliftMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(13),
	}

	kettlebellSumoDeadlift := model.BaseExercise{
		Name:               "Kettlebell Sumo Deadlift",
		Description:        "This version of the deadlift is a good way of drilling a movement pattern where you engage your hamstrings and hinge at the hips to perform the move. You’re using a fairly light weight so it won’t fatigue your muscles before you get to the heavy deadlift sets in the next exercise.",
		VideoURL:           "https://www.youtube.com/watch?v=P-fdRaL-5b4",
		DefaultSeries:      3,
		DefaultRepetitions: 10,
		CategoryID:         1,
		Muscles:            kettlebellSumoDeadliftMuscles,
	}

	kettlebellSumoDeadliftHighMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(12),
		GetMuscle(13),
	}

	kettlebellSumoDeadliftHigh := model.BaseExercise{
		Name:               "Kettlebell Sumo Deadlift High Pull",
		Description:        "The sumo deadlift high pull builds on the deadlift, but we widen the stance and bring the grip inside the knees to facilitate a longer pulling motion. We also add velocity to the movement. The sumo deadlift high pull replicates the upward movement pattern of a clean or snatch and serves as a bridge between the deadlift and the faster lifts.",
		VideoURL:           "https://www.youtube.com/watch?v=0xpSOScnr1E",
		DefaultSeries:      3,
		DefaultRepetitions: 10,
		CategoryID:         1,
		Muscles:            kettlebellSumoDeadliftHighMuscles,
	}

	deadliftMuscles := []*model.Muscle{
		GetMuscle(2),
		GetMuscle(3),
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(12),
		GetMuscle(13),
	}

	deadlift := model.BaseExercise{
		Name:               "Deadlift",
		Description:        "The deadlift is arguably the most effective whole-body strength and muscle builder. It also focuses on your posterior chain - the muscles on the back of your body, which often end up undertrained but play a key role in promoting good posture and keeping you injury-free.",
		VideoURL:           "https://www.youtube.com/watch?v=op9kVnSso6Q",
		DefaultSeries:      5,
		DefaultRepetitions: 5,
		CategoryID:         1,
		Muscles:            deadliftMuscles,
	}

	romanianDeadliftMuscles := []*model.Muscle{
		GetMuscle(2),
		GetMuscle(3),
		GetMuscle(7),
		GetMuscle(8),
		GetMuscle(12),
		GetMuscle(13),
	}

	romanianDeadlift := model.BaseExercise{
		Name:               "Romanian Deadlift",
		Description:        "After your heavy deadlift sets this is an ideal way of developing the hamstring strength needed to improve your deadlift. It’s technically a lot easier than the deadlift so you’ll be able to go reasonably heavy even when your muscles are tired.",
		VideoURL:           "https://www.youtube.com/watch?v=iiStbsdh4s0",
		DefaultSeries:      4,
		DefaultRepetitions: 6,
		CategoryID:         1,
		Muscles:            romanianDeadliftMuscles,
	}

	inclineBenchPressMuscles := []*model.Muscle{
		GetMuscle(1),
		GetMuscle(5),
		GetMuscle(6),
	}

	inclineBenchPress := model.BaseExercise{
		Name:               "Incline Bench Press",
		Description:        "The Incline Bench Press is a version of the traditional Bench Press in which the bench is positioned at about a 45-degree angle. The resulting inclined position targets your upper chest and the frontside of your shoulders more the the standard flat bench.",
		VideoURL:           "https://www.youtube.com/watch?v=0G2_XV7slIg",
		DefaultSeries:      3,
		DefaultRepetitions: 10,
		CategoryID:         1,
		Muscles:            inclineBenchPressMuscles,
	}

	pressUpMuscles := []*model.Muscle{
		GetMuscle(1),
		GetMuscle(5),
		GetMuscle(6),
		GetMuscle(9),
	}

	pressUp := model.BaseExercise{
		Name:               "Press Up",
		Description:        "Press-ups are exercises to strengthen your arms and chest muscles. They are done by lying with your face towards the floor and pushing with your hands to raise your body until your arms are straight.",
		VideoURL:           "https://www.youtube.com/watch?v=0pkjOk0EiAk",
		DefaultSeries:      3,
		DefaultRepetitions: 10,
		CategoryID:         1,
		Muscles:            pressUpMuscles,
	}

	dumbbellBenchPressMuscles := []*model.Muscle{
		GetMuscle(1),
		GetMuscle(5),
		GetMuscle(6),
		GetMuscle(12),
	}

	dumbbellBenchPress := model.BaseExercise{
		Name:               "Dumbbell Bench Press",
		Description:        "The dumbbell bench press is a variation of the barbell bench press and an exercise used to build the muscles of the chest. Often times, the dumbbell bench press is recommended after reaching a certain point of strength on the barbell bench press to avoid pec and shoulder injuries.",
		VideoURL:           "https://www.youtube.com/watch?v=X3YrlBmjWrY",
		DefaultSeries:      3,
		DefaultRepetitions: 10,
		CategoryID:         1,
		Muscles:            dumbbellBenchPressMuscles,
	}

	benchPressMuscles := []*model.Muscle{
		GetMuscle(1),
		GetMuscle(5),
		GetMuscle(6),
		GetMuscle(12),
	}

	benchPress := model.BaseExercise{
		Name:               "Bench Press",
		Description:        "The bench press is an upper-body weight training exercise in which the trainee presses a weight upwards while lying on a weight training bench. The bench press is one of three lifts in the sport of powerlifting alongside the deadlift and squat.",
		VideoURL:           "https://www.youtube.com/watch?v=SCVCLChPQFY",
		DefaultSeries:      5,
		DefaultRepetitions: 5,
		CategoryID:         1,
		Muscles:            benchPressMuscles,
	}

	cyclingMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
	}

	cycling := model.BaseExercise{
		Name:               "Cycling",
		Description:        "Cycling is a popular exercise that improves your fitness and can help you lose weight. Although cycling is traditionally done outdoors, many gyms and fitness centers have stationary bikes that allow you to cycle while staying indoors. Harvard Health estimates that a 155-pound (70-kg) person burns around 260 calories per 30 minutes of cycling on a stationary bike at a moderate pace, or 298 calories per 30 minutes on a bicycle at a moderate pace of 12–13.9 mph (19–22.4 km/h)",
		VideoURL:           "https://www.youtube.com/watch?v=4Hl1WAGKjMc",
		DefaultSeries:      1,
		DefaultRepetitions: 20,
		CategoryID:         3,
		Muscles:            cyclingMuscles,
	}

	runningMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
	}

	running := model.BaseExercise{
		Name:               "Running",
		Description:        "Running is a popular exercise that improves your fitness and can help you lose weight. Harvard Health estimates that a 155-pound (70-kg) person burns approximately 372 calories per 30 minutes of running at a 6-mph (9.7-km/h) pace",
		VideoURL:           "https://www.youtube.com/watch?v=kVnyY17VS9Y",
		DefaultSeries:      1,
		DefaultRepetitions: 20,
		CategoryID:         3,
		Muscles:            runningMuscles,
	}

	walkingMuscles := []*model.Muscle{
		GetMuscle(3),
		GetMuscle(7),
	}

	walking := model.BaseExercise{
		Name:               "Walking",
		Description:        "Walking is one of the best exercises for weight loss — and for good reason. It’s convenient and an easy way for beginners to start exercising without feeling overwhelmed or needing to purchase equipment. Also, it’s a lower-impact exercise, meaning it doesn’t stress your joints. According to Harvard Health, it’s estimated that a 155-pound (70-kg) person burns around 167 calories per 30 minutes of walking at a moderate pace of 4 mph (6.4 km/h)",
		VideoURL:           "https://www.youtube.com/watch?v=GHldM-dSUSI",
		DefaultSeries:      1,
		DefaultRepetitions: 60,
		CategoryID:         3,
		Muscles:            walkingMuscles,
	}

	strength := []model.BaseExercise{
		kettlebellGlobletSquat,
		backSquat,
		bulgarianSplitSquat,
		snatchGripDeadlift,
		kettlebellSwing,
		kettlebellSumoDeadlift,
		kettlebellSumoDeadliftHigh,
		deadlift,
		romanianDeadlift,
		inclineBenchPress,
		pressUp,
		dumbbellBenchPress,
		benchPress,
	}

	weightLoss := []model.BaseExercise{
		cycling,
		running,
		walking,
	}

	storage.DB.Create(&strength)
	storage.DB.Create(&weightLoss)
}
