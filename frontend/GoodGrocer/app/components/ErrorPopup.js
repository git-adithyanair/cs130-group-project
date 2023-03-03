import React from "react";
import {
  View,
  Animated,
  TouchableWithoutFeedback,
  Modal,
  Text,
} from "react-native";
import { useSelector } from "react-redux";

import Button from "./Button";
import { Dim, Colors, Font } from "../Constants";

const ErrorPopup = (props) => {
  const panY = new Animated.Value(-10);
  const top = panY.interpolate({
    inputRange: [-1, 0, 1],
    outputRange: [-2, 0, 0],
  });

  const resetPositionAnim = Animated.timing(panY, {
    toValue: 0,
    duration: 300,
    useNativeDriver: false,
  });
  const closeAnim = Animated.timing(panY, {
    toValue: -10,
    duration: 300,
    useNativeDriver: false,
  });

  const visible = useSelector((state) => state.user.errorPopupVisible) || false;
  const message = useSelector((state) => state.user.errorMessageText) || "";

  if (visible) {
    resetPositionAnim.start();
  }

  return (
    <Modal animated animationType="fade" visible={visible} transparent>
      <TouchableWithoutFeedback
        onPress={() => {
          closeAnim.start(props.onPress);
        }}
      >
        <View
          style={{
            backgroundColor: "rgba(0,0,0,0.4)",
            flex: 1,
            justifyContent: "center",
          }}
        >
          <Animated.View style={[{ top }]}>
            <View
              style={{
                justifyContent: "center",
                alignItems: "center",
              }}
            >
              <View
                style={{
                  width: Dim.width * 0.75,
                  borderRadius: 20,
                  backgroundColor: Colors.lightGreen,
                  paddingTop: 30,
                  paddingBottom: 25,
                  justifyContent: "space-around",
                  alignItems: "center",
                  ...props.style,
                }}
              >
                <Text
                  style={{
                    marginBottom: 15,
                    textAlign: "center",
                    fontSize: Font.s2.size,
                    fontFamily: Font.s2.family,
                    fontWeight: Font.s2.weight,
                    ...props.textStyle,
                  }}
                >
                  Something went wrong!
                </Text>
                <Text
                  style={{
                    textAlign: "center",
                    marginBottom: 20,
                    width: "85%",
                    ...props.messageStyle,
                  }}
                >
                  {message}
                </Text>
                <View
                  style={{
                    flexDirection: "row",
                    justifyContent: "space-around",
                  }}
                >
                  <Button
                    onPress={() => {
                      closeAnim.start(props.onPress);
                    }}
                    title="OK"
                    appButtonText={{
                      fontSize: undefined,
                      fontFamily: undefined,
                      fontWeight: undefined,
                    }}
                    appButtonContainer={{
                      width: Dim.width * 0.3,
                    }}
                  />
                </View>
              </View>
            </View>
          </Animated.View>
        </View>
      </TouchableWithoutFeedback>
    </Modal>
  );
};

export default ErrorPopup;
