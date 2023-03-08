import React from "react";
import { TextInput as RN_TextInput, StyleSheet } from "react-native";
import { Colors } from "../Constants";

const TextInput = (props) => {
  return (
    <RN_TextInput
      style={{ ...styles.input, ...props.style }}
      onChangeText={props.onChange}
      value={props.value}
      autoCapitalize="none"
      placeholder={props.placeholder || "Enter text here..."}
      placeholderTextColor={Colors.darkGreen}
      secureTextEntry={props.secureTextEntry}
    />
  );
};

const styles = StyleSheet.create({
  input: {
    height: 40,
    marginTop: 5,
    marginBottom: 20,
    padding: 10,
    borderRadius: 10,
    backgroundColor: Colors.lightGreen,
  },
});

export default TextInput;
