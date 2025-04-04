import { SimpleGrid } from "@mantine/core"
import { RecipeCard } from "../../components/recipeCard/recipeCard"
import { MetaFunction } from "@remix-run/react"
import { LoaderFunctionArgs } from "@remix-run/node"
import { ClientError } from "graphql-request"
import { getSessionOrThrow } from "~/.server/session"
import { getSDK } from "~/gql/client"
import { environment } from "~/.server/env"
import { useLoaderData } from "@remix-run/react"

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

export async function loader(args: LoaderFunctionArgs) {
  const session = await getSessionOrThrow(args, false)
  console.log("Session Token: ", session.get("sessionToken"))
  const auth = session.get("sessionToken")
  const sdk = getSDK(`${environment.BACKEND_URL}`, auth)
  try {
    const data = await sdk.Recipe().catch(console.error)
    console.log(data?.recipe?.list?.items || null)
    return data?.recipe?.list?.items
  } catch (err) {
    if (err instanceof ClientError && err.message === "missing auth") {
      return null
    }
    throw err
  }
}

export default function Index() {
  const data = useLoaderData<typeof loader>()
  console.log(data)

  return (
    <>
      {/* recipe component */}
      <SimpleGrid
        cols={{ base: 1, sm: 2, md: 2, lg: 4 }}
        pb={40}
        pt={40}
        style={styles.grid}
      >
        {recipes.map((recipe) => (
          <div key={recipe.title} style={styles.col}>
            <RecipeCard recipe={recipe} />
          </div>
        ))}
      </SimpleGrid>
    </>
  )
}

const styles = {
  grid: {
    background: "#fff",
    width: "90%",
    margin: "auto",
  },
  col: {
    padding: 10,
    margin: 10,
    borderWidth: 0,
    borderColor: "black",
    borderStyle: "solid",
    justifyContent: "space-evenly",
    boxShadow: "0px  5px 15px #bfbfbf",
  },
}
