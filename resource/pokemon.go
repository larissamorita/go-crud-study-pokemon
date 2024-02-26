package resource

type ResponsePKM struct {
	PkmEntries []DataPKM `json:"pokemon_entries"`
}
type DataPKM struct {
	EntryNumber int        `json:"entry_number"`
	PkmSpecies  PkmSpecies `json:"pokemon_species"`
}
type PkmSpecies struct {
	PkmName string `json:"name"`
	PkmURL  string `json:"url"`
}

//adicao de novo GET do conteudo de color

type URLcolor struct {
	Color URLcolorName `json:"color"`
}

type URLcolorName struct {
	ColorName string `json:"name"`
}

// Formato Output
type Output struct {
	Data []DataPkm2 `json:"data"`
}
type DataPkm2 struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}
