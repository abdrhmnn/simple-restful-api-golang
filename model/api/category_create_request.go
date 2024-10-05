package api

type CategoryCreateRequest struct {
	Name string `validate:"required,max=200,mix=1"`
}
