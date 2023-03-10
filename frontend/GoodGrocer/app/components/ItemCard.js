import React from "react";
import { Image, Text, View, StyleSheet, ImageBackground } from "react-native";
import { Dim, Colors, Font } from "../Constants";

const ItemCard = (props) => {
  return (
    <View style={styles.container}>
      <ImageBackground
        source={
          props.imageUri
            ? {
                uri: "data:image/png;base64," + props.imageUri,
              }
            : require("../assets/grocery-item.png")
        }
        style={{
          width: Dim.width * 0.15,
          height: Dim.width * 0.15,
          borderRadius: 100,
          // marginBottom: 10,
        }}
      />
      <View style={{ marginLeft: 20 }}>
        <Text
          style={{
            fontSize: Font.s3.size,
            fontWeight: Font.s3.weight,
          }}
        >
          {props.name}
        </Text>
        <View style={{ height: 4 }} />
        {props.quantityType !== "numerical" ? (
          <Text>
            Amount: {props.quantity.toFixed(2)} {props.quantityType}
          </Text>
        ) : (
          <Text>Amount: {Math.round(props.quantity)} count</Text>
        )}
        <View style={{ height: 4 }} />
        {props.preferredBrand ? (
          <Text>Prefers the {props.preferredBrand} brand.</Text>
        ) : null}
        <View style={{ height: 4 }} />
        {props.extraNotes ? <Text>Extra notes: {props.extraNotes}</Text> : null}
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    backgroundColor: Colors.lightGreen,
    borderRadius: 10,
    padding: 20,
    flexDirection: "row",
    alignItems: "center",
  },
  text: {
    marginTop: 30,
  },
});

export default ItemCard;
