import {
  SET_TOKEN,
  SET_ALL_DETAILS,
  UPDATE_DETAILS,
  SET_ERROR_POPUP,
} from "./actions";

const reducer = (
  state = { errorPopupVisible: false, errorMessageText: "" },
  action
) => {
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

    case SET_ERROR_POPUP:
      return {
        ...state,
        errorPopupVisible: true,
        errorMessageText: action.errorMessageText,
      };

    default:
      return state;
  }
};

export default reducer;
