import React from "react";
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

function Profile({ setPage }) {
  const dispatch = useDispatch();
  const sections = [
    "Change Address",
    "Change Name",
    "Change Profile Picture",
    "Join Community",
    "Create Community",
  ];

  const handleNavigation = (section) => {
    // handle navigation for pages here based on section name
    console.log(section);
  };

  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.content}>
        <Image source={require("../assets/logo.png")} />
        <View style={styles.listOfRequests}>
          <View>
            <Image
              style={styles.profileImage}
              source={{
                uri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png",
              }}
            />
          </View>
          <View style={styles.requestDetails}>
            <Text style={styles.titleText}>Angela</Text>
            <Text>Number of neighborhoods: 5</Text>
          </View>
        </View>
        <FlatList
          style={{ paddingVertical: 20 }}
          data={sections}
          scrollEnabled={false}
          renderItem={(itemData) => (
            <TouchableOpacity onPress={() => handleNavigation(itemData.item)}>
              <Text
                style={{
                  fontSize: 18,
                  fontWeight: "bold",
                  paddingVertical: 25,
                }}
              >
                {itemData.item}
              </Text>
            </TouchableOpacity>
          )}
          ItemSeparatorComponent={() => (
            <Divider
              orientation="horizontal"
              style={{ width: Dim.width * 0.8 }}
            ></Divider>
          )}
        ></FlatList>
        <Button
          title={"Sign out"}
          width={200}
          appButtonContainer={{ backgroundColor: Colors.lightGreen }}
          onPress={() => dispatch(setToken(""))}
        />
      </View>
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
    marginTop: 40,
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
