package models

//User type for rendering the user Model
type User struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	DisplayName  string `json:"display_name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	ProfileImage string `json:"profile_image_url"`
	ViewCount    int64  `json:"view_count"`
	Email        string `json:"email,omitempty"`
}

//UserData to User Struct
type UserData struct {
	Data []User `json:"data"`
}

//Stream type of API_Authorization data
type Stream struct {
	ID           string   `json:"id"`
	Type         string   `json:"type"`
	Language     string   `json:"language"`
	GameID       string   `json:"game_id"`
	ThumbnailURL string   `json:"thumbnail_url"`
	Title        string   `json:"title"`
	UserID       string   `json:"user_id"`
	StartedAt    string   `json:"started_at"`
	CommunityIDs []string `json:"community_ids,omitempty"`
}

//Streams strcutre with array
type Streams struct {
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
	Data []Stream `json:"data,omitempty"`
}

//UserGorm Model
type UserGorm struct {
	ID           string
	Login        string
	DisplayName  string
	Type         string `gorm:"default:'NULL'"`
	Description  string `gorm:"default:'NULL'"`
	ProfileImage string
	ViewCount    int64
	Email        string `gorm:"default:'NULL'"`
}
