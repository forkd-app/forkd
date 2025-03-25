import { create } from "zustand"
import type { Maybe, CurrentUserQuery } from "~/gql/forkd.g"

type CurrentUser = Exclude<
  CurrentUserQuery["user"],
  null | undefined
>["current"]

interface GlobalValues {
  user: Maybe<CurrentUser>
  setUser: (user: Maybe<CurrentUser>) => void
}

export const useGlobals = create<GlobalValues>()((set) => ({
  user: null,
  setUser: (user) => set({ user }),
}))
