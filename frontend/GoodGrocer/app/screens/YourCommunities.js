import React from "react";
import Button from "../components/Button";
import { SafeAreaView, StyleSheet, FlatList, View } from "react-native";
import CommunityCard from "../components/CommunityCard";
import { Dim, Colors, Font } from "../Constants";

const YourCommunities = (props, navigation) => {
  const data = [
    { communityName: "Westwood", distance: 0.5, members: 10 },
    { communityName: "Brentwood", distance: 2, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
    { communityName: "Beverly Hills", distance: 5, members: 10 },
  ];
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
        data={data}
        renderItem={(itemData) => (
          <CommunityCard
            communityName={itemData.item.communityName}
            distanceFromUser={itemData.item.distance}
            numberOfMembers={itemData.item.members}
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
              title={"Join More!"}
            ></Button>
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
