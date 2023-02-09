import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, TextInput, View, Pressable } from 'react-native';






function Requests({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
            <Image source={require("../assets/logo.png")}/>
            <Text>Requests</Text>
            <View style={styles.listOfRequests}>
            <View><Text>Angela</Text></View>
            <View style={styles.requestDetails}>
                <View><Text>Neighborhood: Westwood</Text></View>
                <View><Text>Store: Trader Joes</Text></View>
                <View><Text>Items: 15</Text></View>
            </View>
            </View>
        </SafeAreaView>
    );

}


const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#fff',
      alignItems: 'center'
    },
    listOfRequests: {
      display: "flex",
      flexDirection: "row",
      paddingTop: 20
    },
    requestDetails:{
      flexDirection: "column",
      paddingLeft: 10
    }
  });

export default Requests;