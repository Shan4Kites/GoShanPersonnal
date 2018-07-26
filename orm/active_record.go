package orm

import "fmt"

type ActiveRecord struct {
	ObjectMapper ObjectMapper
}

func (self ActiveRecord) Create()  {
	mapper := self.ObjectMapper.GetMapping()
	fmt.Println("creating record uses the mapper : ", mapper)
}

func (self ActiveRecord) Update()  {
	mapper := self.ObjectMapper.GetMapping()
	fmt.Println("updating record uses the mapper : ", mapper)
}

