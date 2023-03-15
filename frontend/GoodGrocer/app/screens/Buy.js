import React, { useState } from "react";
import {
  SafeAreaView,
  Alert,
  StyleSheet,
  FlatList,
  Text,
  View,
  TouchableOpacity,
  ImageBackground,
  Image,
} from "react-native";
import Modal from "react-native-modal";
import { Picker } from "@react-native-picker/picker";
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view";
import { Ionicons } from "@expo/vector-icons";
import * as ImagePicker from "expo-image-picker";

import Button from "../components/Button";
import TextInput from "../components/TextInput";
import ItemCard from "../components/ItemCard";
import { Colors, Font, Dim } from "../Constants";
import useRequest from "../hooks/useRequest";

const Buy = ({ navigation, route }) => {
  const [item, setItem] = useState("");
  const [numItems, setNumItems] = useState("");
  const [type, setType] = useState("numerical");
  const [brand, setBrand] = useState("");
  const [notes, setNotes] = useState("");
  const [image, setImage] = useState("");
  const [items, setItems] = useState([]);
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleModal = () => setIsModalVisible(() => !isModalVisible);

  const addItem = () => {
    setIsModalVisible(() => !isModalVisible);
    if (!item || !numItems || !type) {
      Alert.alert(
        "Oops!",
        "You need the name, quantity and quantity type of an item to add it to your list!"
      );
    } else {
      var amount = parseFloat(numItems);
      let individualItem = {
        name: item,
        quantity_type: type,
        quantity: amount,
        preferred_brand: brand,
        image: image,
        extra_notes: notes,
      };
      setItem("");
      setType("");
      setNumItems("");
      setBrand("");
      setNotes("");
      setImage("");
      setItems((prev) => [...prev, individualItem]);
    }
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
        setImage(result.assets[0].base64);
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

  console.log(type);

  return (
    <SafeAreaView style={styles.container}>
      <View style={{ marginTop: 20, marginBottom: 30 }}>
        <View style={{ marginLeft: 20 }}>
          <Text style={styles.title}>
            Complete your order in {route.params.communityName}.
          </Text>
        </View>
        <View style={styles.content}>
          <Button
            title={"Add Items"}
            onPress={handleModal}
            width={Dim.width * 0.5}
            appButtonContainer={styles.button}
          />
          <Modal
            isVisible={isModalVisible}
            transparent={true}
            style={styles.modalStyle}
          >
            <TouchableOpacity
              style={{ alignSelf: "flex-end" }}
              onPress={() => setIsModalVisible(false)}
            >
              <Ionicons
                name={"close-circle"}
                size={35}
                color={Colors.darkGreen}
              />
            </TouchableOpacity>
            <KeyboardAwareScrollView
              showsVerticalScrollIndicator={false}
              extraScrollHeight={30}
              keyboardShouldPersistTaps="handled"
              style={styles.innerModal}
            >
              <View style={{ marginTop: 10 }}>
                <Text style={styles.modalFont}>Item name</Text>
              </View>
              <View style={styles.modalTextinput}>
                <TextInput
                  onChange={(item) => setItem(item)}
                  placeholder="Enter item name..."
                />
              </View>
              <View>
                <View style={styles.content}>
                  <Text style={{ ...styles.modalFont, marginLeft: 0 }}>
                    Quantity Type
                  </Text>
                </View>
                <Picker
                  selectedValue={type}
                  onValueChange={(itemValue, itemIndex) => setType(itemValue)}
                  testID={'picker-select'}
                >
                  <Picker.Item label="count" value="numerical" />
                  <Picker.Item label="lbs" value="lbs" />
                  <Picker.Item label="oz" value="oz" />
                  <Picker.Item label="fl_oz" value="fl_oz" />
                  <Picker.Item label="gal" value="gal" />
                  <Picker.Item label="litres" value="litres" />
                </Picker>
              </View>
              <View>
                <Text style={styles.modalFont}>Quantity of Item</Text>
                <View style={styles.modalTextinput}>
                  <TextInput
                    onChange={(numItems) => setNumItems(numItems)}
                    placeholder="Enter item quantity..."
                  />
                </View>
              </View>
              <View>
                <Text style={styles.modalFont}>Preferred Brand</Text>
                <View style={styles.modalTextinput}>
                  <TextInput
                    onChange={(brand) => setBrand(brand)}
                    placeholder="Enter preferred brand..."
                  />
                </View>
              </View>
              <View>
                <Text style={styles.modalFont}>Notes</Text>
                <View style={styles.modalTextinput}>
                  <TextInput
                    onChange={(notes) => setNotes(notes)}
                    placeholder="Enter extra notes..."
                  />
                </View>
              </View>
              <View style={{ marginTop: 20 }}>
                {image ? (
                  <ImageBackground
                    style={{
                      width: Dim.width * 0.25,
                      height: Dim.width * 0.25,
                      alignSelf: "center",
                    }}
                    source={{
                      uri: "data:image/png;base64," + image,
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
              <View
                style={{
                  alignItems: "center",
                  marginBottom: 10,
                  marginTop: 30,
                }}
              >
                <Button
                  title={"Add Item"}
                  onPress={addItem}
                  textColor={"white"}
                  backgroundColor={Colors.lightGreen}
                  width={200}
                />
              </View>
            </KeyboardAwareScrollView>
          </Modal>
        </View>
      </View>

      <View style={{ marginLeft: 20, marginBottom: 10 }}>
        <Text style={styles.title}>Your Items</Text>
      </View>
      <View style={{ flex: 1, ...styles.minWrapper }}>
        <FlatList
          data={items}
          contentContainerStyle={{ paddingBottom: 20 }}
          renderItem={({ item }) => (
            <ItemCard
              name={item.name}
              quantity={item.quantity}
              quantityType={item.quantity_type}
              preferredBrand={item.preferred_brand}
              extraNotes={item.extra_notes}
              imageUri={item.image}
            />
          )}
          keyExtractor={(item) => item.name}
          ItemSeparatorComponent={() => (
            <View
              style={{
                height: 10,
              }}
            />
          )}
          ListFooterComponent={() => (
            <View style={{ alignItems: "center", marginTop: 10 }}>
              <Button
                title={"Complete your Order"}
                onPress={async () => {
                  if (items.length !== 0) {
                    setLoading(true);
                    await createRequest.doRequest();}
                }}
                textColor={"white"}
                backgroundColor={Colors.darkGreen}
                width={300}
                loading={loading}
              />
            </View>
          )}
        />
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
  },
  content: {
    alignItems: "center",
  },
  title: {
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
  },
  modalFont: {
    fontSize: Font.s3.size,
    fontFamily: Font.s3.family,
    fontWeight: Font.s3.weight,
    marginLeft: 20,
  },
  modalStyle: {
    backgroundColor: Colors.cream,
    position: "absolute",
    width: Dim.width,
    height: Dim.height * 0.9,
    bottom: 0,
    borderRadius: 15,
    padding: 20,
    marginLeft: 0,
    marginBottom: 0,
  },
  innerModal: {
    width: "100%",
  },
  modalTextinput: {
    marginLeft: 20,
    marginRight: 20,
  },
  item: {
    backgroundColor: "#f9c2ff",
    padding: 20,
    marginVertical: 8,
    marginHorizontal: 16,
  },
  button: {
    alignSelf: "center",
    backgroundColor: Colors.lightGreen,
    marginTop: 30,
  },
  minWrapper: {
    width: Dim.width * 0.9,
    alignSelf: "center",
  },
});

export default Buy;
