package repo

type NewRepo struct {
	Authorization
	Task
}

type Authorization interface {
	SignUp()
	SignIn()
}

type Task interface {
}
