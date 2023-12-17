package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/enum"
)

// Keywords holds the schema definition for the Keywords entity.
type Keywords struct {
	ent.Schema
}

// Fields of the Keywords.
func (Keywords) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values(
				enum.StatusPending,
				enum.StatusFailed,
				enum.StatusProcessing,
				enum.StatusFinished,
			),
		field.Int("ads_amount"),
		field.Int("links"),
		field.Int("search_result_amount"),
		field.Int("html_code"),
	}
}

// Edges of the Keywords.
func (Keywords) Edges() []ent.Edge {
	return nil
}
