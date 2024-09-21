-- Write your up sql migration here
CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  username varchar(50) NOT NULL UNIQUE,
  email varchar(255) NOT NULL UNIQUE,
  join_date timestamp NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS recipes (
  id bigserial PRIMARY KEY,
  author_id bigint NOT NULL CONSTRAINT fk_recipe_author REFERENCES users(id),
  forked_from bigint CONSTRAINT fk_recipe_fork REFERENCES recipes(id),
  slug varchar(75) NOT NULL UNIQUE,
  description text,
  initial_publish_date timestamp NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS recipe_comments (
  id bigserial PRIMARY KEY,
  recipe_id bigint NOT NULL CONSTRAINT fk_recipe_comment REFERENCES recipes(id),
  author_id bigint NOT NULL CONSTRAINT fk_recipe_comment_author REFERENCES users(id),
  content text NOT NULL,
  post_date timestamp NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS recipe_revisions (
  id bigserial PRIMARY KEY,
  recipe_id bigint NOT NULL CONSTRAINT fk_recipe_revisions REFERENCES recipes(id),
  -- Free form content, maybe like an "about" section. Maybe this should be like explaining the changes made
  description text,
  publish_date timestamp NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS recipe_revision_delta (
  id bigserial PRIMARY KEY,
  -- TODO: Add check constraint that these are not the same
  from_recipe_revision_id bigint NOT NULL CONSTRAINT fk_recipe_revision_deltas_to REFERENCES recipe_revisions(id),
  to_recipe_revision_id bigint NOT NULL CONSTRAINT fk_recipe_revision_deltas_from REFERENCES recipe_revisions(id)
  -- TODO: Figure out delta shape. Maybe something like:
  -- added: new ingredients
  -- removed: removed ingredients
  -- changed: ingredients that have changes to measurements (either quantity and/or unit) or the comment
);
CREATE TABLE IF NOT EXISTS tags (
  name varchar(255) PRIMARY KEY,
  description text
);
CREATE TABLE IF NOT EXISTS measurement_units (
  name varchar(255) PRIMARY KEY,
  description text
);
CREATE TABLE IF NOT EXISTS measurement_units_tags (
  measurement varchar(255) NOT NULL CONSTRAINT fk_measurement_tags_measurement REFERENCES measurement_units(name),
  tag varchar(255) NOT NULL CONSTRAINT fk_measurement_tags_tag REFERENCES tags(name)
);
CREATE TABLE IF NOT EXISTS ingredients (
  name varchar(255) PRIMARY KEY,
  description text
);
CREATE TABLE IF NOT EXISTS ingredient_tags (
  ingredient varchar(255) NOT NULL CONSTRAINT fk_ingredient_tags_ingredient REFERENCES ingredients(name),
  tag varchar(255) NOT NULL CONSTRAINT fk_ingredient_tags_tag REFERENCES tags(name)
);
CREATE TABLE IF NOT EXISTS recipe_ingredients (
  id bigserial PRIMARY KEY,
  revision_id bigint NOT NULL CONSTRAINT fk_recipe_revision_ingredients REFERENCES recipe_revisions(id),
  ingredient varchar(255) NOT NULL CONSTRAINT fk_recipe_ingredient REFERENCES ingredients(name),
  quantity real NOT NULL,
  unit varchar(255) NOT NULL CONSTRAINT fk_recipe_ingredient_quantity REFERENCES measurement_units(name),
  comment text
);
