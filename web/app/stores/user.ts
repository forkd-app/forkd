import { createSlice, PayloadAction } from "@reduxjs/toolkit"
import type { Maybe, CurrentUserQuery } from "~/gql/forkd.g"

export type CurrentUser = Exclude<
  CurrentUserQuery["user"],
  null | undefined
>["current"]

export interface UserState {
  value: Maybe<CurrentUser>
}

const initialState: UserState = {
  value: null,
}

export const globalSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUser: (state, action: PayloadAction<Maybe<CurrentUser>>) => {
      state.value = action.payload
    },
  },
})

export const { setUser } = globalSlice.actions
export default globalSlice.reducer
