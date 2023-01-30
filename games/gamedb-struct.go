package games

type GamesDB []GamesDBElement

type GamesDBElement struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Developer string `json:"developer"`
	//Category    *Category `json:"category"`
	ReleaseYear *int64 `json:"releaseYear"`
}

type Category string

const (
	Action_Category      = "Action"
	Adult_Category       = "Adult"
	Adventure_Category   = "Adventure"
	Educational_Category = "Educational"
	Fighting_Category    = "Fighting"
	Puzzle_Category      = "Puzzle"
	Racing_Category      = "Racing"
	RolePlaying_Category = "Role-Playing"
	Shooter_Category     = "Shooter"
	Simulation_Category  = "Simulation"
	Sports_Category      = "Sports"
	Strategy_Category    = "Strategy"
	Traditional_Category = "Traditional"
	Various_Category     = "Various"
)
