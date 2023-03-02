import React from "react";
import { View, StyleSheet } from "react-native";
import { Entypo } from "@expo/vector-icons";
import { SearchBar as ElementsSearch } from "react-native-elements";
import { Dim, Colors, Font, BorderRadius } from "../Constants";

const SearchBar = (props) => {
  return (
    <View
      style={{
        alignSelf: "center",
        flexDirection: "row",
        width: Dim.width * 0.85,
        backgroundColor: Colors.darkGreen,
        paddingLeft: 10,
        borderRadius: BorderRadius,
        ...props.style,
      }}
    >
      <View style={{ justifyContent: "center", marginLeft: 10 }}>
        <Entypo name="magnifying-glass" color={Colors.white} size={20} />
      </View>
      <ElementsSearch
        placeholder={
          props.placeholder == undefined ? "Search" : props.placeholder
        }
        placeholderTextColor={Colors.white}
        value={props.value}
        onChangeText={props.onChangeText}
        editable={true}
        clearIcon={{color: Colors.white}}
        searchIcon={null}
        containerStyle={{
          paddingTop: 0,
          backgroundColor: "transparent",
          paddingBottom: 0,
          flex: 1,
          borderBottomColor: "transparent",
          borderTopColor: "transparent"
        }}
        inputContainerStyle={{
          borderRadius: BorderRadius,
          backgroundColor: Colors.darkGreen
        }}
        inputStyle={{
          color: Colors.white,
          fontSize: Font.s2.size,
          fontFamily: Font.s2.family,
          fontWeight: Font.s2.weight,
        }}
        autoFocus={props.autoFocus}
        onSubmitEditing={props.onSubmitEditing}
      />
    </View>
  );
};

export default SearchBar;
