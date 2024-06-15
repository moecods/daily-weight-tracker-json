package main

type Goals struct {
	TargetWeightKg float64 `json:"target_weight_kg"`
	TargetDate     string  `json:"target_date"`
}

func (user *User) saveGoals(targetWeightKg float64, targetDate string) error {
	user.Goals = Goals{
		TargetWeightKg: targetWeightKg,
		TargetDate:     targetDate,
	}

	return nil
}