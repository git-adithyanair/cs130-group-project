import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, TextInput, View, Pressable } from 'react-native';

function Signup({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <View> 
        <Image source={require("../assets/logo.png")}/>
        <Text style={styles.titleText}>Signup</Text>
        <Text>Email</Text>
        <TextInput style={styles.input}/>
        <Text>Phone Number</Text>
        <TextInput style={styles.input}/>
        <Text>Name</Text>
        <TextInput style={styles.input}/>
        <Text>Password</Text>
        <TextInput style={styles.input}/>
        <Pressable onPress={() => navigation.navigate('AddressSignup')}> 
        <Image source={require("../assets/continueaddress.png")}/>
        </Pressable>
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
  input: {
    height: 40,
    marginTop: 12,
    marginBottom: 12, 
    borderWidth: 1,
    padding: 10,
  }, 
  titleText: {
    fontSize: 25
  }
});


export default Signup;