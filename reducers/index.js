import { combineReducers } from 'redux';
import { streamReducer, pagination, streamError, selectedID } from './streamReducer';
import {UserReducer} from './userReducer';

const rootReducer = combineReducers({
  streams: streamReducer,
  pagination,
  error: streamError,
  ID : selectedID,
  User: UserReducer
});

export default rootReducer;
