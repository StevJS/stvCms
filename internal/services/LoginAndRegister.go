package services

type ILoginAndRegisterService interface {
}

func NewLoginAndRegisterService() ILoginAndRegisterService {
	return &postService{}
}
