package entity

type Person struct {
	ID    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Age   int    `json:"age"`
}
