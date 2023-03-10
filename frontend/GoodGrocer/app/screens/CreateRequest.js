import React, { useState } from "react";
import {
  View,
  Text,
  SafeAreaView,
  TouchableOpacity,
  StyleSheet,
  Image,
  ImageBackground,
  FlatList,
  Alert,
} from "react-native";
import * as ImagePicker from "expo-image-picker";
import { KeyboardAwareFlatList } from "react-native-keyboard-aware-scroll-view";

import TextInput from "../components/TextInput";
import { Colors, Dim, Font } from "../Constants";
import ItemCard from "../components/ItemCard";
import Button from "../components/Button";
import useRequest from "../hooks/useRequest";

const QuantityButton = (props) => {
  return (
    <TouchableOpacity
      style={{
        padding: 20,
        backgroundColor: props.selected ? Colors.darkGreen : Colors.cream,
        borderRadius: 15,
        width: Dim.width * 0.4,
        borderWidth: 0.5,
      }}
      onPress={props.onPress}
    >
      <Text
        style={{
          fontSize: Font.s3.size,
          fontWeight: Font.s3.weight,
          textAlign: "center",
        }}
      >
        {props.type}
      </Text>
    </TouchableOpacity>
  );
};

const CreateOrder = ({ navigation, route }) => {
  const [currentItem, setCurrentItem] = useState({
    name: "",
    quantity: 0,
    quantity_type: "numerical",
    preferred_brand: "",
    extra_notes: "",
    image: "",
  });
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);

  const setQuantityType = (type) => {
    setCurrentItem({ ...currentItem, quantity_type: type });
  };

  const pickImage = async () => {
    try {
      const result = await ImagePicker.launchImageLibraryAsync({
        mediaTypes: ImagePicker.MediaTypeOptions.Images,
        allowsEditing: true,
        aspect: [1, 1],
        quality: 0.8,
        base64: true,
      });
      if (!result.canceled) {
        setCurrentItem({ ...currentItem, image: result.assets[0].base64 });
      }
    } catch (err) {
      console.log(err);
    }
  };

  const getImagePickerPermissionAsync = async () => {
    const { status } = await ImagePicker.requestMediaLibraryPermissionsAsync();
    if (status !== "granted") {
      Alert.alert(
        "Oops!",
        "We need access to your photo library to assign a profile picture for you!"
      );
    } else {
      pickImage();
    }
  };

  const requiredBody = {
    community_id: route.params.communityId,
    items: items,
  };

  const createRequest = useRequest({
    url: "/request",
    method: "post",
    body: {
      ...requiredBody,
      ...(route.params.storeId ? { store_id: route.params.storeId } : {}),
    },
    onSuccess: () => {
      setLoading(false);
      navigation.navigate("OrderCreated");
    },
    onFail: () => {
      setLoading(false);
    },
  });

  return (
    <SafeAreaView style={{ width: "90%", alignSelf: "center" }}>
      <KeyboardAwareFlatList
        extraScrollHeight={30}
        keyboardShouldPersistTaps="handled"
        data={items}
        renderItem={({ item }) => (
          <ItemCard
            name={item.name}
            quantity={item.quantity}
            quantityType={item.quantity_type}
            preferredBrand={item.preferred_brand ? item.preferred_brand : null}
            extraNotes={item.extra_notes ? item.extra_notes : null}
            imageUri={item.image ? item.image : null}
          />
        )}
        showsVerticalScrollIndicator={false}
        keyExtractor={(item) => items.indexOf(item).toString()}
        ItemSeparatorComponent={<View style={{ height: 10 }} />}
        ListEmptyComponent={<Text>Begin adding items to get started!</Text>}
        ListHeaderComponent={
          <View style={{ marginVertical: 20 }}>
            <Text
              style={{
                fontSize: Font.s1.size,
                fontFamily: Font.s1.family,
                fontWeight: Font.s1.weight,
              }}
            >
              Complete your request
            </Text>
            <Text style={{ marginTop: 10 }}>
              Add all the items you want your shopper to collect for you!
            </Text>
            <Text style={{ fontWeight: "bold", marginTop: 20 }}>
              Added items ({items.length})
            </Text>
          </View>
        }
        ListFooterComponent={
          <View style={{ marginTop: 30, marginBottom: 40 }}>
            <Text style={{ fontWeight: "bold", marginBottom: 10 }}>
              Add an item
            </Text>
            <TextInput
              placeholder="Enter item name..."
              style={{ marginBottom: 5 }}
              onChange={(text) =>
                setCurrentItem({ ...currentItem, name: text })
              }
              value={currentItem.name}
            />
            <TextInput
              placeholder="Enter item quantity..."
              style={{ marginBottom: 15 }}
              onChange={(text) =>
                setCurrentItem({ ...currentItem, quantity: parseFloat(text) })
              }
              keyboardType="numeric"
              value={
                currentItem.quantity > 0 ? currentItem.quantity.toString() : ""
              }
            />
            <View>
              <View style={styles.buttonContainer}>
                <QuantityButton
                  type="numerical"
                  onPress={() => setQuantityType("numerical")}
                  selected={currentItem.quantity_type === "numerical"}
                />
                <QuantityButton
                  type="oz"
                  onPress={() => setQuantityType("oz")}
                  selected={currentItem.quantity_type === "oz"}
                />
              </View>
              <View style={styles.buttonContainer}>
                <QuantityButton
                  type="lbs"
                  onPress={() => setQuantityType("lbs")}
                  selected={currentItem.quantity_type === "lbs"}
                />
                <QuantityButton
                  type="fl oz"
                  onPress={() => setQuantityType("fl_oz")}
                  selected={currentItem.quantity_type === "fl_oz"}
                />
              </View>
              <View style={styles.buttonContainer}>
                <QuantityButton
                  type="gal"
                  onPress={() => setQuantityType("gal")}
                  selected={currentItem.quantity_type === "gal"}
                />
                <QuantityButton
                  type="litres"
                  onPress={() => setQuantityType("litres")}
                  selected={currentItem.quantity_type === "litres"}
                />
              </View>
            </View>
            <TextInput
              placeholder="Enter preferred brand (empty if none)..."
              style={{ marginBottom: 5 }}
              onChange={(text) =>
                setCurrentItem({ ...currentItem, preferred_brand: text })
              }
              value={currentItem.preferred_brand}
            />
            <TextInput
              placeholder="Enter extra notes (empty if none)..."
              style={{ marginBottom: 5 }}
              onChange={(text) =>
                setCurrentItem({ ...currentItem, extra_notes: text })
              }
              value={currentItem.extra_notes}
            />
            <View style={{ marginTop: 20 }}>
              {currentItem.image ? (
                <ImageBackground
                  style={{
                    width: Dim.width * 0.25,
                    height: Dim.width * 0.25,
                    alignSelf: "center",
                  }}
                  source={{
                    uri: "data:image/png;base64," + currentItem.image,
                  }}
                />
              ) : (
                <View style={{ alignItems: "center" }}>
                  <Image
                    source={require("../assets/grocery-item.png")}
                    style={{
                      width: Dim.width * 0.25,
                      height: Dim.width * 0.25,
                    }}
                  />
                </View>
              )}
            </View>

            <TouchableOpacity
              onPress={async () => await getImagePickerPermissionAsync()}
              style={{
                marginVertical: 20,
                width: Dim.width * 0.5,
                alignSelf: "center",
              }}
            >
              <Text
                style={{
                  color: Colors.darkGreen,
                  textAlign: "center",

                  fontWeight: "bold",
                }}
              >
                Click here to add an item picture..
              </Text>
            </TouchableOpacity>
            <Button
              title="Add item"
              onPress={() => {
                if (currentItem.name === "") {
                  Alert.alert("Oops!", "Please enter an item name.");
                  return;
                }
                if (currentItem.quantity === 0) {
                  Alert.alert("Oops!", "Please enter an item quantity.");
                  return;
                }
                if (currentItem.quantity_type === "") {
                  Alert.alert("Oops!", "Please enter an item quantity type.");
                  return;
                }
                setItems([...items, currentItem]);
                setCurrentItem({
                  name: "",
                  quantity: 0,
                  quantityType: "numerical",
                  preferredBrand: "",
                  extraNotes: "",
                  image: "",
                });
              }}
              appButtonContainer={{
                marginTop: 10,
                width: Dim.width * 0.4,
                alignSelf: "center",
              }}
            />
            <Button
              title="Create Request"
              onPress={async () => {
                setLoading(true);
                await createRequest.doRequest();
              }}
              appButtonContainer={{
                marginTop: 40,
                backgroundColor: Colors.lightGreen,
              }}
              loading={loading}
            />
          </View>
        }
      />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  buttonContainer: {
    flexDirection: "row",
    marginBottom: 10,
    justifyContent: "space-around",
  },
});

export default CreateOrder;
