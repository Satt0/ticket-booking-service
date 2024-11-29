package entities

type Users struct {
	BaseModel
	Email    string `gorm:"size:255;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Age      string `gorm:"size:255;not null" json:"age"`
	Gender   string `gorm:"size:255;not null" json:"gender"`
	Balance  string `gorm:"size:255;not null" json:"balance"`
}
