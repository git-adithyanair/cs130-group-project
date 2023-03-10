import React from "react";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  Image,
  View,
  FlatList,
} from "react-native";
import ItemCard from "../components/ItemCard";
import { Dim, Font, Colors } from "../Constants";

const RequestDetail = (props) => {
  const { items, user } = props.route.params;

  return (
    <SafeAreaView style={styles.wrapper}>
      <FlatList
        contentContainerStyle={styles.container}
        style={styles.list}
        data={items}
        renderItem={({ item }) => (
          <ItemCard
            name={item.name}
            quantity={item.quantity}
            quantityType={item.quantity_type}
            preferredBrand={
              item.preferred_brand.Valid ? item.preferred_brand.String : null
            }
            extraNotes={item.extra_notes.Valid ? item.extra_notes.String : null}
            imageUri={item.image.Valid ? item.image.String : null}
          />
        )}
        keyExtractor={(item) => item.id}
        ListHeaderComponent={() => (
          <View style={{ alignItems: "center", paddingBottom: 20, flex: 1 }}>
            <Text style={styles.title}>{user.name}'s Request</Text>
            <Image
              source={{
                uri: user.profileImage,
              }}
              style={styles.profilePic}
            />
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
      />
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
    paddingBottom: 30,
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
  phoneNumberText: {
    paddingBottom: 10,
    fontWeight: "bold",
    color: Colors.darkGreen,
  },
});

export default RequestDetail;
