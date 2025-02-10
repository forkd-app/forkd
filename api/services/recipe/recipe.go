package recipe

import (
	"context"
	"fmt"
	"forkd/db"
	"forkd/graph/model"
	"forkd/services/auth"
	"forkd/util"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecipeService interface {
	ListRecipes(ctx context.Context, input *model.ListRecipeInput) (*model.PaginatedRecipes, error)
	GetRecipeByID(ctx context.Context, id uuid.UUID) (*model.Recipe, error)
	GetRecipeBySlug(ctx context.Context, slug string) (*model.Recipe, error)
	ListRevisions(ctx context.Context, input *model.ListRevisionsInput) (*model.PaginatedRecipeRevisions, error)
	GetRevisionById(ctx context.Context) (*model.RecipeRevision, error)
	ListRecipeIngredientsForRevision(ctx context.Context) ([]model.RecipeIngredient, error)
	ListRecipeStepsForRevision(ctx context.Context) ([]model.RecipeStep, error)
	CreateRecipe(ctx context.Context, params db.CreateRecipeParams) (*model.Recipe, error)
	AddRevision(ctx context.Context, params db.CreateRevisionParams) (*model.RecipeRevision, error)
}

type recipeService struct {
	queries *db.Queries
	conn    *pgxpool.Pool
	auth    auth.AuthService
}

// AddRevision implements RecipeService.
func (r recipeService) AddRevision(ctx context.Context, params db.CreateRevisionParams) (*model.RecipeRevision, error) {
	panic("unimplemented")
}

// CreateRecipe implements RecipeService.
func (r recipeService) CreateRecipe(ctx context.Context, params db.CreateRecipeParams) (*model.Recipe, error) {
	panic("unimplemented")
}

// GetRecipeByID implements RecipeService.
func (r recipeService) GetRecipeByID(ctx context.Context, id uuid.UUID) (*model.Recipe, error) {
	pgId := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := r.queries.GetRecipeById(ctx, pgId)
	return util.HandleNoRowsOnNullableType(result, err, model.RecipeFromDBType)
}

// GetRecipeBySlug implements RecipeService.
func (r recipeService) GetRecipeBySlug(ctx context.Context, slug string) (*model.Recipe, error) {
	result, err := r.queries.GetRecipeBySlug(ctx, slug)
	return util.HandleNoRowsOnNullableType(result, err, model.RecipeFromDBType)
}

// GetRevisionById implements RecipeService.
func (r recipeService) GetRevisionById(ctx context.Context) (*model.RecipeRevision, error) {
	panic("unimplemented")
}

// ListRecipeIngredientsForRevision implements RecipeService.
func (r recipeService) ListRecipeIngredientsForRevision(ctx context.Context) ([]model.RecipeIngredient, error) {
	panic("unimplemented")
}

// ListRecipeStepsForRevision implements RecipeService.
func (r recipeService) ListRecipeStepsForRevision(ctx context.Context) ([]model.RecipeStep, error) {
	panic("unimplemented")
}

// ListRevisionsForRecipe implements RecipeService.
func (r recipeService) ListRevisions(ctx context.Context, input *model.ListRevisionsInput) (*model.PaginatedRecipeRevisions, error) {
	var params db.ListRevisionsParams
	if input == nil {
		params.Limit = 20
		params.SortDir = true
		params.SortCol = "publish_date"
	} else {
		params.Limit = int32(*input.Limit)
		switch *input.SortCol {
		case model.ListRecipeSortColPublishDate:
			params.SortCol = "publish_date"
		}
		switch *input.SortDir {
		case model.SortDirDesc:
			params.SortDir = true
		case model.SortDirAsc:
			params.SortDir = false
		}
		if input.RecipeID != nil {
			params.RecipeID = pgtype.UUID{
				Bytes: *input.RecipeID,
				Valid: true,
			}
		}
		if input.ParentID != nil {
			params.ParentID = pgtype.UUID{
				Bytes: *input.ParentID,
				Valid: true,
			}
		}
		if input.PublishStart != nil {
			params.PublishStart = pgtype.Timestamp{
				Time:  *input.PublishStart,
				Valid: true,
			}
		}
		if input.PublishEnd != nil {
			params.PublishEnd = pgtype.Timestamp{
				Time:  *input.PublishEnd,
				Valid: true,
			}
		}
		if input.NextCursor != nil {
			cursor := new(ListRevisionsCursor)
			err := cursor.Decode(*input.NextCursor)
			if err != nil {
				return nil, err
			}
			if !cursor.Validate(ListRevisionsCursor{
				ListRevisionsInput: *input,
			}) {
				return nil, fmt.Errorf("invalid cursor")
			}
			params.PublishCursor = cursor.PublishCursor
		}
	}

	result, err := r.queries.ListRevisions(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch revisions: %w", err)
	}

	revisions := make([]*model.RecipeRevision, len(result))
	for i, rev := range result {
		revisions[i] = model.RevisionFromDBType(rev)
	}

	var NextCursor *string = nil
	if len(revisions) == int(params.Limit) {
		cursor := ListRevisionsCursor{
			ListRevisionsInput: *input,
		}

		switch params.SortCol {
		case "publish_date":
			fallthrough
		default:
			cursor.PublishCursor = pgtype.Timestamp{
				Time:  revisions[len(revisions)-1].PublishDate,
				Valid: true,
			}
		}

		encoded, err := cursor.Encode()

		if err != nil {
			return nil, err
		}

		NextCursor = &encoded
	}

	return &model.PaginatedRecipeRevisions{
		Items: revisions,
		Pagination: &model.PaginationInfo{
			Count:      len(revisions),
			NextCursor: NextCursor,
		},
	}, nil
}

// ListRecipes implements RecipeService.
func (r recipeService) ListRecipes(ctx context.Context, input *model.ListRecipeInput) (*model.PaginatedRecipes, error) {
	var params db.ListRecipesParams
	user, _ := r.auth.GetUserSessionFromCtx(ctx)
	params.CurrentUser = user.ID
	if input == nil {
		params.Limit = 20
		params.SortDir = true
		params.SortCol = "publish_date"
	} else {
		params.Limit = int32(*input.Limit)
		switch *input.SortCol {
		case model.ListRecipeSortColPublishDate:
			params.SortCol = "publish_date"
		case model.ListRecipeSortColSlug:
			params.SortCol = "slug"
		}
		switch *input.SortDir {
		case model.SortDirDesc:
			params.SortDir = true
		case model.SortDirAsc:
			params.SortDir = false
		}
		if input.AuthorID != nil {
			params.AuthorID = pgtype.UUID{
				Bytes: *input.AuthorID,
				Valid: true,
			}
		}
		if input.PublishStart != nil {
			params.PublishStart = pgtype.Timestamp{
				Time:  *input.PublishStart,
				Valid: true,
			}
		}
		if input.PublishEnd != nil {
			params.PublishEnd = pgtype.Timestamp{
				Time:  *input.PublishEnd,
				Valid: true,
			}
		}
		if input.NextCursor != nil {
			cursor := new(ListRecipesCursor)
			err := cursor.Decode(*input.NextCursor)
			if err != nil {
				return nil, err
			}
			if !cursor.Validate(ListRecipesCursor{
				ListRecipeInput: *input,
			}) {
				return nil, fmt.Errorf("invalid cursor")
			}
			params.PublishCursor = cursor.PublishCursor
			params.SlugCursor = cursor.SlugCursor
		}
	}
	result, err := r.queries.ListRecipes(ctx, params)
	if err != nil {
		return nil, err
	}

	count := len(result)
	recipes := make([]*model.Recipe, count)
	for i, recipe := range result {
		recipes[i] = model.RecipeFromDBType(recipe)
	}

	var NextCursor *string = nil
	if count == int(params.Limit) {
		cursor := ListRecipesCursor{
			ListRecipeInput: *input,
		}

		switch params.SortCol {
		case "slug":
			cursor.SlugCursor = pgtype.Text{
				String: recipes[len(recipes)-1].Slug,
				Valid:  true,
			}
		case "publish_date":
			fallthrough
		default:
			cursor.PublishCursor = pgtype.Timestamp{
				Time:  recipes[len(recipes)-1].InitialPublishDate,
				Valid: true,
			}
		}

		encoded, err := cursor.Encode()

		if err != nil {
			return nil, err
		}

		NextCursor = &encoded
	}

	return &model.PaginatedRecipes{
		Items: recipes,
		Pagination: &model.PaginationInfo{
			Count:      count,
			NextCursor: NextCursor,
		},
	}, nil
}

func New(queries *db.Queries, conn *pgxpool.Pool, authService auth.AuthService) RecipeService {
	return recipeService{
		queries,
		conn,
		authService,
	}
}

type ListRecipesCursor struct {
	model.ListRecipeInput
	PublishCursor pgtype.Timestamp
	SlugCursor    pgtype.Text
}

func (cursor *ListRecipesCursor) Decode(encoded string) error {
	return util.DecodeBase64StringToStruct(encoded, cursor)
}

func (cursor ListRecipesCursor) Encode() (string, error) {
	return util.EncodeStructToBase64String(cursor)
}

func (cursor ListRecipesCursor) Validate(input ListRecipesCursor) bool {
	return util.ComparePointerValues(cursor.Limit, input.Limit) &&
		util.ComparePointerValues(cursor.SortCol, input.SortCol) &&
		util.ComparePointerValues(cursor.SortDir, input.SortDir) &&
		util.ComparePointerValues(cursor.AuthorID, input.AuthorID) &&
		util.ComparePointerValues(cursor.PublishStart, input.PublishStart) &&
		util.ComparePointerValues(cursor.PublishEnd, input.PublishEnd)
}

type ListRevisionsCursor struct {
	model.ListRevisionsInput
	PublishCursor pgtype.Timestamp
}

func (cursor *ListRevisionsCursor) Decode(encoded string) error {
	return util.DecodeBase64StringToStruct(encoded, cursor)
}

func (cursor ListRevisionsCursor) Encode() (string, error) {
	return util.EncodeStructToBase64String(cursor)
}

func (cursor ListRevisionsCursor) Validate(input ListRevisionsCursor) bool {
	return util.ComparePointerValues(cursor.Limit, input.Limit) &&
		util.ComparePointerValues(cursor.SortCol, input.SortCol) &&
		util.ComparePointerValues(cursor.SortDir, input.SortDir) &&
		util.ComparePointerValues(cursor.RecipeID, input.RecipeID) &&
		util.ComparePointerValues(cursor.ParentID, input.ParentID) &&
		util.ComparePointerValues(cursor.PublishStart, input.PublishStart) &&
		util.ComparePointerValues(cursor.PublishEnd, input.PublishEnd)
}
