import { React, useEffect, useState } from "react";
import { SafeAreaView, StyleSheet, FlatList, View, Text } from "react-native";
import CommunityCard from "../components/CommunityCard";
import SearchBar from "../components/SearchBar";
import { Dim, Colors, Font } from "../Constants";
import useRequest from "../hooks/useRequest";

const JoinCommunity = (props) => {
  const [allCommunities, setAllCommunities] = useState([]);
  const [userCommunities, setUserCommunities] = useState([]);
  const [communities, setCommunities] = useState([]);
  const [searchedCommunities, setSearchedCommunities] = useState([]);
  const [community, setCommunity] = useState("");

  const getAllCommunities = useRequest({
    url: "/community",
    method: "get",
    onSuccess: (data) => {
      setAllCommunities(data);
    },
  });

  const getUserCommunities = useRequest({
    url: "/user/community",
    method: "get",
    onSuccess: (data) => {
      setUserCommunities(data);
    },
  });

  const allComm = async () => await getAllCommunities.doRequest();
  const userComm = async () => await getUserCommunities.doRequest();

  useEffect(() => {
    allComm();
    userComm();
  }, []);

  useEffect(() => {
    setCommunities(
      allCommunities.filter(
        (community) =>
          !userCommunities.find((comm) => comm.place_id === community.place_id)
      )
    );

    setSearchedCommunities(
      allCommunities.filter(
        (community) =>
          !userCommunities.find((comm) => comm.place_id === community.place_id)
      )
    );
  }, [allCommunities, userCommunities]);

  const searchCommunities = (text) => {
    setCommunity(text);
    if (!text) {
      setSearchedCommunities(communities);
    } else {
      setSearchedCommunities(
        communities.filter((item) => {
          return item.community.name
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
        data={searchedCommunities}
        renderItem={(itemData) => (
          <CommunityCard
            communityNameStyle={{
              fontSize:
                itemData.item.community.name.length > 10
                  ? Font.s2.size
                  : Font.s1.size,
            }}
            communityName={itemData.item.community.name}
            distanceFromUser={
              Math.round((itemData.item.community.range / 1609.344) * 100) / 100
            }
            numberOfMembers={itemData.item.member_count}
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
