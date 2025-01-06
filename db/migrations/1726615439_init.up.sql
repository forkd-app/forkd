-- Write your up sql migration here
CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  username varchar(50) NOT NULL UNIQUE,
  email varchar(255) NOT NULL UNIQUE,
  join_date timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NULL 
);
CREATE TABLE IF NOT EXISTS recipes (
  id bigserial PRIMARY KEY,
  featured_revision bigint NOT NULL CONSTRAINT fk_recipe_revision_id REFERENCES recipe_revisions(id),
  author_id bigint NOT NULL CONSTRAINT fk_recipe_author REFERENCES users(id),
  forked_from bigint CONSTRAINT fk_recipe_fork REFERENCES recipes(id),
  slug varchar(75) NOT NULL UNIQUE,
  private boolean NOT NULL,
  description text,
  initial_publish_date timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS recipe_revisions (
  id bigserial PRIMARY KEY,
  recipe_id bigint NOT NULL CONSTRAINT fk_recipe_revisions REFERENCES recipes(id),
  parent_id bigint NOT NULL CONSTRAINT fk_recipe_revision_parent REFERENCES recipe_revisions(id),
  child_id bigint CONSTRAINT fk_recipe_revision_child REFERENCES recipe_revisions(id),
  -- Free form content, maybe like an "about" section. Maybe this should be like explaining the changes made
  description text,
  title text,
  publish_date timestamp NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS linked_recipes (
  from_recipe_id bigint NOT NULL CONSTRAINT fk_linked_recipes_from REFERENCES recipes(id),
  to_recipe_id bigint NOT NULL CONSTRAINT fk_linked_recipes_to REFERENCES recipes(id),
  CONSTRAINT linked_recipe_pk PRIMARY KEY(from_recipe_id, to_recipe_id)
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
CREATE TABLE IF NOT EXISTS recipe_steps (
  id bigserial PRIMARY KEY,
  revision_id bigint NOT NULL CONSTRAINT fk_recipe_revision_ingredients REFERENCES recipe_revisions(id),
  content text NOT NULL,
  index int NOT NULL
);
