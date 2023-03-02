import React from "react";
import { Image, Text, View, StyleSheet, TouchableOpacity } from "react-native";
import { Card, Button, Title, Paragraph } from "react-native-paper";
import { Dim, Colors, Font, BorderRadius } from "../Constants";

const RequestCard = (props) => {
  return (
    <TouchableOpacity onPress={props.onPress}>
      <Card style={Styles.container}>
        <View style={Styles.user}>
          <Card.Content>
            <Image
              source={{
                uri: props.imageUri,
              }}
              style={Styles.image}
            />
            <Title>{props.name}</Title>
          </Card.Content>
          <Card.Content style={Styles.text}>
            <Text>
              <Paragraph>Store: </Paragraph>
              <Text>{props.storeName}</Text>
            </Text>
            <Text>
              <Paragraph>Items: </Paragraph>
              <Text>{props.numItems}</Text>
            </Text>
            {props.requestComplete ? (
              <Text style={Styles.completeText}>Complete</Text>
            ) : null}
          </Card.Content>
        </View>
      </Card>
    </TouchableOpacity>
  );
};

const Styles = StyleSheet.create({
  container: {
    backgroundColor: Colors.white,
    alignContent: "center",
    flexDirection: "row",
  },
  user: {
    flexDirection: "row",
    marginTop: 20,
  },
  image: {
    width: 75,
    height: 75,
    borderRadius: 75 / 2,
  },
  text: {
    marginTop: 30,
  },
  completeText: {
    color: Colors.darkGreen,
    fontWeight: "bold",
  },
});

export default RequestCard;
