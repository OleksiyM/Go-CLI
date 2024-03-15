package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	quitCommand      = "quit"
	coffeeSimulation = "coffee"
	capitalsQuiz     = "capitals"
	dateCommand      = "date"
)

var countryCapitals = map[string]string{
	"France":                       "Paris",
	"Germany":                      "Berlin",
	"Italy":                        "Rome",
	"Spain":                        "Madrid",
	"United Kingdom":               "London",
	"India":                        "New Delhi",
	"Japan":                        "Tokyo",
	"Brazil":                       "Brasília",
	"Nigeria":                      "Abuja",
	"Ghana":                        "Accra",
	"Ethiopia":                     "Addis Ababa",
	"Algeria":                      "Algiers",
	"Jordan":                       "Amman",
	"Netherlands":                  "Amsterdam", // Official seat of government
	"Andorra":                      "Andorra la Vella",
	"Turkey":                       "Ankara",
	"Madagascar":                   "Antananarivo",
	"Samoa":                        "Apia",
	"Turkmenistan":                 "Ashgabat",
	"Eritrea":                      "Asmara",
	"Kazakhstan":                   "Astana", // Renamed back to Astana in 2022
	"Paraguay":                     "Asunción",
	"Greece":                       "Athens",
	"Cook Islands":                 "Avarua", // Self-governing in association with New Zealand
	"Iraq":                         "Baghdad",
	"Azerbaijan":                   "Baku",
	"Mali":                         "Bamako",
	"Brunei":                       "Bandar Seri Begawan",
	"Thailand":                     "Bangkok",
	"Central African Republic":     "Bangui",
	"Gambia":                       "Banjul",
	"Saint Kitts and Nevis":        "Basseterre",
	"China":                        "Beijing",
	"Lebanon":                      "Beirut",
	"Serbia":                       "Belgrade",
	"Belize":                       "Belmopan", // Replaced Belize City as capital in 1970
	"Germany (West)":               "Bonn",     // Former capital of West Germany
	"Switzerland":                  "Bern",     // De facto capital (seat of government)
	"Kyrgyzstan":                   "Bishkek",
	"Guinea-Bissau":                "Bissau",
	"South Africa":                 "Pretoria", // Executive capital
	"Colombia":                     "Bogotá",
	"Montserrat":                   "Brades", // De facto capital after volcano eruption
	"Venezuela":                    "Caracas",
	"Saint Lucia":                  "Castries",
	"United States Virgin Islands": "Charlotte Amalie",
	"Moldova":                      "Chișinău",
	"Turks and Caicos Islands":     "Cockburn Town",
	"Sri Lanka":                    "Sri Jayawardenepura Kotte", // Official capital
	// "Colombo"       // Commercial capital (de facto)
	"Guinea":  "Conakry",
	"Denmark": "Copenhagen",
	"Benin":   "Cotonou", // De facto capital
	// "Porto-Novo"    // Official capital
	"Senegal":    "Dakar",
	"Syria":      "Damascus",
	"Bangladesh": "Dhaka",
	"East Timor": "Dili",
	"Djibouti":   "Djibouti",
	"Tanzania":   "Dodoma", // Became capital in 1996
	// "Dar es Salaam"  // Former capital, still has some government functions
	"Qatar":            "Doha",
	"Isle of Man":      "Douglas",
	"Ireland":          "Dublin",
	"Tajikistan":       "Dushanbe",
	"Christmas Island": "Flying Fish Cove",
	"Sierra Leone":     "Freetown",
	"Tuvalu":           "Funafuti",
	"Botswana":         "Gaborone",
	"Guyana":           "Georgetown",
}

func greetUser() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What's your name? ")
	scanned := scanner.Scan()
	if !scanned {
		fmt.Println("Error reading input:", scanner.Err())
		return ""
	}
	name := strings.TrimSpace(scanner.Text())
	fmt.Printf("Hello, %s!\n", name)
	return name
}

func getTodayInfo() (string, string, string) {
	now := time.Now()
	dayOfWeek := now.Weekday().String()
	date := now.Format("2006-01-02")
	currTime := now.Format("15:04:05")
	return dayOfWeek, date, currTime
}

func brewCoffee(sugar int, milk string, strength int) {
	fmt.Println("Starting coffee machine...")
	fmt.Println("Grinding coffee beans...")
	time.Sleep(1 * time.Second)

	fmt.Println("Adding water...")
	time.Sleep(1 * time.Second)

	var sugarText string
	switch sugar {
	case 0:
		sugarText = "without sugar"
	case 1:
		sugarText = "with 1 sugar"
	case 2:
		sugarText = "with 2 sugars"
	case 3:
		sugarText = "with 3 sugars"
	default:
		sugarText = "with some amount of sugar"
	}

	var strengthText string
	switch strength {
	case 1:
		strengthText = "weak"
	case 2:
		strengthText = "medium"
	case 3:
		strengthText = "strong"
	}

	var milkText string
	if milk == "no" {
		milkText = "black without milk"
	} else {
		milkText = fmt.Sprintf("and %s milk", milk)
	}

	fmt.Printf("Preparing your %s coffee %s %s...\n", strengthText, sugarText, milkText)
	time.Sleep(3 * time.Second)
	fmt.Println("Your coffee is ready! Enjoy!")
}

