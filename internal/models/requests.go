package models

type ShortenRequest struct {
	URL		string `json:"url" validate:"required,url"`
	ExpireIn int    `json:"expire_in"` //seconds, optional
	RedirectType string `json:"redirect_type"` //optional: 301, 302, 307
}