ALTER TABLE recipe_revisions
ADD CONSTRAINT title_length_check
CHECK (char_length(trim(title)) >= 3 AND char_length(trim(title)) <= 150);
