import {SEARCH_RESULTS} from '../actionTypes'

export default (state=[],{type,payload}) => {
    switch (type) {
        case SEARCH_RESULTS:
            return payload
        default:
            return state;
    }
}