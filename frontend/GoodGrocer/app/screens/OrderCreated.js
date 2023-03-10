import React, { useEffect } from "react";
import { SafeAreaView, StyleSheet, Text } from "react-native";
import Button from "../components/Button";
import { Colors, Font } from "../Constants";

const OrderCreated = ({ navigation }) => {
  useEffect(() => {
    navigation.setOptions({
      headerLeft: () => null,
    });
  }, [navigation]);

  return (
    <SafeAreaView style={styles.container}>
      <Text style={styles.title}>Your request has been created!</Text>
      <Text style={{ textAlign: "center", marginTop: 15 }}>
        We'll let you know once it's in progress...
      </Text>
      <Button
        title={"Back to Home"}
        onPress={() => navigation.popToTop()}
        appButtonContainer={{
          marginTop: 30,
          width: "80%",
          backgroundColor: Colors.lightGreen,
          alignSelf: "center",
        }}
      />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    justifyContent: "center",
  },
  title: {
    fontSize: Font.s2.size,
    fontFamily: Font.s2.family,
    fontWeight: Font.s2.weight,
    textAlign: "center",
  },
});

export default OrderCreated;
