/**
 * 能发送异步ajax请求的函数模块
 * 封装的时axios库
 * 函数返回值时promise对象
 */

import axios from 'axios'

export default function ajax(url, data={}, type='GET') {
    if(type==='GET') {//发送GET请求
        return axios.get(url, {// 配置对象params
            params:data//指定请求参数
        })
    } else {//POST请求
        return axios.post(url, data)
    }
}

//请求登陆接口
ajax('/userapp/user/login',{username:'Tom', password:'12345'}, 'GET').then()