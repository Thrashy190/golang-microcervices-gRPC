package models

type Student struct {
	// Id, Name, Age
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Test struct {
	// Id, Name
	Id   string `json:"id"`
	Name string `json:"name"`
}
