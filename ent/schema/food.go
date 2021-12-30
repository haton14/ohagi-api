package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Food holds the schema definition for the Food entity.
type Food struct {
	ent.Schema
}

// Fields of the Food.
func (Food) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(100),
		field.String("unit").MaxLen(50),
	}
}

// Edges of the Food.
func (Food) Edges() []ent.Edge {
	return nil
}
