package models
import "time"

type URL struct {
	ID        		int    	`json:"id"`
	Code 	  		string  `json:"code"`
	TargetUrl 		string  `json:"target_url"`
	RedirectType 	string  `json:"redirect_type"`
	Clicks    		int       `json:"clicks"`
	CreatedAt 		time.Time `json:"created_at"`
	ExpireAt  		time.Time `json:"expire_at"`
}