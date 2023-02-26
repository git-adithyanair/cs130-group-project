import React from 'react';
import { SafeAreaView, StyleSheet, Text, Image, ScrollView, View, TouchableOpacity } from 'react-native';
import ItemCard from '../components/ItemCard'
import Button from '../components/Button'
import {Colors} from '../Constants'

function Errand({setPage}) {
  const testData = [
    {
      itemName: "Pinklady Apples",
      numOfItem: 10,
      id: 1
    },
    {
      itemName: "Celery",
      numOfItem: 20,
      id: 2
    },
    {
      itemName: "Celery",
      numOfItem: 20,
      id: 2
    },
    {
      itemName: "Celery",
      numOfItem: 20,
      id: 2
    },
    {
      itemName: "Celery",
      numOfItem: 20,
      id: 2
    },
    {
      itemName: "Celery",
      numOfItem: 20,
      id: 2
    },
    {
      itemName: "Celery",
      numOfItem: 20,
      id: 2
    },
    {
      itemName: "Celery",
      numOfItem: 20,
      id: 2
    }
  ]

  const itemList = testData.map((item) => <View style={styles.itemCard}><ItemCard itemName={item.itemName} numOfItem={item.numOfItem} key={item.key}/></View>);
    return <SafeAreaView style={styles.container}>
    <View style={styles.content}>
    <Image source={require("../assets/logo.png")}/>
    <Image style={styles.profileImage} source={{uri: "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460__340.png"}}/>
    <Text>Store: Ralphs</Text>
    <Text>Item Count: 20</Text>
    <Button title="Go Back" width="30%" backgroundColor={Colors.darkGreen} onPress={()=>setPage(0)}/> 
    <ScrollView style={styles.listOfItems}>
        {itemList}
      </ScrollView>
    </View>
  </SafeAreaView>; 

}


const styles = StyleSheet.create({
    container: {
      flex: 1,
      backgroundColor: '#fff'
    },
    content: {
      alignItems: 'center',
      marginTop: 40
    }, 
    titleText:{
      fontSize: 25
    },
    subtitleText:{
      fontSize: 14
    },
    profileImage: {
      width: 75,
      height: 75,
      borderRadius: 75 / 2,
      marginTop: 15
    },
    itemCard: {
      marginTop: 20
    },
    listOfItems:{
      marginBottom: 270,
      width: '75%',
      marginTop: 15
    },
    titleCard: {
      display: 'flex',
      flexDirection: 'row',
      backgroundColor: 'red'
    }
  }); 

export default Errand;