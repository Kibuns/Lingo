package Models

type Letter struct {
	Char      string `json:"char"`
	IsCorrect bool   `json:"iscorrect"`
	IsPresent bool   `json:"ispresent"`
}

type ParsedInput struct {
	Letters     []Letter `json:"letters"`
	GuessedWord bool     `json:"guessedword"`
}
