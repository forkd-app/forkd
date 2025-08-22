import { Center, Loader, Paper, SimpleGrid } from "@mantine/core"
import { ReactNode, useEffect, useState } from "react"
import { ListRecipesQuery } from "~/gql/forkd.g"
import { RecipeCard } from "../recipeCard/recipeCard"
import { useFetch, useIntersection } from "@mantine/hooks"

type Props = {
  recipes: Exclude<ListRecipesQuery["recipe"], null | undefined>["list"]
}

export default function RecipeList({ recipes }: Props): ReactNode {
  const [recipeList, setRecipeList] = useState(recipes.items)
  const [cursor, setCursor] = useState(recipes.pagination.nextCursor)
  const { refetch, data, loading, abort } = useFetch<Props["recipes"]>(
    `/api/recipes/paginate/${cursor ?? ""}`,
    { autoInvoke: false }
  )
  useEffect(() => {
    return () => {
      abort()
    }
  }, [abort])
  useEffect(() => {
    if (data) {
      setRecipeList((r) => r.concat(data.items))
      setCursor(data.pagination.nextCursor)
    }
  }, [data])
  const { ref, entry } = useIntersection({
    rootMargin: "1200px",
  })
  useEffect(() => {
    if (entry?.isIntersecting && !loading && cursor) {
      refetch()
    }
  }, [entry?.isIntersecting, loading, refetch, cursor])
  return (
    <Paper>
      <SimpleGrid
        cols={{ base: 1, sm: 2, md: 2, lg: 4 }}
        pb={40}
        pt={40}
        style={styles.grid}
      >
        {recipeList.map((recipe) => (
          <div key={recipe.id} style={styles.col}>
            <RecipeCard recipe={recipe || {}} />
          </div>
        ))}
      </SimpleGrid>
      {loading ? (
        <Center>
          <Loader size="xl" />{" "}
        </Center>
      ) : (
        <div ref={ref}></div>
      )}
    </Paper>
  )
}

const styles = {
  grid: {
    background: "#fff",
    width: "90%",
    margin: "auto",
  },
  col: {
    padding: 10,
    margin: 10,
    borderWidth: 0,
    borderColor: "black",
    borderStyle: "solid",
    justifyContent: "space-evenly",
    boxShadow: "0px  5px 15px #bfbfbf",
  },
}
