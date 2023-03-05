import React from "react";
import {
  TouchableOpacity,
  StyleSheet,
  Text,
  ActivityIndicator,
} from "react-native";
import { Dim, Colors, Font } from "../Constants";

const Button = (props) => {
  const styles = StyleSheet.create({});

  return (
    <TouchableOpacity
      onPress={props.onPress}
      disabled={props.disabled}
      style={{
        elevation: 8,
        backgroundColor: Colors.darkGreen,
        borderRadius: 10,
        paddingVertical: 10,
        paddingHorizontal: 12,
        width: props.width,
        marginTop: 10,
        ...props.appButtonContainer,
      }}
    >
      {props.loading ? (
        <ActivityIndicator
          color={Colors.white}
          size={Font.s3.size}
        />
      ) : (
        <Text
          style={{
            fontSize: 18,
            color: props.textColor,
            fontWeight: "bold",
            alignSelf: "center",
            textTransform: "uppercase",
            ...props.appButtonText,
          }}
        >
          {props.title}
        </Text>
      )}
    </TouchableOpacity>
  );
};

export default Button;
