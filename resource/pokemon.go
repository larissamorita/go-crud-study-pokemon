package resource

type ResponsePKM struct {
	PkmEntries []DataPkm `json:"pokemon_entries"`
}
type DataPKM struct {
	entryNumber int           `json:"entry_number"`
	pkmSpecies  []DataSpecies `json:"pokemon_species"`
}
type DataSpecies struct {
	pkmName string `json:"name"`
	pkmURL  string `json:"url"`
}
