package service

type AdminService interface {
	DeleteById(id int) error
}
