import React from "react";
import { render, fireEvent, waitFor } from "@testing-library/react-native";
import { Provider } from "react-redux";
import '@testing-library/jest-dom'
import { GOOGLE_MAPS_API_KEY } from "../app/Constants";
import { store } from "../app/store/config";
import { setToken } from "../app/store/actions";
import useRequest from "../app/hooks/useRequest";
import Signup from "../app/screens/Signup";
import AddressSignup from "../app/screens/AddressSignup";
import axios from 'axios'; 

jest.mock("axios")

const testAddressSignup = '423 Kelton Ave'
const testAddressFromGoogleMapsApi = '640 Veteran Ave'

const mockGoogleApiResponse = {
    html_attributions: [],
    results: [
        {
        formatted_address: '',
        geometry: {location: {lat: 0}},
        icon: '',
        icon_background_color: '',
        icon_mask_base_uri: '',
        name: testAddressFromGoogleMapsApi,
        place_id: '',
        reference: '',
        types: ['']
        }
    ],
    status: 'OK'
}

const testAddressSignupRoute = {params: {
    email: "email@test.com",
    password: "password",
    name: "Test User",
    phoneNumber: "1234567890"
}}

const testNavigation = {navigate: jest.fn()}


// Test Signup Component
test("given valid email, password, name, phone number, the user proceeds to AddressSignup page", async () => {  
    const signupComponent = render(
        <Provider store={store}>
            <Signup navigation={testNavigation} /> 
        </Provider>
    );
    
    fireEvent.changeText(
        signupComponent.getByPlaceholderText("Enter your email..."),
        "email@test.com"
    );
    fireEvent.changeText(
        signupComponent.getByPlaceholderText("Enter a password..."),
        "password"
    );
    
    fireEvent.changeText(
        signupComponent.getByPlaceholderText("Enter your name..."),
        "Test User"
    );
    
    fireEvent.changeText(
        signupComponent.getByPlaceholderText("Enter your phone number..."),
        "1234567890"
    );
    
    fireEvent.press(signupComponent.queryByText("Continue"));
    
    // Expect to navigate to AddressSignup once 
    expect(testNavigation.navigate.mock.calls).toHaveLength(1) 
})


// Test AddressSignup Component 
test("given valid address, the signup is successful", async () => {
    axios.get.mockResolvedValue({ data: mockGoogleApiResponse });
    const addressSignupComponent = render(
    <Provider store={store}>
        <AddressSignup navigation={testNavigation} route={testAddressSignupRoute}/> 
    </Provider>
    );

    fireEvent.changeText(
    addressSignupComponent.getByPlaceholderText("Enter your address..."),
    testAddressSignup
    );

    fireEvent.press(addressSignupComponent.queryByText("Search"));

    const mockedDoRequest = jest.fn(() => {
        store.dispatch(setToken("token"));
      });
    
      useRequest.mockImplementation(() => {
        return {
          doRequest: mockedDoRequest,
        };
      });
    
    const { doRequest } = useRequest();

    await waitFor(async () => {
        // Expect for axios calls to Google Maps API (when picking address)
        expect(axios.get).toHaveBeenCalledTimes(1)
        expect(axios.get).toHaveBeenCalledWith(`https://maps.googleapis.com/maps/api/place/textsearch/json`,  {"params": {"fields": "formatted_address,place_id,geometry,name", "key": GOOGLE_MAPS_API_KEY, "query": testAddressSignup}});
 

        // Expect for signup to be successful
        fireEvent.press(addressSignupComponent.queryByTestId(testAddressFromGoogleMapsApi))
        fireEvent.press(addressSignupComponent.queryByText("Sign Up"));
        expect(doRequest).toHaveBeenCalled();
        expect(store.getState().user.token).toBe("token");
    })
})