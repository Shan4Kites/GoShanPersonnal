package orm

import (
	"fmt"
	"database/sql"
)

type ActiveRecord struct {
	DB *sql.DB
	ObjectMapper ObjectMapper
}

func convertToInsertFields(mapping map[string]interface{}) (string,string,[]interface{}) {
	fields := ""
	placeHolders := ""
	i := 1
	var values []interface{}

	for k,v := range mapping {
		if len(fields) > 0 {
			fields += ","
			placeHolders += ","
		}
		placeHolders += fmt.Sprintf("$%d", i)
		i++
		fields += k
		values = append(values, v)
	}
	return fields,placeHolders,values
}

func (self ActiveRecord) Create()  {
	mapper := self.ObjectMapper.GetMapping()
	fields,placeHolders,values := convertToInsertFields(mapper)
	sqlStatement := "INSERT INTO "+ self.ObjectMapper.GetTableName() + "(" + fields + ") VALUES ("+ placeHolders+ ")"
	_, err := self.DB.Exec(sqlStatement, values...)
	if err != nil {
		panic(err)
	}
}

func (self ActiveRecord) Update()  {
	mapper := self.ObjectMapper.GetMapping()
	fmt.Println("updating record uses the mapper : ", mapper)
}

