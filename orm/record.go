package orm

type Record interface {
	Create()
	Update()
	Delete()
	Where(query string, args ...interface{})  (objects []interface{})
}