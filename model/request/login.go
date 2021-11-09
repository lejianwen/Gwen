package request

type Login struct {
	Username string `json:"username" validate:"required,gte=5,lte=10" label:"用户名"`
	Password string `json:"password,omitempty" validate:"required,gte=5,lte=10"`
}
