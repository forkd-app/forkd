-- Write your up sql migration here
CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  display_name varchar(50) NOT NULL UNIQUE,
  email varchar(255) NOT NULL UNIQUE,
  join_date timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NULL 
);

CREATE TABLE IF NOT EXISTS recipes (
  id bigserial PRIMARY KEY,
  author_id bigint NOT NULL CONSTRAINT fk_recipe_author REFERENCES users(id),
  forked_from bigint CONSTRAINT fk_recipe_fork REFERENCES recipes(id),
  slug varchar(75) NOT NULL UNIQUE,
  private boolean NOT NULL,
  initial_publish_date timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS recipe_revisions (
  id bigserial PRIMARY KEY,
  recipe_id bigint NOT NULL CONSTRAINT fk_recipe_revisions REFERENCES recipes(id),
  parent_id bigint NULL CONSTRAINT fk_recipe_revision_parent REFERENCES recipe_revisions(id),
  -- About section for the recipe.
  recipe_description text,
  --change description should be a note describing how description changed
  change_comment text,
  title text NOT NULL,
  publish_date timestamp NOT NULL DEFAULT now()
);

ALTER TABLE IF EXISTS recipes
ADD featured_revision bigint NULL CONSTRAINT fk_recipe_revision_id REFERENCES recipe_revisions(id);

   
CREATE TABLE IF NOT EXISTS linked_recipes (
  from_recipe_id bigint NOT NULL CONSTRAINT fk_linked_recipes_from REFERENCES recipes(id),
  to_recipe_id bigint NOT NULL CONSTRAINT fk_linked_recipes_to REFERENCES recipes(id),
  CONSTRAINT linked_recipe_pk PRIMARY KEY(from_recipe_id, to_recipe_id)
);
CREATE TABLE IF NOT EXISTS tags (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL UNIQUE,
  description text,
  --True if by user , False if internal keyword
  user_generated boolean NOT NULL
);
CREATE TABLE IF NOT EXISTS measurement_units (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL UNIQUE,
  description text
);
CREATE TABLE IF NOT EXISTS measurement_units_tags (
  id bigserial PRIMARY KEY,
  measurement bigint NOT NULL UNIQUE CONSTRAINT fk_measurement_tags_measurement REFERENCES measurement_units(id),
  tag bigint NOT NULL CONSTRAINT fk_measurement_tags_tag REFERENCES tags(id)
);
CREATE TABLE IF NOT EXISTS ingredients (
  id bigserial PRIMARY KEY,
  name varchar(255) UNIQUE NOT NULL ,
  description text
);
CREATE TABLE IF NOT EXISTS ingredient_tags (
  id bigserial PRIMARY KEY,
  ingredient bigint NOT NULL CONSTRAINT fk_ingredient_tags_ingredient REFERENCES ingredients(id),
  tag bigint NOT NULL CONSTRAINT fk_ingredient_tags_tag REFERENCES tags(id)
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

CREATE TABLE IF NOT EXISTS ratings(
  id bigserial PRIMARY KEY,
  revision_id bigint NOT NULL CONSTRAINT fk_recipe_revision_ingredients REFERENCES recipe_revisions(id),
  --should be 1-5
  star_value smallint CONSTRAINT check_star_Number CHECK(star_value >= 1 AND star_value  <= 5)
);
CREATE TABLE IF NOT EXISTS magic_links(
  id bigserial PRIMARY KEY,
  author_id bigint NOT NULL CONSTRAINT fk_magic_link REFERENCES users(id),
  token varchar (255) NOT NULL,
  expiry timestamp NOT NULL CHECK (expiry > now())
);
CREATE TABLE IF NOT EXISTS sessions(
  id bigserial PRIMARY KEY,
  user_id bigint NOT NULL CONSTRAINT fk_session_user REFERENCES users(id),
  expiry  timestamp NOT NULL CHECK (expiry > now())
);
