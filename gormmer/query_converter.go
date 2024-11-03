package gormmer

import (
	"net/url"
	"strconv"
	"strings"

	"gorm.io/gorm/clause"
)

/**
Converter HTTP request query to db query structure
*/

func ConvertQueryToLimit(url *url.URL) int {
	var limit = 6

	limitQuery := url.Query().Get("limit")

	if url.Query().Get("limit") != "" {
		limitQueryInt, err := strconv.Atoi(limitQuery)
		if err != nil {
			return limit
		}

		return limitQueryInt
	}

	return limit
}

func ConvertQueryToFilter(url *url.URL, allowFilterQuery []string) map[string]string {
	var filter = make(map[string]string)

	for key, val := range url.Query() {
		for _, allowKey := range allowFilterQuery {
			if allowKey == key {
				_, err := strconv.Atoi(val[0])
				if err == nil {
					filter[key+" = ?"] = val[0]
				} else {
					_, err := strconv.ParseBool(val[0])
					if err == nil {
						filter[key+" = ?"] = val[0]
					} else {
						filter[key+" LIKE ?"] = "%" + val[0] + "%"
					}
				}
			}
		}
	}

	return filter
}

func ConvertQueryToOrder(url *url.URL, selectOrder string) clause.OrderByColumn {
	var isDescOrder = true
	var orderBy = "id"

	// helper to fixing join query
	if selectOrder != "" {
		orderBy = selectOrder
	}

	order := strings.ToLower(url.Query().Get("order"))
	by := strings.ToLower(url.Query().Get("by"))

	if order == "asc" {
		isDescOrder = false
	}

	if by != "" {
		orderBy = by
	}

	return clause.OrderByColumn{Column: clause.Column{Name: orderBy}, Desc: isDescOrder}
}
