package usr_service

type Usr_service struct {
	repo Repo
}

func Init(repo Repo) *Usr_service {

	return &Usr_service{repo}
}
