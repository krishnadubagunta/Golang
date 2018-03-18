import { AUTH } from '../actionTypes';

export default (state, {type, payload}) => {
    switch (type) {
        case AUTH:
            return payload
        default:
            return ""
    }
}