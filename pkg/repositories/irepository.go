package repositories

type Repository interface {
	FindAll() interface{}
	FindById(id int64) interface{}
	Create() interface{}
	Update() interface{}
	Delete(id int64) interface{}
}
