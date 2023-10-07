package sqlclient

import "strings"

func buildWhereClause(whereCondition interface{}) string {
	var whereItems []string
	var result = ""
	if items, ok := whereCondition.(map[string]string); ok {
		for key, value := range items {
			whereItems = append(whereItems, key+" = "+value)
		}
	}
	if len(whereItems) > 0 {
		result += " WHERE " + strings.Join(whereItems, " AND ")
	}
	return result
}

func buildClause(conditions map[string]interface{}) string {
	var result = ""
	if groupCondition, ok := conditions["group"]; ok {
		result += " GROUP BY " + buildGroupClause(groupCondition)
	}
	if orderCondition, ok := conditions["order"]; ok {
		result += " ORDER BY " + buildOrderClause(orderCondition)
	}
	if limitCondition, ok := conditions["limit"]; ok {
		result += " LIMIT " + buildLimitClause(limitCondition)
	}
	return result
}

func buildGroupClause(groupCondition interface{}) string {
	var result = ""
	if field, ok := groupCondition.(string); ok {
		result += " " + field
	}
	return result
}

func buildOrderClause(orderCondition interface{}) string {
	var result = ""
	if items, ok := orderCondition.(map[string]string); ok {
		result += " " + items["field"]
		if orderType, ok := items["orderType"]; ok {
			result += " " + orderType
		}
	}
	return result
}

func buildLimitClause(whereCondition interface{}) string {
	var result = ""
	if limit, ok := whereCondition.(string); ok {
		result += " " + limit
	}
	return result
}
