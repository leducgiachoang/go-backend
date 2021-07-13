package model

type SignUp struct {
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	Address  string `json:"address,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Token    string `json:"token,omitempty"`
}
