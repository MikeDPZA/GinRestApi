package models

type Book struct {
	ID       int64 `gorm:"primary_key;auto_increment;not null;unique;"`
	Name     string
	AuthorID int64 `gorm:"constraint:OnDelete:SET NULL;"`
	Author   Author
}

func (b *Book) FindAll() []Book {
	var Books []Book
	ctx.Where("author_id is not null").Preload("Author").Find(&Books)
	return Books
}

func (b *Book) FindById(id int64) *Book {
	ctx.Where("ID=?", id).Preload("Author").Find(&b)
	return b
}

func (b *Book) Create() *Book {
	ctx.Create(&b)
	return b
}

func (b *Book) Update() *Book {
	ctx.Save(&b)
	return b
}

func (b *Book) Delete(id int64) *Book {
	ctx.Where("ID=?", id).Delete(&b)
	return b
}
