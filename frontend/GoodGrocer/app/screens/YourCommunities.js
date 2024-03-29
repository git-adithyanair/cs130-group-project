import React, { useEffect, useState, useRef } from "react";
import Button from "../components/Button";
import { SafeAreaView, StyleSheet, FlatList, View, Text } from "react-native";
import CommunityCard from "../components/CommunityCard";
import { Dim, Colors, Font } from "../Constants";
import useRequest from "../hooks/useRequest";
import Loading from "./Loading";
import getDistanceFromLatLonInMi from "../helpers/distance";

const YourCommunities = (props) => {
  const [communityData, setCommunityData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [loadingUserCoords, setLoadingUserCoords] = useState(true);
  const [userXCoord, setUserXCoord] = useState(0.0);
  const [userYCoord, setUserYCoord] = useState(0.0);

  const getUserInfo = useRequest({
    url: "/user",
    method: "get",
    onSuccess: (data) => {
      setUserXCoord(data.x_coord);
      setUserYCoord(data.y_coord);
      setLoadingUserCoords(false);
    },
    onFail: () => setLoadingUserCoords(false),
  });

  const userInfo = async () => await getUserInfo.doRequest();

  const getCommunities = useRequest({
    url: "/user/community",
    method: "get",
    onSuccess: (data) => {
      let communities = [];
      data.forEach((community) => {
        communities.push({
          members: community.member_count,
          communityId: community.community.id,
          communityName: community.community.name,
          distance: getDistanceFromLatLonInMi(
            community.community.center_x_coord,
            community.community.center_y_coord,
            userXCoord,
            userYCoord
          ),
        });
      });
      setCommunityData(communities);
      setLoading(false);
    },
    onFail: () => setLoading(false),
  });

  const getUserCommunities = async () => getCommunities.doRequest();

  useEffect(() => {
    userInfo();
    if (!loadingUserCoords) {
      getUserCommunities();
    }
  }, [props.navigation, loadingUserCoords]);

  useEffect(() => {
    const unsubscribe = props.navigation.addListener("focus", () => {
      setCommunityData([]);
      if (!loadingUserCoords) {
        getUserCommunities();
      }
    });
    return unsubscribe;
  }, [props.navigation, loadingUserCoords]);

  // useEffect(() => {
  //     if(!loadingUserCoords){
  //       setLoading(true)
  //       setCommunityData([]);
  //       getUserCommunities();
  //     }
  // }, [props.navigation,loadingUserCoords]);

  if (loading || loadingUserCoords) {
    return <Loading />;
  }

  return (
    <SafeAreaView style={styles.wrapper}>
      <FlatList
        horizontal={false}
        numColumns={2}
        style={styles.list}
        contentContainerStyle={styles.container}
        columnWrapperStyle={{ justifyContent: "space-between" }}
        showsVerticalScrollIndicator={false}
        keyExtractor={(item) => item.communityId}
        data={communityData}
        renderItem={(itemData) => (
          <CommunityCard
            communityName={itemData.item.communityName}
            distanceFromUser={itemData.item.distance}
            numberOfMembers={itemData.item.members}
            onPressCommunity={() =>
              props.navigation.navigate("RequestList", {
                communityId: itemData.item.communityId,
                communityName: itemData.item.communityName,
              })
            }
          />
        )}
        ItemSeparatorComponent={() => (
          <View
            style={{
              height: 15,
              width: Dim.width,
            }}
          />
        )}
        ListFooterComponent={() => (
          <View style={{ alignItems: "center", paddingBottom: 20 }}>
            <Button
              width={200}
              appButtonContainer={{
                backgroundColor: Colors.lightGreen,
                marginTop: 20,
              }}
              appButtonText={{ textTransform: "none" }}
              title={"Join More"}
              onPress={() =>
                props.navigation.navigate("JoinCommunity", {
                  userXCoord,
                  userYCoord,
                })
              }
            />
            <Button
              width={200}
              appButtonContainer={{
                backgroundColor: Colors.lightGreen,
                marginTop: 20,
              }}
              appButtonText={{ textTransform: "none" }}
              title={"Create a Community"}
              onPress={() => props.navigation.navigate("CreateCommunity")}
            />
          </View>
        )}
        ListEmptyComponent={() => (
          <View
            style={{ alignItems: "center", height: "100%", paddingTop: "50%" }}
          >
            <Text
              style={{
                fontFamily: Font.s1.family,
                fontSize: Font.s1.size,
                alignSelf: "center",
              }}
            >
              Join a community to get started! 🧸
            </Text>
          </View>
        )}
      ></FlatList>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
    backgroundColor: Colors.white,
  },
  list: {
    flex: 1,
  },
  container: {
    width: Dim.width * 0.9,
    paddingTop: 10,
    alignSelf: "center",
  },
});

export default YourCommunities;
