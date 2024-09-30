package models


type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Role     string `json:"role" gorm:"not null"` 
	Posts    []Post `json:"posts" gorm:"foreignKey:AuthorID"`
}

type Post struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"not null"`
	Content  string `json:"content" gorm:"not null"`
	AuthorID uint   `json:"author_id"`
	Author   User   `json:"author" gorm:"foreignKey:AuthorID"`
}
