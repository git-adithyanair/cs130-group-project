import React from "react";
import { Image, Text, View, StyleSheet, TouchableOpacity } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { Dim, Colors, Font } from "../Constants";

const RequestCard = (props) => {
  return (
    <View
      style={{
        ...styles.container,
        backgroundColor: props.selected ? Colors.darkGreen : Colors.lightGreen,
      }}
    >
      <TouchableOpacity
        style={{
          flexDirection: "row",
          flex: 1,
          marginRight: 30,
        }}
        onPress={props.onPress}
      >
        <View style={{ alignItems: "center" }}>
          <Image
            source={{
              uri: props.imageUri,
            }}
            style={styles.image}
          />
          <Text>{props.name}</Text>
        </View>
        <View
          style={{
            marginLeft: 20,
            justifyContent: "center",
          }}
        >
          <Text style={styles.storeText}>{props.storeName}</Text>
          <Text style={styles.storeText}>{props.storeAddress}</Text>
          {props.isUserRequest ? (
            <Text style={{ fontSize: 12, marginBottom: 5 }}>
              Community: {props.communityName}
            </Text>
          ) : null}
          <Text style={{ fontSize: 12 }}>{props.numItems} item(s) to get.</Text>
        </View>
      </TouchableOpacity>
      {!props.isUserRequest ? (
        <TouchableOpacity
          style={{ alignSelf: "center", padding: 10 }}
          onPress={props.onPressSelect}
        >
          <Ionicons
            name={props.selected ? "remove-circle" : "add-circle"}
            size={35}
            color={props.selected ? Colors.lightGreen : Colors.darkGreen}
          />
        </TouchableOpacity>
      ) : null}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    alignContent: "center",
    flexDirection: "row",
    width: Dim.width * 0.9,
    alignSelf: "center",
    borderRadius: 10,
    padding: 20,
    marginTop: 15,
    justifyContent: "space-between",
  },
  storeText: {
    fontSize: 14,
    fontWeight: Font.s3.weight,
    marginBottom: 5,
  },
  image: {
    width: Dim.width * 0.15,
    height: Dim.width * 0.15,
    borderRadius: 100,
    marginBottom: 10,
  },
  completeText: {
    color: Colors.darkGreen,
    fontWeight: "bold",
  },
});

export default RequestCard;
