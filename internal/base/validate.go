package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)
	if req == nil {
		res = append(res, "Переменная не инициализирована")
		return res
	}
	if req.UserID == "" {
		res = append(res, "Поле 'UserID' не инициализировано")
	}
	if req.Title == "" {
		res = append(res, "Поле 'Title' не инициализировано")
	}
	if req.Description == "" {
		res = append(res, "Поле 'Description' не инициализировано")
	}
	return res
}
