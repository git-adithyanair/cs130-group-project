import React, { useState } from "react";
import {
  TouchableOpacity,
  StyleSheet,
  Text,
  View,
  FlatList,
} from "react-native";
import { Dim, Colors, GOOGLE_MAPS_API_KEY, Font } from "../Constants";
import TextInput from "../components/TextInput";
import Button from "./Button";
import axios from "axios";

// pass in onSelectLocation(data) to get data for a location when it is selected
// data is in form {address: "", place_id: "", x_coord: 0.0, y_coord: 0.0}
const LocationFinderCard = (props) => {
  const [address, setAddress] = useState("");
  const [data, setData] = useState([]);
  const [selectedLocation, setSelectedLocation] = useState({});

  const search = async () => {
    axios
      .get(
        `https://maps.googleapis.com/maps/api/place/findplacefromtext/json`,
        {
          params: {
            fields: "formatted_address,place_id,geometry,name",
            input: address,
            inputtype: "textquery",
            key: GOOGLE_MAPS_API_KEY,
          },
        }
      )
      .then(({ data }) => {
        if (data.status != "INVALID_REQUEST") {
          console.log(data);
          setData(data.candidates);
        }
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <View style={styles.container}>
      <Text style={{ fontWeight: "bold" }}>{props.searchLabel}</Text>
      <TextInput
        style={{ width: props.width }}
        onChange={(address) => setAddress(address.nativeEvent.text)}
        placeholder="Enter your address..."
      />
      <Button
        title={"Search"}
        onPress={() => {
          search();
        }}
        textColor={"white"}
        width={props.width}
        appButtonContainer={{ marginVertical: 30 }}
      />
      <Text style={{ opacity: data.length === 0 ? 0 : 100 }}>
        Select the correct address below:
      </Text>
      <FlatList
        contentContainerStyle={{ marginTop: 10, width: props.width }}
        style={styles.list}
        data={data}
        renderItem={(itemData) => (
          <AddressCard
            name={itemData.item.name}
            address={itemData.item.formatted_address}
            selected={itemData.item.place_id === selectedLocation.place_id}
            onPress={(isSelected) => {
              console.log(selectedLocation);
              const sendData = isSelected
                ? {
                    address: itemData.item.formatted_address,
                    place_id: itemData.item.place_id,
                    x_coord: itemData.item.geometry.location.lat,
                    y_coord: itemData.item.geometry.location.lng,
                  }
                : {};
              setSelectedLocation(sendData);
              console.log(selectedLocation);
              props.onSelectLocation(sendData);
            }}
          />
        )}
        keyExtractor={() => Math.random().toString()}
        ItemSeparatorComponent={() => (
          <View
            style={{
              height: 15,
              width: Dim.width,
            }}
          />
        )}
      ></FlatList>
    </View>
  );
};

const AddressCard = (props) => {
  const [selected, setSelected] = useState(props.selected);

  const selectItem = () => {
    props.onPress(!selected);
    setSelected(!selected);
  };

  return (
    <View>
      <TouchableOpacity onPress={selectItem}>
        <View
          style={
            selected
              ? styles.addressContainerSelected
              : styles.addressContainerUnselected
          }
        >
          <Text
            style={{
              fontWeight: "bold",
              color: selected ? "black" : Colors.darkGreen,
              fontSize: Font.s2.size,
            }}
          >
            {props.name}
          </Text>
          <Text style={{ color: "black" }}>{props.address}</Text>
        </View>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    backgroundColor: "#fff",
    alignItems: "flex-start",
  },
  addressContainerUnselected: {
    backgroundColor: Colors.cream,
    alignContent: "center",
    paddingVertical: 15,
    paddingHorizontal: 40,
    borderRadius: 10,
    flex: 1,
  },
  addressContainerSelected: {
    backgroundColor: Colors.lightGreen,
    alignContent: "center",
    paddingVertical: 15,
    paddingHorizontal: 40,
    borderRadius: 10,
  },
});

export default LocationFinderCard;