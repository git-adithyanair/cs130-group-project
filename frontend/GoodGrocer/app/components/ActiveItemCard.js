import React, { useState } from "react";
import Button from "../components/Button";
import { ImageBackground, Text, View, StyleSheet } from "react-native";
import { Card } from "react-native-paper";
import { Dim, Colors } from "../Constants";

const ActiveItemCard = (props) => {
  const [found, setFound] = useState(
    props.item.found.Valid ? props.item.found.Bool : null
  );

  return (
    <Card style={Styles.container} elevation={0}>
      <Card.Content style={{ flexDirection: "row" }}>
        <View>
          <Text style={{ fontSize: 16, marginBottom: 10 }}>
            {props.item.name}
          </Text>
          <Text>
            <Text style={{ marginTop: 20, fontWeight: "bold" }}>Amount: </Text>
            <Text>
              {props.item.quantity} ({props.item.quantity_type})
            </Text>
          </Text>
          <Text>
            <Text style={{ marginTop: 20, fontWeight: "bold" }}>
              Preferred Brand:{" "}
            </Text>
            <Text>{props.item.preferred_brand.String}</Text>
          </Text>
          <Text>
            <Text style={{ marginTop: 20, fontWeight: "bold" }}>
              Extra Notes:{" "}
            </Text>
            <Text>{props.item.extra_notes.String}</Text>
          </Text>
        </View>
        <View style={{ marginLeft: "auto" }}>
          <ImageBackground
            source={
              props.item.image.Valid
                ? {
                    uri: "data:image/png;base64," + props.item.image.String,
                  }
                : require("../assets/grocery-item.png")
            }
            style={{
              width: Dim.width * 0.15,
              height: Dim.width * 0.15,
              borderRadius: 10,
            }}
          />
        </View>
      </Card.Content>
      <Card.Content style={{ marginTop: 10 }}>
        {found != null ? (
          <View>
            <Text style={Styles.itemFoundText}>
              Item {!found ? "Not " : ""}Found
            </Text>
          </View>
        ) : (
          <View
            style={{ flexDirection: "row", justifyContent: "space-evenly" }}
          >
            <Button
              width={Dim.width * 0.35}
              appButtonContainer={{
                backgroundColor: "#dce0de",
                height: 35,
                paddingVertical: 10,
              }}
              appButtonText={{
                textTransform: "none",
                fontWeight: "normal",
                fontSize: 14,
              }}
              title={"Not Found"}
              onPress={() => {
                setFound(false);
                props.onPressUpdateFound(false);
              }}
            />
            <Button
              width={Dim.width * 0.35}
              appButtonContainer={{
                backgroundColor: "#606e66",
                height: 35,
                paddingVertical: 10,
              }}
              appButtonText={{
                textTransform: "none",
                color: "white",
                fontSize: 14,
              }}
              title={"Found"}
              onPress={() => {
                setFound(true);
                props.onPressUpdateFound(true);
              }}
            />
          </View>
        )}
      </Card.Content>
    </Card>
  );
};

const Styles = StyleSheet.create({
  container: {
    backgroundColor: Colors.lightGreen,
  },
  text: {
    marginTop: 30,
  },
  itemFoundText: {
    textAlign: "center",
    color: "#606e66",
    fontWeight: "bold",
    paddingTop: 20,
  },
});

export default ActiveItemCard;
