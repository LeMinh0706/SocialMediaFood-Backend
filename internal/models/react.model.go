package models

type ReactResponse struct {
	ID      int64          `json:"id"`
	Account AccountForPost `json:"account"`
}
