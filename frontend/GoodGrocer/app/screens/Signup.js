import React from 'react';
import TextInput from '../components/TextInput';
import Button from '../components/Button';
import { SafeAreaView, StyleSheet, Text, Image, View, Pressable } from 'react-native';

function Signup({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <View> 
        <Image source={require("../assets/logo.png")}/>
        <Text style={styles.titleText}>Signup</Text>
        <Text>Email</Text>
        <TextInput/>
        <Text>Phone Number</Text>
        <TextInput/>
        <Text>Name</Text>
        <TextInput/>
        <Text>Password</Text>
        <TextInput/>
        <Button title={"Continue with Address"} onPress={() => navigation.navigate('AddressSignup')} textColor={"white"} backgroundColor={"#0070CA"} width={300} />
        </View>
        </SafeAreaView>
    );

}



const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    justifyContent: 'center',
    alignItems: 'center'
  },
  titleText: {
    fontSize: 25
  }
});


export default Signup;