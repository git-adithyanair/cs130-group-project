import React from "react";
import "@testing-library/jest-dom";
import { render, fireEvent } from "@testing-library/react-native";
import { Provider } from "react-redux";

import { store } from "../app/store/config";
import { setToken } from "../app/store/actions";
import useRequest from "../app/hooks/useRequest";
import Login from "../app/screens/Login";

test("given valid email and password, the login is successful", () => {
  const component = render(
    <Provider store={store}>
      <Login />
    </Provider>
  );

  const mockedDoRequest = jest.fn(() => {
    store.dispatch(setToken("token"));
  });

  useRequest.mockImplementation(() => {
    return {
      doRequest: mockedDoRequest,
    };
  });

  const { doRequest } = useRequest();

  fireEvent.changeText(
    component.getByPlaceholderText("Enter your email..."),
    "email"
  );
  fireEvent.changeText(
    component.getByPlaceholderText("Enter your password..."),
    "password"
  );

  fireEvent.press(component.queryByText("Log In"));

  expect(doRequest).toHaveBeenCalled();
  expect(store.getState().user.token).toBe("token");
});
