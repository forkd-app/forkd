ALTER TABLE users
ADD COLUMN photo text;

ALTER TABLE recipe_revisions
ADD COLUMN photo text;

ALTER TABLE recipe_steps
ADD COLUMN photo text;
