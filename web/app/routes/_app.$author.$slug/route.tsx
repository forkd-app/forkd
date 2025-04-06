import {
  Flex,
  Text,
  Image,
  Rating,
  Pill,
  Button,
  Checkbox,
  List,
  Title,
} from "@mantine/core"
import { IconArrowLeft, IconShare } from "@tabler/icons-react"
import { useMediaQuery } from "@mantine/hooks"
import { LoaderFunctionArgs } from "@remix-run/node"
import { getSessionOrThrow } from "~/.server/session"
import { getSDK } from "~/gql/client"
import { environment } from "~/.server/env"

export async function loader(args: LoaderFunctionArgs) {
  const session = await getSessionOrThrow(args, false)
  const auth = session.get("sessionToken")
  const sdk = getSDK(`${environment.BACKEND_URL}`, auth)
  try {
    if (args.params.author && args.params.slug) {
      const data = await sdk.RecipeBySlug({
        slug: args.params.slug,
        authorDisplayName: args.params.author,
      })
      console.log("recipe by author/ slug", data)
      return data
    }
  } catch (error) {
    return null
  }
  return null
}

export default function Recipe() {
  const isMobile = useMediaQuery("(max-width: 1199px)")

  return (
    <Flex style={styles.container} direction="column">
      <Flex align="center" style={styles.header}>
        <IconArrowLeft size={30} />
        <Title order={1}> Lasagna </Title>
      </Flex>
      <Flex direction={isMobile ? "column" : "row"}>
        <Flex style={styles.column} direction="column">
          <Image src="" alt="recipe" />
          <Flex justify="space-between">
            <Text style={styles.text}> Chris Burger </Text>
            <Text style={styles.text}> Posted on 01/01/1001 </Text>
          </Flex>
          <Flex justify="space-between">
            <Text style={styles.text}> 0 Forks </Text>
            <Flex align="center">
              <Rating />
              <Text>(0 ratings)</Text>
            </Flex>
          </Flex>
          <Flex>
            <Pill style={styles.pill} size="sm">
              Original Recipe
            </Pill>
            <Pill style={styles.pill} size="sm" withRemoveButton>
              Vegan
            </Pill>
            <Pill style={styles.pill} size="sm" withRemoveButton>
              Italian
            </Pill>
            <Pill style={styles.pill} size="sm" withRemoveButton>
              Saucy
            </Pill>
          </Flex>
          <Text style={styles.text}>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua.
            <span style={{ textDecoration: "underline" }}> READ MORE</span>
          </Text>
          <Flex align="center">
            <Text style={styles.text}>
              Share this recipe <IconShare />
            </Text>
          </Flex>
          <Flex justify="space-between">
            <Button color="gray" size="xl">
              {" "}
              Fork Recipe
            </Button>
            <Text style={styles.text}>Save Recipe</Text>
          </Flex>
        </Flex>
        <Flex style={styles.column} direction="column">
          <Flex direction="column">
            <Title order={2}>Ingredients</Title>
            <Checkbox
              style={styles.text}
              defaultChecked
              label="2 1/2 cups small shell pasta"
              color="gray"
            />
            <Checkbox
              style={styles.text}
              defaultChecked
              label="1 tablespoon extra vigin olive oil"
              color="gray"
            />
            <Checkbox
              style={styles.text}
              defaultChecked
              label="1 small yellow onion"
              color="gray"
            />
          </Flex>
          <Flex direction="column">
            <Title order={2}>Instructions</Title>
            <div>
              <Title order={3}>Part One</Title>
              <List type="ordered">
                <List.Item style={styles.text}>
                  Boil pasta: Lorem ipsum dolor sit amet, sed do eiusmod tempor
                  incididunt ut labore et dolore magna aliqua.
                </List.Item>
                <List.Item style={styles.text}>
                  Saute onion: Lorem ipsum dolor sit amet, consectetur
                  adipiscing elit.
                </List.Item>
                <List.Item style={styles.text}>Plate with olive oil</List.Item>
              </List>
            </div>
          </Flex>

          <Flex direction="column">
            <Title order={2}>Revisions</Title>
            <div>
              <List type="unordered">
                <List.Item style={styles.text}>
                  01/01/2025 I forgot to add sugar to the recipe. I added sugar
                  in my revision.{" "}
                </List.Item>
              </List>
            </div>
          </Flex>
        </Flex>
      </Flex>
    </Flex>
  )
}

const styles = {
  container: {
    padding: 40,
    paddingBottom: 200,
  },
  text: {
    paddingTop: 5,
    paddingBottom: 5,
  },
  column: {
    padding: 15,
  },
  header: {
    padding: 10,
  },
  pill: {
    margin: 1.5,
  },
}
