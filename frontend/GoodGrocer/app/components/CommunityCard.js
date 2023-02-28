import { React, useState } from "react";
import { View, StyleSheet, TouchableOpacity, Text } from "react-native";
import Button from "./Button";
import { Entypo, Ionicons } from "@expo/vector-icons";
import { Dim, Colors, Font, BorderRadius } from "../Constants";

const CommunityCard = (props) => {
  const [joined, setJoined] = useState(false);
  return (
    <TouchableOpacity
      style={styles.main}
      onPress={props.onPressCommunity}
      disabled={props.joinCommunity && !joined ? true : false}
    >
      <View style={styles.mainView}>
        <Text style={styles.communityName}>{props.communityName}</Text>
        <View style={styles.mainInfoWrapper}>
          <View style={styles.infoWrapper}>
            <Entypo
              name="location-pin"
              color={Colors.lightGreen}
              size={20}
            ></Entypo>
            <Text style={styles.info}>{props.distanceFromUser}mi</Text>
          </View>
          <View style={{ ...styles.infoWrapper, paddingLeft: 20 }}>
            <Ionicons
              name="people-sharp"
              color={Colors.lightGreen}
              size={20}
            ></Ionicons>
            <Text style={styles.info}>{props.numberOfMembers}</Text>
          </View>
        </View>
      </View>
      {props.joinCommunity ? (
        <Button
          onPress={() => setJoined(true)}
          disabled={joined ? true : false}
          title={joined ? "joined!" : "join"}
          appButtonContainer={{
            width: "85%",
            height: "27%",
            alignSelf: "center",
            backgroundColor: joined ? Colors.lightGreen : Colors.darkGreen,
          }}
          appButtonText={{
            textTransform: "none",
            color: joined ? Colors.darkGreen : Colors.white,
            fontSize: Font.s3.size,
          }}
        ></Button>
      ) : null}
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  main: {
    backgroundColor: Colors.cream,
    height: Dim.height * 0.15,
    width: Dim.width * 0.4,
    borderRadius: BorderRadius,
    justifyContent: "center",
  },
  mainView: {
    justifyContent: "center",
    paddingLeft: 15,
  },
  communityName: {
    color: Colors.darkGreen,
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
  },
  mainInfoWrapper: {
    flexDirection: "row",
    marginTop: 12,
  },
  infoWrapper: {
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center",
  },
  info: {
    fontSize: Font.s3.size,
    fontFamily: Font.s3.family,
    fontWeight: Font.s3.weight,
    color: Colors.darkGreen,
  },
});

export default CommunityCard;
