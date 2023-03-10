import React, { useEffect, useState } from "react";
import { useDispatch } from "react-redux";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  Image,
  View,
  TouchableOpacity,
  FlatList,
} from "react-native";
import Button from "../components/Button";
import { setToken } from "../store/actions";
import { Colors, Dim } from "../Constants";
import { Divider } from "react-native-elements";
import useRequest from "../hooks/useRequest";

function Profile({ navigation }) {
  const dispatch = useDispatch();
  const sections = [
    "My Requests",
    "Change Address",
    "Change Name",
    "Change Profile Picture",
    "Join Community",
    "Create Community",
  ];
  const [userData, setUserData] = useState({});

  const getUserInfo = useRequest({
    url: "/user",
    method: "get",
    onSuccess: (data) => {
      setUserData(data);
    },
  });

  const userInfo = async () => await getUserInfo.doRequest();

  useEffect(() => {
    const unsubscribe = navigation.addListener("focus", () => {
      userInfo();
    });
    return unsubscribe;
  }, [navigation]);

  const handleNavigation = (section) => {
    switch (section) {
      case "My Requests":
        navigation.navigate("UserRequests", {
          user: userData,
        });
        break;
      case "Change Profile Picture":
        navigation.navigate("UpdateProfilePicture", {
          user: userData,
        });
        break;
      case "Join Community":
        navigation.navigate("JoinCommunity", {
          userXCoord: userData.x_coord,
          userYCoord: userData.y_coord
        });
        break;
      case "Create Community":
        navigation.navigate("CreateCommunity");
        break;
      case "Change Name":
        navigation.navigate("ChangeName");
        break;
      case "Change Address":
        navigation.navigate("ChangeAddress");
        break;
      default:
        break;
    }
    console.log(section);
  };

  return (
    <SafeAreaView style={styles.container}>
      <FlatList
        data={sections}
        renderItem={(itemData) => (
          <TouchableOpacity
            onPress={() => handleNavigation(itemData.item)}
            style={{ alignItems: "center" }}
          >
            <Text
              style={{
                fontSize: 18,
                fontWeight: "bold",
                paddingVertical: 25,
                alignContent: "center",
                width: Dim.width * 0.8,
              }}
            >
              {itemData.item}
            </Text>
          </TouchableOpacity>
        )}
        ItemSeparatorComponent={() => (
          <View style={{ alignItems: "center" }}>
            <Divider
              orientation="horizontal"
              style={{ width: Dim.width * 0.8 }}
            ></Divider>
          </View>
        )}
        ListHeaderComponent={() => (
          <View style={styles.content}>
            <Image
              style={{ alignItems: "center" }}
              source={require("../assets/logo.png")}
            />
            <View style={styles.listOfRequests}>
              <View>
                <Image
                  style={styles.profileImage}
                  source={{
                    uri: userData.profile_picture,
                  }}
                />
              </View>
              <View style={styles.requestDetails}>
                <Text style={styles.titleText}>{userData.full_name}</Text>
                <Text>Number of communities: {userData.community_count}</Text>
              </View>
            </View>
          </View>
        )}
        ListFooterComponent={() => (
          <View style={{ alignItems: "center", paddingVertical: 10 }}>
            <Button
              title={"Sign out"}
              width={200}
              appButtonContainer={{ backgroundColor: Colors.lightGreen }}
              onPress={() => dispatch(setToken(""))}
            />
          </View>
        )}
      ></FlatList>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
  },
  content: {
    alignItems: "center",
    marginTop: 20,
  },
  listOfRequests: {
    display: "flex",
    flexDirection: "row",
    paddingTop: 20,
  },
  requestDetails: {
    paddingLeft: 10,
    alignItems: "center",
    marginTop: 10,
  },
  profileImage: {
    width: 80,
    height: 80,
    borderRadius: 80 / 2,
  },
  titleText: {
    fontSize: 25,
    textAlign: "left",
    alignSelf: "flex-start",
  },
});

export default Profile;
