import { FC } from "react"
import { Flex, Text, Button, Image, Rating, Title } from "@mantine/core"
import { ListRecipesQuery } from "~/gql/forkd.g"
import { Link } from "@remix-run/react"

interface Props {
  recipe: Recipe
}

type Recipe = Exclude<
  ListRecipesQuery["recipe"],
  null | undefined
>["list"]["items"][0]

export const RecipeCard: FC<Props> = ({ recipe }) => {
  return (
    <Flex
      direction={"column"}
      justify={"space-between"}
      style={styles.flexContainer}
    >
      <Flex direction={"column"}>
        <Image src="images/image.jpg" alt="recipe" />
        <div>
          <Title style={styles.text} order={4}>
            {recipe?.featuredRevision?.title || "No Title"}
          </Title>
          <Rating
            style={styles.text}
            defaultValue={recipe?.featuredRevision?.rating || 0}
          />
          <Text style={styles.text}>
            posted by {recipe.author.displayName || "No Name"} on{" "}
            {recipe?.featuredRevision?.publishDate || "unknown"}
          </Text>
        </div>
        <div>
          <Text style={styles.text}>
            {recipe?.featuredRevision?.recipeDescription || "No Description"}
          </Text>
        </div>
      </Flex>
      <Button
        component={Link}
        to={`/${recipe?.author?.displayName}/${recipe?.slug}`}
      >
        View Recipe
      </Button>
    </Flex>
  )
}

const styles = {
  flexContainer: {
    height: "100%",
  },
  text: {
    paddingTop: 5,
    paddingBottom: 5,
  },
}
