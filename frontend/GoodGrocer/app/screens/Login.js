import React from 'react';
import { Pressable, SafeAreaView, StyleSheet, Text, Image, TextInput, View } from 'react-native';

function Login({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <View>
        <Image source={require("../assets/logo.png")}/>
        <Text style={styles.titleText}>Welcome Back</Text>
        <Text>Email or Phone Number</Text>
        <TextInput style={styles.input}/>
        <Text>Password</Text> 
        <TextInput style={styles.input}/>
        <Pressable onPress={() => navigation.navigate('Requests')}> 
        <Image source={require("../assets/signinbutton.png")}/>
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

export default Login;