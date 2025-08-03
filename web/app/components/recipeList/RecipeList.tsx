import { SimpleGrid } from "@mantine/core"
import { ReactNode } from "react"
import { ListRecipesQuery } from "~/gql/forkd.g"
import { RecipeCard } from "../recipeCard/recipeCard"

type Props = {
  recipes: Exclude<ListRecipesQuery["recipe"], null | undefined>["list"]
}

export default function RecipeList({ recipes }: Props): ReactNode {
  return (
    <>
      {/* recipe component */}
      <SimpleGrid
        cols={{ base: 1, sm: 2, md: 2, lg: 4 }}
        pb={40}
        pt={40}
        style={styles.grid}
      >
        {recipes?.items.map((recipe) => (
          <div key={recipe.id} style={styles.col}>
            <RecipeCard recipe={recipe || {}} />
          </div>
        ))}
      </SimpleGrid>
    </>
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
