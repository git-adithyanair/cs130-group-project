import React from "react";
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { createStackNavigator } from "@react-navigation/stack";
import { Image, StyleSheet } from 'react-native'; 
import Ionicons from 'react-native-vector-icons/Ionicons';
import Buy from '../screens/Buy'; 
import Shop from '../screens/Shop';
import JoinCommunity from '../screens/JoinCommunity';
import YourCommunities from "../screens/YourCommunities";
import {Dim, Colors} from "../Constants"

const HomeStack = createStackNavigator();

const HomeStackScreen = () => {
  return (
    <HomeStack.Navigator>
      <HomeStack.Screen name="YourCommunities" component={YourCommunities} options={{title: "Your Communities"}}/>
      <HomeStack.Screen name="JoinCommunity" component={JoinCommunity} options={{title: "Join Community"}}/>
    </HomeStack.Navigator>
  );
}

const Tab = createBottomTabNavigator(); 

const tabBarPages = [
    {"name": "Shop",
    "component": Shop},
    {"name": "Home",
    "component": HomeStackScreen},
    {"name": "Buy",
    "component": Buy}
]

const TabBar = (props) => {
    const components = tabBarPages.map((page, index) => <Tab.Screen name={page.name} component={page.component} key={index}/>)
    if(components.length === 0){
        return null; 
    }
    return  <Tab.Navigator screenOptions={({ route }) => ({
        headerShown: false,
        tabBarShowLabel: false,
        tabBarIcon: ({ focused, color, size }) => {
            let iconName;
            if (route.name === 'Shop') {
              iconName = focused
                ? 'ios-cart'
                : 'ios-cart-outline';
            } else if (route.name === 'Buy') {
              iconName = focused ? 'ios-pizza' : 'ios-pizza-outline';
            }
            else if (route.name === 'Home') {
                return <Image source={{uri: props.imageUri}} style={styles.logo} />;
            }
            return <Ionicons name={iconName} size={size} color={'white'} />;
          },
        tabBarStyle: {
          height: Dim.width * 0.15,
          width: Dim.width, 
          backgroundColor: Colors.darkGreen,
          color: Colors.white,
          position: 'absolute',
          marginBottom: 0
      },
    })}>
        {components}
    </Tab.Navigator>; 
};


const styles = StyleSheet.create({
    logo: {
      width: 35,
      height: 35
    },
  });


export default TabBar;