import { RECEIVE_STREAMS, RECEIVE_PAGINATION, ID , USER, SEARCH_RESULTS,AUTH} from '../actionTypes';
import axios from 'axios';
import {HOST} from 'react-native-dotenv'
import {Platform} from 'react-native';

const BASE_URL = `${HOST}/api`;

export const fetchStreams = () => async dispatch => {
  const req = await requests(`${BASE_URL}/streams`);
  dispatch({
      type: RECEIVE_STREAMS,
      payload: req.data
    });
  dispatch({
    type: RECEIVE_PAGINATION,
    payload: req.pagination
  });
};

const parsePromise = (res) => {
  return new Promise((resolve, reject)=>{
    var result = {}
    try{
      result = JSON.parse(res)
      resolve(result)
    }
    catch(e){
      reject(result)
    }
  })
}

export const fetchUsers = (q) => async dispatch => {
  var results =[]
  if(q.length > 0){
    const req = await requests(`${BASE_URL}/search/user?q=${q}`)
    results = req.data
  }
  dispatch({
    type: SEARCH_RESULTS,
    payload : results
  })
}

export const selectID = (id) => async dispatch => {
  dispatch({
    type : ID,
    payload : id
  })
}

export const fetchVideoURL = (id) => async dispatch => {
  const req = await requests(`${BASE_URL}/user/${id}`)
  dispatch({
    type : USER,
    payload : req.data.data[0]
  })
}

export const login = () => async dispatch => {
  const req = await requests(`${BASE_URL}/auth`)
  dispatch({
    type: AUTH,
    payload : req.data
  })
}

const requests = async(url) => {
 const res = await axios.get(url,{
    method:"GET",
    headers : {
      "Accept-Encoding" : ["gzip","deflate"]
    }
  });
  return res;
}