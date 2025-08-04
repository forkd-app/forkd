import {
  Flex,
  Pill,
  Button,
  Title,
  Anchor,
  Center,
  ActionIcon,
  AspectRatio,
  TextInput,
  Textarea,
  PillsInput,
} from "@mantine/core"
import { IconArrowLeft, IconPlus } from "@tabler/icons-react"
import { useMediaQuery } from "@mantine/hooks"
import { LoaderFunctionArgs } from "@remix-run/node"
import { getSessionOrThrow } from "~/.server/session"
import { getSDK } from "~/gql/client"
import { environment } from "~/.server/env"
import { useLoaderData } from "@remix-run/react"
import { RecipeBySlugQuery } from "~/gql/forkd.g"

type Recipe = Exclude<RecipeBySlugQuery["recipe"], null | undefined>["bySlug"]

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
      return data?.recipe?.bySlug
    }
  } catch (error) {
    return null
  }
  return null
}

export default function AddRecipe() {
  const isMobile = useMediaQuery("(max-width: 900px)")

  return (
    <Flex style={styles.container} direction="column">
      <Flex align="center" justify="space-between" style={styles.header}>
        <Flex gap="sm">
          <IconArrowLeft size={30} />
          <Title order={1} size="h4">
            {" "}
            Add Recipe{" "}
          </Title>
        </Flex>
        <Flex gap="xl">
          <Button color="gray">Publish Recipe</Button>
          <Center>
            <Anchor c="black">Save to Drafts</Anchor>
          </Center>
        </Flex>
      </Flex>
      <Flex direction={isMobile ? "column" : "row"} gap="md">
        <Flex style={styles.column} direction="column" gap="md">
          <AspectRatio
            ratio={1080 / 720}
            bg="gray.3"
            miw={isMobile ? "100%" : "500px"}
          >
            <Center style={{ height: "100%" }}>
              <ActionIcon
                radius="lg"
                variant="filled"
                color="gray"
                aria-label="Plus"
                style={{ width: "80px", height: "80px" }}
              >
                <IconPlus
                  style={{ width: "20px", height: "20px" }}
                  stroke={1.5}
                />
              </ActionIcon>
            </Center>
          </AspectRatio>
          <TextInput
            w="100%"
            size="md"
            label="Recipe Name"
            placeholder="Add recipe name (10-75 characters)"
          />
          <Textarea
            w="100%"
            size="md"
            label="Description"
            description="Tell us about the recipe"
          />
        </Flex>
        <Flex style={styles.column} direction="column">
          <Flex direction="column">
            <PillsInput label="Tags">
              <Pill.Group>
                <Pill withRemoveButton>React</Pill>
                <Pill>Vue</Pill>
                <Pill>Svelte</Pill>
                <PillsInput.Field placeholder="Enter tags" />
              </Pill.Group>
            </PillsInput>
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
