package entity

// Create Book struct representing the book table in the database
type Book struct {
	Id          uint64 `gorm:"primary_key;auto_increment" json:"id"` // Primary key, auto-increment id with json tag id for json marshalling
	Title       string `gorm:"type:varchar(255)" json:"title"`       // Data type varchar with json tag name for json marshalling
	Author      string `gorm:"type:varchar(255)" json:"author"`      // Data type varchar with json tag name for json marshalling
	Price       int64  `gorm:"type:int(11)" json:"price"`            // Data type varchar with json tag name for json marshalling
	Description string `gorm:"type:varchar(255)" json:"description"` // Data type varchar with json tag name for json marshalling
	UserId      uint64 `gorm:"not_null" json:"-"`

	// Create a foreign key to user table
	User User `gorm:"foreignkey:UserId;constrait:onUpdate:cascade;onDelete:cascade" json:"user"`
}
