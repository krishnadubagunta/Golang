import { RECEIVE_STREAMS, RECEIVE_PAGINATION, ID , USER} from '../actionTypes';
import axios from 'axios';
import {Platform} from 'react-native';

const BASE_URL = 'http://10.0.0.92:8080/api';

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


const requests = async(url) => {
  var res = {}
  if(Platform.OS === "android"){
    res = await axios.get(url,{
      method: "GET",
      headers : {
        "Accept-Encoding" : ["gzip","deflate"]
      }
    });
  }
  else{
    res = await axios.get(url)
  }
  return res;
}