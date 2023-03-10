import { React, useEffect, useState } from "react";
import { SafeAreaView, StyleSheet, FlatList, View, Text } from "react-native";
import CommunityCard from "../components/CommunityCard";
import SearchBar from "../components/SearchBar";
import { Dim, Colors, Font } from "../Constants";
import useRequest from "../hooks/useRequest";

const JoinCommunity = (props) => {
  const getDistance = (x1, y1, x2, y2) => {
    return Math.round(Math.sqrt(Math.pow(y2-y1,2)+Math.pow(x2-x1,2))*100)/100
  }
  
  const [allCommunities, setAllCommunities] = useState([]);
  const [userCommunities, setUserCommunities] = useState([]);
  const [communities, setCommunities] = useState([]);
  const [searchedCommunities, setSearchedCommunities] = useState([]);
  const [community, setCommunity] = useState("");

  const userXCoord = props.route.params.userXCoord
  const userYCoord = props.route.params.userYCoord

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
          !userCommunities.find((comm) => comm.community.id === community.community.id)
      )
    );

    setSearchedCommunities(
      allCommunities.filter(
        (community) =>
          !userCommunities.find((comm) => comm.community.id === community.community.id)
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
              getDistance(userXCoord,userYCoord, itemData.item.community.center_x_coord, itemData.item.community.center_y_coord)
            }
            numberOfMembers={itemData.item.member_count}
            joinCommunity={true}
            communityId={itemData.item.community.id}
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
              No communities to join at the moment
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

export default JoinCommunity;
