import React from "react";
import { TextInput as RN_TextInput, StyleSheet } from "react-native";
import { Colors } from "../Constants";

const TextInput = (props) => {
  return (
    <RN_TextInput
      style={styles.input}
      onChange={props.onChange}
      value={props.value}
      autoCapitalize="none"
      placeholder="Enter text here..."
      placeholderTextColor={Colors.lightGreen}
    />
  );
};

const styles = StyleSheet.create({
  input: {
    height: 40,
    marginTop: 12,
    marginBottom: 12,
    borderWidth: 1,
    padding: 10,
    borderRadius: 10,
  },
});

export default TextInput;
