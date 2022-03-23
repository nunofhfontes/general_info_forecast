class User {
    id: string;
    userName: string;
    role: string;
  
    constructor(userName: string, role: string) {
      this.userName = userName;
      this.id = new Date().toISOString();
      this.role = role;
    }
  }
  
  export default User;

  interface UserModel {
    id: string
    userName: string
    role: string
  }