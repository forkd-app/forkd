-- name: UpsertIngredient :one
WITH upsert AS (
    INSERT INTO
    ingredients (
        name
    )
    VALUES (
        sqlc.arg('name')
    )
    ON CONFLICT (name)
    DO NOTHING
    RETURNING
        ingredients.id,
        ingredients.name
)

SELECT
    upsert.id,
    upsert.name
FROM
    upsert
UNION
SELECT
    ingredients.id,
    ingredients.name
FROM
    ingredients
WHERE
    ingredients.name = sqlc.arg('name');
