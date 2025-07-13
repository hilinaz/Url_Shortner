package db
import "log"
func CreateTable(){
	Table:= `CREATE TABLE IF NOT EXISTS UrlPaths(
	Id int PRIMARY KEY AUTO_INCREMENT,
	OriginalUrl VARCHAR(255),
	ShortUrl VARCHAR(255),
	CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP);`
	_,err:=DB.Exec(Table)
	if err!=nil{
		log.Fatal("Table could not be created",err)
	}
	log.Println("table created")
}