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
	GetRecipeByID(ctx context.Context, id uuid.UUID) (*model.Recipe, error)
	GetRecipeBySlug(ctx context.Context, slug string) (*model.Recipe, error)
	GetRecipeForkedFromRevision(ctx context.Context, id uuid.UUID) (*model.RecipeRevision, error)
	GetRecipeFeaturedRevision(ctx context.Context, id uuid.UUID) (*model.RecipeRevision, error)
	GetRevisionRecipe(ctx context.Context, id uuid.UUID) (*model.Recipe, error)
	GetRevisionParent(ctx context.Context, id uuid.UUID) (*model.RecipeRevision, error)
	GetRevisionForStep(ctx context.Context, id int64) (*model.RecipeRevision, error)
	GetRevisionForIngredient(ctx context.Context, id int64) (*model.RecipeRevision, error)
	GetRevisionById(ctx context.Context, id uuid.UUID) (*model.RecipeRevision, error)
	GetUnitForRecipeIngredient(ctx context.Context, id int64) (*model.MeasurementUnit, error)
	GetIngredientForRecipeIngredient(ctx context.Context, id int64) (*model.Ingredient, error)
	ListRecipes(ctx context.Context, input *model.ListRecipeInput) (*model.PaginatedRecipes, error)
	ListRevisions(ctx context.Context, input *model.ListRevisionsInput) (*model.PaginatedRecipeRevisions, error)
	ListRecipeIngredients(ctx context.Context, id uuid.UUID) ([]*model.RecipeIngredient, error)
	ListRecipeSteps(ctx context.Context, id uuid.UUID) ([]*model.RecipeStep, error)
	CreateRecipe(ctx context.Context, input model.CreateRecipeInput) (*model.Recipe, error)
	AddRevision(ctx context.Context, input model.AddRevisionInput) (*model.RecipeRevision, error)
}

type recipeService struct {
	queries     *db.Queries
	conn        *pgxpool.Pool
	authService auth.AuthService
}

// GetRevisionRecipe implements RecipeService.
func (r recipeService) GetRevisionRecipe(ctx context.Context, id uuid.UUID) (*model.Recipe, error) {
	uuid := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := r.queries.GetRecipeByRevisionID(ctx, uuid)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch recipe for revision %s: %w", id, err)
	}
	return model.RecipeFromDBType(result), nil
}

