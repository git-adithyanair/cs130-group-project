import React from "react";
import Button from "../components/Button";
import { SafeAreaView, StyleSheet, Pressable, Image } from "react-native";

function Landing({ navigation }) {
  return (
    <SafeAreaView style={styles.container}>
      <Image source={require("../assets/logo.png")} />
      <Image
        source={require("../assets/slogan.png")}
        style={{
          marginVertical: 20,
        }}
      />
      <Button
        title={"Login"}
        onPress={() => navigation.navigate("Login")}
        textColor={"white"}
        backgroundColor={"#7B886B"}
        width={200}
      />
      <Button
        title={"Sign Up"}
        onPress={() => navigation.navigate("Signup")}
        textColor={"white"}
        backgroundColor={"#7B886B"}
        width={200}
      />
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
});

export default Landing;
