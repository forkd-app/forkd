import { configureStore } from "@reduxjs/toolkit"
import userReducer, { UserState } from "./user"

export function getStore(state: { user: UserState }) {
  return configureStore({
    reducer: {
      user: userReducer,
    },
    preloadedState: state,
  })
}

export type RootState = ReturnType<ReturnType<typeof getStore>["getState"]>
export type AppDispatch = ReturnType<typeof getStore>["dispatch"]