// GetIngredientForRecipeIngredient implements RecipeService.
func (r recipeService) GetIngredientForRecipeIngredient(ctx context.Context, id int64) (*model.Ingredient, error) {
	result, err := r.queries.GetIngredientFromRecipeIngredientId(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.IngredientFromDBType(result), nil
}

// GetUnitForRecipeIngredient implements RecipeService.
func (r recipeService) GetUnitForRecipeIngredient(ctx context.Context, id int64) (*model.MeasurementUnit, error) {
	result, err := r.queries.GetMeasurementUnitFromIngredientId(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.MeasurementUnitFromDBType(result), nil
}

// GetRevisionForIngredient implements RecipeService.
func (r recipeService) GetRevisionForIngredient(ctx context.Context, id int64) (*model.RecipeRevision, error) {
	result, err := r.queries.GetRecipeRevisionByIngredientId(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.RevisionFromDBType(result), nil
}

// GetRevisionForStep implements RecipeService.
func (r recipeService) GetRevisionForStep(ctx context.Context, id int64) (*model.RecipeRevision, error) {
	result, err := r.queries.GetRecipeRevisionByStepId(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.RevisionFromDBType(result), nil
}

// GetRevisionParent implements RecipeService.
func (r recipeService) GetRevisionParent(ctx context.Context, id uuid.UUID) (*model.RecipeRevision, error) {
	uuid := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := r.queries.GetRecipeRevisionByParentID(ctx, uuid)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch parent for revision %s: %w", id, err)
	}
	return model.RevisionFromDBType(result), nil
}

// GetRecipeFeaturedRevision implements RecipeService.
func (r recipeService) GetRecipeFeaturedRevision(ctx context.Context, id uuid.UUID) (*model.RecipeRevision, error) {
	uuid := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	data, err := r.queries.GetFeaturedRevisionByRecipeId(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch recipe with featured revision: %w", err)
	}

	return model.RevisionFromDBType(data), nil
}

// GetRecipeForkedFromRevision implements RecipeService.
func (r recipeService) GetRecipeForkedFromRevision(ctx context.Context, id uuid.UUID) (*model.RecipeRevision, error) {
	uuid := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	data, err := r.queries.GetForkedFromRevisionByRecipeId(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch recipe with forkedFrom revision: %w", err)
	}

	return model.RevisionFromDBType(data), nil
}

// AddRevision implements RecipeService.
func (r recipeService) AddRevision(ctx context.Context, input model.AddRevisionInput) (*model.RecipeRevision, error) {
	if input.Revision == nil {
		// TODO: Write an actual error here
		return nil, nil
	}
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	qtx := r.queries.WithTx(tx)
	defer tx.Rollback(ctx) //nolint:all
	recipeParams := db.UpdateRecipeParams{
		ID: pgtype.UUID{
			Bytes: input.ID,
			Valid: true,
		},
		Slug: input.Slug,
	}
	recipe, err := qtx.UpdateRecipe(ctx, recipeParams)
	if err != nil {
		return nil, err
	}
	revisionParams := db.CreateRevisionParams{
		RecipeID: recipe.ID,
		Title:    input.Revision.Title,
		ParentID: pgtype.UUID{
			Bytes: input.Parent,
			Valid: true,
		},
	}

	if input.Revision.Description != nil {
		revisionParams.RecipeDescription = pgtype.Text{
			String: *input.Revision.Description,
			Valid:  true,
		}
	}
	if input.Revision.ChangeComment != nil {
		revisionParams.ChangeComment = pgtype.Text{
			String: *input.Revision.ChangeComment,
			Valid:  true,
		}
	}
	revision, err := qtx.CreateRevision(ctx, revisionParams)
	if err != nil {
		return nil, err
	}
	for _, ingredient := range input.Revision.Ingredients {
		if ingredient == nil {
			// TODO: Write an actual error here
			return nil, nil
		}
		db_ingredient, err := qtx.UpsertIngredient(ctx, ingredient.Ingredient)
		if err != nil {
			return nil, err
		}
		db_unit, err := qtx.UpsertMeasurement(ctx, ingredient.Unit)
		if err != nil {
			return nil, err
		}
		params := db.CreateRevisionIngredientParams{
			RevisionID:        revision.ID,
			IngredientID:      db_ingredient.ID,
			MeasurementUnitID: db_unit.ID,
			Quantity:          float32(ingredient.Quantity),
		}
		_, err = qtx.CreateRevisionIngredient(ctx, params)
		if err != nil {
			return nil, err
		}
	}
	for _, step := range input.Revision.Steps {
		if step == nil {
			// TODO: Write an actual error here
			return nil, nil
		}
		params := db.CreateRevisionStepParams{
			RevisionID: revision.ID,
			Content:    step.Instruction,
			Index:      int32(step.Step),
		}
		_, err = qtx.CreateRevisionStep(ctx, params)
		if err != nil {
			return nil, err
		}
	}
	return model.RevisionFromDBType(revision), tx.Commit(ctx)
}

// CreateRecipe implements RecipeService.
func (r recipeService) CreateRecipe(ctx context.Context, input model.CreateRecipeInput) (*model.Recipe, error) {
	if input.Revision == nil {
		// TODO: Write an actual error here
		return nil, nil
	}
	user, _ := r.authService.GetUserSessionFromCtx(ctx)
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	qtx := r.queries.WithTx(tx)
	defer tx.Rollback(ctx) //nolint:all
	var forkdFrom pgtype.UUID
	if input.ForkdFrom != nil {
		forkdFrom.Bytes = *input.ForkdFrom
		forkdFrom.Valid = true
	}
	recipeParams := db.CreateRecipeParams{
		AuthorID:   user.ID,
		ForkedFrom: forkdFrom,
		Slug:       input.Slug,
		Private:    input.Private,
	}
	recipe, err := qtx.CreateRecipe(ctx, recipeParams)
	if err != nil {
		return nil, err
	}
	revisionParams := db.CreateRevisionParams{
		RecipeID: recipe.ID,
		Title:    input.Revision.Title,
	}

	if input.Revision.Description != nil {
		revisionParams.RecipeDescription = pgtype.Text{
			String: *input.Revision.Description,
			Valid:  true,
		}
	}
	if input.Revision.ChangeComment != nil {
		revisionParams.ChangeComment = pgtype.Text{
			String: *input.Revision.ChangeComment,
			Valid:  true,
		}
	}
	revision, err := qtx.CreateRevision(ctx, revisionParams)
	if err != nil {
		return nil, err
	}
	for _, ingredient := range input.Revision.Ingredients {
		if ingredient == nil {
			// TODO: Write an actual error here
			return nil, nil
		}
		db_ingredient, err := qtx.UpsertIngredient(ctx, ingredient.Ingredient)
		if err != nil {
			return nil, err
		}
		db_unit, err := qtx.UpsertMeasurement(ctx, ingredient.Unit)
		if err != nil {
			return nil, err
		}
		params := db.CreateRevisionIngredientParams{
			RevisionID:        revision.ID,
			IngredientID:      db_ingredient.ID,
			MeasurementUnitID: db_unit.ID,
			Quantity:          float32(ingredient.Quantity),
		}
		_, err = qtx.CreateRevisionIngredient(ctx, params)
		if err != nil {
			return nil, err
		}
	}
	for _, step := range input.Revision.Steps {
		if step == nil {
			// TODO: Write an actual error here
			return nil, nil
		}
		params := db.CreateRevisionStepParams{
			RevisionID: revision.ID,
			Content:    step.Instruction,
			Index:      int32(step.Step),
		}
		_, err = qtx.CreateRevisionStep(ctx, params)
		if err != nil {
			return nil, err
		}
	}
	return model.RecipeFromDBType(recipe), tx.Commit(ctx)
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
func (r recipeService) GetRevisionById(ctx context.Context, id uuid.UUID) (*model.RecipeRevision, error) {
	pgId := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := r.queries.GetRecipeRevisionById(ctx, pgId)
	return util.HandleNoRowsOnNullableType(result, err, model.RevisionFromDBType)
}

// ListRecipeIngredientsForRevision implements RecipeService.
func (r recipeService) ListRecipeIngredients(ctx context.Context, id uuid.UUID) ([]*model.RecipeIngredient, error) {
	uuid := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := r.queries.ListIngredientsByRecipeRevisionID(ctx, uuid)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch ingredients for revision %s: %w", id, err)
	}
	return model.ListIngredientsFromDBType(result), nil
}

// ListRecipeStepsForRevision implements RecipeService.
func (r recipeService) ListRecipeSteps(ctx context.Context, id uuid.UUID) ([]*model.RecipeStep, error) {
	uuid := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := r.queries.ListStepsByRecipeRevisionID(ctx, uuid)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch steps for revision %s: %w", id, err)
	}
	return model.ListStepsFromDBType(result), nil
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
	user, _ := r.authService.GetUserSessionFromCtx(ctx)
	if user != nil {
		params.CurrentUser = user.ID
	}
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
