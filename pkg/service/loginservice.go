package service

//go:generate mockgen -destination mock/mock_login.go -source loginservice.go ILogin
type ILogin interface {
	Login(user, password string) error
}

func (q InfoEventService) Login(user, password string) error {
	err := q.SQLite3Repo.LoginInfo(user, password)
	if err != nil {
		return err
	}
	return nil
}
