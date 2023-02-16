import React from 'react';
import TextInput from '../components/TextInput';
import Button from '../components/Button'; 
import { SafeAreaView, StyleSheet, Text, Image, View, Pressable } from 'react-native';

function AddressSignup({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <View>
        <Image source={require("../assets/logo.png")}/>
        <Text style={styles.titleText}>Address</Text>
        <Text>Find your address</Text>
        <TextInput/>
        <Text style={styles.titleText}>Add a picture</Text>
        <View style={styles.defaultPic}>
        <Image source={require("../assets/default-profile-pic.png")}/>
        </View>
        <Button title={"Sign Up"} onPress={() => navigation.navigate('LoggedInHome')} textColor={"white"} backgroundColor={"#0070CA"} width={300} />
        </View>
        </SafeAreaView>
    );

}

const styles = StyleSheet.create({
  names: {
    flexDirection: 'row', 
    flexWrap: 'wrap',
    justifyContent: 'space-between'
  }, 
  container: {
    flex: 1,
    backgroundColor: '#fff',
    justifyContent: 'center',
    alignItems: 'center'
  },
  titleText: {
    fontSize: 25
  },
  defaultPic:{
    alignItems: 'center'
  }
});


export default AddressSignup;