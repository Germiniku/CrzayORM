package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string	// 将Go语言的类型转换为该数据库的数据类型
	TableExistSQL(tablename string) (string, []interface{}) // 返回某个表是否存在的SQL语句
}

func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialet Dialect, ok bool) {
	dialet, ok = dialectsMap[name]
	return
}
