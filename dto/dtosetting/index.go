package dtosetting

type Business_str struct {
	Business_str_id string          `json:"id"`
	Parentid        string          `json:"parentid"`
	Name            string          `json:"name"`
	Dtime           string          `json:"dtime"`
	Desc            string          `json:"desc"`
	Status          int             `json:"status"`
	Sort            int             `json:"sort"`
	Children        []*Business_str `json:"children"`
	Cmd_            string          `json:"cmd_"`
}
type Organize_str struct {
	Organize_str_id string          `json:"id"`
	Parentid        string          `json:"parentid"`
	Name            string          `json:"name"`
	Dtime           string          `json:"dtime"`
	Desc            string          `json:"desc"`
	Status          int             `json:"status"`
	Sort            int             `json:"sort"`
	Children        []*Organize_str `json:"children"`
	Cmd_            string          `json:"cmd_"`
}
type User_str struct {
	User_id         string `json:"id"`
	Username        string `json:"username"`
	Avatar          string `json:"avatar"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Organize_str_id string `json:"organize_str_id"`
	Business_str_id string `json:"business_str_id"`
	Cmd_            string `json:"cmd_"`
}
type Role_str struct {
	Role_id  string `json:"id"`
	Rolename string `json:"rolename"`
	Sort     string `json:"sort"`
	Status   string `json:"status"`
	Dtime    string `json:"dtime"`
	Cmd_     string `json:"cmd_"`
}
