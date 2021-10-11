import { userAdmin } from './mock'

export const loginByEmail = async (email, pwd) => {
    console.log(`[loginByEmail] email ${email}`);
    let user = {};
    if (email == userAdmin.email && pwd == userAdmin.password) {
        user = userAdmin;
    }
    return user
}

export const getUserInfo = async (token) => {
    console.log(`[getUserInfo] token ${token}`);
    let user = {};
    if (token == userAdmin.token) {
        user = userAdmin;
    }
    return user
}