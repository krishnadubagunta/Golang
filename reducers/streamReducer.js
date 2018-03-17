import { RECEIVE_STREAMS, RECEIVE_PAGINATION } from '../actionTypes';

export function streamReducer(state = [], { type, payload }) {
  switch (type) {
    case RECEIVE_STREAMS:
      return { streams : payload };
    default:
      return state;
  }
}

export function pagination(state = {}, { type, payload }) {
  switch (type) {
    case RECEIVE_PAGINATION:
      return { payload };
    default:
      return null;
  }
}

export function streamError(state = {}, { type, payload }) {
  switch (type) {
    case 'Error':
      return { payload };
    default:
      return null;
  }
}

export function selectedID(state={},{type,payload}){
  switch(type){
    case "ID":
      return {ID : payload}
    default:
      return null;
  }
}