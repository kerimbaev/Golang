package models

type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Reciept struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	CategoryID string `json:"category_id" db:"category_id"`
	Description string `json:"description" db:"description"`
	DatePublished string `json:"datePublished" db:"datePublished"`
	Ingridient string `json:"ingridient" db:"ingridient"`
}
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`

	Email    string `json:"email"`
	Password string `json:"password"`
}