package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/dialect"

)

// Bouncer holds the schema definition for the Bouncer entity.
type Bouncer struct {
	ent.Schema
}

// Fields of the Bouncer.
func (Bouncer) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
		field.String("name").Unique().SchemaType(map[string]string{
			dialect.MySQL: "varchar(190)",
			dialect.SQLite: "varchar(190)",
		}),
		field.String("api_key"), // hash of api_key
		field.Bool("revoked"),
		field.String("ip_address").Default("").Optional(),
		field.String("type").Optional(),
		field.String("version").Optional(),
		field.Time("until").Default(time.Now).Optional(),
		field.Time("last_pull").
			Default(time.Now),
	}
}

// Edges of the Bouncer.
func (Bouncer) Edges() []ent.Edge {
	return nil
}
