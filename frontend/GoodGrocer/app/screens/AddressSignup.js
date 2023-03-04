import React, { useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  Image,
  View,
  TouchableOpacity,
  ImageBackground,
} from "react-native";
import * as ImagePicker from "expo-image-picker";
import { useDispatch } from "react-redux";

import Button from "../components/Button";
import { Colors, Dim, Font } from "../Constants";
import useRequest from "../hooks/useRequest";
import { setToken } from "../store/actions";
import LocationFinderCard from "../components/LocationFinderCard";

function AddressSignup({ route, navigation }) {
  const { email, name, phoneNumber, password } = route.params;
  const [pictureUri, setPictureUri] = useState("");
  const [locationData, setLocationData] = useState({});

  const dispatch = useDispatch();

  const signup = useRequest({
    url: "/user",
    method: "post",
    body: {
      email,
      password,
      full_name: name,
      phone_number: phoneNumber,
      address: locationData.address,
      place_id: locationData.place_id,
      x_coord: locationData.x_coord,
      y_coord: locationData.y_coord,
      profile_picture: pictureUri || "DEFAULT",
    },
    onSuccess: (data) => {
      dispatch(setToken(data.token));
    },
  });

  const pickImage = async () => {
    try {
      const result = await ImagePicker.launchImageLibraryAsync({
        mediaTypes: ImagePicker.MediaTypeOptions.Images,
        allowsEditing: true,
        aspect: [1, 1],
        quality: 0.8,
        base64: true,
      });
      if (!result.canceled) {
        setPictureUri(result.assets[0].base64);
      }
    } catch (err) {
      console.log(err);
    }
  };

  const getImagePickerPermissionAsync = async () => {
    const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
    if (status !== "granted") {
      Alert.alert(
        "Oops!",
        "We need access to your photo library to assign a profile picture for you!"
      );
    } else {
      pickImage();
    }
  };

  return (
    <SafeAreaView style={styles.container}>
      <View style={{ width: 300 }}>
        <Text style={styles.titleText}>Final steps!</Text>
        <LocationFinderCard
          searchLabel="Find your address"
          width={Dim.width * 0.7}
          onSelectLocation={(data) => {
            setLocationData(data);
          }}
        />
        <View style={{ marginTop: 30 }}>
          {pictureUri ? (
            <ImageBackground
              style={{
                width: Dim.width * 0.5,
                height: Dim.width * 0.5,
                alignSelf: "center",
              }}
              source={{ uri: "data:image/png;base64," + pictureUri }}
            />
          ) : (
            <View style={styles.defaultPic}>
              <Image
                source={require("../assets/default-profile-pic.png")}
                style={{
                  width: Dim.width * 0.5,
                  height: Dim.width * 0.5,
                }}
              />
            </View>
          )}
        </View>

        <TouchableOpacity
          onPress={async () => await getImagePickerPermissionAsync()}
        >
          <Text
            style={{
              color: Colors.darkGreen,
              textAlign: "center",
              marginTop: 10,
              fontWeight: "bold",
            }}
          >
            Click here to add a profile picture...
          </Text>
        </TouchableOpacity>
      </View>
      <Button
        title={"Sign Up"}
        onPress={async () => await signup.doRequest()}
        textColor={"white"}
        backgroundColor={"#0070CA"}
        width={Dim.width * 0.7}
        appButtonContainer={{
          backgroundColor:
            JSON.stringify(locationData) === "{}" ? "#d3d3d3" : "#0070CA",
          alignSelf: "center",
          marginBottom: 40,
        }}
        disabled={JSON.stringify(locationData) === "{}"}
      />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "space-between",
  },
  titleText: {
    marginVertical: 20,
    fontFamily: Font.s1.family,
    fontSize: 30,
    fontWeight: Font.s1.weight,
  },
  defaultPic: {
    alignItems: "center",
  },
});

export default AddressSignup;
