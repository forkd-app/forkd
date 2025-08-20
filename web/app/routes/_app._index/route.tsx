import { MetaFunction, useLoaderData, LoaderFunctionArgs } from "react-router"
import { ClientError } from "graphql-request"
import { getSessionOrThrow } from "~/.server/session"
import { getSDK } from "~/gql/client"
import { environment } from "~/.server/env"
import RecipeList from "../../components/recipeList/RecipeList"

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

  if (!recipes) return null

  return <RecipeList recipes={recipes} />
}
