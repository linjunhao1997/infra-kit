package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type Namespace struct {
	ent.Schema
}

func (Namespace) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return xid.New().String()
			}),
		field.String("code"),
		field.String("name"),
		field.String("description"),
		field.String("label"),
		field.Time("ctime"),
		field.Time("mtime"),
		field.Bool("deleted"),
	}
}

func (Namespace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org", Org.Type).Ref("namespaces").Unique(),
		edge.To("users", User.Type),
	}
}
