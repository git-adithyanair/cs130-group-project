import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, TextInput, View, Pressable } from 'react-native';

function AddressSignup({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <View>
        <Image source={require("../assets/logo.png")}/>
        <Text style={styles.titleText}>Address</Text>
        <Text>Find your address</Text>
        <TextInput style={styles.input}/>
        <Text style={styles.titleText}>Add a picture</Text>
        <View style={styles.defaultPic}>
        <Image source={require("../assets/default-profile-pic.png")}/>
        </View>
        <Pressable onPress={() => navigation.navigate('Requests')}> 
        <Image source={require("../assets/signup2.png")}/>
        </Pressable>
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
  input: {
    height: 40,
    marginTop: 12,
    marginBottom: 12,
    borderWidth: 1,
    padding: 10,
  }, 
  titleText: {
    fontSize: 25
  },
  defaultPic:{
    alignItems: 'center'
  }
});


export default AddressSignup;