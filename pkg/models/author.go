package models

type Author struct {
	ID    int64 `gorm:"primary_key;auto_increment;not null;unique;"`
	Name  string
	Books []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (a *Author) FindAll() []Author {
	var authors []Author
	ctx.Preload("Books").Find(&authors)
	return authors
}

func (a *Author) FindById(id int64) *Author {
	ctx.Where("ID=?", id).Preload("Books").Find(&a)
	return a
}

func (a *Author) Create() *Author {
	ctx.Create(&a)
	return a
}

func (a *Author) Update() *Author {
	ctx.Save(&a)
	return a
}

func (a *Author) Delete(id int64) *Author {
	ctx.Where("ID=?", id).Delete(&a)
	return a
}
