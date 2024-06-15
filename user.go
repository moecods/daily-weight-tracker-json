package main

type User struct {
	Username string        `json:"username"`
	Profile  Profile       `json:"profile"`
	Weights  []WeightEntry `json:"weights"`
	Goals    Goals         `json:"goals"`
}

func (d *Data) findUser(username string) (*User, int) {
	for index, user := range d.Users {
		if user.Username == username {
			return &user, index
		}
	}

	return nil, -1
}

func (data *Data) registerUser(username string) error {
	newUser := User {
		Username: username,
	}

	data.Users = append(data.Users, newUser)

	return nil
}