package models

// type Usuario struct {
// 	ID                  int    `json:"id"`
// 	User_login          string `json:"user_login"`
// 	User_pass           string `json:"user_pass"`
// 	User_nicename       string `json:"user_nicename"`
// 	User_email          string `json:"user_email"`
// 	User_url            string `json:"user_url"`
// 	User_activation_key string `json:"user_activation_key"`
// }

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Rut      string `json:"rut"`
	Role     string `json:"role"`
}

type Users []User
