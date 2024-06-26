// Code generated by ent, DO NOT EDIT.

package book

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsbn holds the string denoting the isbn field in the database.
	FieldIsbn = "isbn"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldGenre holds the string denoting the genre field in the database.
	FieldGenre = "genre"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldPublicizedAt holds the string denoting the publicized_at field in the database.
	FieldPublicizedAt = "publicized_at"
	// Table holds the table name of the book in the database.
	Table = "books"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldIsbn,
	FieldTitle,
	FieldAuthor,
	FieldGenre,
	FieldQuantity,
	FieldPublicizedAt,
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

// OrderOption defines the ordering options for the Book queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIsbn orders the results by the isbn field.
func ByIsbn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsbn, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByGenre orders the results by the genre field.
func ByGenre(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGenre, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByPublicizedAt orders the results by the publicized_at field.
func ByPublicizedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicizedAt, opts...).ToFunc()
}
