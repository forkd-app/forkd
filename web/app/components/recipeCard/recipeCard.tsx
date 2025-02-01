import { Flex, Text, Button } from "@mantine/core"
import ImgFiller from "../../../public/Image Item.jpg"

interface Recipe {
  title: string
}

export function RecipeCard({ title }: Recipe) {
  return (
    <Flex
      direction={"column"}
      justify={"space-evenly"}
      style={styles.flexContainer}
    >
      <img src={ImgFiller} alt="recipe" height={"60%"} width={"100%"} />
      <div>
        <Text>{title}</Text>
        <Text>stars</Text>
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
