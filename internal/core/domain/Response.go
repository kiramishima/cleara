package domain

type Response[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data"`
}
