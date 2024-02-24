package model

type Response struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Data       []Data `json:"data"`
}

type Data struct {
	City          string     `json:"city"`
	Name          string     `json:"name"`
	EstimatedCost int        `json:"estimated_cost"`
	UserRating    UserRating `json:"user_rating"`
	Id            int        `json:"id"`
}

type UserRating struct {
	AverageRating float64 `json:"average_rating"`
	Votes         int     `json:"votes"`
}
