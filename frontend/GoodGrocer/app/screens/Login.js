import React, {useState, useEffect} from 'react';
import TextInput from '../components/TextInput'
import Button from '../components/Button'
import { SafeAreaView, StyleSheet, Text, Image, View } from 'react-native';
import { useIsFocused } from "@react-navigation/native";

const Login =  ({navigation}) => {
    const isFocused = useIsFocused(); 
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState(''); 
    const [failedLogIn, setFailedLogIn] = useState(false); 

    useEffect(() => {
      setFailedLogIn(false);
    }, [isFocused])

    const handleLogin = () => {
      console.log(email)
      console.log(password)
      fetch('http://api.good-grocer.click/user/login', {
        method: 'POST',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          email,
          password
        }),
      }).then(response => response.json())
      .then(json => {
        if(json.token){
          console.log(json.token)
          navigation.navigate("LoggedInHome")
        }
        else{
          setFailedLogIn(true)
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
        <Text style={styles.titleText}>Welcome Back</Text>
        <Text>Email or Phone Number</Text>
        <TextInput onChange={email=>setEmail(email.nativeEvent.text)}/>
        <Text>Password</Text> 
        <TextInput onChange={password => setPassword(password.nativeEvent.text)}/>
        <Button title={"Sign In"} onPress={() => handleLogin()} textColor={"white"} backgroundColor={"#0070CA"} width={300} />
        <Text style={styles.errorMessageText}>{failedLogIn ? "Invalid Credentials" : ""}</Text>
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
  },
  errorMessageText:{
    color: "red",
    textAlign: "center",
    paddingTop: 10
  }
});

export default Login;