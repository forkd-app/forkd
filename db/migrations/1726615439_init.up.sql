-- Write your up sql migration here
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    display_name varchar(50) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    join_date timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS recipes (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    author_id uuid NOT NULL REFERENCES users (id),
    slug varchar(75) NOT NULL UNIQUE,
    private boolean NOT NULL,
    initial_publish_date timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS recipe_revisions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    recipe_id uuid NOT NULL REFERENCES recipes (id),
    parent_id uuid REFERENCES recipe_revisions (id),
    -- About section for the recipe.
    recipe_description text,
    --change description should be a note describing how description changed
    change_comment text,
    title text NOT NULL,
    publish_date timestamp NOT NULL DEFAULT now()
);

ALTER TABLE IF EXISTS recipes
ADD forked_from uuid REFERENCES recipe_revisions (id),
ADD featured_revision uuid REFERENCES recipe_revisions (id);

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
    measurement_unit_id bigint NOT NULL REFERENCES measurement_units (id),
    tag_id bigint NOT NULL REFERENCES tags (id),
    PRIMARY KEY (measurement_unit_id, tag_id)
);

CREATE TABLE IF NOT EXISTS ingredients (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    description text
);

CREATE TABLE IF NOT EXISTS ingredient_tags (
    ingredient_id bigint NOT NULL REFERENCES ingredients (id),
    tag_id bigint NOT NULL REFERENCES tags (id),
    PRIMARY KEY (ingredient_id, tag_id)
);

CREATE TABLE IF NOT EXISTS recipe_ingredients (
    id bigserial PRIMARY KEY,
    revision_id uuid NOT NULL REFERENCES recipe_revisions (id),
    ingredient_id bigint NOT NULL REFERENCES ingredients (id),
    quantity real NOT NULL,
    measurement_unit_id bigint NOT NULL REFERENCES measurement_units (id),
    comment text
);

CREATE TABLE IF NOT EXISTS recipe_steps (
    id bigserial PRIMARY KEY,
    revision_id uuid NOT NULL REFERENCES recipe_revisions (id),
    content text NOT NULL,
    index int NOT NULL
);

CREATE TABLE IF NOT EXISTS ratings (
    revision_id uuid NOT NULL REFERENCES recipe_revisions (id),
    user_id uuid NOT NULL REFERENCES users (id),
    --should be 1-5
    star_value smallint CONSTRAINT check_star_number CHECK (
        star_value >= 1 AND star_value <= 5
    ),
    PRIMARY KEY (revision_id, user_id)
);

CREATE TABLE IF NOT EXISTS magic_links (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL REFERENCES users (id),
    token uuid NOT NULL DEFAULT gen_random_uuid(),
    expiry timestamp NOT NULL CHECK (expiry > now())
);

CREATE TABLE IF NOT EXISTS sessions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL REFERENCES users (id),
    expiry timestamp NOT NULL CHECK (expiry > now())
);
