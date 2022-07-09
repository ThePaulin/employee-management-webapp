

export const usernameValidator = username =>{
    if(!username){
        return "username is required"
    }else if(username.length < 8){
        return "username must have a minimum of 8 characters"
    }
    return "";
}

export const passwordValidator = password =>{
    if(!password){
        return "Password is required"
    }else if(password.length < 8){
        return "Password must have a minimum of 8 characters"
    }
    return "";
}