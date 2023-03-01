import { SET_TOKEN, SET_ALL_DETAILS, UPDATE_DETAILS } from "./actions";

const reducer = (state = {}, action) => {
  switch (action.type) {
    case SET_TOKEN:
      return { ...state, token: action.token };

    case SET_ALL_DETAILS:
      return {
        ...state,
        ...action.user,
        token: action.token,
      };

    case UPDATE_DETAILS:
      return { ...state, ...action.details };

    default:
      return state;
  }
};

export default reducer;
