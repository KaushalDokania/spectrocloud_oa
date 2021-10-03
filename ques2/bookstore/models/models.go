package models

type Book struct {
	ID     uint64  `gorm:"primary_key;auto_increment" json:"id" bson:"_id"`
	Isbn   string  `json:"isbn" bson:"isbn"`
	Title  string  `json:"title" bson:"title"`
	Author string  `json:"author" bson:"author"`
	Price  float32 `json:"price" bson:"price"`
}

