import React from "react";
import { Image, Text, View, StyleSheet, TouchableOpacity } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { Dim, Colors, Font } from "../Constants";

const RequestCard = (props) => {
  return (
    <View
      style={{
        ...styles.container,
        backgroundColor: props.requestComplete
          ? Colors.darkGreen
          : Colors.lightGreen,
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
          <Text style={{ fontSize: 12 }}>{props.numItems} item(s) to get.</Text>
        </View>
      </TouchableOpacity>
      <View style={{ alignSelf: "center", padding: 10 }}>
        <Ionicons
          name={props.requestComplete ? "checkmark-done-circle" : "basket"}
          size={35}
          color={props.requestComplete ? Colors.lightGreen : Colors.darkGreen}
        />
      </View>
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
