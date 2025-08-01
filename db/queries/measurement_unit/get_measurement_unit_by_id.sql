-- name: GetMeasurementUnitById :one
SELECT
    measurement_units.id,
    measurement_units.name,
    measurement_units.description
FROM
    measurement_units
WHERE
    measurement_units.id = sqlc.arg('id')
LIMIT 1;
