import React from "react";
import { Dimensions } from "react-native";

export const Dim = {
  height: Dimensions.get("window").height,
  width: Dimensions.get("window").width,
};

export const Colors = {
  darkGreen: "#7B886B",
  lightGreen: "#C0C6B9",
  white: "#FFFFFF",
  cream:"#F1ECEC",
  blue: "#0070CA",
};

export const Font = {
  s1: {
    size: 20,
    family: "Inter_600SemiBold",
    weight: "600",
  },
  s2: {
    size: 18,
    family: "Inter_600SemiBold",
    weight: "600",
  },
  s3: {
    size: 13,
    family: "Inter_600SemiBold",
    weight: "600",
  },
};

export const BorderRadius = 10;
