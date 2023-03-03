import React from "react";
import {
  Text,
  View,
  StyleSheet,
  TouchableOpacity,
  Linking,
} from "react-native";
import { Title } from "react-native-paper";
import { Dim, Colors } from "../Constants";

const LocationCard = (props) => {
  const openMaps = () => {
    const scheme = Platform.select({
      ios: "maps:0,0?q=",
      android: "geo:0,0?q=",
    });
    const latLng = `${props.lat},${props.long}`;
    const label = props.mapsLabel;
    const url = Platform.select({
      ios: `${scheme}${label}@${latLng}`,
      android: `${scheme}${latLng}(${label})`,
    });

    Linking.openURL(url);
  };
  return (
    <TouchableOpacity onPress={openMaps}>
      <View style={Styles.container} noShadow={true}>
        <Title>{props.title}</Title>
        <Text>{props.address}</Text>
      </View>
    </TouchableOpacity>
  );
};

const Styles = StyleSheet.create({
  container: {
    backgroundColor: Colors.cream,
    alignContent: "center",
    width: Dim.width * 0.9,
    paddingVertical: 15,
    paddingHorizontal: 30,
    borderRadius: 10,
  },
  text: {
    marginTop: 30,
  },
});

export default LocationCard;
