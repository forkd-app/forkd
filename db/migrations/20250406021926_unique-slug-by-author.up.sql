ALTER TABLE recipes
DROP CONSTRAINT recipes_slug_key;
ALTER TABLE recipes
ADD CONSTRAINT recipes_slug_author_id_key UNIQUE (slug, author_id);
