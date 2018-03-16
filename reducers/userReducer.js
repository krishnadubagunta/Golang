import {USER} from '../actionTypes'

export function UserReducer(state,{type,payload}){
    switch (type) {
        case USER:
            return {user : payload}
        default:
            return null
    }
}