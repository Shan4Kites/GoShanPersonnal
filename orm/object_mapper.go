package orm

type ObjectMapper interface {
	GetMapping() map[string]interface{}
	GetTableName() string
	GetObject(mapper map[string][]byte) (object interface{})
}