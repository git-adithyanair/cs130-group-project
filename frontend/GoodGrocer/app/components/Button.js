import React from "react";
import {
  TouchableOpacity,
  StyleSheet,
  Text
} from "react-native";


const Button = (props) => {
    const styles = StyleSheet.create({
        appButtonContainer: {
          elevation: 8,
          backgroundColor: props.backgroundColor,
          borderRadius: 10,
          paddingVertical: 10,
          paddingHorizontal: 12,
          width: props.width,
          marginTop: 10
        },
        appButtonText: {
          fontSize: 18,
          color: props.textColor,
          fontWeight: "bold",
          alignSelf: "center",
          textTransform: "uppercase"
        }
      });

  return <TouchableOpacity onPress={props.onPress} style={styles.appButtonContainer}>
<Text style={styles.appButtonText}>{props.title}</Text>
</TouchableOpacity>
};




export default Button;