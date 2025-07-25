ALTER TABLE recipes
DROP CONSTRAINT recipes_slug_author_id_key;
ALTER TABLE recipes
ADD CONSTRAINT recipes_slug_key UNIQUE (slug);
