package api

type UserCreateRequest struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}
