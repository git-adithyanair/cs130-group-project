import React from 'react';
import TextInput from '../components/TextInput'
import Button from '../components/Button'
import { SafeAreaView, StyleSheet, Text, Image, View } from 'react-native';

function Login({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <View>
        <Image source={require("../assets/logo.png")}/>
        <Text style={styles.titleText}>Welcome Back</Text>
        <Text>Email or Phone Number</Text>
        <TextInput/>
        <Text>Password</Text> 
        <TextInput/>
        <Button title={"Sign In"} onPress={() => navigation.navigate('LoggedInHome')} textColor={"white"} backgroundColor={"#0070CA"} width={300} />
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

export default Login;