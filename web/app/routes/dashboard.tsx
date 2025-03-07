import { SimpleGrid } from "@mantine/core"
import { RecipeCard } from "../components/recipeCard/recipeCard"

const recipes = [
  { title: "pancakes" },
  { title: "chickpea salad" },
  { title: "chocolate mousse" },
  { title: "rice and beans" },
  { title: "chai latte" },
  { title: "grape leaves" },
]

export default function Dashboard() {
  return (
    <>
      {/* recipe component */}
      <SimpleGrid
        cols={{ base: 1, sm: 2, md: 2, lg: 4 }}
        pb={40}
        pt={40}
        style={{ background: "#fffaf5", width: "90%", margin: "auto" }}
      >
        {recipes.map((recipe) => (
          <div key={recipe.title} style={styles.grid}>
            <RecipeCard {...recipe} />
          </div>
        ))}
      </SimpleGrid>
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
