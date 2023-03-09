import React from "react";
import { TouchableOpacity, Text } from "react-native";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { createStackNavigator } from "@react-navigation/stack";
import { useNavigation, CommonActions } from "@react-navigation/native";
import Ionicons from "react-native-vector-icons/Ionicons";
import JoinCommunity from "../screens/JoinCommunity";
import YourCommunities from "../screens/YourCommunities";
import Profile from "../screens/Profile";
import { Dim, Colors } from "../Constants";
import ActiveErrand from "../screens/ActiveErrand";
import ActiveRequest from "../screens/ActiveRequest";
import RequestList from "../screens/RequestList";
import RequestDetail from "../screens/RequestDetail";
import CreateCommunity from "../screens/CreateCommunity";
import AddStores from "../screens/AddStores";
import Buy from "../screens/Buy";
import OrderCreated from "../screens/OrderCreated";
import UserRequests from "../screens/UserRequests";

const HomeStack = createStackNavigator();
const ErrandStack = createStackNavigator();
const ProfileStack = createStackNavigator();

const HomeStackScreen = () => {
  return (
    <HomeStack.Navigator screenOptions={{ headerTintColor: Colors.darkGreen }}>
      <HomeStack.Screen
        name="YourCommunities"
        component={YourCommunities}
        options={{ title: "Your Communities" }}
      />
      <HomeStack.Screen
        name="RequestList"
        component={RequestList}
        options={({ navigation }) => ({
          title: "Community Requests",
          headerRight: () => (
            <TouchableOpacity onPress={() => navigation.navigate("Buy")}>
              <Ionicons
                name={"add-circle"}
                size={30}
                color={Colors.darkGreen}
              />
            </TouchableOpacity>
          ),
        })}
      />
      <HomeStack.Screen
        name="JoinCommunity"
        component={JoinCommunity}
        options={{
          title: "Join Community",
        }}
      />
      <HomeStack.Screen
        name="RequestDetail"
        component={RequestDetail}
        options={{ title: "Request Details" }}
      />
      <HomeStack.Screen
        name="CreateCommunity"
        component={CreateCommunity}
        options={{ title: "Create a Community" }}
      />
      <HomeStack.Screen
        name="AddStores"
        component={AddStores}
        options={{ title: "Add Stores" }}
      />
      <HomeStack.Screen
        name="Buy"
        component={Buy}
        options={{ title: "Create Order" }}
      />
      <HomeStack.Screen
        name="OrderCreated"
        component={OrderCreated}
        options={{ title: "Order Created" }}
      />
    </HomeStack.Navigator>
  );
};

const ProfileStackScreen = () => {
  return (
    <ProfileStack.Navigator
      screenOptions={{ headerTintColor: Colors.darkGreen }}
    >
      <ProfileStack.Screen name="Profile" component={Profile} />
      <ProfileStack.Screen
        name="UserRequests"
        component={UserRequests}
        options={{
          title: "My Requests",
        }}
      />
      <ProfileStack.Screen
        name="RequestDetail"
        component={RequestDetail}
        options={{ title: "Request Details" }}
      />
      <ProfileStack.Screen
        name="JoinCommunity"
        component={JoinCommunity}
        options={{
          title: "Join Community",
        }}
      />
      <ProfileStack.Screen
        name="CreateCommunity"
        component={CreateCommunity}
        options={{ title: "Create a Community" }}
      />
    </ProfileStack.Navigator>
  );
};

const ErrandStackScreen = () => {
  return (
    <ErrandStack.Navigator
      screenOptions={{ headerTintColor: Colors.darkGreen }}
    >
      <ErrandStack.Screen
        name="ActiveErrand"
        component={ActiveErrand}
        options={{ title: "Active Errand" }}
      />
      <ErrandStack.Screen
        name="ActiveRequest"
        component={ActiveRequest}
        options={{ title: "" }}
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
        tabBarIcon: ({ focused, size }) => {
          let iconName;
          if (route.name === "Home") {
            iconName = focused
              ? "ios-people-circle"
              : "ios-people-circle-outline";
          } else if (route.name === "MyProfile") {
            iconName = focused ? "person" : "person-outline";
          } else if (route.name === "Errand") {
            iconName = focused ? "cart" : "cart-outline";
          }
          return <Ionicons name={iconName} size={size} color={"white"} />;
        },
        tabBarStyle: {
          height: Dim.height * 0.1,
          width: Dim.width,
          backgroundColor: Colors.darkGreen,
          color: Colors.white,
        },
      })}
    >
      {components}
    </Tab.Navigator>
  );
};

export default TabBar;
