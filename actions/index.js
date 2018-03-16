import { RECEIVE_STREAMS, RECEIVE_PAGINATION, ID , USER} from '../actionTypes';
import axios from 'axios';

const BASE_URL = 'https://cebitqqhnz.localtunnel.me/api';

export const fetchStreams = () => async dispatch => {
  const req = await axios.get(`${BASE_URL}/streams`);
  dispatch({
    type: RECEIVE_STREAMS,
    payload: req.data.data
  });
  dispatch({
    type: RECEIVE_PAGINATION,
    payload: req.data.pagination
  });
};

export const selectID = (id) => async dispatch => {
  dispatch({
    type : ID,
    payload : id
  })
}

export const fetchVideoURL = (id) => async dispatch => {
  const req = await axios.get(`${BASE_URL}/user/${id}`)
  dispatch({
    type : USER,
    payload : req.data.data[0]
  })
}

// function RNFecthWithURL(url) {
//   RNFetchBlob.fetch('GET',`${BASE_URL}/${url}`).then( (res) => {
//     return res.json()
//   }).catch((err,message)=>{
//     console.log(err,message);
//   })
// }