package main

type Profile struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	HeightCm int    `json:"height_cm"`
}

func (user *User) saveProfile(name string, age int, gender string, heightCm int) error {
	user.Profile = Profile{
		Name:     name,
		Age:      age,
		Gender:   gender,
		HeightCm: heightCm,
	}

	return nil
}

