import React, { useState } from "react";
import TextInput from "../components/TextInput";
import Button from "../components/Button";
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  Image,
  View,
  Alert,
} from "react-native";
import { Font } from "../Constants";

function Signup({ navigation }) {
  const [email, setEmail] = useState("");
  const [phoneNumber, setPhoneNumber] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");

  return (
    <SafeAreaView style={styles.container}>
      <KeyboardAwareScrollView
        showsVerticalScrollIndicator={false}
        extraScrollHeight={30}
        keyboardShouldPersistTaps="handled"
      >
        <View>
          <Text style={styles.titleText}>Welcome!</Text>
          <Text>Email</Text>
          <TextInput
            onChange={(email) => setEmail(email)}
            placeholder="Enter your email..."
          />
          <Text>Password</Text>
          <TextInput
            onChange={(password) => setPassword(password)}
            placeholder="Enter a password..."
            secureTextEntry={true}
          />
          <Text>Name</Text>
          <TextInput
            onChange={(name) => setName(name)}
            placeholder="Enter your name..."
          />
          <Text>Phone Number</Text>
          <TextInput
            onChange={(phoneNumber) => setPhoneNumber(phoneNumber)}
            placeholder="Enter your phone number..."
          />
          <Button
            title={"Continue"}
            onPress={() => {
              if (!email || !phoneNumber || !name || !password) {
                Alert.alert("Oops!", "Please fill out all fields.");
              } else {
                navigation.navigate("AddressSignup", {
                  email,
                  phoneNumber,
                  name,
                  password,
                });
              }
            }}
            textColor={"white"}
            backgroundColor={"#0070CA"}
            width={300}
            appButtonContainer={{ marginTop: 20 }}
          />
        </View>
      </KeyboardAwareScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
  },
  titleText: {
    marginVertical: 20,
    fontFamily: Font.s1.family,
    fontSize: 30,
    fontWeight: Font.s1.weight,
  },
});

export default Signup;
