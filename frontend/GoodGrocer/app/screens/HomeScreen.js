import React from 'react';
import { SafeAreaView, StyleSheet, Pressable, Image } from 'react-native';

function HomeScreen({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <Image source={require("../assets/logo.png")}/>
        <Image source={require("../assets/slogan.png")}/>
        <Pressable onPress={() => navigation.navigate('Login')}> 
          <Image source={require("../assets/login.png")} /> 
        </Pressable>
        <Pressable onPress={() => navigation.navigate('Signup')}> 
          <Image source={require("../assets/signup.png")}/>
        </Pressable>
        </SafeAreaView>
    );

}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});

export default HomeScreen;