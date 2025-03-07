import { Grid } from "@mantine/core"
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
      <Grid style={styles.gridContainer} justify="center">
        {recipes.map((recipe) => (
          <Grid.Col
            key={recipe.title}
            style={styles.grid}
            span={{ base: 12, md: 5.5, lg: 3.5 }}
          >
            <RecipeCard {...recipe} />
          </Grid.Col>
        ))}
      </Grid>
    </>
  )
}

const styles = {
  gridContainer: {
    padding: 50,
  },
  grid: {
    padding: 10,
    margin: 10,
    borderWidth: 0,
    borderColor: "black",
    borderStyle: "solid",
    justifyContent: "space-evenly",
    boxShadow: "0px  5px 15px #bfbfbf",
  },
}
