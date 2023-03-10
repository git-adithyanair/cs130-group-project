import React, { useState } from "react";
import {
  TouchableOpacity,
  StyleSheet,
  Text,
  View,
  FlatList,
} from "react-native";
import { Colors, GOOGLE_MAPS_API_KEY } from "../Constants";
import TextInput from "../components/TextInput";
import Button from "./Button";
import axios from "axios";

// pass in onSelectLocation(data) to get data for a location when it is selected
// data is in form {address: "", place_id: "", x_coord: 0.0, y_coord: 0.0}
const LocationFinderCard = (props) => {
  const [address, setAddress] = useState("");
  const [data, setData] = useState([]);
  const [selectedLocation, setSelectedLocation] = useState({});
  const [noResults, setNoResults] = useState(false);

  const search = async () => {
    axios
      .get(`https://maps.googleapis.com/maps/api/place/textsearch/json`, {
        params: {
          fields: "formatted_address,place_id,geometry,name",
          query: address,
          key: GOOGLE_MAPS_API_KEY,
        },
      })
      .then(({ data }) => {
        if (data.status == "ZERO_RESULTS") {
          setData([]);
          setNoResults(true);
        } else if (data.status != "INVALID_REQUEST") {
          setData(data.results);
          setNoResults(false);
        }
        setSelectedLocation({});
        props.onSelectLocation({});
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
        onChange={(address) => setAddress(address)}
        placeholder={
          props.placeholder ? props.placeholder : "Enter your address..."
        }
      />
      <Button
        title={"Search"}
        onPress={() => {
          search();
        }}
        textColor={"white"}
        width={props.width}
        appButtonContainer={{ marginBottom: 20 }}
      />
      <Text>
        {noResults
          ? "No results."
          : data.length !== 0
          ? "Select the correct address →"
          : null}
      </Text>
      <FlatList
        horizontal
        showsHorizontalScrollIndicator={true}
        contentContainerStyle={{ marginTop: 10 }}
        pagingEnabled={true}
        style={styles.list}
        data={data}
        renderItem={(itemData) => (
          <AddressCard
            name={itemData.item.name}
            address={itemData.item.formatted_address}
            selected={itemData.item.place_id === selectedLocation.place_id}
            width={props.width - 9}
            onPress={(isSelected) => {
              const sendData = isSelected
                ? {
                    address: itemData.item.formatted_address,
                    place_id: itemData.item.place_id,
                    x_coord: itemData.item.geometry.location.lat,
                    y_coord: itemData.item.geometry.location.lng,
                  }
                : {};
              setSelectedLocation(sendData);
              props.onSelectLocation(sendData);
            }}
          />
        )}
        keyExtractor={() => Math.random().toString()}
        ItemSeparatorComponent={() => <View style={{ width: 10 }} />}
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
          style={{
            ...(selected
              ? styles.addressContainerSelected
              : styles.addressContainerUnselected),
            width: props.width,
          }}
        >
          <Text
            style={{
              fontWeight: "bold",
              color: selected ? "black" : Colors.darkGreen,
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
    paddingHorizontal: 20,
    borderRadius: 10,
  },
  addressContainerSelected: {
    backgroundColor: Colors.lightGreen,
    alignContent: "center",
    paddingVertical: 15,
    paddingHorizontal: 20,
    borderRadius: 10,
  },
});

export default LocationFinderCard;
