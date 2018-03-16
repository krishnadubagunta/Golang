import React, { Component } from 'react';
import { createStore, applyMiddleware } from 'redux';
import { Provider } from 'react-redux';
import thunkMiddleware from 'redux-thunk';
import createLogger from 'redux-logger';
import rootReducer from './reducers';
import {Tabs} from './navigation';

const createStoreWithMiddleware = applyMiddleware(thunkMiddleware)(createStore);

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
