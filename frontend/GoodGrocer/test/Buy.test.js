import React from "react";
import { render, fireEvent, waitFor } from "@testing-library/react-native";
import { Provider } from "react-redux";
import { store } from "../app/store/config";
import useRequest from "../app/hooks/useRequest";
import Buy from "../app/screens/Buy";
import '@testing-library/jest-dom';
import { cleanup } from '@testing-library/react-native';

afterEach(cleanup);

const mockedDoRequest = jest.fn();

jest.mock("../app/hooks/useRequest", () =>
  jest.fn(() => {
    return {
      doRequest: mockedDoRequest,
    };
  })
);

const testBuyRoute = {
    params: {
      communityId: 0,
      storeId: 1,
    },
};

// Test adding an item and completing the order
test("given a valid item the item is added and the order is completed", () => {
    const BuyComponent = render(
        <Provider store={store}>
            <Buy
                route={testBuyRoute}
            />
        </Provider>
    );

    const { doRequest } = useRequest();

    fireEvent.press(BuyComponent.queryByText("Add Items"));

    fireEvent.changeText(
        BuyComponent.getByPlaceholderText("Enter item name..."),
        "apple"
    );

    fireEvent(BuyComponent.getByTestId('picker-select'), 'onValueChange', 'oz');

    fireEvent.changeText(
        BuyComponent.getByPlaceholderText("Enter item quantity..."),
        "5"
    );

    fireEvent.press(BuyComponent.queryByText("Add Item"));

    fireEvent.press(BuyComponent.queryByText("Complete your Order"));

    expect(doRequest).toHaveBeenCalled();
});

// Test pop up shows up if not all fields are inputed
test("given invalid information no item is added", () => {
  jest.clearAllMocks();
  const BuyComponent = render(
      <Provider store={store}>
          <Buy
              route={testBuyRoute}
          />
      </Provider>
  );

  const { doRequest } = useRequest();

  fireEvent.press(BuyComponent.queryByText("Add Items"));

  fireEvent.changeText(
      BuyComponent.getByPlaceholderText("Enter item name..."),
      "apple"
  );

  fireEvent.press(BuyComponent.queryByText("Add Item"));

  fireEvent.press(BuyComponent.queryByText("Complete your Order"));

  expect(doRequest).toHaveBeenCalledTimes(0);
});

