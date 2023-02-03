import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, TextInput, View, Pressable } from 'react-native';

function AddressSignup({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
        <Image source={require("../assets/logo.png")}/>
        <Text>Address</Text>
        <Text>Address line 1</Text>
        <TextInput style={styles.input}/>
        <Text>Address line 2</Text>
        <TextInput style={styles.input}/>
        <Text>City</Text>
        <TextInput style={styles.input}/>
        
        <View style={styles.names}>
            <Text>State</Text>
            <Text>Zipcode                                     </Text>
        <TextInput style={halfInputStyle}/>
        <TextInput style={halfInputStyle}/>
        </View>


        <Pressable onPress={() => navigation.navigate('Requests')}> 
        <Image source={require("../assets/signup2.png")}/>
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


export default AddressSignup;