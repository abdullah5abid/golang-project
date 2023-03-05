package response

type WordClass struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Words       []string `json:"words"`
}
