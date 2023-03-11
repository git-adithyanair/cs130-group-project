import { React, useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  View,
  Text,
  Alert,
  FlatList,
  TouchableOpacity,
} from "react-native";
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view";
import { Dim, Colors, GOOGLE_MAPS_API_KEY } from "../Constants";
import TextInput from "../components/TextInput";
import LocationFinderCard from "../components/LocationFinderCard";
import Button from "../components/Button";
import axios from "axios";
import useRequest from "../hooks/useRequest";

const CreateCommunity = (props) => {
  const [name, setName] = useState("");
  const [range, setRange] = useState("");
  const [locationData, setLocationData] = useState({});
  const [stores, setStores] = useState([]);
  const [noResults, setNoResults] = useState(false);
  const [selectedStoreLocation, setStoreSelectedLocation] = useState({});
  const [storesSelected, setStoresSelected] = useState([]);

  const handleName = (e) => {
    setName(e);
  };

  const handleRange = (e) => {
    setRange(e);
  };

  const handleLocation = (e) => {
    setLocationData(e);
  };

  const searchStores = async () => {
    axios
      .get(`https://maps.googleapis.com/maps/api/place/textsearch/json`, {
        params: {
          fields: "formatted_address,place_id,geometry,name",
          query: "grocery store",
          location: `${locationData.x_coord},${locationData.y_coord}`,
          radius: 100,
          key: GOOGLE_MAPS_API_KEY,
        },
      })
      .then(({ data }) => {
        if (data.status == "ZERO_RESULTS") {
          setStores([]);
          setNoResults(true);
        } else if (data.status != "INVALID_REQUEST") {
          setStores(data.results);
          setNoResults(false);
        }
        setStoreSelectedLocation({});
      })
      .catch((error) => {
        console.error("ERROR", error);
      });
  };

  const createCommunity = useRequest({
    url: "/community",
    method: "post",
    body: {
      name: name,
      place_id: locationData.place_id,
      center_x_coord: locationData.x_coord,
      center_y_coord: locationData.y_coord,
      range: parseInt(range),
      address: locationData.address,
      stores: storesSelected,
    },
    onSuccess: (data) => {
      console.log("AHH", data);
      props.navigation.navigate("YourCommunities");
    },
  });

  return (
    <SafeAreaView style={styles.wrapper}>
      <KeyboardAwareScrollView
        showsVerticalScrollIndicator={false}
        extraScrollHeight={30}
        keyboardShouldPersistTaps="handled"
      >
        <View style={styles.minWrapper}>
          <Text style={{ ...styles.title, marginTop: 20 }}>Community Name</Text>
          <TextInput
            placeholder="Enter Community Name"
            onChange={(e) => handleName(e)}
          />
        </View>
        <View style={styles.minWrapper}>
          <Text style={styles.title}>Community Range</Text>
          <TextInput
            placeholder="Enter Community Range (m)"
            onChange={(e) => handleRange(e)}
            keyboardType={"numeric"}
          />
        </View>
        <View style={styles.minWrapper}>
          <LocationFinderCard
            searchLabel="Community address"
            placeholder={"Enter Community Address"}
            width={Dim.width * 0.9}
            onSelectLocation={(e) => handleLocation(e)}
          />
        </View>
        <Button
          title="Add Stores"
          appButtonContainer={styles.addStoreButton}
          width={Dim.width * 0.5}
          onPress={() => {
            if (!locationData) {
              Alert.alert("Oops!", "Please fill out community location field.");
            } else {
              searchStores();
            }
          }}
        />
        <Text style={{ paddingLeft: 20, paddingTop: 10 }}>
          {noResults
            ? "No results."
            : stores.length !== 0
            ? "Select stores →"
            : null}
        </Text>
        <FlatList
          horizontal
          showsHorizontalScrollIndicator={true}
          contentContainerStyle={{ marginTop: 10 }}
          style={styles.list}
          data={stores}
          renderItem={(itemData) => (
            <StoresAddressCard
              name={itemData.item.name}
              address={itemData.item.formatted_address}
              selected={
                itemData.item.place_id === selectedStoreLocation.place_id
              }
              onPress={(isSelected) => {
                const sendData = isSelected
                  ? {
                      name: itemData.item.name,
                      place_id: itemData.item.place_id,
                      x_coord: itemData.item.x_coord,
                      y_coord: itemData.item.y_coord,
                      address: itemData.item.formatted_address,
                    }
                  : {};
                setStoreSelectedLocation(sendData);
                if (isSelected) {
                  setStoresSelected([...storesSelected, sendData]);
                } else {
                  setStoresSelected(
                    storesSelected.filter(
                      (store) => store.place_id !== itemData.item.place_id
                    )
                  );
                }
              }}
            />
          )}
          keyExtractor={(item) => item.place_id}
          ItemSeparatorComponent={() => <View style={{ width: 10 }} />}
        />
        <Button
          title="Create Community"
          appButtonContainer={{
            alignSelf: "center",
            marginBottom: 20,
            marginTop: 40,
          }}
          width={Dim.width * 0.7}
          onPress={async () => await createCommunity.doRequest()}
        />
      </KeyboardAwareScrollView>
    </SafeAreaView>
  );
};

const StoresAddressCard = (props) => {
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
            width: Dim.width * 0.9 - 50,
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
  wrapper: { flex: 1, backgroundColor: Colors.white },
  minWrapper: {
    width: Dim.width * 0.9,
    alignSelf: "center",
  },
  title: {
    marginBottom: 5,
    fontWeight: "bold",
  },
  addStoreButton: {
    alignSelf: "center",
    backgroundColor: Colors.lightGreen,
    marginTop: 40,
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
  list: {
    paddingHorizontal: 20,
  },
});

export default CreateCommunity;