func askCoffeeOptions() (int, string, int, bool) {
	sugar := 0
	milk := "no"
	strength := 0 // variable to store coffee strength

	fmt.Println("How many sugars would you like? (0, 1, 2 or 3)")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		fmt.Println("Error reading input:", scanner.Err())
		return 0, "", 0, false
	}
	sugarStr := scanner.Text()
	var err error
	sugar, err = strconv.Atoi(sugarStr)
	if err != nil || (sugar != 0 && sugar != 1 && sugar != 2 && sugar != 3) {
		fmt.Println("Invalid input. Please enter 0, 1, 2 or 3 for sugar.")
		return 0, "", 0, false
	}

	fmt.Println("Would you like milk? (no, regular, kokos)")
	scanned = scanner.Scan()
	if !scanned {
		fmt.Println("Error reading input:", scanner.Err())
		return 0, "", 0, false
	}
	milk = strings.ToLower(scanner.Text())
	if milk != "no" && milk != "regular" && milk != "kokos" {
		fmt.Println("Invalid input. Please enter 'no', 'regular', or 'kokos' for milk.")
		return 0, "", 0, false
	}

	// Get coffee strength
	fmt.Print("Enter coffee strength (1 = weak, 2 = medium, 3 = strong): ")
	scanned = scanner.Scan()
	if !scanned {
		fmt.Println("Error reading input:", scanner.Err())
		return 0, "", 0, false
	}
	strengthStr := scanner.Text()
	var errStrength error
	strength, errStrength = strconv.Atoi(strengthStr)
	if errStrength != nil || (strength != 1 && strength != 2 && strength != 3) {
		fmt.Println("Invalid input. Please enter 1 (weak), 2 (medium), or 3 (strong) for coffee strength.")
		return 0, "", 0, false
	}

	var strengthText string
	switch strength {
	case 1:
		strengthText = "weak"
	case 2:
		strengthText = "medium"
	case 3:
		strengthText = "strong"
	}

	fmt.Println("Summary:")
	fmt.Printf("  Sugar: %d\n", sugar)
	fmt.Printf("  Milk: %s\n", milk)
	fmt.Printf("  Strength: %s\n", strengthText)
	fmt.Println("Is this correct? (yes/no)")
	scanned = scanner.Scan()
	if !scanned {
		fmt.Println("Error reading input:", scanner.Err())
		return 0, "", 0, false
	}
	confirmation := strings.ToLower(scanner.Text())
	if confirmation != "yes" {
		fmt.Println("Alright, starting over.")
		return 0, "", 0, false
	}

	return sugar, milk, strength, true
}

func capitalsQuizF() {
	// Prompt user for number of questions (default 10)
	var numQuestions int
	for {
		fmt.Println("How many countries would you like to be tested on? (default 10)")
		fmt.Scanf("%d\n", &numQuestions)

		// Handle empty input (pressing Enter without typing a number):
		if numQuestions == 0 {
			numQuestions = 10 // Set default value
			break             // Exit the loop immediately to proceed with the quiz
		} else if numQuestions < 1 || numQuestions > len(countryCapitals) {
			// Handle invalid input (non-numeric or out-of-range):
			fmt.Printf("Invalid input. Please enter a number between 1 and %d.\n", len(countryCapitals))
			numQuestions = 0 // Reset numQuestions to handle empty input correctly
		} else {
			// Handle valid input:
			break // Exit the loop to proceed with the quiz
		}
	}

	score := 0
	questionsAsked := 0

	// Shuffle the country keys
	countries := make([]string, 0, len(countryCapitals))
	for country := range countryCapitals {
		countries = append(countries, country)
	}
	for i := range countries {
		j := rand.Intn(i + 1)
		countries[i], countries[j] = countries[j], countries[i]
	}

	// Ask questions for chosen number of countries (without repetition)
	usedCountries := make(map[string]bool)
	for questionsAsked < numQuestions { // Corrected indentation here
		country := countries[rand.Intn(len(countries))]
		if !usedCountries[country] {
			usedCountries[country] = true
			capital := countryCapitals[country]

			questionsAsked++

			fmt.Printf("Question %d from %d: What is the capital of %s? ", questionsAsked, numQuestions, country) // Corrected indentation here
			reader := bufio.NewReader(os.Stdin)
			answer, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			answer = strings.TrimSpace(answer)

			if strings.EqualFold(answer, capital) {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Printf("Incorrect. The capital of %s is %s.\n", country, capital)
			}
		}
	}

	fmt.Printf("You answered %d out of %d questions correctly.\n", score, questionsAsked)
}

func displayCommands() {
	fmt.Println(" Available commands:")
	fmt.Println("--------------------")
	fmt.Println("- quit: Close the application.")
	fmt.Println("- coffee: Simulate brewing coffee (weak, medium, or strong).")
	fmt.Println("- capitals: Test your knowledge of world capitals with a quiz.")
	fmt.Println("- date: Display the current date, time, and day of the week.")
	fmt.Println("") // Add some blank lines for visual separation
}

func main() {
	name := greetUser()

	displayCommands()

	for {
		fmt.Print("Enter a command (quit, coffee, capitals, date): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanned := scanner.Scan()
		if !scanned {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		command := strings.ToLower(scanner.Text())

		switch command {
		case quitCommand:
			fmt.Println("Goodbye, ", name, "!")
			return

		case coffeeSimulation:
			sugar, milk, strength, confirmed := askCoffeeOptions()
			if !confirmed {
				continue
			}

			// TODO: Use the retrieved strength to potentially alter coffee preparation logic here

			brewCoffee(sugar, milk, strength)

		case capitalsQuiz:
			capitalsQuizF()

		case dateCommand:
			var dayOfWeek, date, currTime = getTodayInfo() // Now the function is called!
			fmt.Printf("Today is %s, %s, %s\n", dayOfWeek, date, currTime)

		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}
