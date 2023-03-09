import React, {useState} from 'react';
import { SafeAreaView, StyleSheet, Pressable, Image, Text, Title, View, ScrollView } from 'react-native';
import TextInput from '../components/TextInput';
import Button from '../components/Button';
import {Colors, Font, Dim} from '../Constants';
import LocationFinderCard from "../components/LocationFinderCard";
import useRequest from "../hooks/useRequest";

function ChangeAddress({navigation}) {
    const [locationData, setLocationData] = useState({});

    const handleLocation = (e) => {
      setLocationData(e)
    };

    const updateAddress = useRequest({
      url: "/user/update-location",
      method: "post",
      body: {
        address: locationData.address,
        place_id: locationData.place_id,
        x_coord: locationData.x_coord,
        y_coord: locationData.y_coord,
      },
      onSuccess: (data) => {
          props.navigation.navigate("YourCommunities");
      },
      onFail: (data) => {
          console.log(locationData);
      },
    });
    return (
        <SafeAreaView style={styles.container}>
            <View style={{marginTop: 20, marginLeft: 20}}>
                <Text style={styles.title}>Change your Address</Text>
            </View>
            <View style={{ marginTop: 30, ...styles.minWrapper }}>
            <LocationFinderCard
              searchLabel="Your Address"
              placeholder={"Enter your new address"}
              width={Dim.width * 0.9}
              onSelectLocation={(e) => handleLocation(e)}
            />
            </View>
            <Button
              title={"Submit"}
              appButtonContainer={styles.button}
              width={Dim.width * 0.5}
              onPress={async () => await updateAddress.doRequest()}
            />
              {/* <View style ={{marginTop: 20}}>
              <LocationFinderCard
                searchLabel="You Address"
                placeholder={"Enter your New Address"}
                width={Dim.width * 0.9}
                onSelectLocation={(e) => handleLocation(e)}
              />

            </View>
            <View style={{alignSelf: 'center'}}>
            <Button
                    title={"Submit"}
                    onPress={async () => await updateAddress.doRequest()}
                    textColor={"white"}
                    backgroundColor={Colors.lightGreen}
                    width={250}>
                </Button>
            </View> */}
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

export default ChangeAddress;