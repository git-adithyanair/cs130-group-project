import React, { useState } from "react";
import { SafeAreaView, StyleSheet, Text, View } from "react-native";
import Button from "../components/Button";
import { Colors, Font, Dim } from "../Constants";
import StoresCard from "../components/StoresCard";

function PickStore({ navigation, route }) {
  const [storeName, setStoreName] = useState("");
  const [storeId, setStoreId] = useState("");

  return (
    <SafeAreaView style={styles.container}>
      <View style={{ marginTop: 20, width: "90%", alignSelf: "center" }}>
        <Text style={styles.title}>
          Pick your Store in {route.params.communityName}
        </Text>
        <Text style={{ marginTop: 10 }}>
          If you don't choose a store, the shopper will choose one for you!
        </Text>
      </View>
      <View style={{ marginTop: 30, ...styles.minWrapper }}>
        <StoresCard
          communityId={route.params.communityId}
          onSelectStore={(data) => {
            setStoreName(data.name);
            setStoreId(data.store_id);
          }}
          width={Dim.width * 0.9}
        />
      </View>
      <Button
        title={"Pick your Items"}
        appButtonContainer={styles.button}
        width={Dim.width * 0.5}
        onPress={() => {
          navigation.navigate("CreateRequest", {
            communityName: route.params.communityName,
            communityId: route.params.communityId,
            storeName: storeName,
            storeId: storeId,
          });
        }}
      />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
  },
  title: {
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
  },
  minWrapper: {
    width: Dim.width * 0.9,
    alignSelf: "center",
  },
  button: {
    alignSelf: "center",
    backgroundColor: Colors.lightGreen,
    marginTop: 30,
  },
});

export default PickStore;
