import axios from "axios";
import { useDispatch, useSelector } from "react-redux";

import { API_URL } from "../Constants";
import { setToken, setErrorPopup } from "../store/actions";

const useRequest = ({
  url,
  method,
  body,
  params = {},
  headers,
  onSuccess,
  onFail,
}) => {
  const token = useSelector((state) => state.user.token);

  const dispatch = useDispatch();

  const doRequest = async (dynamicBody = {}) => {
    try {
      const response = await axios({
        method,
        url: API_URL + url,
        data: method === "get" ? null : { ...body, ...dynamicBody },
        headers: { ...headers, Authorization: `Bearer ${token}` },
        params,
      });
      if (onSuccess) {
        onSuccess(response.data);
      }
      return response.data;
    } catch (err) {
      if (err.response && err.response.status === 401) {
        dispatch(setToken(""));
      }
      const error = err.response
        ? {
            message: err.response.data.error,
            httpCode: err.response.status,
            internalCode: err.response.data.code,
          }
        : {
            message: "Couldn't connect to server! Please try again later.",
            httpCode: 0,
            internalCode: "SERVER_ERROR",
          };
      dispatch(setErrorPopup(error.message));
      if (onFail) {
        onFail(error);
      }
      console.log(error);
    }
  };

  return { doRequest };
};

export default useRequest;
