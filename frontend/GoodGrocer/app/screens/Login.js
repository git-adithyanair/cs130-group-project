import React, { useState, useEffect } from "react";
import TextInput from "../components/TextInput";
import Button from "../components/Button";
import { SafeAreaView, StyleSheet, Text, Image, View } from "react-native";
import { useIsFocused } from "@react-navigation/native";
import { API_URL, Font } from "../Constants";
import axios from "axios";
import { useDispatch } from "react-redux";
import { setToken } from "../store/actions";
import useRequest from "../hooks/useRequest";

const Login = ({ navigation }) => {
  const isFocused = useIsFocused();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const dispatch = useDispatch();

  const login = useRequest({
    url: "/user/login",
    method: "post",
    body: {
      email,
      password,
    },
    onSuccess: (data) => {
      dispatch(setToken(data.token));
    },
  });

  return (
    <SafeAreaView style={styles.container}>
      <View style={{ paddingTop: 50 }}>
        <Text style={styles.titleText}>Welcome Back</Text>
        <Text>Email</Text>
        <TextInput
          onChange={(email) => setEmail(email.nativeEvent.text)}
          placeholder="Enter your email or phone number..."
        />
        <Text>Password</Text>
        <TextInput
          onChange={(password) => setPassword(password.nativeEvent.text)}
          placeholder="Enter your password..."
        />
        <Button
          title={"Log In"}
          onPress={async () => {
            if (!email || !password) {
              Alert.alert("Oops!", "Please fill out all fields.");
            } else {
              await login.doRequest();
            }
          }}
          textColor={"white"}
          backgroundColor={"#0070CA"}
          width={300}
          appButtonContainer={{ marginTop: 20 }}
        />
      </View>
    </SafeAreaView>
  );
};

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

export default Login;
