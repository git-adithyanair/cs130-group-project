import React from "react";
import { Dimensions } from "react-native";

export const API_URL = "http://api.good-grocer.click";
// export const API_URL = "http://localhost:8080";

export const GOOGLE_MAPS_API_KEY = "AIzaSyA1nWQO2sDr2MVzipOLMvApGd3cGWUgc1s";

export const Dim = {
  height: Dimensions.get("window").height,
  width: Dimensions.get("window").width,
};

export const Colors = {
  darkGreen: "#7B886B",
  lightGreen: "#C0C6B9",
  white: "#FFFFFF",
  cream: "#F1ECEC",
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
