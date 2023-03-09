import React, { useState, useEffect } from "react";
import {
  TouchableOpacity,
  StyleSheet,
  Text,
  View,
  FlatList,
} from "react-native";
import { Colors } from "../Constants";
import useRequest from "../hooks/useRequest";

// required things to pass in:
// - width (width of component)
// - onSelectStore(data) to get data for a store when it is selected
//                 data is in form {address: "", place_id: "", name: 0.0, store_id: 0.0}
// - communityId to know what community to get stores for
/* Example: 
    <StoresCard
        communityId={props.communityId}
        onSelectStore={(data) => {
          //do something with the data
        }}
        width={Dim.width * 0.9}
    ></StoresCard>
*/
const StoresCard = (props) => {
  const [data, setData] = useState([]);
  const [selectedStore, setSelectedStore] = useState({});
  const [loading, setLoading] = useState(true);

  const getCommunityStores = useRequest({
    url: `/community/stores/${props.communityId}`,
    method: "get",
    onSuccess: (data) => {
      setData(data);
    },
  });

  useEffect(() => {
    if (loading) {
      const getStores = async () => getCommunityStores.doRequest();
      getStores();
      setLoading(false);
    }
  }, [loading]);

  return (
    <View style={{ ...styles.container, width: props.width }}>
      <Text style={{ fontWeight: "bold" }}>Pick your store</Text>
      <FlatList
        horizontal
        showsHorizontalScrollIndicator={true}
        contentContainerStyle={{ marginTop: 10 }}
        pagingEnabled={true}
        style={styles.list}
        data={data}
        renderItem={(itemData) => (
          <StoreCard
            name={itemData.item.name}
            address={itemData.item.address}
            selected={itemData.item.place_id === selectedStore.place_id}
            width={props.width - 9}
            onPress={(isSelected) => {
              console.log(selectedStore);
              const sendData = isSelected
                ? {
                    address: itemData.item.address,
                    place_id: itemData.item.place_id,
                    name: itemData.item.name,
                    store_id: itemData.item.id,
                  }
                : {};
              setSelectedStore(sendData);
              console.log(selectedStore);
              props.onSelectStore(sendData);
            }}
          />
        )}
        keyExtractor={() => Math.random().toString()}
        ItemSeparatorComponent={() => <View style={{ width: 10 }} />}
      ></FlatList>
    </View>
  );
};

const StoreCard = (props) => {
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

export default StoresCard;
