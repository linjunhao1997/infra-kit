package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type Credential struct {
	ent.Schema
}

func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return xid.New().String()
			}),
		field.String("ak"),
		field.String("sk"),
		field.Time("ctime"),
		field.Time("mtime"),
		field.Bool("deleted"),
	}
}

func (Credential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("credential").Unique(),
	}
}
