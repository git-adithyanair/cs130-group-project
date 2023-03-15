import React from "react";
import { render, fireEvent } from "@testing-library/react-native";
import { Provider } from "react-redux";

import { store } from "../app/store/config";
import useRequest from "../app/hooks/useRequest";
import JoinCommunity from "../app/screens/JoinCommunity";

const mockedDoRequest = jest.fn();

jest.mock("../app/hooks/useRequest", () =>
  jest.fn(() => {
    return {
      doRequest: mockedDoRequest,
    };
  })
);

const testNavigation = { navigate: jest.fn() };

const testJoinCommunityRoute = {
  params: {
    userXCoord: 0.0,
    userYCoord: 0.0,
  },
};

test("two endpoints are called as soon as screen is rendered", () => {
  render(
    <JoinCommunity navigation={testNavigation} route={testJoinCommunityRoute} />
  );

  const { doRequest } = useRequest();

  expect(doRequest).toHaveBeenCalledTimes(2);
});

test("upon entering something in search, the search query updates accordingly", () => {
  const component = render(
    <Provider store={store}>
      <JoinCommunity
        navigation={testNavigation}
        route={testJoinCommunityRoute}
      />
    </Provider>
  );

  const initialDataLength =
    component.getByTestId("communitiesToJoin").props["data"].length;

  fireEvent.changeText(
    component.getByPlaceholderText("Search..."),
    "searchString"
  );

  expect(component.getByTestId("searchInput").props["value"]).toBe(
    "searchString"
  );
  expect(
    component.getByTestId("communitiesToJoin").props["data"].length
  ).toBeLessThanOrEqual(initialDataLength);
});
