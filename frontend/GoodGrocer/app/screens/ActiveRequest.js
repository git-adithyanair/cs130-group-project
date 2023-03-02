import React, { useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  Image,
  View,
  FlatList,
} from "react-native";
import axios from "axios";
import { Dim, Colors, Font, API_URL } from "../Constants";
import ActiveItemCard from "../components/ActiveItemCard";
import StoreCard from "../components/StoreCard";
import { useSelector } from "react-redux";

const ActiveRequest = ({ route, navigation }) => {
  const { name, profileImage, items, store } = route.params;
  const token = useSelector((state) => state.user.token);

  const updateFound = async (found, id) => {
    axios
      .post(
        `${API_URL}/item/update-status`,
        {
          id,
          found,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      )
      .then(({ data }) => {
        console.log(data);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <SafeAreaView style={styles.wrapper}>
      <FlatList
        contentContainerStyle={styles.container}
        style={styles.list}
        data={items}
        renderItem={(itemData) => (
          <ActiveItemCard
            item={itemData.item}
            onPressUpdateFound={(found) => updateFound(found, itemData.item.id)}
          />
        )}
        keyExtractor={(item) => item.id}
        ListHeaderComponent={() => (
          <View style={{ alignItems: "center", paddingBottom: 20, flex: 1 }}>
            <Text style={styles.title}>{name}'s Request</Text>
            <Image
              source={{
                uri: profileImage,
              }}
              style={styles.profilePic}
            />
            {store != null ? (
              <StoreCard
                style={{ flex: 1 }}
                store={store.name}
                address={store.address}
                lat={store.x_coord}
                long={store.y_coord}
              />
            ) : null}
          </View>
        )}
        ItemSeparatorComponent={() => (
          <View
            style={{
              height: 15,
              width: Dim.width,
            }}
          />
        )}
      ></FlatList>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    flex: 1,
    backgroundColor: "#fff",
  },
  container: {
    width: Dim.width * 0.9,
    alignSelf: "center",
    paddingTop: 10,
    paddingBottom: 80,
  },
  title: {
    paddingTop: 10,
    paddingBottom: 20,
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
  },
  content: {
    alignItems: "center",
  },
  list: {
    flex: 1,
  },
  noErrandText: {
    fontSize: Font.s2.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s3.weight,
    color: Colors.darkGreen,
    paddingHorizontal: 10,
  },
  profilePic: {
    width: 100,
    height: 100,
    borderRadius: 100 / 2,
    marginBottom: 20,
  },
});

export default ActiveRequest;
