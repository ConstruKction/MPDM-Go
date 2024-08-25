package main

type Song struct {
	PvID			string
	Name			string
	LineNumberEN	int
	LineNumberJP	int
	DifficultyEX	float64
	DifficultyExtra	float64
	ModPVDBPath		string
	State			int
	Pack			string
}

func NewSong(
	pvID string,
	name string,
	lineNumberEN int,
	lineNumberJP int,
	difficultyEX float64,
	difficultyExtra float64,
	modPVDBPath string,
	state int,
	pack string,
) *Song {
	return &Song{
		PvID:			 pvID,
		Name:			 name,
		LineNumberEN:	 lineNumberEN,
		LineNumberJP:	 lineNumberJP,
		DifficultyEX:	 difficultyEX,
		DifficultyExtra: difficultyExtra,
		ModPVDBPath:	 modPVDBPath,
		State:			 state,
		Pack:			 pack,
	}
}