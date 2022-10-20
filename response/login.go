package response

type LoginSuccess struct {
	Account string `json:"account"`
	Name    string `json:"name"`
	Token   string `json:"token"`
}

