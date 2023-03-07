import { React, useState } from "react";
import { View, StyleSheet, TouchableOpacity, Text } from "react-native";
import Button from "./Button";
import { Entypo, Ionicons } from "@expo/vector-icons";
import { Dim, Colors, Font, BorderRadius } from "../Constants";
import useRequest from "../hooks/useRequest";

const CommunityCard = (props) => {
  const id = props.communityId;
  const [loading, setLoading] = useState(false);
  const [joined, setJoined] = useState(false);

  const joinCommunity = useRequest({
    url: "/community/join",
    method: "post",
    body: {
      id: id,
    },
    onSuccess: (data) => {
      setLoading(false);
      setJoined(true);
    },
    onFail: () => setLoading(false),
  });

  const handleJoined = async () => {
    setLoading(true);
    await joinCommunity.doRequest();
  };

  return (
    <TouchableOpacity
      style={styles.main}
      onPress={props.onPressCommunity}
      disabled={props.joinCommunity ? true : false}
    >
      <View style={styles.mainView}>
        <Text
          style={{
            color: Colors.darkGreen,
            fontSize: Font.s1.size,
            fontFamily: Font.s1.family,
            fontWeight: Font.s1.weight,
            ...props.communityNameStyle,
          }}
        >
          {props.communityName}
        </Text>
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
          onPress={() => handleJoined()}
          disabled={joined ? true : false}
          title={joined ? "joined!" : "join"}
          appButtonContainer={{
            width: "85%",
            height: "27%",
            alignSelf: "center",
            backgroundColor: joined ? Colors.lightGreen : Colors.darkGreen,
            marginTop: 8,
          }}
          appButtonText={{
            textTransform: "none",
            color: joined ? Colors.darkGreen : Colors.white,
            fontSize: Font.s3.size,
          }}
          loading={loading}
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
