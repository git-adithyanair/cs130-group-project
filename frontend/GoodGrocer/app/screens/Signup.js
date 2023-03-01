import React, { useState } from "react";
import TextInput from "../components/TextInput";
import Button from "../components/Button";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  Image,
  View,
  Pressable,
} from "react-native";

function Signup({ navigation }) {
  const [email, setEmail] = useState("");
  const [phoneNumber, setPhoneNumber] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");

  return (
    <SafeAreaView style={styles.container}>
      <View>
        <Image source={require("../assets/logo.png")} />
        <Text style={styles.titleText}>Signup</Text>
        <Text>Email</Text>
        <TextInput onChange={(email) => setEmail(email.nativeEvent.text)} />
        <Text>Phone Number</Text>
        <TextInput
          onChange={(phoneNumber) =>
            setPhoneNumber(phoneNumber.nativeEvent.text)
          }
        />
        <Text>Name</Text>
        <TextInput onChange={(name) => setName(name.nativeEvent.text)} />
        <Text>Password</Text>
        <TextInput
          onChange={(password) => setPassword(password.nativeEvent.text)}
        />
        <Button
          title={"Continue with Address"}
          onPress={() =>
            navigation.navigate("AddressSignup", {
              email,
              phoneNumber,
              name,
              password,
            })
          }
          textColor={"white"}
          backgroundColor={"#0070CA"}
          width={300}
        />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    justifyContent: "center",
    alignItems: "center",
  },
  titleText: {
    fontSize: 25,
  },
});

export default Signup;
