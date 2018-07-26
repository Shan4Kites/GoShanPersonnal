package orm

type ObjectMapper interface {
	GetMapping() map[string]interface{}
	GetTableName() string
}