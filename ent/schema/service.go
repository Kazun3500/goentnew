package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("owner_id"),
		field.Enum("type").NamedValues("svh", "взх", "mail_lite", "ЭПЛ"),
	}
}

func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Required().Unique().
			Ref("services").Field("owner_id"),
	}
}
