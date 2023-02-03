import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, TextInput, View, Pressable } from 'react-native';

function Signup({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <Image source={require("../assets/logo.png")}/>
        <Text>Signup</Text>
        <Text>Email</Text>
        <TextInput style={styles.input}/>
        <Text>Phone Number</Text>
        <TextInput style={styles.input}/>
        
        <View style={styles.names}>
            <Text>First Name</Text>
            <Text>Last Name                                     </Text>
        <TextInput style={halfInputStyle}/>
        <TextInput style={halfInputStyle}/>
        </View>

        <Text>Password</Text>
        <TextInput style={styles.input}/>
        <Pressable onPress={() => navigation.navigate('AddressSignup')}> 
        <Image source={require("../assets/continueaddress.png")}/>
        </Pressable>
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
    justifyContent: 'center'
  },
  input: {
    height: 40,
    margin: 12,
    borderWidth: 1,
    padding: 10,
  }, 
  halfInput: {
    width: "40%"
  }
});
const halfInputStyle = {...styles.input, ...styles.halfInput}


export default Signup;