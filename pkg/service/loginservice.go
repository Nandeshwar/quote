package service

type ILogin interface {
	Login(user, password string) error
}

func (q QuoteService) Login(user, password string) error {
	err := q.SQLite3Repo.LoginInfo(user, password)
	if err != nil {
		return err
	}
	return nil
}
