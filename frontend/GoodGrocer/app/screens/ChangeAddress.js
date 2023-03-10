import React, { useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Pressable,
  Image,
  Text,
  Title,
  View,
  ScrollView,
} from "react-native";
import TextInput from "../components/TextInput";
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view";
import Button from "../components/Button";
import { Colors, Font, Dim } from "../Constants";
import LocationFinderCard from "../components/LocationFinderCard";
import useRequest from "../hooks/useRequest";

function ChangeAddress({ navigation }) {
  const [locationData, setLocationData] = useState({});
  const [loading, setLoading] = useState(false);

  const handleLocation = (e) => {
    setLocationData(e);
  };

  const updateAddress = useRequest({
    url: "/user/update-location",
    method: "post",
    body: {
      address: locationData.address,
      place_id: locationData.place_id,
      x_coord: locationData.x_coord,
      y_coord: locationData.y_coord,
    },
    onSuccess: (data) => {
      setLoading(false);
      navigation.goBack();
      // props.navigation.navigate("YourCommunities");
    },
    onFail: () => setLoading(false),
  });

  return (
    <SafeAreaView style={styles.container}>
      <KeyboardAwareScrollView
        showsVerticalScrollIndicator={false}
        extraScrollHeight={30}
        keyboardShouldPersistTaps="handled"
      >
        <View style={{ marginTop: 20, marginLeft: 20 }}>
          <Text style={styles.title}>Change your Address</Text>
        </View>
        <View style={{ marginTop: 30, ...styles.minWrapper }}>
          <LocationFinderCard
            searchLabel="Your Address"
            placeholder={"Enter your new address"}
            width={Dim.width * 0.9}
            onSelectLocation={(e) => handleLocation(e)}
          />
        </View>
        <Button
          title={"Submit"}
          appButtonContainer={styles.button}
          width={Dim.width * 0.5}
          onPress={async () => {
            setLoading(true);
            await updateAddress.doRequest();
          }}
        />
      </KeyboardAwareScrollView>
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

export default ChangeAddress;
