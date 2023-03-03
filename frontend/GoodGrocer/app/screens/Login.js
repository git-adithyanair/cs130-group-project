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
  const [failedLogIn, setFailedLogIn] = useState(false);

  const dispatch = useDispatch();

  useEffect(() => {
    setFailedLogIn(false);
  }, [isFocused]);

  const login = useRequest({
    url: "/user/login",
    method: "post",
    body: {
      email,
      password,
    },
    onSuccess: (data) => {
      if (data.token) {
        dispatch(setToken(data.token));
      } else {
        setFailedLogIn(true);
      }
      console.log(data);
    },
    onFail: () => {
      setFailedLogIn(true);
    },
  });

  return (
    <SafeAreaView style={styles.container}>
      <View style={{ paddingTop: 50 }}>
        {/* <Image source={require("../assets/logo.png")} /> */}
        <Text style={styles.titleText}>Welcome Back</Text>
        <Text>Email or Phone Number</Text>
        <TextInput onChange={(email) => setEmail(email.nativeEvent.text)} />
        <Text>Password</Text>
        <TextInput
          onChange={(password) => setPassword(password.nativeEvent.text)}
        />
        <Button
          title={"Sign In"}
          onPress={async () => await login.doRequest()}
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
