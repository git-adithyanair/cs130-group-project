import React, { useState } from "react";
import Button from "../components/Button";
import { Image, Text, View, StyleSheet } from "react-native";
import { Card, Title } from "react-native-paper";
import { Dim, Colors } from "../Constants";

const ActiveItemCard = (props) => {
  const [found, setFound] = useState(
    props.item.found.Valid ? props.item.found.Bool : null
  );

  return (
    <Card style={Styles.container}>
      <Card.Content style={{ flexDirection: "row" }}>
        <View>
          <Text>
            <Title style={{ marginTop: 20, fontWeight: "bold" }}>Item: </Title>
            <Title>{props.item.name}</Title>
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
        <View style={{ width: 75, height: 75, marginLeft: "auto" }}>
          <Image
            source={{
              uri: props.item.image.String,
            }}
            style={{ flex: 1 }}
          />
        </View>
      </Card.Content>
      <Card.Content>
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
              appButtonContainer={{ backgroundColor: Colors.cream }}
              appButtonText={{
                textTransform: "none",
                fontWeight: "normal",
                fontSize: 16,
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
                backgroundColor: Colors.lightGreen,
                fontWeight: "normal",
                fontSize: 16,
              }}
              appButtonText={{ textTransform: "none" }}
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
    backgroundColor: Colors.white,
  },
  text: {
    marginTop: 30,
  },
  itemFoundText: {
    textAlign: "center",
    color: Colors.darkGreen,
    fontWeight: "bold",
    paddingTop: 20,
  },
});

export default ActiveItemCard;
