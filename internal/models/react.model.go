package models

type ReactResponse struct {
	ID      int64          `json:"id"`
	Account AccountForPost `json:"account"`
}

type ListReactResponse struct {
	React []ReactResponse `json:"react"`
	Total int64           `json:"total"`
}
