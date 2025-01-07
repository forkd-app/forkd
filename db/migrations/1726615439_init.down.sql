-- Write your down sql migration here
DROP TABLE IF EXISTS magic_links;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS ratings;
DROP TABLE IF EXISTS recipe_ingredients;
DROP TABLE IF EXISTS ingredient_tags;
DROP TABLE IF EXISTS measurement_units_tags;
DROP TABLE IF EXISTS recipe_steps;
ALTER TABLE recipes
DROP COLUMN IF EXISTS featured_revision;
DROP TABLE IF EXISTS recipe_revisions;
DROP TABLE IF EXISTS measurement_units;
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS linked_recipes;
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS users;

