import React from "react";
import TabBar from "../components/TabBar";
import defaultProfilePic from "../assets/default-profile-pic.png";
import { Image } from "react-native";

function Tabs({ navigation }) {
  const defaultImageUri = Image.resolveAssetSource(defaultProfilePic).uri;
  return <TabBar imageUri={defaultImageUri} />;
}

export default Tabs;
