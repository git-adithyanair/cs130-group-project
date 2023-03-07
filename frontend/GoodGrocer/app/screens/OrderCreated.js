import React from 'react';
import { SafeAreaView, StyleSheet, Image, Text, View } from 'react-native';
import {Font} from '../Constants';

function OrderCreated({navigation}) {
    return (
        <SafeAreaView style={styles.container}>
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
    fontSize: Font.s2.size,
    fontFamily: Font.s2.family,
    fontWeight: Font.s2.weight,
    textAlign: 'center',
  }
});

export default OrderCreated;