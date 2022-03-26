import { configureStore } from '@reduxjs/toolkit'

import authenticationReducer from "./authSlice"

const store = configureStore({
  reducer: {
    // weather: weatherReducer,
    //users: usersReducer,
    // stocks: stocksReducer
    auth: authenticationReducer
  }
})

export default store;

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch