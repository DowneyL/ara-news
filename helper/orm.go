package helper

import (
	"errors"
	"strings"
)

const (
	desc string = "desc"
	asc         = "asc"
)

func GetOrmOrders(orderBy string) ([]string, error) {
	var orders []string
	split := strings.Split(strings.Trim(orderBy, " "), ",")
	for _, orderStatement := range split {
		orderStatement = strings.Trim(orderStatement, " ")
		ss := strings.Split(orderStatement, " ")
		var key, code string
		if len(ss) == 2 {
			key = ss[0]
			code = strings.ToLower(ss[1])
		} else if len(ss) == 1 {
			key = ss[0]
			code = "asc"
		} else {
			return nil, errors.New("invalid order statement")
		}
		if code != desc && code != asc {
			return nil, errors.New("invalid order code")
		}

		if code == desc {
			key = "-" + key
		}

		orders = append(orders, key)
	}

	return orders, nil
}
