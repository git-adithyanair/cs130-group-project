

import HomeScreen from './app/screens/HomeScreen';
import Login from './app/screens/Login'
import Signup from './app/screens/Signup'
import AddressSignup from './app/screens/AddressSignup';
// import Testing from './app/screens/Testing';
import LoggedInHome from './app/screens/LoggedInHome';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';

const Stack = createStackNavigator();
export default function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator>
        {/* <Stack.Screen name="Testing" component={Testing} /> */}
        <Stack.Screen name="Home" component={HomeScreen} />
        <Stack.Screen name="Login" component={Login} />
        <Stack.Screen name="Signup" component={Signup} />
        <Stack.Screen name="AddressSignup" component={AddressSignup} />
        <Stack.Screen name="LoggedInHome" component={LoggedInHome} options={{ headerShown: false }} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}


