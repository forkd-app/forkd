package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"forkd/graph/model"
)

// Revision is the resolver for the revision field.
func (r *recipeIngredientResolver) Revision(ctx context.Context, obj *model.RecipeIngredient) (*model.RecipeRevision, error) {
	panic(fmt.Errorf("not implemented: Revision - revision"))
}

// Unit is the resolver for the unit field.
func (r *recipeIngredientResolver) Unit(ctx context.Context, obj *model.RecipeIngredient) (*model.MeasurementUnit, error) {
	panic(fmt.Errorf("not implemented: Unit - unit"))
}

// Ingredient is the resolver for the ingredient field.
func (r *recipeIngredientResolver) Ingredient(ctx context.Context, obj *model.RecipeIngredient) (*model.Ingredient, error) {
	panic(fmt.Errorf("not implemented: Ingredient - ingredient"))
}

// RecipeIngredient returns RecipeIngredientResolver implementation.
func (r *Resolver) RecipeIngredient() RecipeIngredientResolver { return &recipeIngredientResolver{r} }

type recipeIngredientResolver struct{ *Resolver }
