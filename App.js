import React, { Component } from 'react';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import thunkMiddleware from 'redux-thunk';
import createLogger from 'redux-logger';
import rootReducer from './reducers';
import {Tabs} from './navigation';
import Reactotron, {networking} from 'reactotron-react-native'

const createStoreWithMiddleware = applyMiddleware(thunkMiddleware)(createStore);

Reactotron
  .configure()
  .use(networking()) // <--- here we go!
  .connect()

store = createStoreWithMiddleware(rootReducer, {});

export default class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <Tabs />
      </Provider>
    );
  }
}
