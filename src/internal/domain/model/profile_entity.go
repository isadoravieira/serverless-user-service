package model

type Profile struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	Bio               string `json:"bio"`
	Birthdate         string `json:"birthday"`
	Gender            string `json:"gender"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	Location          string `json:"location"`
	CreatedAt         string `json:"createdAt"`
	UpdatedAt         string `json:"updatedAt"`
}
