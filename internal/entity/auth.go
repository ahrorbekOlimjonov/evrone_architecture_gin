package entity

type (
	// User -.
	User struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		BirthDate  string `json:"birth_date"`
		TgUserName string `json:"tg_user_name"`
		Phone      string `json:"phone"`
		Instagram  string `json:"instagram"`
		ClientFrom string `json:"client_from"`
		RoleID     string `json:"role_id"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}
)

type (
	// Role -.
	Role struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		ClientTypeId string `json:"client_type_id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}
)

type (
	// ClientType -.
	ClientType struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)
