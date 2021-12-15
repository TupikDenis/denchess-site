package tournaments

import (
	"bufio"
	"log"
	"os"

	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
	"github.com/jinzhu/gorm"
)

func AddTournament(url string, name string, year uint) {
	file, err := os.Open(url)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	jsonContent := "["

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		jsonContent += fileScanner.Text() + ","
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	if last := len(jsonContent) - 1; last >= 0 && jsonContent[last] == ',' {
		jsonContent = jsonContent[:last]
	}
	jsonContent += "]"

	log.Println(jsonContent)

	file.Close()
	tournament := models.Tournament{
		Model: gorm.Model{},
		Year:  year,
		Round: name,
		JSON:  jsonContent,
	}

	db := handlers.Database()
	db.AutoMigrate(&models.Tournament{})
	db.Create(&tournament)
	db.Close()
}

func checkUsername(username string) bool {
	var users []models.Total

	db := handlers.Database()
	db.Find(&users)
	var flag = true

	for _, item := range users {
		if item.Username == username {
			flag = false
		}
	}

	return flag
}

func updateUsername(total models.Total, id uint) {
	var username models.Total

	db := handlers.Database()
	db.First(&username, id)

	username.TotalFirstPlace += total.TotalFirstPlace
	username.TotalSecondtPlace += total.TotalSecondtPlace
	username.TotalThirdPlace += total.TotalThirdPlace
	username.TotalPoints += total.TotalPoints
	username.TotalScore += total.TotalScore

	db.Save(&username)

	db.Close()
}

func getUserNameId(username string) uint {
	var user []models.Total
	var id uint

	db := handlers.Database()
	db.Find(&user)
	db.Close()

	for _, item := range user {
		if item.Username == username {
			id = item.ID
		}
	}

	return id
}

func AddTotal(year uint, username string, first int, second int, third int, score int, totalScore int) {
	total := models.Total{
		Year:              year,
		Username:          username,
		TotalFirstPlace:   first,
		TotalSecondtPlace: second,
		TotalThirdPlace:   third,
		TotalPoints:       score,
		TotalScore:        totalScore,
	}

	flag := checkUsername(total.Username)
	if flag == true {
		db := handlers.Database()
		db.AutoMigrate(&models.Total{})
		db.Create(&total)
		db.Close()
	} else {
		id := getUserNameId(total.Username)
		updateUsername(total, id)
	}

}
