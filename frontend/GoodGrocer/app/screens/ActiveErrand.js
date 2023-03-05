import React, { useEffect, useState } from "react";
import Button from "../components/Button";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  Image,
  View,
  FlatList,
} from "react-native";
import axios from "axios";
import { Dim, Colors, Font, API_URL } from "../Constants";
import ErrandRequestCard from "../components/ErrandRequestCard";
import { useSelector } from "react-redux";

const ActiveErrand = ({ navigation }) => {
  const [data, setData] = useState({});
  const [completeErrandEnabled, setCompleteErrandEnabled] = useState(true);
  const token = useSelector((state) => state.user.token);

  const getData = async () => {
    axios
      .get(`${API_URL}/errand/active`, {
        headers: { Authorization: `Bearer ${token}` },
      })
      .then(({ data }) => {
        setData(data);
        setCompleteErrandEnabled(checkRequestCompletion(data));
      })
      .catch((error) => {
        console.error(error);
      });
  };

  const completeErrand = async () => {
    axios
      .post(
        `${API_URL}/errand/update-status`,
        {
          id: data.errand.id,
          is_complete: true,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      )
      .then(({ data }) => {
        setData({});
      })
      .catch((error) => {
        console.error(error);
      });
  };

  const checkRequestCompletion = (data) => {
    if (JSON.stringify(data) === "{}") {
      return false;
    }
    for (const request of data.requests) {
      for (const item of request.items) {
        if (!item.found.Valid) {
          return false;
        }
      }
    }
    return true;
  };

  const requestComplete = (items) => {
    for (const item of items) {
      if (!item.found.Valid) {
        return false;
      }
    }
    return true;
  };

  useEffect(() => {
    const unsubscribe = navigation.addListener("focus", () => {
      getData();
    });
    return unsubscribe;
  }, [navigation]);

  return (
    <SafeAreaView style={styles.wrapper}>
      <FlatList
        contentContainerStyle={styles.container}
        style={styles.list}
        data={data.requests}
        renderItem={(itemData) => (
          <ErrandRequestCard
            imageUri="https://i.pinimg.com/236x/10/f4/a9/10f4a952ddf8e6828ae6833b3088dfa0.jpg"
            name={itemData.item.user.full_name}
            storeName={itemData.item.store.name}
            storeAddress={itemData.item.store.address}
            numItems={itemData.item.items.length}
            requestComplete={requestComplete(itemData.item.items)}
            onPress={() =>
              navigation.navigate("ActiveRequest", {
                user: itemData.item.user,
                profileImage:
                  "https://i.pinimg.com/236x/10/f4/a9/10f4a952ddf8e6828ae6833b3088dfa0.jpg",
                items: itemData.item.items,
                store: itemData.item.store,
              })
            }
            key={itemData.item.id}
          />
        )}
        keyExtractor={(item) => item.id}
        ListHeaderComponent={() => (
          <View style={{ alignItems: "center" }}>
            <Text style={styles.title}>Errand for {data.community_name}</Text>
          </View>
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
          <View style={{ alignItems: "center", paddingTop: 20 }}>
            {JSON.stringify(data) === "{}" ? (
              <Text style={styles.noErrandText}>
                You currently do not have an active errand. To create an errand,
                go to the home tab and select some requests!
              </Text>
            ) : (
              <Button
                width={Dim.width * 0.9}
                appButtonContainer={{
                  backgroundColor: Colors.lightGreen,
                  opacity: completeErrandEnabled ? 100 : 0,
                }}
                appButtonText={{ textTransform: "none" }}
                title={"Complete Errand"}
                onPress={completeErrand}
                isDisabled={!completeErrandEnabled}
              />
            )}
          </View>
        )}
      ></FlatList>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    flex: 1,
    backgroundColor: "#fff",
  },
  container: {
    width: Dim.width * 0.9,
    alignSelf: "center",
    paddingTop: 10,
    paddingBottom: 80,
  },
  title: {
    marginTop: 10,
    marginBottom: 20,
    fontSize: Font.s2.size,
    fontFamily: Font.s2.family,
    fontWeight: Font.s2.weight,
  },
  content: {
    alignItems: "center",
  },
  list: {
    flex: 1,
  },
  noErrandText: {
    fontSize: Font.s2.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s3.weight,
    color: Colors.darkGreen,
    paddingHorizontal: 10,
  },
});

export default ActiveErrand;
