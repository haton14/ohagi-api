package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RecordFood holds the schema definition for the RecordFood entity.
type RecordFood struct {
	ent.Schema
}

// Fields of the RecordFood.
func (RecordFood) Fields() []ent.Field {
	return []ent.Field{
		field.Int("record_id").Positive().Unique(),
		field.Int("food_id").Positive(),
		field.Int("amount").Positive().Default(1),
	}
}

// Edges of the RecordFood.
func (RecordFood) Edges() []ent.Edge {
	return nil
}
