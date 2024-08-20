package orm

type Model interface{}

func Select(m Model) {

	// v := reflect.ValueOf(m)
	// vtype := reflect.TypeOf(v)

	// columnNames := make([]string, v.NumField())
	// values := make([]any)
	// for i := 0; i < v.NumField(); i++ {
	// 	v.Field(i).
	// 		columnNames[i] = vtype.Field(i).Name
	// }

	// log.Println(columnNames)

	// query := strings.Join(columnNames, ", ")
	// tableName := strings.ToLower(v.Type().Name())

}

func Insert(m Model) string {
	return ""
}

func Delete(m Model) string {
	return ""
}

func Update(m Model) string {
	return ""
}
