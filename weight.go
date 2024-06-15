package main

type WeightEntry struct {
	Date     string  `json:"date"`
	WeightKg float64 `json:"weight_kg"`
	Notes    string  `json:"notes,omitempty"`
}


func (user *User) saveWeightEntry(weight float64, notes string, date string) error {
	weightData := WeightEntry{
		Date:     date,
		WeightKg: weight,
		Notes:    notes,
	}

	updated := false
	for i, entry := range user.Weights {
		if entry.Date == date {
			user.Weights[i] = weightData
			updated = true
			break
		}
	}

	if !updated {
		user.Weights = append(user.Weights, weightData)
	}

	return nil
}


func (user *User) removeWeightEntry(date string) error {
	for index, entry := range user.Weights {
		if entry.Date == date {
			user.Weights = append(user.Weights[:index], user.Weights[index+1:]...)
		}
	}

	return nil
}