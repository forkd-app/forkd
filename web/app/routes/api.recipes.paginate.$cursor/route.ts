import { recipesLoader } from "~/.server/loaders/recipes"
import type { Route } from "./+types/route"
export const loader = recipesLoader<Route.LoaderArgs>
