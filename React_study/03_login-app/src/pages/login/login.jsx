import React, {Component} from 'react';
import { Form, Icon, Input, Button,message } from 'antd';

import './login.css';
import logo from './images/logo.jpg'

import {reqLogin} from "../../api";

/*
后台管理的路由组件
*/

class Login extends Component {


    handleSubmit = (event) => {
        //阻止事件的默认行为
        event.preventDefault()

        //得到form对象
        const form = this.props.form
        const values = form.getFieldsValue()//获得表单数据
        const {username, password} = values//获得表单数据单个数值

        //调用ajax函数，请求登陆,验证登陆，ES6语法
        reqLogin(username, password).then(response => {
            console.log(response.data)
            //请求成功，但并不一定登陆成功，校验服务器返回的code, 正规叫状态码
            const result = response.data
            if(result.Code==='200') {
                message.success('登陆成功aaaaaaa！')

                //跳转到管理界面
                // this.props.history.push()
                this.props.history.replace('/')
            } else {
                message.error(result.Msg)
            }
        }).catch(error => {
            alert('登陆请求出错，请检查...')
            console.log(error)
        })

        //获取表单项的输入数据

    }


    render () {

        //强大的form对象，给表单每个属性定义名字的
        const { getFieldDecorator } = this.props.form;

        return (
            <div className="login">
                <header className="login-header">
                    <h1>登陆注册模块</h1>
                </header>
                <section className="login-content">
                    <h1>用户登陆</h1>
                    <Form onSubmit={this.handleSubmit} className="login-form">
                        <Form.Item>
                            {
                                getFieldDecorator('username', {})(
                                    <Input
                                        prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />}
                                        placeholder="Username"
                                    />
                                )
                            }
                        </Form.Item>
                        <Form.Item>
                            {
                                getFieldDecorator('password', {})(
                                    <Input
                                        prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />}
                                        type="password"
                                        placeholder="Password"
                                    />
                                )
                            }
                        </Form.Item>
                        <Form.Item>
                            <Button type="primary" htmlType="submit" className="login-form-button">
                                登陆
                            </Button>
                        </Form.Item>
                    </Form>
                </section>

            </div>
        )
    }
}


/**
 * 1、高阶函数
 * 1).一类特别的函数
 *  a.接收函数类型的参数
 *  b.返回值是函数
 * 2).常见
 *  a.定时器：setTimeout()/setInterval()
 *  b.Promise: Promise()
 *  c.函数对象的bind
 *
 *
 *
 * 2、高阶组件
 *  1).本质就是一个组件
 *  2).接收一个组件（被包装组件），返回一个新的组件， 包装组件会向被包装组件传入特定的属性
 *  3).作用：扩展组件的功能
 *  4).高阶组件也就是高阶函数：接收一个组件函数，返回时一个新的组件函数
 */

 /**
  * 包装Form组件生成一个新的组件；Form(form)
  * 新的组件会向Form组件传递一个强大的对象属性：form
  */

const WrapLogin = Form.create()(Login)//高阶函数
export default WrapLogin

export default Form.create()(Login)


/*
1、前台表单验证
2、收集表单输入数据
*/