// PROBLEM ! - ANALYSIS
/*
class user{
  var id;
  var username;
  var password;
}
class userservice {
  user[] users = [];
  user[] getallusers() {
    return users;
  }
  user getuserbyid(userid){
    return this.users.filter(userid);
  }
}
*/

// Yang benar

class User{
  id;
  username;
  password;
}
class UserService {
  user[] users = [];
  user[] GetAllUsers(){
    return users
  }
  User GetUserById(userId){
    return this.users.filter(userId);
  }
}