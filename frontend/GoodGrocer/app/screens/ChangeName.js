import React, { useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Pressable,
  Image,
  Text,
  Title,
  View,
  ScrollView,
} from "react-native";
import TextInput from "../components/TextInput";
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view";

import Button from "../components/Button";
import { Colors, Font } from "../Constants";
import useRequest from "../hooks/useRequest";

function ChangeName({ navigation }) {
  const [name, setName] = useState("");
  const [loading, setLoading] = useState(false);

  const changeUserName = useRequest({
    url: "/user/update-name",
    method: "post",
    body: {
      name: name,
    },
    onSuccess: (data) => {
      setLoading(false);
      navigation.goBack();
    },
    onFail: () => setLoading(false),
  });

  return (
    <SafeAreaView style={styles.container}>
      <KeyboardAwareScrollView
        showsVerticalScrollIndicator={false}
        extraScrollHeight={30}
        keyboardShouldPersistTaps="handled"
      >
        <View style={{ marginTop: 20, marginLeft: 20 }}>
          <Text style={styles.title}>Change your Name</Text>
        </View>
        <View style={{ marginTop: 10, marginLeft: 30, marginRight: 30 }}>
          <TextInput onChange={(name) => setName(name)}></TextInput>
        </View>
        <View style={{ alignItems: "center" }}>
          <Button
            title={"Submit"}
            onPress={async () => {
              setLoading(true);
              await changeUserName.doRequest();
            }}
            textColor={"white"}
            backgroundColor={Colors.lightGreen}
            width={250}
          ></Button>
        </View>
      </KeyboardAwareScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
  },
  title: {
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
  },
});

export default ChangeName;
