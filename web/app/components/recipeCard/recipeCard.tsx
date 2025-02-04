import { Flex, Text, Button, Image } from "@mantine/core"

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
      <Image src="images/image.jpg" alt="recipe" />
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
