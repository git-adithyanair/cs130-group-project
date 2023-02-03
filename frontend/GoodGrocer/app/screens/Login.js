import React from 'react';
import { Pressable, SafeAreaView, StyleSheet, Text, Image, TextInput } from 'react-native';

function Login({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <Image source={require("../assets/logo.png")}/>
        <Text>Welcome Back</Text>
        <Text>Email or Phone Number</Text>
        <TextInput style={styles.input}/>
        <Text>Password</Text> 
        <TextInput style={styles.input}/>
        <Pressable onPress={() => navigation.navigate('Requests')}> 
        <Image source={require("../assets/signinbutton.png")}/>
        </Pressable>
        </SafeAreaView>
    );

}


const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    justifyContent: 'center'
    

  },
  input: {
    height: 40,
    margin: 12,
    borderWidth: 1,
    padding: 10,
  }
});

export default Login;