export const SET_TOKEN = "SET_TOKEN";
export const SET_ALL_DETAILS = "SET_ALL_DETAILS";
export const UPDATE_DETAILS = "UPDATE_DETAILS";
export const SET_ERROR_POPUP = "SET_ERROR";

export const setToken = (token) => {
  return { type: SET_TOKEN, token };
};

export const setAllDetails = (user, token) => {
  return { type: SET_ALL_DETAILS, user, token };
};

export const updateDetails = (details) => {
  return { type: UPDATE_DETAILS, details };
};

export const setErrorPopup = (errorMessageText) => {
  return { type: SET_ERROR_POPUP, errorMessageText };
};
