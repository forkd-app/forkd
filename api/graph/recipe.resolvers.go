package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"forkd/graph/model"

	"github.com/jackc/pgx/v5/pgtype"
)

// Author is the resolver for the author field.
func (r *recipeResolver) Author(ctx context.Context, obj *model.Recipe) (*model.User, error) {
	if obj == nil {
		return nil, fmt.Errorf("missing recipe object")
	}
	uuid := pgtype.UUID{
		Bytes: obj.ID,
		Valid: true,
	}
	data, err := r.Queries.GetAuthorByRecipeId(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch author: %w", err)
	}

	return model.UserFromDBType(data), nil
}

// ForkedFrom is the resolver for the forkedFrom field.
func (r *recipeResolver) ForkedFrom(ctx context.Context, obj *model.Recipe) (*model.RecipeRevision, error) {
	return r.RecipeService.GetRecipeForkedFromRevision(ctx, obj.ID)
}

// Revisions is the resolver for the revisions field.
func (r *recipeResolver) Revisions(ctx context.Context, obj *model.Recipe, input *model.ListRevisionsInput) (*model.PaginatedRecipeRevisions, error) {
	if input == nil {
		sortCol := model.ListRecipeSortColPublishDate
		sortDir := model.SortDirDesc
		limit := 20
		input = &model.ListRevisionsInput{
			RecipeID: &obj.ID,
			Limit:    &limit,
			SortDir:  &sortDir,
			SortCol:  &sortCol,
		}
	} else {
		input.RecipeID = &obj.ID
	}
	return r.RecipeService.ListRevisions(ctx, input)
}

// FeaturedRevision is the resolver for the featuredRevision field.
func (r *recipeResolver) FeaturedRevision(ctx context.Context, obj *model.Recipe) (*model.RecipeRevision, error) {
	return r.RecipeService.GetRecipeFeaturedRevision(ctx, obj.ID)
}

// Recipe is the resolver for the recipe field.
func (r *recipeRevisionResolver) Recipe(ctx context.Context, obj *model.RecipeRevision) (*model.Recipe, error) {
	return r.RecipeService.GetRevisionRecipe(ctx, obj.ID)
}

// Parent is the resolver for the parent field.
func (r *recipeRevisionResolver) Parent(ctx context.Context, obj *model.RecipeRevision) (*model.RecipeRevision, error) {
	return r.RecipeService.GetRevisionParent(ctx, obj.ID)
}

// Ingredients is the resolver for the ingredients field.
func (r *recipeRevisionResolver) Ingredients(ctx context.Context, obj *model.RecipeRevision) ([]*model.RecipeIngredient, error) {
	return r.RecipeService.ListRecipeIngredients(ctx, obj.ID)
}

// Steps is the resolver for the steps field.
func (r *recipeRevisionResolver) Steps(ctx context.Context, obj *model.RecipeRevision) ([]*model.RecipeStep, error) {
	return r.RecipeService.ListRecipeSteps(ctx, obj.ID)
}

// Rating is the resolver for the rating field.
func (r *recipeRevisionResolver) Rating(ctx context.Context, obj *model.RecipeRevision) (*float64, error) {
	// TODO: Implement logic for computing the rating. This might be best done as a computed field inside the db, but might also be good to have a dedicated resolver for
	rating := float64(0)

	return &rating, nil
}

// Revision is the resolver for the revision field.
func (r *recipeStepResolver) Revision(ctx context.Context, obj *model.RecipeStep) (*model.RecipeRevision, error) {
	return r.RecipeService.GetRevisionForStep(ctx, int64(obj.ID))
}

// Recipe returns RecipeResolver implementation.
func (r *Resolver) Recipe() RecipeResolver { return &recipeResolver{r} }

// RecipeRevision returns RecipeRevisionResolver implementation.
func (r *Resolver) RecipeRevision() RecipeRevisionResolver { return &recipeRevisionResolver{r} }

// RecipeStep returns RecipeStepResolver implementation.
func (r *Resolver) RecipeStep() RecipeStepResolver { return &recipeStepResolver{r} }

type recipeResolver struct{ *Resolver }
type recipeRevisionResolver struct{ *Resolver }
type recipeStepResolver struct{ *Resolver }
