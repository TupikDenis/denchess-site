package tournaments

import (
	"github.com/TupikDenis/denchess-site.git/pkg/handlers"
	"github.com/TupikDenis/denchess-site.git/pkg/models"
)

func GetPlayersFromTournament(year uint) []models.TransformedTotal {
	var tournaments []models.Total
	var _tournaments []models.TransformedTotal

	db := handlers.Database()
	db.Order("total_score desc").Find(&tournaments)
	for _, item := range tournaments {
		if item.Year == year {
			_tournaments = append(
				_tournaments, models.TransformedTotal{
					Id:                item.ID,
					Rank:              0,
					Year:              item.Year,
					Username:          item.Username,
					TotalFirstPlace:   item.TotalFirstPlace,
					TotalSecondtPlace: item.TotalSecondtPlace,
					TotalThirdPlace:   item.TotalThirdPlace,
					TotalPoints:       item.TotalPoints,
					TotalScore:        item.TotalScore,
				})
		}
	}

	return _tournaments
}

func GetAllTournaments(year uint) []models.TransformedTournament {
	var tournaments []models.Tournament
	var _tournaments []models.TransformedTournament

	db := handlers.Database()
	db.Find(&tournaments)
	for _, item := range tournaments {
		if item.Year == year {
			_tournaments = append(
				_tournaments, models.TransformedTournament{
					Id:    item.ID,
					Year:  item.Year,
					Round: item.Round,
					JSON:  item.JSON,
				})
		}
	}

	return _tournaments
}
