package swagger

type SwaggerAccessResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type SwaggerValidateData struct {
	Err    string   `json:"error"`
	Fields []string `json:"fields"`
}
type SwaggerNewError struct {
	Err string `json:"error"`
}
type SwaggerRefreshRequest struct {
	Refresh string `json:"refresh" validate:"required"`
}
