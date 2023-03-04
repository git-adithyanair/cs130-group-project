import { React, useEffect, useState } from "react";
import { SafeAreaView, StyleSheet, FlatList, View, Text } from "react-native";
import CommunityCard from "../components/CommunityCard";
import SearchBar from "../components/SearchBar";
import { Dim, Colors, Font } from "../Constants";
import useRequest from "../hooks/useRequest";

const JoinCommunity = (props) => {
  let allCommunities = [];
  let userCommunities = [];
  const [communities, setCommunities] = useState([]);
  const [community, setCommunity] = useState("");

  // const data = [
  //   { communityName: "Westwood", distance: 0.5, members: 10 },
  //   { communityName: "Brentwood", distance: 2, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  //   { communityName: "Beverly Hills", distance: 5, members: 10 },
  // ];

  const getAllCommunities = useRequest({
    url: "/community",
    method: "get",
    onSuccess: (data) => {
      allCommunities = data;
      console.log("all communities", allCommunities);
      // data.forEach((community) => {
      //   setCommunityData((oldArray) => [
      //     ...oldArray,
      //     {
      //       members: community.member_count,
      //       communityId: community.community.id,
      //       communityName: community.community.name,
      //       distance:
      //         Math.round((community.community.range / 1609.344) * 100) / 100,
      //     },
      //   ]);
      // });
    },
  });

  const getUserCommunities = useRequest({
    url: "/user/community",
    method: "get",
    onSuccess: (data) => {
      userCommunities = data;
      console.log("user communities", userCommunities);
      // data.forEach((community) => {
      //   setCommunityData((oldArray) => [
      //     ...oldArray,
      //     {
      //       members: community.member_count,
      //       communityId: community.community.id,
      //       communityName: community.community.name,
      //       distance:
      //         Math.round((community.community.range / 1609.344) * 100) / 100,
      //     },
      //   ]);
      // });
    },
  });

  const allComm = async () => await getAllCommunities.doRequest();
  const userComm = async () => await getUserCommunities.doRequest();
  
  useEffect(() => {
    allComm();
    userComm();

    allCommunities.filter((community) => {
      !userCommunities.find((comm) => comm.place_id === community.place_id);
    });

    console.log("filtered communities", allCommunities);
  }, []);

  const searchCommunities = (text) => {
    setCommunity(text);
    if (!text) {
      setCommunities(data);
    } else {
      setCommunities(
        data.filter((item) => {
          return item.communityName
            .toLowerCase()
            .startsWith(text.toLowerCase());
        })
      );
    }
  };

  return (
    <SafeAreaView style={styles.wrapper}>
      <SearchBar
        style={{ marginVertical: 10 }}
        placeholder={"Search..."}
        value={community}
        onChangeText={(text) => searchCommunities(text)}
      />
      <FlatList
        horizontal={false}
        numColumns={2}
        style={styles.list}
        contentContainerStyle={styles.container}
        columnWrapperStyle={{ justifyContent: "space-between" }}
        showsVerticalScrollIndicator={false}
        keyExtractor={(item) => Math.random().toString()}
        data={communities}
        renderItem={(itemData) => (
          <CommunityCard
            communityName={itemData.item.communityName}
            distanceFromUser={itemData.item.distance}
            numberOfMembers={itemData.item.members}
            joinCommunity={true}
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
          <View style={{ height: Dim.width * 0.05 }}></View>
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

export default JoinCommunity;
