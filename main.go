package main

import (
	"log"
	"os"
	"fmt"
	"flag"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	// Define subcommands
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	saveWeightCmd := flag.NewFlagSet("saveWeight", flag.ExitOnError)
	saveProfileCmd := flag.NewFlagSet("saveProfile", flag.ExitOnError)
	saveGoalsCmd := flag.NewFlagSet("saveGoals", flag.ExitOnError)
	removeWeightCmd := flag.NewFlagSet("removeWeight", flag.ExitOnError)

	// Register command flags
	registerUsername := registerCmd.String("username", "", "Username to register")

	// Save weight command flags
	saveWeightUsername := saveWeightCmd.String("username", "", "Username")
	saveWeightDate := saveWeightCmd.String("date", "", "Date of the weight entry")
	saveWeightKg := saveWeightCmd.Float64("weight", 0, "Weight in kg")
	saveWeightNotes := saveWeightCmd.String("notes", "", "Notes for the weight entry")

	// Remove Weight command flags
	removeWeightUsername := removeWeightCmd.String("username", "", "Username")
	removeWeightDate := removeWeightCmd.String("date", "", "date")

	// Save profile command flags
	saveProfileUsername := saveProfileCmd.String("username", "", "Username")
	saveProfileName := saveProfileCmd.String("name", "", "Name")
	saveProfileAge := saveProfileCmd.Int("age", 0, "Age")
	saveProfileGender := saveProfileCmd.String("gender", "", "Gender")
	saveProfileHeight := saveProfileCmd.Int("height", 0, "Height in cm")

	// Save goals command flags
	saveGoalsUsername := saveGoalsCmd.String("username", "", "Username")
	saveGoalsTargetWeight := saveGoalsCmd.Float64("targetWeight", 0, "Target weight in kg")
	saveGoalsTargetDate := saveGoalsCmd.String("targetDate", "", "Target date")

	// Check the first argument
	if len(os.Args) < 2 {
		fmt.Println("Expected 'register', 'saveWeight', 'saveProfile' or 'saveGoals' subcommands")
		os.Exit(1)
	}

	// Switch on the first argument
	switch os.Args[1] {
	case "register":
		registerCmd.Parse(os.Args[2:])
		if *registerUsername == "" {
			fmt.Println("Username is required for registration")
			os.Exit(1)
		}
		data, err := loadData("./data.json")
		checkError(err)
		err = data.registerUser(*registerUsername)
		checkError(err)
		err = saveData("./data.json", data)
		checkError(err)
		fmt.Println("User registered successfully")
	case "saveWeight":
		saveWeightCmd.Parse(os.Args[2:])
		data, err := loadData("./data.json")
		checkError(err)
		user, index := data.findUser(*saveWeightUsername)
		if user == nil {
			fmt.Println("User not found")
			os.Exit(1)
		}
		err = user.saveWeightEntry(*saveWeightKg, *saveWeightNotes, *saveWeightDate)
		checkError(err)
		data.Users[index] = *user
		err = saveData("./data.json", data)
		checkError(err)
		fmt.Println("Weight entry saved successfully")
	case "removeWeight":
		saveWeightCmd.Parse(os.Args[2:])
		data, err := loadData("./data.json")
		checkError(err)
		user, index := data.findUser(*removeWeightUsername)
		if user == nil {
			fmt.Println("User not found")
			os.Exit(1)
		}
		err = user.removeWeightEntry(*removeWeightDate)
		checkError(err)
		data.Users[index] = *user
		err = saveData("./data.json", data)
		checkError(err)
		fmt.Println("Weight entry saved successfully")
	case "saveProfile":
		saveProfileCmd.Parse(os.Args[2:])
		data, err := loadData("./data.json")
		checkError(err)
		user, index := data.findUser(*saveProfileUsername)
		if user == nil {
			fmt.Println("User not found")
			os.Exit(1)
		}
		err = user.saveProfile(*saveProfileName, *saveProfileAge, *saveProfileGender, *saveProfileHeight)
		checkError(err)

		data.Users[index] = *user
		err = saveData("./data.json", data)
		checkError(err)
		fmt.Println("Profile saved successfully")
	case "saveGoals":
		saveGoalsCmd.Parse(os.Args[2:])
		data, err := loadData("./data.json")
		checkError(err)
		user, index := data.findUser(*saveGoalsUsername)
		if user == nil {
			fmt.Println("User not found")
			os.Exit(1)
		}
		err = user.saveGoals(*saveGoalsTargetWeight, *saveGoalsTargetDate)
		checkError(err)
		data.Users[index] = *user
		err = saveData("./data.json", data)
		checkError(err)
		fmt.Println("Goals saved successfully")
	default:
		fmt.Println("Expected 'register', 'saveWeight', 'saveProfile' or 'saveGoals' subcommands")
		os.Exit(1)
	}
}