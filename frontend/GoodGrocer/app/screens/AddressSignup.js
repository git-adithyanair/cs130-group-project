import React, {useState} from 'react';
import TextInput from '../components/TextInput';
import Button from '../components/Button'; 
import { SafeAreaView, StyleSheet, Text, Image, View } from 'react-native';

function AddressSignup({route, navigation}) {
    const { email, name, phoneNumber, password } = route.params; 
    const [address, setAddress] = useState('')
    const [signupErrorMsg, setSignupErrorMsg] = useState("") 

    const handleSignup = () => {
      fetch('http://api.good-grocer.click/user', {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email,
          password,
          full_name: name,
          phone_number: phoneNumber,
          address, 
          place_id: "test",
          x_coord: 2.4933,
          y_coord: 3.359
        }),
      }).then(response => response.json())
      .then(json => {
        if(json.error){
          console.log(json.error)
          setSignupErrorMsg(json.error)
        }
        else{
          navigation.navigate("LoggedInHome")
        }
      })
      .catch(error => {
        console.error(error);
      });
  }

    return (
        <SafeAreaView style={styles.container}>
        <View>
        <Image source={require("../assets/logo.png")}/>
        <Text style={styles.titleText}>Address</Text>
        <Text>Find your address</Text>
        <TextInput onChange={address => setAddress(address.nativeEvent.text)}/>
        <Text style={styles.titleText}>Add a picture</Text>
        <View style={styles.defaultPic}>
        <Image source={require("../assets/default-profile-pic.png")}/>
        </View>
        <Button title={"Sign Up"} onPress={() => handleSignup()} textColor={"white"} backgroundColor={"#0070CA"} width={300} />
        </View>
        <Text style={styles.errorMessageText}>{signupErrorMsg}</Text>
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
  },
  errorMessageText:{
    color: "red",
    textAlign: "center",
    paddingTop: 10
  }
});


export default AddressSignup;