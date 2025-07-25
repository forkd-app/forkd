import { SimpleGrid } from "@mantine/core"
import { RecipeCard } from "../../components/recipeCard/recipeCard"
import { MetaFunction, useLoaderData } from "@remix-run/react"
import { LoaderFunctionArgs } from "@remix-run/node"
import { ClientError } from "graphql-request"
import { getSessionOrThrow } from "~/.server/session"
import { getSDK } from "~/gql/client"
import { environment } from "~/.server/env"

export const meta: MetaFunction = () => {
  return [
    { title: "Forkd App" },
    {
      name: "Create new recipes and add your spin on existing recipes",
      content: "Welcome to Forkd!",
    },
  ]
}

export async function loader(args: LoaderFunctionArgs) {
  const session = await getSessionOrThrow(args, false)
  const auth = session.get("sessionToken")
  const sdk = getSDK(`${environment.BACKEND_URL}`, auth)
  try {
    const data = await sdk.ListRecipes()
    // console.log(data?.recipe?.list || null)
    return data?.recipe?.list ?? null
  } catch (err) {
    if (err instanceof ClientError && err.message === "missing auth") {
      return null
    }
    throw err
  }
}

export default function Index() {
  const recipes = useLoaderData<typeof loader>()

  return (
    <>
      {/* recipe component */}
      <SimpleGrid
        cols={{ base: 1, sm: 2, md: 2, lg: 4 }}
        pb={40}
        pt={40}
        style={styles.grid}
      >
        {recipes?.items.map((recipe) => (
          <div key={recipe.id} style={styles.col}>
            <RecipeCard recipe={recipe || {}} />
          </div>
        ))}
      </SimpleGrid>
    </>
  )
}
