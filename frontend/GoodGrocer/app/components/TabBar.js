import React from "react";
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { Image, StyleSheet } from 'react-native'; 
import Ionicons from 'react-native-vector-icons/Ionicons';
import Buy from '../screens/Buy'; 
import Shop from '../screens/Shop';
import Profile from '../screens/Profile';

const Tab = createBottomTabNavigator(); 

const tabBarPages = [
    {"name": "Shop",
    "component": Shop},
    {"name": "Profile",
    "component": Profile},
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
            else if (route.name === 'Profile') {
                return <Image source={{uri: props.imageUri}} style={styles.logo} />;
            }
            return <Ionicons name={iconName} size={size} color={'white'} />;
          },
        tabBarStyle: {
          height: 80,
          width: '100%', 
          backgroundColor: '#7B886B',
          color: 'white',
          position: 'absolute',
          marginBottom: 0
      },
    })}>
        {components}
    </Tab.Navigator>; 
};


const styles = StyleSheet.create({
    logo: {
      width: 66,
      height: 50
    },
  });


export default TabBar;