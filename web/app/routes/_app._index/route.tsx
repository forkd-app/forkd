import { SimpleGrid } from "@mantine/core"
import { RecipeCard } from "../../components/recipeCard/recipeCard"
import { MetaFunction } from "@remix-run/react"

export const meta: MetaFunction = () => {
  return [
    { title: "Forkd App" },
    {
      name: "Create new recipes and add your spin on existing recipes",
      content: "Welcome to Forkd!",
    },
  ]
}

const recipes = [
  { title: "pancakes" },
  { title: "chickpea salad" },
  { title: "chocolate mousse" },
  { title: "rice and beans" },
  { title: "chai latte" },
  { title: "grape leaves" },
]

export default function Index() {
  return (
    <>
      {/* recipe component */}
      <SimpleGrid
        cols={{ base: 1, sm: 2, md: 2, lg: 4 }}
        pb={40}
        pt={40}
        style={styles.grid}
      >
        {recipes.map((recipe) => (
          <div key={recipe.title} style={styles.col}>
            <RecipeCard recipe={recipe} />
          </div>
        ))}
      </SimpleGrid>
    </>
  )
}

const styles = {
  grid: {
    background: "#fffaf5",
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
