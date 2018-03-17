import { combineReducers } from 'redux';
import { streamReducer, pagination, streamError, selectedID } from './streamReducer';
import {UserReducer} from './userReducer';
import searchResults from './searchReducer';
import auth from './auth'
const rootReducer = combineReducers({
  streams: streamReducer,
  pagination,
  error: streamError,
  ID : selectedID,
  User: UserReducer,
  searchResults ,
  auth
});

export default rootReducer;
