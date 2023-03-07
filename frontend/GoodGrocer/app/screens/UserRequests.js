import React, { useEffect, useState } from "react";
import { SafeAreaView, StyleSheet, Text, FlatList, View } from "react-native";
import RequestCard from "../components/RequestCard";
import { Colors, Font } from "../Constants";
import useRequest from "../hooks/useRequest";
import { createMaterialTopTabNavigator } from "@react-navigation/material-top-tabs";

const UserRequests = (props) => {
  const { user } = props.route.params;
  const Tab = createMaterialTopTabNavigator();

  const [userRequestData, setUserRequestData] = useState({});
  const [loading, setLoading] = useState(true);

  const getUserRequests = useRequest({
    url: "/user/requests",
    method: "get",
    onSuccess: (data) => {
      const requests = { pending: [], in_progress: [], complete: [] };
      ["pending", "in_progress", "complete"].map((status, _) =>
        data[status].forEach((requestData) => {
          requests[status].push({
            name: user.full_name,
            storeName: requestData.store ? requestData.store.name : "Any Store",
            storeAddress: requestData.store ? requestData.store.address : "N/A",
            id: requestData.request.id,
            numItems: requestData.items.length,
            imageUri: user.profile_picture,
            items: requestData.items,
            communityName: requestData.community_name,
          });
        })
      );
      setUserRequestData(requests);
    },
  });

  useEffect(() => {
    if (loading) {
      const getRequests = async () => getUserRequests.doRequest();
      getRequests();
      setLoading(false);
    }
  }, [loading]);

  return (
    <SafeAreaView style={styles.container}>
      <Tab.Navigator
        screenOptions={{
          tabBarLabelStyle: {
            fontFamily: Font.s1.family,
            textTransform: "none",
            fontWeight: Font.s1.weight,
          },
          tabBarIndicatorStyle: { backgroundColor: Colors.darkGreen },
        }}
      >
        <Tab.Screen
          name="Pending"
          children={() => (
            <UserRequestTab
              userRequestData={userRequestData.pending}
              navigation={props.navigation}
            ></UserRequestTab>
          )}
        />
        <Tab.Screen
          name="In Progress"
          children={() => (
            <UserRequestTab
              userRequestData={userRequestData.in_progress}
              navigation={props.navigation}
            ></UserRequestTab>
          )}
        />
        <Tab.Screen
          name="Complete"
          children={() => (
            <UserRequestTab
              userRequestData={userRequestData.complete}
              navigation={props.navigation}
            ></UserRequestTab>
          )}
        />
      </Tab.Navigator>
    </SafeAreaView>
  );
};

const UserRequestTab = (props) => {
  return (
    <FlatList
      data={props.userRequestData}
      contentContainerStyle={{ paddingBottom: 20 }}
      renderItem={({ item }) => (
        <RequestCard
          key={item.id}
          name={item.name}
          storeName={item.storeName}
          storeAddress={item.storeAddress}
          numItems={item.numItems}
          imageUri={item.imageUri}
          communityName={item.communityName}
          isUserRequest={true}
          onPress={() =>
            props.navigation.navigate("RequestDetail", {
              requestId: item.id,
              storeName: item.storeName,
              items: item.items,
              user: {
                name: item.name,
                profileImage: item.imageUri,
              },
            })
          }
        />
      )}
      keyExtractor={(item) => item.id}
    />
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.white,
  },
});

export default UserRequests;
