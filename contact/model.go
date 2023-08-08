package contact

type Contact struct {
	Id      int   `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"not null" binding:"required"`
	Phone   string `json:"phone" gorm:"not null" binding:"required"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
