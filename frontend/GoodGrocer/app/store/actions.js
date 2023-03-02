export const SET_TOKEN = "SET_TOKEN";
export const SET_ALL_DETAILS = "SET_ALL_DETAILS";
export const UPDATE_DETAILS = "UPDATE_DETAILS";

export const setToken = (token) => {
  return { type: SET_TOKEN, token };
};

export const setAllDetails = (user, token) => {
  return { type: SET_ALL_DETAILS, user, token };
};

export const updateDetails = (details) => {
  return { type: UPDATE_DETAILS, details };
};
