interface AuthModel {
    isAuthtenticated: boolean,
    userName: string,
    password: string,
    jwtToken: string,
    role: string
}

  export default AuthModel;