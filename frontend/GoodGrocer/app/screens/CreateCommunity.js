import { React, useState } from "react";
import { SafeAreaView, StyleSheet, View, Text, Alert } from "react-native";
import { Dim, Colors, Font } from "../Constants";
import TextInput from "../components/TextInput";
import LocationFinderCard from "../components/LocationFinderCard";
import Button from "../components/Button";

const CreateCommunity = (props) => {
  const [name, setName] = useState("");
  const [range, setRange] = useState("");
  const [locationData, setLocationData] = useState({});

  const handleName = (text) => {
    setName(text);
  };

  const handleRange = (text) => {
    setRange(text);
  };

  const handleLocation = (text) => {
    setLocationData(text);
  };

  return (
    <SafeAreaView style={styles.wrapper}>
      <View style={styles.minWrapper}>
        <Text style={styles.title}>Community Name</Text>
        <TextInput
          placeholder="Enter Community Name"
          onChange={(text) => handleName(text)}
        />
      </View>
      <View style={styles.minWrapper}>
        <Text style={styles.title}>Community Range</Text>
        <TextInput
          placeholder="Enter Community Range (m)"
          onChange={(text) => handleRange(text)}
        />
      </View>
      <View style={{ marginTop: 30, ...styles.minWrapper }}>
        <LocationFinderCard
          searchLabel="Community address"
          placeholder={"Enter Community Address"}
          width={Dim.width * 0.9}
          onSelectLocation={(text) => handleLocation(text)}
        />
      </View>
      <Button
        title="Add Stores"
        appButtonContainer={styles.addStoreButton}
        width={Dim.width * 0.5}
        onPress={() => {
          if (!name || !range || !locationData) {
            Alert.alert("Oops!", "Please fill out all fields.");
          } else {
            props.navigation.navigate("AddStores", {
              communityName: name,
              communityRange: range,
              communityAddr: locationData,
            });
          }
        }}
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
  addStoreButton: {
    alignSelf: "center",
    backgroundColor: Colors.lightGreen,
    marginTop: 70,
  },
});

export default CreateCommunity;
