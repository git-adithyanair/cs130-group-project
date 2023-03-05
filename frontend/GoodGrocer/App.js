import React from "react";
import { View } from "react-native";
import { useFonts, Inter_600SemiBold } from "@expo-google-fonts/inter";
import { Provider, useSelector, useDispatch } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";

import { NavigationContainer } from "@react-navigation/native";
import { createStackNavigator } from "@react-navigation/stack";

import Landing from "./app/screens/Landing";
import Login from "./app/screens/Login";
import Signup from "./app/screens/Signup";
import AddressSignup from "./app/screens/AddressSignup";
// import Testing from "./app/screens/Testing";
import Tabs from "./app/screens/Tabs";
import ErrorPopup from "./app/components/ErrorPopup";

import { store, persistor } from "./app/store/config";
import { updateDetails } from "./app/store/actions";

const Stack = createStackNavigator();

const App = () => {
  const dispatch = useDispatch();

  const token = useSelector((state) => state.user.token);

  // return (
  //   <NavigationContainer>
  //     <ErrorPopup
  //       message={errorMessageText}
  //       onPress={() => dispatch(updateDetails({ errorPopupVisible: false }))}
  //     />
  //     <Stack.Navigator>
  //       <Stack.Screen
  //         name="Testing"
  //         component={Testing}
  //         options={{ headerShown: false }}
  //       />
  //     </Stack.Navigator>
  //   </NavigationContainer>
  // );

  // The token could be retrieved from the store, so we show the tabs for the user.
  if (token) {
    return (
      <NavigationContainer>
        <ErrorPopup
          onPress={() =>
            dispatch(
              updateDetails({ errorPopupVisible: false, errorMessageText: "" })
            )
          }
        />
        <Stack.Navigator>
          <Stack.Screen
            name="Tabs"
            component={Tabs}
            options={{ headerShown: false }}
          />
        </Stack.Navigator>
      </NavigationContainer>
    );
    // The token is empty, so the user is not logged in, so we show the login/signup pages.
  } else {
    return (
      <NavigationContainer>
        <ErrorPopup
          onPress={() =>
            dispatch(
              updateDetails({ errorPopupVisible: false, errorMessageText: "" })
            )
          }
        />
        <Stack.Navigator>
          <Stack.Screen
            name="Landing"
            component={Landing}
            options={{ headerShown: false }}
          />
          <Stack.Screen name="Login" component={Login} />
          <Stack.Screen
            name="Signup"
            component={Signup}
            options={{ headerTitle: "Sign Up" }}
          />
          <Stack.Screen
            name="AddressSignup"
            component={AddressSignup}
            options={{ headerTitle: "Sign Up" }}
          />
        </Stack.Navigator>
      </NavigationContainer>
    );
  }
};

const AppWrapper = () => {
  let [fontsLoaded] = useFonts({
    Inter_600SemiBold,
  });

  if (!fontsLoaded) {
    return <View />;
  }
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <App />
      </PersistGate>
    </Provider>
  );
};

export default AppWrapper;
