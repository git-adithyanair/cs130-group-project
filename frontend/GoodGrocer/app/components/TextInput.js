import React from "react";
import {
  TextInput as RN_TextInput,
  StyleSheet
} from "react-native";


const TextInput = (props) => {
  return <RN_TextInput style={styles.input} onChange={props.onChange} value={props.value} autoCapitalize='none'/>; 
};


const styles = StyleSheet.create({
    input: {
      height: 40,
      marginTop: 12,
      marginBottom: 12,
      borderWidth: 1,
      padding: 10,
    }
});

export default TextInput;