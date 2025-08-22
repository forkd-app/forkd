import { MetaFunction, useLoaderData } from "react-router"
import { recipesLoader } from "~/.server/loaders/recipes"
import RecipeList from "../../components/recipeList/RecipeList"
import type { Route } from "./+types/route"

export const meta: MetaFunction = () => {
  return [
    { title: "Forkd App" },
    {
      name: "Create new recipes and add your spin on existing recipes",
      content: "Welcome to Forkd!",
    },
  ]
}

export const loader = recipesLoader<Route.LoaderArgs>

export default function Index() {
  const recipes = useLoaderData<typeof loader>()

  if (!recipes) return null

  return <RecipeList recipes={recipes} />
}
