package models

type GetFriendsResponse struct {
	Friends []Friend `json:"friends"`
	Error   Error    `json:"error"`
}
