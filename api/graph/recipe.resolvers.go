package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"forkd/graph/model"
)

// Author is the resolver for the author field.
func (r *recipeResolver) Author(ctx context.Context, obj *model.Recipe) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Author - author"))
}

// Revisions is the resolver for the revisions field.
func (r *recipeResolver) Revisions(ctx context.Context, obj *model.Recipe, limit *int, nextCursor *string) (*model.PaginatedRecipeRevisions, error) {
	panic(fmt.Errorf("not implemented: Revisions - revisions"))
}

// Revision is the resolver for the revision field.
func (r *recipeCommentResolver) Revision(ctx context.Context, obj *model.RecipeComment) (*model.RecipeRevision, error) {
	panic(fmt.Errorf("not implemented: Revision - revision"))
}

// Recipe is the resolver for the recipe field.
func (r *recipeCommentResolver) Recipe(ctx context.Context, obj *model.RecipeComment) (*model.Recipe, error) {
	panic(fmt.Errorf("not implemented: Recipe - recipe"))
}

// Author is the resolver for the author field.
func (r *recipeCommentResolver) Author(ctx context.Context, obj *model.RecipeComment) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Author - author"))
}

// Ingredients is the resolver for the ingredients field.
func (r *recipeRevisionResolver) Ingredients(ctx context.Context, obj *model.RecipeRevision) ([]*model.RecipeIngredient, error) {
	panic(fmt.Errorf("not implemented: Ingredients - ingredients"))
}

// Steps is the resolver for the steps field.
func (r *recipeRevisionResolver) Steps(ctx context.Context, obj *model.RecipeRevision) ([]*model.RecipeStep, error) {
	panic(fmt.Errorf("not implemented: Steps - steps"))
}

// Revision is the resolver for the revision field.
func (r *recipeStepResolver) Revision(ctx context.Context, obj *model.RecipeStep) (*model.RecipeRevision, error) {
	panic(fmt.Errorf("not implemented: Revision - revision"))
}

// Recipe returns RecipeResolver implementation.
func (r *Resolver) Recipe() RecipeResolver { return &recipeResolver{r} }

// RecipeComment returns RecipeCommentResolver implementation.
func (r *Resolver) RecipeComment() RecipeCommentResolver { return &recipeCommentResolver{r} }

// RecipeRevision returns RecipeRevisionResolver implementation.
func (r *Resolver) RecipeRevision() RecipeRevisionResolver { return &recipeRevisionResolver{r} }

// RecipeStep returns RecipeStepResolver implementation.
func (r *Resolver) RecipeStep() RecipeStepResolver { return &recipeStepResolver{r} }

type recipeResolver struct{ *Resolver }
type recipeCommentResolver struct{ *Resolver }
type recipeRevisionResolver struct{ *Resolver }
type recipeStepResolver struct{ *Resolver }
