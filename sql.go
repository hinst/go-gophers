package gophers

import "strconv"

type SqlValues struct {
	key    string
	values []any
}

func (me *SqlValues) Add(value any) string {
	me.values = append(me.values, value)
	return me.getKey() + strconv.Itoa(len(me.values))
}

func (me SqlValues) getKey() string {
	return IfElse(me.key != "", me.key, "$")
}

func (me SqlValues) Values() []any {
	return me.values
}
