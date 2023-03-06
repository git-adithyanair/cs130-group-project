import React, { useEffect, useState } from "react";
import Button from "../components/Button";
import { SafeAreaView, StyleSheet, FlatList, View, Text } from "react-native";
import CommunityCard from "../components/CommunityCard";
import { Dim, Colors, Font } from "../Constants";
import useRequest from "../hooks/useRequest";

const YourCommunities = (props) => {
  const [communityData, setCommunityData] = useState([]);

  const getCommunities = useRequest({
    url: "/user/community",
    method: "get",
    onSuccess: (data) => {
      data.forEach((community) => {
        setCommunityData((oldArray) => [
          ...oldArray,
          {
            members: community.member_count,
            communityId: community.community.id,
            communityName: community.community.name,
            distance:
              Math.round((community.community.range / 1609.344) * 100) / 100,
          },
        ]);
      });
    },
  });

  const func = async () => getCommunities.doRequest();

  useEffect(() => {
    const unsubscribe = props.navigation.addListener('focus', () => {
      setCommunityData([])
      func();
    });

    return unsubscribe;
  }, [props.navigation]);


  return (
    <SafeAreaView style={styles.wrapper}>
      <FlatList
        horizontal={false}
        numColumns={2}
        style={styles.list}
        contentContainerStyle={styles.container}
        columnWrapperStyle={{ justifyContent: "space-between" }}
        showsVerticalScrollIndicator={false}
        keyExtractor={(item) => Math.random().toString()}
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
          <View style={{ alignItems: "center" }}>
            <Button
              width={200}
              appButtonContainer={{ backgroundColor: Colors.lightGreen }}
              appButtonText={{ textTransform: "none" }}
              title={"Join More"}
              onPress={() => props.navigation.navigate("JoinCommunity")}
            ></Button>
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
    paddingBottom: 80,
    paddingTop: 10,
    alignSelf: "center",
  },
});

export default YourCommunities;
