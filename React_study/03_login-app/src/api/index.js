/**
 * 包含应用中所有接口请求函数的模块
 * 每个函数的返回值都是promise
 */

import ajax from "./ajax";

//登陆
export const reqLogin = (username, password) => ajax('/userapp/user/login', {username, password}, 'GET')

