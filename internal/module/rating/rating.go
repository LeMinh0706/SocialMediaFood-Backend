package rating

type RatingRequest struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Star          int32  `json:"star"`
	Content       string `json:"content"`
}
