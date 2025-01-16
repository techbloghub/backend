// Code generated by ent, DO NOT EDIT.

package company

import (
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
)

const (
	// Label holds the string label denoting the company type in the database.
	Label = "company"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLogoURL holds the string denoting the logo_url field in the database.
	FieldLogoURL = "logo_url"
	// FieldBlogURL holds the string denoting the blog_url field in the database.
	FieldBlogURL = "blog_url"
	// FieldRssURL holds the string denoting the rss_url field in the database.
	FieldRssURL = "rss_url"
	// Table holds the table name of the company in the database.
	Table = "companies"
)

// Columns holds all SQL columns for company fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldLogoURL,
	FieldBlogURL,
	FieldRssURL,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// ValueScanner of all Company fields.
	ValueScanner struct {
		LogoURL field.TypeValueScanner[*url.URL]
		BlogURL field.TypeValueScanner[*url.URL]
		RssURL  field.TypeValueScanner[*url.URL]
	}
)

// OrderOption defines the ordering options for the Company queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLogoURL orders the results by the logo_url field.
func ByLogoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogoURL, opts...).ToFunc()
}

// ByBlogURL orders the results by the blog_url field.
func ByBlogURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlogURL, opts...).ToFunc()
}

// ByRssURL orders the results by the rss_url field.
func ByRssURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRssURL, opts...).ToFunc()
}
