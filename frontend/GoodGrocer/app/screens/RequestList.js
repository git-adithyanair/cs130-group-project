import React, { useEffect, useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  FlatList,
  View,
  TouchableOpacity,
} from "react-native";
import RequestCard from "../components/RequestCard";
import { Colors, Dim, Font } from "../Constants";
import useRequest from "../hooks/useRequest";
import Button from "../components/Button";
import Ionicons from "react-native-vector-icons/Ionicons";
import Loading from "./Loading";

function RequestList(props) {
  const [communityRequestData, setCommunityRequestData] = useState([]);
  const [selectedRequests, setSelectedRequests] = useState([]);
  const [loading, setLoading] = useState(true);
  const [creatingErrand, setCreatingErrand] = useState(false);

  useEffect(() => {
    props.navigation.setOptions({
      headerRight: () => (
        <TouchableOpacity
          onPress={() =>
            props.navigation.navigate("PickStore", {
              communityId: props.route.params.communityId,
              communityName: props.route.params.communityName,
            })
          }
          style={{ marginRight: 20 }}
        >
          <Ionicons name={"add"} size={30} color={Colors.darkGreen} />
        </TouchableOpacity>
      ),
    });
  }, [props.navigation]);

  const getCommunityRequests = useRequest({
    url: `/community/requests?id=${props.route.params.communityId}`,
    method: "get",
    onSuccess: (data) => {
      const requests = [];
      data.forEach((requestData) => {
        if (requestData.request.status === "completed") {
          return;
        }
        requests.push({
          name: requestData.user.full_name,
          storeName: requestData.store ? requestData.store.name : "Any Store",
          storeAddress: requestData.store ? requestData.store.address : "N/A",
          id: requestData.request.id,
          numItems: requestData.items.length,
          imageUri: requestData.user.profile_picture,
          available: requestData.request.status === "pending",
          items: requestData.items,
        });
      });
      setCommunityRequestData(requests);
      setLoading(false);
    },
  });

  const createErrand = useRequest({
    url: "/errand",
    method: "post",
    body: {
      community_id: props.route.params.communityId,
      request_ids: selectedRequests,
    },
    onSuccess: () => {
      setSelectedRequests([]);
      setCreatingErrand(false);
      props.navigation.navigate("Errand");
    },
    onFail: () => {
      setCreatingErrand(false);
    },
  });

  useEffect(() => {
    if (loading) {
      const getRequests = async () => getCommunityRequests.doRequest();
      getRequests();
    }
  }, [loading]);

  if (loading) {
    return <Loading />;
  }

  return (
    <SafeAreaView style={styles.container}>
      <FlatList
        data={communityRequestData}
        contentContainerStyle={{ paddingBottom: 20 }}
        renderItem={({ item }) =>
          item.available ? (
            <RequestCard
              key={item.id}
              name={item.name}
              storeName={item.storeName}
              storeAddress={item.storeAddress}
              numItems={item.numItems}
              imageUri={item.imageUri}
              onPress={() =>
                props.navigation.navigate("RequestDetail", {
                  requestId: item.id,
                  storeName: item.storeName,
                  items: item.items,
                  user: {
                    name: item.name,
                    profileImage: item.imageUri,
                  },
                })
              }
              onPressSelect={() => {
                if (selectedRequests.includes(item.id)) {
                  setSelectedRequests(
                    selectedRequests.filter((id) => id !== item.id)
                  );
                } else {
                  setSelectedRequests([...selectedRequests, item.id]);
                }
              }}
              selected={selectedRequests.includes(item.id)}
            />
          ) : null
        }
        keyExtractor={(item) => item.id}
        ListHeaderComponent={
          <Text style={styles.titleText}>
            Requests in {props.route.params.communityName}
          </Text>
        }
        ListEmptyComponent={() => (
          <View
            style={{ alignItems: "center", height: "100%", paddingTop: "10%" }}
          >
            <Text
              style={{
                fontSize: Font.s2.size,
                alignSelf: "center",
              }}
            >
              Create a request to get started! 🧸
            </Text>
          </View>
        )}
      />
      {selectedRequests.length > 0 ? (
        <View
          style={{
            padding: 10,
          }}
        >
          <Button
            width={200}
            appButtonContainer={{
              backgroundColor: Colors.lightGreen,
              alignSelf: "center",
              marginVertical: 10,
            }}
            title="Create Errand"
            loading={creatingErrand}
            onPress={async () => {
              setCreatingErrand(true);
              await createErrand.doRequest();
            }}
          />
        </View>
      ) : null}
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: Colors.white,
  },
  titleText: {
    fontSize: Font.s1.size,
    fontWeight: Font.s1.weight,
    width: Dim.width * 0.85,
    textAlign: "center",
    alignSelf: "center",
    marginTop: 20,
  },
});

export default RequestList;
