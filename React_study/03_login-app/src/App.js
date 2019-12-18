import React, {Component} from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import './App.css';
import '../node_modules/antd/dist/antd.css'

import Login from './pages/login/login';
import Admin from './pages/admin/admin';

/*
应用的根组件
*/

export default class App extends Component {
  render () {
    return (
      <BrowserRouter>
        <Switch >{/*只匹配其中一个路由*/ }
            <Route path='/Login' component={Login}></Route>
            <Route path='/' component={Admin}></Route>
        </Switch>
      </BrowserRouter>
    )
  }
}