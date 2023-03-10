import React, {useState} from 'react';
import { SafeAreaView, StyleSheet, Pressable, Image, Text, Title, View, ScrollView } from 'react-native';
import TextInput from '../components/TextInput';
import Button from '../components/Button';
import {Colors, Font, Dim} from '../Constants';
import StoresCard from "../components/StoresCard";
import useRequest from "../hooks/useRequest";

function PickStore({navigation, route}) {
    const [storeData, setStoresData] = useState({});
    const [storeName, setStoreName] = useState('');
    const [storeId, setStoreId] = useState('');


    // const getStores = useRequest({
    //     url:  `/community/stores/${route.params.communityId}`,
    //     method: "get",
    //     onSuccess: (data) => {
    //       const stores = [];
    //       data.forEach((storeData) => {
    //         stores.push({
    //           name: storeData.name,
    //           id: storeData.id,
    //         });
    //       });
    //       setStoresData(stores);
    //     },
    //   });

    // const allStores = async () => await getStores.doRequest();

    // useEffect(() => {
    //     allStores();
    //   }, []);

    return (
        <SafeAreaView style={styles.container}>
            <View style={{marginTop: 20, marginLeft: 20}}>
                <Text style={styles.title}>Pick your Store in </Text>
                <Text style={styles.title}>{route.params.communityName}</Text>
            </View>
            <View style={{ marginTop: 30, ...styles.minWrapper }}>
            <StoresCard
                communityId={route.params.communityId}
                onSelectStore={(data) => {
                    setStoreName(data.name);
                    setStoreId(data.store_id);
                    console.log(storeId);
                }}
                width={Dim.width * 0.9}
            ></StoresCard>
            </View>
            <Button
              title={"Pick your Items"}
              appButtonContainer={styles.button}
              width={Dim.width * 0.5}
              onPress={() => {
                  navigation.navigate("Buy", {
                    communityName: route.params.communityName,
                    communityId: route.params.communityId,
                    storeName: storeName,
                    storeId: storeId,
                  });
              }}
            />
        </SafeAreaView>

    );

}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  title: {
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
  },
  minWrapper: {
    width: Dim.width * 0.9,
    alignSelf: "center",
  },
  button: {
    alignSelf: "center",
    backgroundColor: Colors.lightGreen,
    marginTop: 30,
  },
});

export default PickStore;