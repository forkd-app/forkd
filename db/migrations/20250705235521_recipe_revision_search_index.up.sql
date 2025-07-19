CREATE INDEX recipe_revisions_search_text
ON recipe_revisions
USING gin (
    (
        setweight(to_tsvector('english', title), 'A')
        || setweight(
            to_tsvector('english', coalesce(recipe_description, '')), 'B'
        )
    )
)
