import React from 'react';
import { SafeAreaView, StyleSheet, Image, Text, View } from 'react-native';
import {Font} from '../Constants';

function OrderCreated({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
            <View style={{alignItems: 'center'}}>
                <Image source={require("../assets/logo.png")}/>
            </View>
            <View style={{flex: 1, justifyContent: 'center', margin: 15}}>
                <Text style={styles.title}>Order complete!</Text>
                <Text style={styles.title}>It is waiting to be picked up...</Text>
            </View>
        </SafeAreaView>


    );

}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  title: {
    fontSize: Font.s1.size,
    fontFamily: Font.s1.family,
    fontWeight: Font.s1.weight,
    textAlign: 'center',
  }
});

export default OrderCreated;