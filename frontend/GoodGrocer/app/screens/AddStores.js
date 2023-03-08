import { React, useState } from "react";
import { SafeAreaView, StyleSheet, View, Text, Alert } from "react-native";
import { Dim, Colors, Font } from "../Constants";
import TextInput from "../components/TextInput";
import LocationFinderCard from "../components/LocationFinderCard";
import Button from "../components/Button";
import useRequest from "../hooks/useRequest";

const AddStores = (props) => {
  const { communityName, communityRange, communityAddr } = props.route.params;
  const [stores, setStores] = useState([]);
  const [storeName, setStoreName] = useState("");
  const [storeLocationData, setStoreLocationData] = useState({});
  const [numStores, setNumStores] = useState(0);

  const onPressAddStore = () => {
    if (!storeName || !storeLocationData) {
      Alert.alert("Oops!", "Please fill out all fields.");
    } else {
      setStores([
        ...stores,
        {
          name: storeName,
          place_id: storeLocationData.place_id,
          x_coord: storeLocationData.x_coord,
          y_coord: storeLocationData.y_coord,
          address: storeLocationData.address,
        },
      ]);
      let curr = numStores;
      setNumStores(curr + 1);
    }
  };

  const createCommunity = useRequest({
    url: "/community",
    method: "post",
    body: {
      name: communityName,
      place_id: communityAddr.place_id,
      center_x_coord: communityAddr.x_coord,
      center_y_coord: communityAddr.y_coord,
      range: parseInt(communityRange),
      address: communityAddr.address,
      stores: stores,
    },
    onSuccess: (data) => {
      props.navigation.navigate("YourCommunities");
    },
  });

  return (
    <SafeAreaView style={styles.wrapper}>
      <View style={styles.minWrapper}>
        <Text style={styles.title}>Store Name</Text>
        <TextInput
          placeholder="Enter Store Name"
          onChange={(n) => setStoreName(n)}
        />
      </View>
      <View style={{ marginTop: 30, ...styles.minWrapper }}>
        <LocationFinderCard
          searchLabel="Store address"
          placeholder={"Enter Store Address"}
          width={Dim.width * 0.9}
          onSelectLocation={(data) => {
            setStoreLocationData(data);
          }}
        />
      </View>
      <Button
        title="Add Store"
        appButtonContainer={{
          alignSelf: "center",
          backgroundColor: Colors.lightGreen,
          marginTop: 10,
        }}
        width={Dim.width * 0.5}
        onPress={() => onPressAddStore()}
      />
      <Text style={{ ...styles.title, alignSelf: "center" }}>
        Number of Stores Added: {numStores}
      </Text>
      <Button
        title="Create Community"
        appButtonContainer={{
          alignSelf: "center",
          marginTop: 50,
        }}
        width={Dim.width * 0.7}
        onPress={async () => await createCommunity.doRequest()}
      />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  wrapper: { flex: 1, backgroundColor: Colors.white },
  minWrapper: {
    width: Dim.width * 0.9,
    alignSelf: "center",
  },
  title: {
    marginTop: 20,
    marginBottom: 5,
    fontWeight: "bold",
  },
  button: {},
});

export default AddStores;
