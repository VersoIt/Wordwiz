package worker

const cronMonthlyReset = "0 0 1 * *"

type Worker struct {
	userRepo userRepo
}

func New(userRepo userRepo) *Worker {
	return &Worker{userRepo: userRepo}
}
