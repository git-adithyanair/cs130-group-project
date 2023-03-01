import React from "react";
import AppLoading from "expo-app-loading";
import { useFonts, Inter_600SemiBold } from "@expo-google-fonts/inter";
import { Provider, useSelector } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";

import { NavigationContainer } from "@react-navigation/native";
import { createStackNavigator } from "@react-navigation/stack";

import Landing from "./app/screens/Landing";
import Login from "./app/screens/Login";
import Signup from "./app/screens/Signup";
import AddressSignup from "./app/screens/AddressSignup";
// import Testing from './app/screens/Testing';
import Tabs from "./app/screens/Tabs";

import { store, persistor } from "./app/store/config";

const Stack = createStackNavigator();

const App = () => {
  const token = useSelector((state) => state.user.token);

  // return <Testing />;

  // The token could be retrieved from the store, so we show the tabs for the user.
  if (token) {
    return (
      <NavigationContainer>
        <Stack.Navigator>
          <Stack.Screen
            name="Tabs"
            component={Tabs}
            options={{ headerShown: false }}
          />
        </Stack.Navigator>
      </NavigationContainer>
    );
    // The token is null, so the app is starting up so we should show the loading page.
  } else if (token === null) {
    return <AppLoading />;
    // The token is undefined, so the user is not logged in, so we show the login/signup pages.
  } else {
    return (
      <NavigationContainer>
        <Stack.Navigator>
          <Stack.Screen
            name="Landing"
            component={Landing}
            options={{ headerShown: false }}
          />
          <Stack.Screen name="Login" component={Login} />
          <Stack.Screen name="Signup" component={Signup} />
          <Stack.Screen name="AddressSignup" component={AddressSignup} />
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
    return <AppLoading />;
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
