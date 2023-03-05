import React from "react";
import {
  Text,
  View,
  StyleSheet,
  TouchableOpacity,
  Linking,
} from "react-native";
import Ionicons from "react-native-vector-icons/Ionicons";
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
        <Ionicons
          name={"navigate-circle"}
          size={30}
          color={"white"}
          style={{ paddingRight: 15, alignSelf: "center" }}
        />
        <View style={{ paddingRight: 25 }}>
          <Text style={{ fontWeight: "bold", color: "white" }}>
            {props.title}
          </Text>
          <Text style={{ color: "white" }}>{props.address}</Text>
        </View>
      </View>
    </TouchableOpacity>
  );
};

const Styles = StyleSheet.create({
  container: {
    backgroundColor: Colors.darkGreen,
    alignContent: "center",
    width: Dim.width * 0.9,
    paddingVertical: 15,
    paddingHorizontal: 30,
    borderRadius: 10,
    flexDirection: "row",
  },
  text: {
    marginTop: 30,
  },
});

export default LocationCard;
