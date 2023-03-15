import React from "react";
import { render, fireEvent } from "@testing-library/react-native";
import { Provider } from "react-redux";

import { store } from "../app/store/config";
import useRequest from "../app/hooks/useRequest";
import CommunityCard from "../app/components/CommunityCard";

const mockedDoRequest = jest.fn();

jest.mock("../app/hooks/useRequest", () =>
  jest.fn(() => {
    return {
      doRequest: mockedDoRequest,
    };
  })
);

test("upon pressing the join button on a community card, the user successfully joins that community", () => {
  const component = render(
    <Provider store={store}>
      <CommunityCard joinCommunity={true} />
    </Provider>
  );

  const { doRequest } = useRequest();
  fireEvent.press(component.queryByText("join"));

  expect(doRequest).toHaveBeenCalled();
});

test("if user tries to join a community they have already joined, button will be disabled", () => {
  const component = render(
    <Provider store={store}>
      <CommunityCard joinCommunity={true} />
    </Provider>
  );

  fireEvent.press(component.queryByText("join"));

  expect(component.queryByText("join")).toBe(null);
});
