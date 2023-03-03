import React, { useEffect, useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  Image,
  ScrollView,
  View,
  TouchableOpacity,
} from "react-native";
import RequestCard from "../components/RequestCard";
import useRequest from "../hooks/useRequest";

function RequestList(props) {
  const [communityRequestData, setCommunityRequestData] = useState([]);

  const getCommunityRequests = useRequest({
    url: `/community/requests?id=${props.route.params.communityId}`,
    method: "get",
    onSuccess: (data) => {
      data.forEach((request) => {
        setCommunityRequestData((oldArray) => [
          ...oldArray,
          {
            name: request.user.full_name,
            storeName: request.store ? request.store.name : "Any Store",
            id: request.request.id,
            numItems: -1,
            imageUri:
              "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png",
          },
        ]);
      });
    },
  });
  const func = async () => getCommunityRequests.doRequest();
  useEffect(() => {
    func();
  }, []);

  const cards = communityRequestData.map((request) => (
    <TouchableOpacity style={styles.requestCard} key={request.id}>
      <RequestCard
        name={request.name}
        storeName={request.storeName}
        numItems={request.numItems}
        imageUri={request.imageUri}
        onPress={() =>
          props.navigation.navigate("RequestDetail", {
            requestId: request.id,
            storeName: request.storeName,
          })
        }
      />
    </TouchableOpacity>
  ));
  return (
    <SafeAreaView style={styles.container}>
      <View style={styles.content}>
        <Image source={require("../assets/logo.png")} />
        <Text style={styles.titleText}>
          Requests in {props.route.params.communityName}
        </Text>
        <ScrollView style={styles.listOfRequests}>{cards}</ScrollView>
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
  },
  content: {
    alignItems: "center",
    marginTop: 40,
  },
  listOfRequests: {
    marginBottom: 150,
    width: "85%%",
  },
  titleText: {
    fontSize: 25,
  },
  requestCard: {
    paddingTop: 20,
  },
});

const cards = testData.map((request) => (
  <TouchableOpacity
    style={styles.requestCard}
    key={request.id}
    onPress={() => setPage(1)}
  >
    <RequestCard
      name={request.name}
      storeName={request.storeName}
      numItems={request.numItems}
      imageUri={request.imageUri}
    />
  </TouchableOpacity>
));
return (
  <SafeAreaView style={styles.container}>
    <View style={styles.content}>
      <Image source={require("../assets/logo.png")} />
      <Text style={styles.titleText}>Requests in Westwood</Text>
      <ScrollView style={styles.listOfRequests}>{cards}</ScrollView>
    </View>
  </SafeAreaView>
);

export default RequestList;
