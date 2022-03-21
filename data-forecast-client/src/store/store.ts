import { configureStore } from '@reduxjs/toolkit'

import authReducer from "./auth"

const store = configureStore({
  reducer: {
    // weather: weatherReducer,
    //users: usersReducer,
    // stocks: stocksReducer
    auth: authReducer
  }
})

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>
// Inferred type: {weather: WeatherState, stocks: StocksState, users: UsersState}
export type AppDispatch = typeof store.dispatch