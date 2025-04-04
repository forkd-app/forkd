import { FC } from "react"
import { Flex, Text, Button, Image, Rating } from "@mantine/core"

interface Props {
  recipe: {
    slug: string
    id: string
    author: { displayName: string }
    featuredRevision: {
      recipeDescription: string | null
      publishDate: string
      rating: number
      title: string
    }
  }
}

export const RecipeCard: FC<Props> = ({ recipe }) => {
  return (
    <Flex
      direction={"column"}
      justify={"space-evenly"}
      style={styles.flexContainer}
    >
      <Image src="images/image.jpg" alt="recipe" />
      <div>
        <Text>{recipe.featuredRevision.title || "No Title"}</Text>
        <Rating defaultValue={recipe.featuredRevision.rating} />
        <Text>
          posted by {recipe.author.displayName || "No Name"} on{" "}
          {recipe.featuredRevision.publishDate || "unknown"}
        </Text>
      </div>
      <div>
        <Text>
          {recipe.featuredRevision.recipeDescription || "No Description"}
        </Text>
      </div>
      <Button>View Recipe</Button>
    </Flex>
  )
}

const styles = {
  flexContainer: {
    height: "100%",
  },
}
