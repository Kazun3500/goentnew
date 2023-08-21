package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Service struct {
	ent.Schema
}

func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").NamedValues("svh", "взх", "mail_lite", "ЭПЛ"),
	}
}

func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Unique().
			Ref("services"),
	}
}
