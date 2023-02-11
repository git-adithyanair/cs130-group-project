import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, TextInput, View, Pressable } from 'react-native';
import Login from './Login'; 
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';



const Tab = createBottomTabNavigator(); 



function Shop({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
          <View style={styles.content}>
            <Image source={require("../assets/logo.png")}/>
            <Text>Shop</Text>
            <View style={styles.listOfRequests}>
            <View><Text>Angela</Text></View>
            <View style={styles.requestDetails}>
                <View><Text>Neighborhood: Westwood</Text></View>
                <View><Text>Store: Trader Joes</Text></View>
                <View><Text>Items: 15</Text></View>
            </View>
            </View>
            </View>
        </SafeAreaView>
    );

}


const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#fff',
    },
    content: {
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

export default Shop;