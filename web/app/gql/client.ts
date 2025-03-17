import { GraphQLClient } from "graphql-request"
import { getSdk } from "./forkd.g"
import { environment } from "~/.server/env"

export function getSDK(endpoint: string, accessToken?: string) {
  const headers = new Headers()
  if (accessToken) {
    headers.append("Authorization", accessToken)
  }
  const client = new GraphQLClient(endpoint, {
    headers,
  })
  return getSdk(client)
}

export const client = getSDK(environment.BACKEND_URL)
