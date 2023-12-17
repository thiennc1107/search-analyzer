// Code generated by ent, DO NOT EDIT.

package keywords

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the keywords type in the database.
	Label = "keywords"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKeyword holds the string denoting the keyword field in the database.
	FieldKeyword = "keyword"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldAdsAmount holds the string denoting the ads_amount field in the database.
	FieldAdsAmount = "ads_amount"
	// FieldLinksAmount holds the string denoting the links_amount field in the database.
	FieldLinksAmount = "links_amount"
	// FieldSearchResultAmount holds the string denoting the search_result_amount field in the database.
	FieldSearchResultAmount = "search_result_amount"
	// FieldHTMLCode holds the string denoting the html_code field in the database.
	FieldHTMLCode = "html_code"
	// Table holds the table name of the keywords in the database.
	Table = "keywords"
)

// Columns holds all SQL columns for keywords fields.
var Columns = []string{
	FieldID,
	FieldKeyword,
	FieldStatus,
	FieldAdsAmount,
	FieldLinksAmount,
	FieldSearchResultAmount,
	FieldHTMLCode,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusFailed     Status = "failed"
	StatusProcessing Status = "processing"
	StatusFinished   Status = "finished"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusFailed, StatusProcessing, StatusFinished:
		return nil
	default:
		return fmt.Errorf("keywords: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Keywords queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKeyword orders the results by the keyword field.
func ByKeyword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKeyword, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByAdsAmount orders the results by the ads_amount field.
func ByAdsAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAdsAmount, opts...).ToFunc()
}

// ByLinksAmount orders the results by the links_amount field.
func ByLinksAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLinksAmount, opts...).ToFunc()
}

// BySearchResultAmount orders the results by the search_result_amount field.
func BySearchResultAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSearchResultAmount, opts...).ToFunc()
}

// ByHTMLCode orders the results by the html_code field.
func ByHTMLCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHTMLCode, opts...).ToFunc()
}
