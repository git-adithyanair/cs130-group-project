import HomeScreen from './app/screens/HomeScreen'; 
import Login from './app/screens/Login'
import Signup from './app/screens/Signup' 
import AddressSignup from './app/screens/AddressSignup';
import Requests from './app/screens/Requests';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';

const Stack = createStackNavigator();

export default function App() {
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Screen name="Home" component={HomeScreen} />
        <Stack.Screen name="Login" component={Login} />
        <Stack.Screen name="Signup" component={Signup} /> 
        <Stack.Screen name="AddressSignup" component={AddressSignup} />
        <Stack.Screen name="Requests" component={Requests} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}


