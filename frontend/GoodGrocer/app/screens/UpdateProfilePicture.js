import React, { useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  View,
  TouchableOpacity,
  ImageBackground,
} from "react-native";
import * as ImagePicker from "expo-image-picker";

import Button from "../components/Button";
import { Colors, Dim } from "../Constants";
import useRequest from "../hooks/useRequest";

function UpdateProfilePicture({ route, navigation }) {
  const [pictureUri, setPictureUri] = useState("");
  const [loading, setLoading] = useState(false);

  const updateProfilePicture = useRequest({
    url: "/user/update-profile-pic",
    method: "post",
    body: {
      image: "data:image/png;base64," + pictureUri,
    },
    onSuccess: (data) => {
      setLoading(false);
      navigation.goBack();
    },
    onFail: () => setLoading(false),
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
      <View style={{ width: 300, justifyContent: "center" }}>
        <View style={{ marginTop: 30 }}>
          <ImageBackground
            style={{
              width: Dim.width * 0.5,
              height: Dim.width * 0.5,
              alignSelf: "center",
            }}
            source={{
              uri: pictureUri
                ? "data:image/png;base64," + pictureUri
                : route.params.user.profile_picture,
            }}
          />
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
            Click here to update your profile picture...
          </Text>
        </TouchableOpacity>
      </View>
      <Button
        title={"Submit"}
        onPress={async () => {
          setLoading(true);
          await updateProfilePicture.doRequest();
        }}
        textColor={"white"}
        backgroundColor={"#0070CA"}
        width={Dim.width * 0.7}
        appButtonContainer={{
          backgroundColor: pictureUri === "" ? "#d3d3d3" : Colors.darkGreen,
          alignSelf: "center",
          marginBottom: 40,
        }}
        disabled={pictureUri === ""}
        loading={loading}
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
  defaultPic: {
    alignItems: "center",
  },
});

export default UpdateProfilePicture;
