import { FC } from "react"
import { Flex, Text, Button, Image, Rating } from "@mantine/core"

interface Props {
  recipe: {
    title: string
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
        <Text>{recipe.title}</Text>
        <Rating defaultValue={4} />
        <Text>posted by author on date</Text>
      </div>
      <div>
        <Text>Lorem ipsum dolor sit amet, consectetur..</Text>
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
