import React from "react";
import { ActivityIndicator, SafeAreaView } from "react-native";
import { Colors } from "../Constants";

const Loading = () => {
  return (
    <SafeAreaView style={{ flex: 1, backgroundColor: Colors.white }}>
      <ActivityIndicator
        size="large"
        color={Colors.darkGreen}
        style={{ alignSelf: "center", flex: 1 }}
      />
    </SafeAreaView>
  );
};

export default Loading;
