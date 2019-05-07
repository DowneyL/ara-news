package mysql

import "github.com/astaxie/beego/orm"

func GetMysqlConnConfigs() map[string]ConnConfig {
	return configs
}

func GetOrmer(genre ...string) orm.Ormer {
	var g string
	if len(genre) > 0 && genre[0] != "" {
		g = genre[0]
	} else {
		g = "default"
	}

	if ormers[g] != nil {
		return ormers[g]
	}

	o := orm.NewOrm()
	_ = o.Using(g)
	ormers[g] = o

	return o
}

// e.g. QueryTable("user"), QueryTable(&user{}) or QueryTable((*User)(nil))
func GetQuerySetter(ptrStructOrTableName interface{}, genre ...string) orm.QuerySeter {
	var o orm.Ormer
	if len(genre) > 0 && genre[0] != "" {
		o = GetOrmer(genre[0])
	} else {
		o = GetOrmer()
	}

	return o.QueryTable(ptrStructOrTableName)
}
