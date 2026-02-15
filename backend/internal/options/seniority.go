package options

type Seniority struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

func GetSeniorities() []Seniority {
	return []Seniority{
		{ID: 1, Data: "Junior"},
		{ID: 2, Data: "Mid-Level"},
		{ID: 3, Data: "Senior"},
		{ID: 4, Data: "Expert"},
	}
}
