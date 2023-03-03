import React from "react";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { createStackNavigator } from "@react-navigation/stack";
import { Image, StyleSheet } from "react-native";
import Ionicons from "react-native-vector-icons/Ionicons";
import JoinCommunity from "../screens/JoinCommunity";
import YourCommunities from "../screens/YourCommunities";
import Profile from "../screens/Profile"; 
import { Dim, Colors } from "../Constants";
import ActiveErrand from "../screens/ActiveErrand";
import ActiveRequest from "../screens/ActiveRequest";
import RequestList from '../screens/RequestList';
import RequestDetail from '../screens/RequestDetail';

const HomeStack = createStackNavigator();
const ErrandStack = createStackNavigator();
const ProfileStack = createStackNavigator(); 

const HomeStackScreen = () => {
  return (
    <HomeStack.Navigator>
      <HomeStack.Screen
        name="YourCommunities"
        component={YourCommunities}
        options={{ title: "Your Communities" }}
      />
      <HomeStack.Screen
        name="RequestList"
        component={RequestList}
        options={{ title: "Community Requests" }}
      />
      <HomeStack.Screen
        name="JoinCommunity"
        component={JoinCommunity}
        options={{ title: "Join Community" }}
      />
      <HomeStack.Screen
        name="RequestDetail"
        component={RequestDetail}
        options={{ title: "Request Detail" }}
      />
    </HomeStack.Navigator>
  );
};

const ProfileStackScreen = () => {
  return (
    <ProfileStack.Navigator>
      <ProfileStack.Screen
        name="Profile"
        component={Profile} 
      />
    </ProfileStack.Navigator>
  );
};

const ErrandStackScreen = () => {
  return (
    <ErrandStack.Navigator>
      <ErrandStack.Screen
        name="ActiveErrand"
        component={ActiveErrand}
        options={{ title: "" }}
      />
      <ErrandStack.Screen
        name="ActiveRequest"
        component={ActiveRequest}
        options={{ title: "", headerTintColor: "green" }}
      />
    </ErrandStack.Navigator>
  );
};


const Tab = createBottomTabNavigator();

const tabBarPages = [
  { name: "Home", component: HomeStackScreen },
  { name: "MyProfile", component: ProfileStackScreen },
  { name: "Errand", component: ErrandStackScreen },
];

const TabBar = (props) => {
  const components = tabBarPages.map((page, index) => (
    <Tab.Screen name={page.name} component={page.component} key={index} />
  ));
  if (components.length === 0) {
    return null;
  }
  return (
    <Tab.Navigator
      screenOptions={({ route }) => ({
        headerShown: false,
        tabBarShowLabel: false,
        tabBarIcon: ({ focused, color, size }) => {
          let iconName;
          if (route.name === "Home") {
            iconName = focused ? "ios-people-circle" : "ios-people-circle-outline";
          }  else if (route.name === "MyProfile") {
            return (
              <Image source={{ uri: props.imageUri }} style={styles.logo} />
            );
          } else if (route.name === "Errand") {
            iconName = focused ? "ios-list" : "ios-list-outline";
          }
          return <Ionicons name={iconName} size={size} color={"white"} />;
        },
        tabBarStyle: {
          height: Dim.width * 0.15,
          width: Dim.width,
          backgroundColor: Colors.darkGreen,
          color: Colors.white,
          position: "absolute",
          marginBottom: 0,
        },
      })}
    >
      {components}
    </Tab.Navigator>
  );
};

const styles = StyleSheet.create({
  logo: {
    width: 35,
    height: 35,
  },
});

export default TabBar;
