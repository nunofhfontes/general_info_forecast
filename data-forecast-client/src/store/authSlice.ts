import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { RootState } from './store';
// import type { RootState } from '../../store/store'

// Define a type for the slice state
interface AuthState {
  isAuthtenticated: boolean,
  userName: string,
  password: string,
  jwtToken: string,
  role: string
}
  
  // Define the initial state using that type
const initialAuthState: AuthState = {
  isAuthtenticated: false,
  userName: "",
  password: "",
  jwtToken: "",
  role: ""
}

const authSlice = createSlice({
  name: 'authentication',
  // `authSlice` will infer the state type from the `initialState` argument
  initialState: initialAuthState,
  reducers: {
    setUser: state => {
      state.isAuthtenticated = true;
    },
    login: (state, action: PayloadAction<AuthState>) => {
      state.isAuthtenticated = true;
      state.userName = action.payload.userName;
      state.password = "";
      state.role = "";
      state.jwtToken = ""
    },
    logout: state => {
      state.isAuthtenticated = false;
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
});

export const authActions = authSlice.actions;

export default authSlice.reducer;

// Other code such as selectors can use the imported `RootState` type
// TODO - is this really needed??
export const selectAuth = (state: RootState) => state.auth;