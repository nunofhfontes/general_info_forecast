import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import type { RootState } from '../../store/store'

// Define a type for the slice state
interface UserState {
    userName: string,
    password: string,
    jwtToken: string,
    role: string
  }
  
  // Define the initial state using that type
  const initialState: UserState = {
    userName: "",
    password: "",
    jwtToken: "",
    role: ""
  }

  export const userSlice = createSlice({
    name: 'user',
    // `userSlice` will infer the state type from the `initialState` argument
    initialState,
    reducers: {
      setUser: state => {
        state.userName = "";
        state.password = "";
        state.role = "";
        state.jwtToken = ""
      },
      // Use the PayloadAction type to declare the contents of `action.payload`
    //   incrementByAmount: (state, action: PayloadAction<number>) => {
    //     state.value += action.payload
    //   }
    }
  })