package orm

import (
	"fmt"
	"database/sql"
	"strconv"
	"strings"
)

type ActiveRecord struct {
	DB *sql.DB
	ObjectMapper ObjectMapper
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
	fields,placeHolders,values := convertToInsertFields(mapper)
	sqlStatement := "UPDATE "+ self.ObjectMapper.GetTableName() +" SET (" + fields +") = ("+ placeHolders + ") WHERE id =" + strconv.Itoa(mapper["id"].(int))
	_, err := self.DB.Exec(sqlStatement, values...)
	if err != nil {
		panic(err)
	}
}

func (self ActiveRecord) Delete()  {
	mapper := self.ObjectMapper.GetMapping()
	sqlStatement := "DELETE FROM "+ self.ObjectMapper.GetTableName() +" WHERE id = "+ strconv.Itoa(mapper["id"].(int))
	_, err := self.DB.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func (self ActiveRecord) Where(query string, args ...interface{})  (objects []interface{}) {
	mapper := self.ObjectMapper.GetMapping()
	fields,_,_ := convertToInsertFields(mapper)

	sqlStatement := "select " + fields + " from "+ self.ObjectMapper.GetTableName() +" where " + query
	rows, err := self.DB.Query(sqlStatement, args...)
	if err != nil {
		panic(err)
	}

	byteRows := convertToByteRows(rows)
	objects = convertToObjects(fields, byteRows, self)
	return
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

/**
	This converts the query results into a [row][column][valueBytes]byte array
 */
func convertToByteRows(rows *sql.Rows) (convertedRowsBytes [][][]byte){
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return
	}

	convertedRowsBytes = [][][]byte{}
	// rawResult is where a single row's byte content would be written.
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each byte array
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}
		bytes := make([][]byte, len(cols))

		for i, raw := range rawResult {
			bytes[i] = raw
		}
		convertedRowsBytes = append(convertedRowsBytes, bytes)
	}
	return
}

func convertToObjects(fields string, byteRows [][][]byte, self ActiveRecord) (objects []interface{}) {
	fieldsSlice := strings.Split(fields, ",")
	for _, byteRow := range byteRows {
		var result map[string][]byte
		result = make(map[string][]byte)
		for j, byteColumn := range byteRow {
			result[fieldsSlice[j]] = byteColumn
		}
		objectReturned := self.ObjectMapper.GetObject(result)
		objects = append(objects, objectReturned)
	}
	return
}

