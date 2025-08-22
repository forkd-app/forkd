import { ClientError } from "graphql-request"
import { environment } from "~/.server/env"
import { getSessionOrThrow } from "~/.server/session"
import { getSDK } from "~/gql/client"
import { LoaderFunctionArgs } from "react-router"

export async function recipesLoader<T extends LoaderFunctionArgs>(args: T) {
  const session = await getSessionOrThrow(args, false)
  const auth = session.get("sessionToken")
  const sdk = getSDK(`${environment.BACKEND_URL}`, auth)
  try {
    const data = await sdk.ListRecipes({
      input: { nextCursor: args.params.cursor },
    })
    return data?.recipe?.list ?? null
  } catch (err) {
    if (err instanceof ClientError && err.message === "missing auth") {
      return null
    }
    throw err
  }
}
