package main

import (
	"context"
	"fmt"
	"forkd/db"
	"forkd/util"
	"log"
	"math/rand/v2"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	fakeData "github.com/brianvoe/gofakeit/v7/data"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

const USER_COUNT = 15
const RECIPE_PER_USER_MIN = 0
const RECIPE_PER_USER_MAX = 8
const REVISION_PER_RECIPE_MIN = 1
const REVISION_PER_RECIPE_MAX = 6
const INGREDIENT_COUNT_MIN = 3
const INGREDIENT_COUNT_MAX = 8
const STEP_COUNT_MIN = 2
const STEP_COUNT_MAX = 5

const RECIPE_REVISION_DESCRIPTION_MIN = 4
const RECIPE_REVISION_DESCRIPTION_MAX = 10
const RECIPE_REVISION_CHANGE_COMMENT_MIN = 2
const RECIPE_REVISION_CHANGE_COMMENT_MAX = 5

const USER_PHOTO_CHANCE = 4.0 / 5.0
const RECIPE_PHOTO_CHANCE = 7.0 / 8.0
const FORK_CHANCE = 1.0 / 3.0
const PRIVATE_CHANCE = 1.0 / 16.0
const REVISION_DIVERGE_CHANCE = 1.0 / 8.0
const REVISION_DESCRIPTION_CHANCE = 3.0 / 4.0
const REVISION_CHANGE_COMMENT_CHANCE = 1.0 / 2.0
const INGREDIENT_COMMENT_CHANCE = 1.0 / 6.0
const REVISION_PHOTO_CHANGE_CHANCE = 1.0 / 16.0
const REVISION_TITLE_CHANGE_CHANCE = 1.0 / 20.0
const REVISION_DESCRIPTION_CHANGE_CHANCE = 1.0 / 14.0
const REVISION_ADD_INGREDIENT_CHANCE = 1.0 / 3.0
const REVISION_REMOVE_INGREDIENT_CHANCE = 1.0 / 4.0
const REVISION_ADD_STEP_CHANCE = 1.0 / 8.0
const REVISION_REMOVE_STEP_CHANCE = 1.0 / 9.0
const REVISION_CHANGE_STEP_CHANCE = 1.0 / 10.0
const INGREDIENT_QUANTITY_CHANGE_CHANCE = 2.0 / 3.0
const INGREDIENT_CHANGE_CHANCE = 2.0 / 3.0
const INGREDIENT_COMMENT_CHANGE_CHANCE = 1.0 / 5.0
const UNIT_CHANGE_CHANCE = 2.0 / 3.0

type RevisionWithIngredients struct {
	revision    db.RecipeRevision
	ingredients []db.RecipeIngredient
	steps       []db.RecipeStep
}

type RecipeWithRevisions struct {
	recipe    db.Recipe
	revisions []RevisionWithIngredients
}

const MINUS_ONE_YEAR = -12 * 24 * 365 * time.Hour

func createUsers(ctx context.Context, tx pgx.Tx, qtx *db.Queries) []db.User {
	log.Println("creating users")
	users := make([]db.User, USER_COUNT)
	for i := 0; i < USER_COUNT; i++ {
		user, err := qtx.SeedUser(ctx, db.SeedUserParams{
			DisplayName: gofakeit.Username(),
			Email:       gofakeit.Email(),
			JoinDate: pgtype.Timestamp{
				Time:  gofakeit.DateRange(time.Now().Add(MINUS_ONE_YEAR), time.Now()),
				Valid: true,
			},
		})
		if err != nil {
			// I know I shouldn't be ignoring this but whatever lol
			_ = tx.Rollback(ctx)
			log.Panicf("error inserting user: %s", err.Error())
		}
		users[i] = db.User(user)

		if flipCoin(USER_PHOTO_CHANCE) {
			updatedUser, err := qtx.UpdateUser(ctx, db.UpdateUserParams{
				ID:          user.ID,
				DisplayName: user.DisplayName,
				Email:       user.Email,
				Photo: pgtype.Text{
					Valid:  true,
					String: fmt.Sprintf("https://picsum.photos/id/%d/500", randBetween(0, 255)),
				},
			})
			if err != nil {
				// I know I shouldn't be ignoring this but whatever lol
				_ = tx.Rollback(ctx)
				log.Panicf("error updating user: %s", err.Error())
			}

			users[i] = updatedUser
		}
		log.Printf("inserted user: %s\n", user.DisplayName)
	}
	return users
}

func main() {
	log.Println("starting seed")
	util.InitEnv()
	env := util.GetEnv()
	dbConnStr := env.GetDbConnStr()
	log.Println("loaded env")

	queries, conn, err := db.GetQueriesWithConnection(dbConnStr)
	if err != nil || queries == nil {
		log.Panicf("unable to connect to db: %s", err.Error())
	}
	defer conn.Close()
	log.Println("connected to db")

	ctx := context.Background()
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Panicf("error starting transaction: %s", err.Error())
	}
	qtx := queries.WithTx(tx)
	log.Println("started transaction")

	log.Println("creating ingredients and units")
	ingredients := make(map[string]int64)
	ingredientNames := make([]string, 0)
	for _, fruit := range fakeData.Food["fruit"] {
		row, err := qtx.UpsertIngredient(ctx, fruit)
		if err != nil {
			// I know I shouldn't be ignoring this but whatever lol
			_ = tx.Rollback(ctx)
			log.Panicf("error upserting ingredient: %s", err.Error())
		}
		ingredients[fruit] = row.ID
		ingredientNames = append(ingredientNames, fruit)
		log.Printf("inserted ingredient: %s\n", fruit)
	}

	for _, veg := range fakeData.Food["vegetable"] {
		row, err := qtx.UpsertIngredient(ctx, veg)
		if err != nil {
			// I know I shouldn't be ignoring this but whatever lol
			_ = tx.Rollback(ctx)
			log.Panicf("error upserting ingredient: %s", err.Error())
		}
		ingredients[veg] = row.ID
		ingredientNames = append(ingredientNames, veg)
		log.Printf("inserted ingredient: %s\n", veg)
	}

	metricPrefixes := []string{"milli", "", "kilo"}
	matricUnits := []string{"gram", "liter"}
	freedomUnits := []string{"teaspoon", "tablespoon", "cup", "ounce", "pound"}
	units := make(map[string]int64)
	unitNames := make([]string, 0)
	for _, unit := range freedomUnits {
		row, err := qtx.UpsertMeasurementUnit(ctx, unit)
		if err != nil {
			// I know I shouldn't be ignoring this but whatever lol
			_ = tx.Rollback(ctx)
			log.Panicf("error upserting measurement: %s", err.Error())
		}
		units[unit] = row.ID
		unitNames = append(unitNames, unit)
		log.Printf("inserted unit: %s\n", unit)
	}

	for _, prefix := range metricPrefixes {
		for _, unit := range matricUnits {
			unit = prefix + unit
			row, err := qtx.UpsertMeasurementUnit(ctx, unit)
			if err != nil {
				// I know I shouldn't be ignoring this but whatever lol
				_ = tx.Rollback(ctx)
				log.Panicf("error upserting measurement: %s", err.Error())
			}
			units[unit] = row.ID
			unitNames = append(unitNames, unit)
			log.Printf("inserted unit: %s\n", unit)
		}
	}
	log.Println("inserted ingredients and units")

	users := createUsers(ctx, tx, qtx)
	log.Println("inserting recipes")
	recipes := make([]RecipeWithRevisions, 0)
	for _, user := range users {
		recipeCount := randBetween(RECIPE_PER_USER_MIN, RECIPE_PER_USER_MAX)
		log.Printf("inserting %d recipes for user %s", recipeCount, user.DisplayName)
		userRecipes := make([]RecipeWithRevisions, 0)
		for i := 0; i < recipeCount; i++ {
			var forkdFrom RevisionWithIngredients
			title := getRandomFood()
			recipeParams := db.SeedRecipeParams{
				AuthorID: user.ID,
				Slug:     strings.ToLower(fmt.Sprintf("%s/%s-%d", url.PathEscape(user.DisplayName), url.PathEscape(strings.ReplaceAll(title, " ", "-")), i)),
				Private:  flipCoin(PRIVATE_CHANCE),
				InitialPublishDate: pgtype.Timestamp{
					Time:  gofakeit.DateRange(user.JoinDate.Time, time.Now()),
					Valid: true,
				},
			}
			if flipCoin(FORK_CHANCE) && len(recipes) > 0 {
				recipe := getRandomVal(recipes)
				if len(recipe.revisions) > 1 {
					forkdFrom = getRandomVal(recipe.revisions)
					recipeParams.ForkedFrom = forkdFrom.revision.ID
					log.Printf("forking from %s", uuid.UUID(forkdFrom.revision.ID.Bytes).String())
				}
			}
			if recipeParams.ForkedFrom.Valid {
				recipeParams.InitialPublishDate = pgtype.Timestamp{
					Time:  gofakeit.DateRange(forkdFrom.revision.PublishDate.Time, time.Now()),
					Valid: true,
				}
			}
			recipe, err := qtx.SeedRecipe(ctx, recipeParams)
			if err != nil {
				// I know I shouldn't be ignoring this but whatever lol
				_ = tx.Rollback(ctx)
				log.Panicf("error creating recipe %x", err)
			}

			revisionCount := randBetween(REVISION_PER_RECIPE_MIN, REVISION_PER_RECIPE_MAX)
			revisions := make([]RevisionWithIngredients, 0)
			if forkdFrom.revision.ID.Valid {
				initialRevision, err := qtx.SeedRevision(ctx, db.SeedRevisionParams{
					RecipeID:          forkdFrom.revision.RecipeID,
					RecipeDescription: forkdFrom.revision.RecipeDescription,
					ChangeComment:     forkdFrom.revision.ChangeComment,
					Photo:             forkdFrom.revision.Photo,
					PublishDate:       recipe.InitialPublishDate,
				})
				if err != nil {
					// I know I shouldn't be ignoring this but whatever lol
					_ = tx.Rollback(ctx)
					log.Panicf("error creating initial revision for forked recipe %x", err)
				}

				revision := RevisionWithIngredients{
					revision:    initialRevision,
					ingredients: make([]db.RecipeIngredient, 0),
					steps:       make([]db.RecipeStep, 0),
				}
				for _, ingredient := range forkdFrom.ingredients {
					ingredient, err := qtx.CreateRecipeIngredient(ctx, db.CreateRecipeIngredientParams{
						RevisionID:        initialRevision.ID,
						IngredientID:      ingredient.IngredientID,
						MeasurementUnitID: ingredient.MeasurementUnitID,
						Quantity:          ingredient.Quantity,
						Comment:           ingredient.Comment,
					})
					if err != nil {
						// I know I shouldn't be ignoring this but whatever lol
						_ = tx.Rollback(ctx)
						log.Panicf("error creating ingredient for forked recipe %x", err)
					}
					revision.ingredients = append(revision.ingredients, ingredient)
				}

				for _, step := range forkdFrom.steps {
					step, err := qtx.CreateRecipeStep(ctx, db.CreateRecipeStepParams{
						RevisionID: forkdFrom.revision.ID,
						Content:    step.Content,
						Index:      step.Index,
					})
					if err != nil {
						// I know I shouldn't be ignoring this but whatever lol
						_ = tx.Rollback(ctx)
						log.Panicf("error creating step for forked recipe %x", err)
					}
					revision.steps = append(revision.steps, step)
				}
				revisions = append(revisions, revision)
				log.Printf("inserted intial revision for forked recipe %s", uuid.UUID(revision.revision.ID.Bytes).String())
			}
			log.Printf("inserted recipe with slug %s", recipe.Slug)
			log.Printf("inserting %d revisions for recipe %s", revisionCount, recipe.Slug)
			for j := 0; j < revisionCount; j++ {
				revisionWithIngredients := RevisionWithIngredients{
					revision:    db.RecipeRevision{},
					ingredients: make([]db.RecipeIngredient, 0),
					steps:       make([]db.RecipeStep, 0),
				}
				var parent RevisionWithIngredients
				revisionParams := db.SeedRevisionParams{
					Title:    title,
					RecipeID: recipe.ID,
					PublishDate: pgtype.Timestamp{
						Valid: true,
						Time:  gofakeit.DateRange(recipe.InitialPublishDate.Time, time.Now()),
					},
				}

				if j > 0 && len(revisions) > 0 {
					if j > 2 && len(revisions) > 2 && flipCoin(REVISION_DIVERGE_CHANCE) {
						parent = getRandomVal(revisions[0 : len(revisions)-2])
						revisionParams.ParentID = parent.revision.ID
					} else {
						parent = revisions[j-1]
						revisionParams.ParentID = parent.revision.ID
					}

					revisionParams.PublishDate.Time = gofakeit.DateRange(parent.revision.PublishDate.Time, time.Now())

					if flipCoin(REVISION_CHANGE_COMMENT_CHANCE) {
						revisionParams.ChangeComment = pgtype.Text{
							Valid:  true,
							String: gofakeit.Paragraph(1, randBetween(RECIPE_REVISION_CHANGE_COMMENT_MIN, RECIPE_REVISION_CHANGE_COMMENT_MAX), 10, "\n"),
						}
					}

					if flipCoin(REVISION_PHOTO_CHANGE_CHANCE) {
						revisionParams.Photo = pgtype.Text{
							Valid:  true,
							String: fmt.Sprintf("https://picsum.photos/id/%d/500", randBetween(0, 255)),
						}
					}

					if flipCoin(REVISION_DESCRIPTION_CHANGE_CHANCE) {
						revisionParams.RecipeDescription = pgtype.Text{
							Valid:  true,
							String: gofakeit.Paragraph(1, randBetween(RECIPE_REVISION_DESCRIPTION_MIN, RECIPE_REVISION_DESCRIPTION_MAX), 10, "\n"),
						}
					}

					if flipCoin(REVISION_TITLE_CHANGE_CHANCE) {
						revisionParams.Title = getRandomFood()
					}

					revision, err := qtx.SeedRevision(ctx, revisionParams)
					if err != nil {
						// I know I shouldn't be ignoring this but whatever lol
						_ = tx.Rollback(ctx)
						log.Panicln(err)
					}

					revisionWithIngredients.revision = revision
				} else {
					if flipCoin(REVISION_DESCRIPTION_CHANCE) {
						revisionParams.RecipeDescription = pgtype.Text{
							Valid:  true,
							String: gofakeit.Paragraph(1, randBetween(RECIPE_REVISION_DESCRIPTION_MIN, RECIPE_REVISION_DESCRIPTION_MAX), 10, "\n"),
						}
					}

					if flipCoin(RECIPE_PHOTO_CHANCE) {
						revisionParams.Photo = pgtype.Text{
							Valid:  true,
							String: fmt.Sprintf("https://picsum.photos/id/%d/500", randBetween(0, 255)),
						}
					}

					revision, err := qtx.SeedRevision(ctx, revisionParams)
					if err != nil {
						// I know I shouldn't be ignoring this but whatever lol
						_ = tx.Rollback(ctx)
						log.Panicf("error creating revision %x", err)
					}

					revisionWithIngredients.revision = revision
				}

				if parent.revision.ID.Valid {
					revisionSteps := parent.steps
					revisionIngredients := parent.ingredients

					if flipCoin(REVISION_REMOVE_STEP_CHANCE) {
						idx := randBetween(0, len(revisionSteps))
						revisionSteps = append(revisionSteps[:idx], revisionSteps[idx+1:]...)
					}

					for _, step := range revisionSteps {
						if flipCoin(REVISION_CHANGE_STEP_CHANCE) {
							step.Content = gofakeit.Paragraph(1, 4, 10, "\n")
						}

						step, err := qtx.CreateRecipeStep(ctx, db.CreateRecipeStepParams{
							RevisionID: revisionWithIngredients.revision.ID,
							Content:    step.Content,
							Index:      step.Index,
						})
						if err != nil {
							// I know I shouldn't be ignoring this but whatever lol
							_ = tx.Rollback(ctx)
							log.Panicf("error creating child revision step %x", err)
						}

						revisionWithIngredients.steps = append(revisionWithIngredients.steps, step)
					}

					if flipCoin(REVISION_ADD_STEP_CHANCE) {
						stepParams := db.CreateRecipeStepParams{
							RevisionID: revisionWithIngredients.revision.ID,
							Content:    gofakeit.Paragraph(1, 4, 10, "\n"),
							Index:      int32(len(revisionWithIngredients.steps)),
						}
						step, err := qtx.CreateRecipeStep(ctx, stepParams)
						if err != nil {
							// I know I shouldn't be ignoring this but whatever lol
							_ = tx.Rollback(ctx)
							log.Panicf("error adding step to child revision %x", err)
						}

						revisionWithIngredients.steps = append(revisionWithIngredients.steps, step)
					}

					if flipCoin(REVISION_REMOVE_INGREDIENT_CHANCE) {
						idx := randBetween(0, len(revisionIngredients))
						revisionIngredients = append(revisionIngredients[:idx], revisionIngredients[idx+1:]...)
					}

					for _, ingredient := range revisionIngredients {
						if flipCoin(INGREDIENT_QUANTITY_CHANGE_CHANCE) {
							ingredient.Quantity = gofakeit.Float32()
						}

						if flipCoin(INGREDIENT_CHANGE_CHANCE) {
							ingredient.IngredientID = ingredients[getRandomVal(ingredientNames)]
						}

						if flipCoin(UNIT_CHANGE_CHANCE) {
							ingredient.MeasurementUnitID = units[getRandomVal(unitNames)]
						}

						if flipCoin(INGREDIENT_COMMENT_CHANGE_CHANCE) {
							if flipCoin(1.0 / 2.0) {
								ingredient.Comment = pgtype.Text{
									Valid:  true,
									String: gofakeit.Sentence(3),
								}
							} else {
								ingredient.Comment = pgtype.Text{
									Valid: false,
								}
							}
						}

						ingredient, err := qtx.CreateRecipeIngredient(ctx, db.CreateRecipeIngredientParams{
							RevisionID:        revisionWithIngredients.revision.ID,
							IngredientID:      ingredient.IngredientID,
							MeasurementUnitID: ingredient.MeasurementUnitID,
							Quantity:          ingredient.Quantity,
							Comment:           ingredient.Comment,
						})
						if err != nil {
							// I know I shouldn't be ignoring this but whatever lol
							_ = tx.Rollback(ctx)
							log.Panicf("error creating ingredient for child revision %x", err)
						}

						revisionWithIngredients.ingredients = append(revisionWithIngredients.ingredients, ingredient)
					}

					if flipCoin(REVISION_ADD_INGREDIENT_CHANCE) {
						ingredientParams := db.CreateRecipeIngredientParams{
							RevisionID:        revisionWithIngredients.revision.ID,
							IngredientID:      ingredients[getRandomVal(ingredientNames)],
							MeasurementUnitID: units[getRandomVal(unitNames)],
							Quantity:          gofakeit.Float32(),
						}

						if flipCoin(INGREDIENT_COMMENT_CHANCE) {
							ingredientParams.Comment = pgtype.Text{
								Valid:  true,
								String: gofakeit.Sentence(3),
							}
						}

						ingredient, err := qtx.CreateRecipeIngredient(ctx, ingredientParams)
						if err != nil {
							// I know I shouldn't be ignoring this but whatever lol
							_ = tx.Rollback(ctx)
							log.Panicf("error adding step to child revision %x", err)
						}

						revisionWithIngredients.ingredients = append(revisionWithIngredients.ingredients, ingredient)
					}
				} else {
					ingredientCount := randBetween(INGREDIENT_COUNT_MIN, INGREDIENT_COUNT_MAX)
					for k := 0; k <= ingredientCount; k++ {
						ingredientParams := db.CreateRecipeIngredientParams{
							RevisionID:        revisionWithIngredients.revision.ID,
							IngredientID:      ingredients[getRandomVal(ingredientNames)],
							MeasurementUnitID: units[getRandomVal(unitNames)],
							Quantity:          gofakeit.Float32(),
						}

						if flipCoin(INGREDIENT_COMMENT_CHANCE) {
							comment := gofakeit.Sentence(3)
							if len(comment) > 75 {
								comment = gofakeit.Sentence(2)
							}
							ingredientParams.Comment = pgtype.Text{
								Valid:  true,
								String: comment,
							}
						}

						ingredient, err := qtx.CreateRecipeIngredient(ctx, ingredientParams)
						if err != nil {
							// I know I shouldn't be ignoring this but whatever lol
							_ = tx.Rollback(ctx)
							log.Panicf("error creating recipe ingredient %x", err)
						}

						revisionWithIngredients.ingredients = append(revisionWithIngredients.ingredients, ingredient)
					}

					stepCount := randBetween(STEP_COUNT_MIN, STEP_COUNT_MAX)
					for k := 0; k <= stepCount; k++ {
						stepParams := db.CreateRecipeStepParams{
							RevisionID: revisionWithIngredients.revision.ID,
							Content:    gofakeit.Sentence(5),
							Index:      int32(k),
						}
						step, err := qtx.CreateRecipeStep(ctx, stepParams)
						if err != nil {
							// I know I shouldn't be ignoring this but whatever lol
							_ = tx.Rollback(ctx)
							log.Panicf("error creating recipe step %x", err)
						}

						revisionWithIngredients.steps = append(revisionWithIngredients.steps, step)
					}
				}

				log.Printf("inserted revision: %s with parent %s", uuid.UUID(revisionWithIngredients.revision.ID.Bytes).String(), uuid.UUID(parent.revision.ID.Bytes).String())
				revisions = append(revisions, revisionWithIngredients)
			}
			userRecipes = append(userRecipes, RecipeWithRevisions{recipe: recipe, revisions: revisions})
		}
		recipes = slices.Concat(recipes, userRecipes)
	}
	_ = tx.Commit(ctx)
}

func randBetween(lower, upper int) int {
	if upper < lower {
		panic(fmt.Sprintf("INVALID RANGE %d, %d", lower, upper))
	}
	return rand.IntN(upper-lower) + lower
}

func flipCoin(odds float32) bool {
	return rand.Float32() <= odds
}

func getRandomVal[T any](arr []T) T {
	size := len(arr)
	if size < 1 {
		panic("EMPTY ARR")
	}

	if size == 1 {
		return arr[0]
	}
	return arr[randBetween(0, size-1)]
}

func getRandomFood() string {
	val := rand.Float32()
	if val >= 0.8 {
		return gofakeit.Breakfast()
	} else if val >= 0.6 {
		return gofakeit.Lunch()
	} else if val >= 0.4 {
		return gofakeit.Dinner()
	} else if val >= 0.2 {
		return gofakeit.Dessert()
	} else {
		return gofakeit.Drink()
	}
}
