package entity

type BatchUpdateInput struct {
	// Put is a slice of entities to be upserted to the entity store.
	Put []Entity
	// PutChildren is a slice of parent/child relationships to be added to the entity store.
	PutChildren []ChildRelation
}
